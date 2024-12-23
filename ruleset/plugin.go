package ruleset

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

// Ruleset is the custom ruleset for this plugin.
type Ruleset struct {
    tflint.BuiltinRuleSet
}

// RuleSetName returns the name of the ruleset.
func (r *Ruleset) RuleSetName() string {
    return "aws_hardcode_check"
}

// RuleSetVersion returns the version of the ruleset.
func (r *Ruleset) RuleSetVersion() string {
    return "0.1.0"
}

// Rules returns the list of rules provided by this plugin.
func (r *Ruleset) Rules() []tflint.Rule {
    return []tflint.Rule{
        NewAwsArnHardcodeRule(),
    }
}

// NewRuleset initializes the plugin with metadata.
func NewRuleset() tflint.Ruleset {
    return &Ruleset{
        BuiltinRuleSet: tflint.BuiltinRuleSet{
            PluginInfo: tflint.PluginInfo{
                Name:    "aws_hardcode_check",
                Version: "0.1.0",
            },
        },
    }
}

