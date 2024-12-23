package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/crazynuxer/tflint-hardcode-check/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "hardcode",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewAwsArnHardcodeRule(),
                                rules.NewAwsNetworkHardcodedIdsRule(),
			},
		},
	})
}
