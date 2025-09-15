package version

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os/exec"
	"regexp"
)

const dueFrameworkTagsUrl = "https://api.github.com/repos/dobyte/due/git/refs/tags"

var (
	defaultGoVersion = "1.22"  // Go版本
	ToolVersion      = "1.0.0" // 工具版本
)

type result struct {
	Status string `json:"status,omitempty"`
	Ref    string `json:"ref,omitempty"`
	Object struct {
		Sha string `json:"sha,omitempty"`
	} `json:"object,omitempty"`
}

// ParseGoVersion 解析go版本
func ParseGoVersion() (string, error) {
	b, err := exec.Command("go", "version").Output()
	if err != nil {
		return "", err
	}

	matches := regexp.MustCompile(`go version go(\d+\.\d+(?:\.\d+)?)`).FindStringSubmatch(string(b))

	if len(matches) > 0 {
		return matches[1], nil
	} else {
		return defaultGoVersion, nil
	}
}

// ParseDueVersion 解析due框架版本
func ParseDueVersion(v string) (string, string, string, error) {
	tag, err := loadTag(v)
	if err != nil {
		return "", "", "", err
	}

	reg := regexp.MustCompile(`refs/tags/((v[0-9]+)\.[0-9]+\.[0-9]+)`)
	rst := reg.FindStringSubmatch(tag.Ref)

	if len(rst) != 3 {
		return "", "", "", errors.New("invalid tag")
	}

	return rst[1], rst[2], tag.Object.Sha[:7], nil
}

// 加载标签
func loadTag(v string) (*result, error) {
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
func fetchTag(v string) (*result, error) {
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

	res := &result{}

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
func fetchAllTags() ([]*result, error) {
	resp, err := http.Get(dueFrameworkTagsUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tags := make([]*result, 0)

	if err = json.Unmarshal(body, &tags); err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, errors.New("the due version not found")
	}

	return tags, nil
}
