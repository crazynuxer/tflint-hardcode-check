package ruleset

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

func New() (tflint.Ruleset, error) {
    ruleset := tflint.NewRuleset()
    ruleset.AddRule(NewAwsArnHardcodeRule())
    return ruleset, nil
}

