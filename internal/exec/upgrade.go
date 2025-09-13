package exec

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/dobyte/due-cli/internal/log"
	"github.com/dobyte/due-cli/internal/os"
)

const (
	dueFrameworkPrefix  = "github.com/dobyte/due"
	dueFrameworkTagsUrl = "https://api.github.com/repos/dobyte/due/git/refs/tags"
)

const (
	upgradeFailure = "upgrade failure"
	upgradeSuccess = "upgrade success"
)

type tag struct {
	Status string `json:"status,omitempty"`
	Ref    string `json:"ref,omitempty"`
	Object struct {
		Sha string `json:"sha,omitempty"`
	} `json:"object,omitempty"`
}

// Upgrade 升级框架
func Upgrade(dir string, version string) {
	if !os.IsDir(dir) {
		log.Fatal(upgradeFailure, "the dir is not a directory")
	}

	mod := filepath.Join(dir, "go.mod")

	if !os.IsFile(mod) {
		log.Fatal(upgradeFailure, "the go.mod file does not exist")
	}

	content, err := os.ReadFile(mod)
	if err != nil {
		log.Fatal(upgradeFailure, err)
	}

	major, full, sha, err := parseTag(version)
	if err != nil {
		log.Fatal(upgradeFailure, err)
	}

	pkg := fmt.Sprintf("%s/%s@%s", dueFrameworkPrefix, major, full)
	cmd := exec.Command("go", "get", pkg)
	cmd.Dir = dir
	cmd.WaitDelay = 30 * time.Second

	if _, err = cmd.Output(); err != nil {
		log.Fatal(upgradeFailure, err)
	}

	reg := regexp.MustCompile(fmt.Sprintf(`(%s/\w+/\w+)/v\d+`, dueFrameworkPrefix))
	rst := reg.FindAllStringSubmatch(string(content), -1)

	for _, group := range rst {
		if len(group) != 2 {
			continue
		}

		pkg = fmt.Sprintf("%s/%s@%s", group[1], major, sha)
		cmd = exec.Command("go", "get", pkg)
		cmd.Dir = dir
		cmd.WaitDelay = 30 * time.Second

		if _, err = cmd.Output(); err != nil {
			log.Fatal(upgradeFailure, err)
		}
	}

	log.Info(upgradeSuccess)
}

// 解析标签
func parseTag(v string) (string, string, string, error) {
	tag, err := loadTag(v)
	if err != nil {
		return "", "", "", err
	}

	reg := regexp.MustCompile(`refs/tags/((v[0-9]+)\.[0-9]+\.[0-9]+)`)
	rst := reg.FindStringSubmatch(tag.Ref)

	if len(rst) != 3 {
		return "", "", "", errors.New("invalid tag")
	}

	return rst[2], rst[1], tag.Object.Sha[:7], nil
}

// 加载标签
func loadTag(v string) (*tag, error) {
	switch v {
	case "":
		return nil, errors.New("the version is empty")
	case "latest":
		tags, err := fetchAllTags()
		if err != nil {
			return nil, err
		}

		return tags[len(tags)-1], nil
	default:
		return fetchTag(v)
	}
}

// 拉取特定标签
func fetchTag(v string) (*tag, error) {
	url := dueFrameworkTagsUrl + "/" + v

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &tag{}

	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	if res.Status == "404" {
		return nil, errors.New("the due version not found")
	}

	if res.Object.Sha == "" {
		return nil, errors.New("the due version not found")
	}

	return res, nil
}

// 拉取所有标签
func fetchAllTags() ([]*tag, error) {
	resp, err := http.Get(dueFrameworkTagsUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tags := make([]*tag, 0)

	if err = json.Unmarshal(body, &tags); err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, errors.New("the due version not found")
	}

	return tags, nil
}
