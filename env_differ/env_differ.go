package env_differ

import (
	"strings"

	"github.com/aryann/difflib"
	"github.com/cloudfoundry/cli/plugin"
)

type EnvDiffer struct {
	connection  plugin.CliConnection
	app1Name    string
	app2Name    string
	errors      []string
	diffStrings []string
}

func NewEnvDiffer(cliConnection plugin.CliConnection, args []string) *EnvDiffer {
	e := new(EnvDiffer)
	e.connection = cliConnection
	e.app1Name = args[1]
	e.app2Name = args[2]
	return e
}

func (e *EnvDiffer) Errors() []string {
	e.buildDiffs()
	return e.errors
}

func (e *EnvDiffer) Diffs() []string {
	e.buildDiffs()
	return e.diffStrings
}

func (e *EnvDiffer) buildDiffs() {
	app1Output, app1err := e.connection.CliCommandWithoutTerminalOutput("env", e.app1Name)
	app2Output, app2err := e.connection.CliCommandWithoutTerminalOutput("env", e.app2Name)
	app1OutputStr := strings.Join(app1Output, "")
	app2OutputStr := strings.Join(app2Output, "")

	errors := []string{}
	diffStrings := []string{}

	if app1err != nil || app2err != nil {
		if app1err != nil {
			errors = append(errors, app1OutputStr)
		}
		if app2err != nil {
			errors = append(errors, app2OutputStr)
		}
	} else {
		if app1OutputStr != app2OutputStr {
			diffs := difflib.Diff(app1Output, app2Output)
			for _, diff := range diffs {
				diffStrings = append(diffStrings, diff.String())
			}
		}
	}

	e.errors = errors
	e.diffStrings = diffStrings
}
