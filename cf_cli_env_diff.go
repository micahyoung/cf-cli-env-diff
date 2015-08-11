package main

import (
	"fmt"
	"strings"

	"github.com/aryann/difflib"
	"github.com/cloudfoundry/cli/plugin"
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
	app1Output, app1err := cliConnection.CliCommandWithoutTerminalOutput("env", args[1])
	app2Output, app2err := cliConnection.CliCommandWithoutTerminalOutput("env", args[2])
	app1OutputStr := strings.Join(app1Output, "\n")
	app2OutputStr := strings.Join(app2Output, "\n")

	if app1err != nil || app2err != nil {
		fmt.Println("There were errors:")
		if app1err != nil {
			fmt.Println(app1OutputStr)
		}
		if app2err != nil {
			fmt.Println(app2OutputStr)
		}
	} else {
		if app1OutputStr == app2OutputStr {
			fmt.Println("Environment variables are identical")
		} else {
			fmt.Println("Environment variable differences:")
			diffs := difflib.Diff(app1Output, app2Output)
			for _, diff := range diffs {
				fmt.Println(diff.String())
			}
		}
	}
}
