package ruleset

import (
    "github.com/hashicorp/hcl/v2"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsArnHardcodeRule struct {
    tflint.DefaultRule
}

func NewAwsArnHardcodeRule() *AwsArnHardcodeRule {
    return &AwsArnHardcodeRule{}
}

func (r *AwsArnHardcodeRule) Name() string {
    return "aws_arn_hardcode_rule"
}

func (r *AwsArnHardcodeRule) Enabled() bool {
    return true
}

func (r *AwsArnHardcodeRule) Severity() string {
    return tflint.ERROR
}

func (r *AwsArnHardcodeRule) Link() string {
    return "https://example.com/aws_arn_hardcode_rule"
}

func (r *AwsArnHardcodeRule) Check(runner tflint.Runner) error {
    return runner.WalkExpressions(func(expr hcl.Expression) error {
        traversal := expr.Variables()
        for _, t := range traversal {
            if t.RootName() == "arn:aws:" {
                runner.EmitIssue(
                    r,
                    "Hardcoded ARN detected. Use variables or data sources instead.",
                    t.SourceRange(),
                )
            }
        }
        return nil
    })
}

