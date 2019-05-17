package main

import "github.com/seoane/terraform-inputs-generator/cmd"

// Version is updated by linker flags during build time
var Version = "dev"

func main() {
	cmd.Execute(Version)
}
