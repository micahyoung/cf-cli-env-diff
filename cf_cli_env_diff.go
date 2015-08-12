package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/micahyoung/cf_cli_env_diff/env_differ"
)

type EnvDiff struct{}

func (c *EnvDiff) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "EnvDiff",
		Commands: []plugin.Command{
			{
				Name:     "env-diff",
				HelpText: "Shows diff in environment variables between two apps",
			},
		},
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
	}
}

func main() {
	plugin.Start(new(EnvDiff))
}

func (c *EnvDiff) Run(cliConnection plugin.CliConnection, args []string) {
	envDiffer := env_differ.New(cliConnection, args)

	if len(envDiffer.Errors()) > 0 {
		fmt.Println("There were errors:")
		fmt.Println(envDiffer.Errors())
		os.Exit(1)
	} else if len(envDiffer.Diffs()) > 0 {
		fmt.Println("Environment variable differences:")
		fmt.Println(envDiffer.Diffs())
		os.Exit(1)
	} else {
		fmt.Println("Environment variables are identical")
		os.Exit(0)
	}
}
