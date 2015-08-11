package main

import (
	"fmt"
	"strings"

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
	}
}

func main() {
	plugin.Start(new(EnvDiff))
}

func (c *EnvDiff) Run(cliConnection plugin.CliConnection, args []string) {
	app1output, app1err := cliConnection.CliCommandWithoutTerminalOutput("env", args[1])
	app2output, app2err := cliConnection.CliCommandWithoutTerminalOutput("env", args[2])

	if app1err != nil || app2err != nil {
		fmt.Println("There were errors:")
		if app1err != nil {
			fmt.Println(strings.Join(app1output, "\n"))
		}
		if app2err != nil {
			fmt.Println(strings.Join(app2output, "\n"))
		}
	}
}
