package question

import "github.com/AlecAivazis/survey/v2"

var (
	Registry = &survey.Question{
		Name: "registry",
		Prompt: &survey.Select{
			Message: "choose a registry component:",
			Options: []string{"nacos", "etcd", "consul"},
			Default: "nacos",
		},
	}

	Config = &survey.Question{
		Name: "config",
		Prompt: &survey.Select{
			Message: "choose a config component:",
			Options: []string{"none", "file", "nacos", "etcd", "consul"},
			Default: "none",
		},
	}

	Cache = &survey.Question{
		Name: "cache",
		Prompt: &survey.Select{
			Message: "choose a cache component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		},
	}

	Lock = &survey.Question{
		Name: "lock",
		Prompt: &survey.Select{
			Message: "choose a lock component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		},
	}

	Transporter = &survey.Question{
		Name: "transporter",
		Prompt: &survey.Select{
			Message: "choose a transporter component:",
			Options: []string{"none", "grpc", "rpcx"},
			Default: "none",
		},
	}
)

func Name(def ...string) *survey.Question {
	prompt := &survey.Input{}
	prompt.Message = "specify the project name:"

	if len(def) > 0 {
		prompt.Default = def[0]
	}

	return &survey.Question{
		Name:     "name",
		Prompt:   prompt,
		Validate: survey.Required,
	}
}

func Network(def ...string) *survey.Question {
	prompt := &survey.Select{}
	prompt.Message = "choose a network component:"
	prompt.Options = []string{"ws", "tcp"}
	prompt.Default = "ws"

	if len(def) > 0 {
		prompt.Default = def[0]
	}

	return &survey.Question{
		Name:     "network",
		Prompt:   prompt,
		Validate: survey.Required,
	}
}

func Locator(def ...string) *survey.Question {
	prompt := &survey.Select{}
	prompt.Message = "choose a locator component:"
	prompt.Options = []string{"redis"}
	prompt.Default = "redis"

	if len(def) > 0 {
		prompt.Default = def[0]
	}

	return &survey.Question{
		Name:     "locator",
		Prompt:   prompt,
		Validate: survey.Required,
	}
}

// func Registry(def ...string) *survey.Question {
// 	prompt := &survey.Select{}
// 	prompt.Message = "choose a registry component:"
// 	prompt.Options = []string{"none", "nacos", "etcd", "consul"}
// 	prompt.Default = "nacos"

// 	if len(def) > 0 {
// 		prompt.Default = def[0]
// 	}

// 	return &survey.Question{
// 		Name:     "registry",
// 		Prompt:   prompt,
// 		Validate: survey.Required,
// 	}
// }

// Locator = &survey.Question{
// 		Name: "locator",
// 		Prompt: &survey.Select{
// 			Message: "choose a locator component:",
// 			Options: []string{"redis"},
// 			Default: "redis",
// 		},
// 	}

// func Transporter() *survey.Question {
// 	return &survey.Question{
// 		Name: "transporter",
// 		Prompt: &survey.Select{
// 			Message: "choose a transporter component:",
// 			Options: []string{"none", "grpc", "rpcx"},
// 			Default: "none",
// 		},
// 	}
// }
