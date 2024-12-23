package main

import (
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/crazynuxer/tflint-hardcode-check/ruleset"
)

func main() {
    tflint.Main(ruleset.New())
}

