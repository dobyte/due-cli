package question

import "github.com/AlecAivazis/survey/v2"

var (
	Network = &survey.Question{
		Name: "network",
		Prompt: &survey.Select{
			Message: "Choose a network:",
			Options: []string{"tcp", "ws"},
			Default: "ws",
		},
	}

	Locator = &survey.Question{
		Name: "locator",
		Prompt: &survey.Select{
			Message: "Choose a locator:",
			Options: []string{"redis"},
			Default: "redis",
		},
	}

	Registry = &survey.Question{
		Name: "registry",
		Prompt: &survey.Select{
			Message: "Choose a registry component:",
			Options: []string{"nacos", "etcd", "consul"},
			Default: "nacos",
		},
	}

	Config = &survey.Question{
		Name: "config",
		Prompt: &survey.Select{
			Message: "Choose a config component:",
			Options: []string{"none", "file", "nacos", "etcd", "consul"},
			Default: "none",
		},
	}

	Cache = &survey.Question{
		Name: "cache",
		Prompt: &survey.Select{
			Message: "Choose a cache component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		},
	}

	Lock = &survey.Question{
		Name: "lock",
		Prompt: &survey.Select{
			Message: "Choose a lock component:",
			Options: []string{"none", "redis", "memcache"},
			Default: "none",
		},
	}

	Transporter = &survey.Question{
		Name: "transporter",
		Prompt: &survey.Select{
			Message: "Choose a transporter component:",
			Options: []string{"none", "grpc", "rpcx"},
			Default: "none",
		},
	}
)
