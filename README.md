# TFLint Hardcode Check Plugin

The **TFLint Hardcode Check Plugin** is a custom TFLint plugin designed to detect hardcoded identifiers in your Terraform configurations. This plugin helps enforce best practices by identifying hardcoded ARNs, VPC IDs, security group IDs, and subnet IDs, encouraging the use of variables or data sources instead.

## Features

- Detect hardcoded ARNs (e.g., `arn:aws:s3:::example-bucket`)
- Detect hardcoded VPC IDs (e.g., `vpc-123abc`)
- Detect hardcoded security group IDs (e.g., `sg-123abc`)
- Detect hardcoded subnet IDs (e.g., `subnet-123abc`)

## Installation

### 1. Add the Plugin to `.tflint.hcl`

To enable the plugin, add the following configuration to your `.tflint.hcl` file:

```hcl
plugin "hardcode" {
    enabled = true
    version = "0.1.0"
    source  = "github.com/crazynuxer/tflint-hardcode-check"
}
```

### 2. Initialize TFLint

Run the following command to initialize TFLint and download the plugin:

```bash
tflint --init
```

## Usage

After the plugin is installed, run TFLint as usual:

```bash
tflint
```

The plugin will scan your Terraform configurations and report any hardcoded identifiers.

## Example Output

For a Terraform file containing:

```hcl
resource "aws_s3_bucket" "example" {
  arn = "arn:aws:s3:::example-bucket" # Hardcoded ARN
}
```

The plugin will emit the following error:

```plaintext
Error: Hardcoded ARN detected in file example.tf on line 2. Use variables or data sources instead. (aws_arn_hardcode)

  on example.tf line 2:
   2: arn = "arn:aws:s3:::example-bucket"

Reference: https://example.com/aws_arn_hardcode_rule
```

## Supported Platforms

This plugin provides prebuilt binaries for the following platforms:

- macOS (arm64)
- Linux (amd64)
- Windows (amd64)

## Contributing

Contributions are welcome! Feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## References

- [TFLint Documentation](https://github.com/terraform-linters/tflint)
- [Custom Plugin Development](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/plugins.md)


