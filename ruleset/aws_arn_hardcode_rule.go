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

func (r *AwsArnHardcodeRule) Severity() tflint.Severity {
    return tflint.ERROR
}

func (r *AwsArnHardcodeRule) Link() string {
    return "https://example.com/docs/rules/aws_arn_hardcode_rule"
}

func (r *AwsArnHardcodeRule) Check(runner tflint.Runner) error {
    return runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
        var value string
        err := runner.EvaluateExpr(expr, &value, nil)
        if err != nil {
            return nil // Skip if not evaluatable
        }
        if value == "arn:aws:" {
            runner.EmitIssue(
                r,
                "Hardcoded ARN found, use variables instead.",
                expr.Range(),
            )
        }
        return nil
    }))
}

