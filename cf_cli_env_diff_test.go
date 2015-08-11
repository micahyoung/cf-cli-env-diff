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

		Describe("when there are no apps", func() {
			It("returns the error message from the CLI", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns([]string{"FAILED\n", "App does not exist"}, errors.New("THIS IS AN ERROR"))
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

		Describe("when there are apps with identical environment variables", func() {
			It("says the environment variables are identical", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputReturns([]string{"FOO=bar\n"}, nil)
				output := io_helpers.CaptureOutput(func() {
					envDiff.Run(fakeCliConnection, []string{"env-diff", "app1", "app2"})
				})

				Expect(output[0]).To(Equal("Environment variables are identical"))
			})
		})

		Describe("when there are apps with different environment variables", func() {
			It("says the environment variables are different", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub = func(args ...string) ([]string, error) {
					if args[1] == "app1" {
						return []string{"FOO=bar\n"}, nil
					} else if args[1] == "app2" {
						return []string{"FOO=qux\n"}, nil
					}
					return []string{"WHAT"}, nil
				}

				output := io_helpers.CaptureOutput(func() {
					envDiff.Run(fakeCliConnection, []string{"env-diff", "app1", "app2"})
				})

				Expect(output[0]).To(Equal("Environment variable differences:"))
				Expect(output[1]).To(Equal("- FOO=bar"))
				Expect(output[2]).To(Equal("+ FOO=qux"))
			})
		})
	})
})
