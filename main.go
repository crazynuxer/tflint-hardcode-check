package main

import (
	"github.com/crazynuxer/tflint-hardcode-check/ruleset"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(ruleset.New())
}

