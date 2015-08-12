package env_differ_test

import (
	"errors"

	"github.com/cloudfoundry/cli/plugin/fakes"
	"github.com/micahyoung/cf_cli_env_diff/env_differ"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EnvDiffer", func() {
	Describe(".Errors", func() {
		var fakeCliConnection *fakes.FakeCliConnection
		var envDiffer *env_differ.EnvDiffer

		BeforeEach(func() {
			fakeCliConnection = &fakes.FakeCliConnection{}
			envDiffer = env_differ.New(fakeCliConnection, []string{"env-diff", "app1", "app2"})
		})

		Describe("when there are no apps", func() {
			It("returns the error message from the CLI", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub = func(args ...string) ([]string, error) {
					switch args[1] {
					case "app1":
						return []string{"FAILED\n", "App app1 does not exist"}, errors.New("ERROR")
					case "app2":
						return []string{"FAILED\n", "App app2 does not exist"}, errors.New("ERROR")
					default:
						return []string{"WHAT"}, nil
					}
				}
				errors := envDiffer.Errors()

				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[0]).To(Equal("env"))
				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)[1]).To(Equal("app1"))
				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[0]).To(Equal("env"))
				Expect(fakeCliConnection.CliCommandWithoutTerminalOutputArgsForCall(1)[1]).To(Equal("app2"))

				Expect(errors[0]).To(Equal("FAILED\nApp app1 does not exist"))
				Expect(errors[1]).To(Equal("FAILED\nApp app2 does not exist"))
			})
		})

		Describe("when there are apps with identical environment variables", func() {
			It("says the environment variables are identical", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub = func(args ...string) ([]string, error) {
					switch args[1] {
					case "app1":
						return []string{"FOO=bar\n"}, nil
					case "app2":
						return []string{"FOO=bar\n"}, nil
					default:
						return []string{"WHAT"}, nil
					}
				}

				Expect(envDiffer.Errors()).To(BeEmpty())
				Expect(envDiffer.Diffs()).To(BeEmpty())
			})
		})

		Describe("when there are apps with different environment variables", func() {
			It("says the environment variables are different", func() {
				fakeCliConnection.CliCommandWithoutTerminalOutputStub = func(args ...string) ([]string, error) {
					switch args[1] {
					case "app1":
						return []string{"FOO=bar\n"}, nil
					case "app2":
						return []string{"FOO=qux\n"}, nil
					default:
						return []string{"WHAT"}, nil
					}
				}

				output := envDiffer.Diffs()

				Expect(envDiffer.Errors()).To(BeEmpty())
				Expect(output[0]).To(Equal("- FOO=bar\n"))
				Expect(output[1]).To(Equal("+ FOO=qux\n"))
			})
		})
	})
})
