package rules

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsArnHardcodeRule checks for hardcoded ARNs in Terraform code.
type AwsArnHardcodeRule struct {
	tflint.DefaultRule
}

// NewAwsArnHardcodeRule creates a new instance of AwsArnHardcodeRule.
func NewAwsArnHardcodeRule() *AwsArnHardcodeRule {
	return &AwsArnHardcodeRule{}
}

// Name returns the rule name.
func (r *AwsArnHardcodeRule) Name() string {
	return "aws_arn_hardcode"
}

// Enabled returns whether the rule is enabled by default.
func (r *AwsArnHardcodeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *AwsArnHardcodeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *AwsArnHardcodeRule) Link() string {
	return "https://example.com/aws_arn_hardcode_rule"
}

// Check scans for hardcoded ARNs in Terraform files.
func (r *AwsArnHardcodeRule) Check(runner tflint.Runner) error {
	// Define the regular expression for matching ARNs
	arnRegex := regexp.MustCompile(`arn:aws:`)

	// Retrieve all Terraform files
	files, err := runner.GetFiles()
	if err != nil {
		return fmt.Errorf("failed to retrieve Terraform files: %w", err)
	}

	// Iterate over each file
	for filename, file := range files {
		logger.Debug(fmt.Sprintf("Scanning file: %s", filename))

		// Read the file content line by line
		scanner := bufio.NewScanner(strings.NewReader(string(file.Bytes)))
		lineNumber := 1
		for scanner.Scan() {
			line := scanner.Text()
			if arnRegex.MatchString(line) {
				// Emit an issue with the specified format
				message := fmt.Sprintf("Error: Hardcoded ARN detected in file %s on line %d. Use variables or data sources instead. (%s)\n\n  on %s line %d:\n   %d: %s\n\nReference: %s\n",
					filename, lineNumber, r.Name(), filename, lineNumber, lineNumber, strings.TrimSpace(line), r.Link())

				runner.EmitIssue(
					r,
					message,
					hcl.Range{
						Filename: filename,
						Start:    hcl.Pos{Line: lineNumber, Column: 1},
						End:      hcl.Pos{Line: lineNumber, Column: len(line) + 1},
					},
				)
			}
			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading file %s: %w", filename, err)
		}
	}

	return nil
}

