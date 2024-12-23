package ruleset

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type AwsArnHardcodeRule struct{}

// NewAwsArnHardcodeRule creates a new instance of the rule
func NewAwsArnHardcodeRule() *AwsArnHardcodeRule {
	return &AwsArnHardcodeRule{}
}

// Name returns the rule name
func (r *AwsArnHardcodeRule) Name() string {
	return "aws_arn_hardcode_check"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsArnHardcodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsArnHardcodeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsArnHardcodeRule) Link() string {
	return "https://example.com/docs/aws_arn_hardcode_check"
}

// Check validates the Terraform configuration
func (r *AwsArnHardcodeRule) Check(runner tflint.Runner) error {
	return runner.WalkExpressions(nil, func(expr hcl.Expression) error {
		value, diags := runner.Evaluate(expr)
		if diags.HasErrors() {
			return nil
		}

		if value.Type().IsPrimitiveType() && value.AsString() == "arn:aws:" {
			runner.EmitIssue(
				r,
				"Hardcoded ARN detected. Consider using variables or data sources instead.",
				expr.Range(),
			)
		}
		return nil
	})
}

