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

// AwsNetworkHardcodedIdsRule checks for hardcoded VPC, security group, and subnet IDs in Terraform code.
type AwsNetworkHardcodedIdsRule struct {
	tflint.DefaultRule
}

// NewAwsNetworkHardcodedIdsRule creates a new instance of AwsNetworkHardcodedIdsRule.
func NewAwsNetworkHardcodedIdsRule() *AwsNetworkHardcodedIdsRule {
	return &AwsNetworkHardcodedIdsRule{}
}

// Name returns the rule name.
func (r *AwsNetworkHardcodedIdsRule) Name() string {
	return "aws_network_hardcoded_ids"
}

// Enabled returns whether the rule is enabled by default.
func (r *AwsNetworkHardcodedIdsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *AwsNetworkHardcodedIdsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *AwsNetworkHardcodedIdsRule) Link() string {
	return "https://example.com/aws_network_hardcoded_ids_rule"
}

// Check scans for hardcoded IDs in Terraform files.
func (r *AwsNetworkHardcodedIdsRule) Check(runner tflint.Runner) error {
	// Define the regular expressions for matching IDs
	vpcRegex := regexp.MustCompile(`vpc-[a-zA-Z0-9]+`)
	sgRegex := regexp.MustCompile(`sg-[a-zA-Z0-9]+`)
	subnetRegex := regexp.MustCompile(`subnet-[a-zA-Z0-9]+`)

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
			// Check for hardcoded VPC IDs
			if vpcRegex.MatchString(line) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Hardcoded VPC ID detected in file %s on line %d. Use variables or data sources instead.", filename, lineNumber),
					hcl.Range{
						Filename: filename,
						Start:    hcl.Pos{Line: lineNumber, Column: 1},
						End:      hcl.Pos{Line: lineNumber, Column: len(line) + 1},
					},
				)
			}
			// Check for hardcoded security group IDs
			if sgRegex.MatchString(line) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Hardcoded security group ID detected in file %s on line %d. Use variables or data sources instead.", filename, lineNumber),
					hcl.Range{
						Filename: filename,
						Start:    hcl.Pos{Line: lineNumber, Column: 1},
						End:      hcl.Pos{Line: lineNumber, Column: len(line) + 1},
					},
				)
			}
			// Check for hardcoded subnet IDs
			if subnetRegex.MatchString(line) {
				runner.EmitIssue(
					r,
					fmt.Sprintf("Hardcoded subnet ID detected in file %s on line %d. Use variables or data sources instead.", filename, lineNumber),
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

