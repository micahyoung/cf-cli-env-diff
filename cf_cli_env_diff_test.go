package main_test

import (
	"errors"

	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/micahyoung/cf_cli_env_diff"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnvDiff", func() {
	Describe(".Run", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var envDiff *EnvDiff

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			envDiff = &EnvDiff{}
		})

		Describe("when there aren no apps", func() {
			It("returns the error message from the CLI", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns([]string{"FAILED", "App does not exist"}, errors.New("THIS IS A STRING"))
				output := io_helpers.CaptureOutput(func() {
					envDiff.Run(fakeCliConnection, []string{"env-diff", "app1", "app2"})
				})

				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[0]).To(Equal("env"))
				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[1]).To(Equal("app1"))
				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[0]).To(Equal("env"))
				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[1]).To(Equal("app2"))

				Expect(output[0]).To(Equal("There were errors:"))
				Expect(output[1]).To(Equal("FAILED"))
				Expect(output[2]).To(Equal("App does not exist"))
				Expect(output[3]).To(Equal("FAILED"))
				Expect(output[4]).To(Equal("App does not exist"))
			})
		})
	})
})
