package main

import (
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "./ruleset"
)

func main() {
    tflint.Main(ruleset.New)
}

