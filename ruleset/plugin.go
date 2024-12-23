package ruleset

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type Plugin struct{}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) PluginInfo() *tflint.PluginInfo {
	return &tflint.PluginInfo{
		Name:    "hardcode-check",
		Version: "0.1.0",
	}
}

func (p *Plugin) Rules() []tflint.Rule {
	return []tflint.Rule{
		NewAwsArnHardcodeRule(), // Reference to your custom rule
	}
}

