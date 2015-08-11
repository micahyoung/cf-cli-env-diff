package main_test

import (
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("main", func() {
	Describe("installation", func() {
		It("installs and uninstalls successfully", func() {
			installResult := Cf("install-plugin", "cf_cli_env_diff.exe")
			Eventually(installResult, 3*time.Second).Should(Say("successfully installed."))

			uninstallResult := Cf("uninstall-plugin", "EnvDiff")
			Eventually(uninstallResult, 3*time.Second).Should(Say("successfully uninstalled."))
		})
	})
})

func Cf(args ...string) *Session {
	path, err := Build("github.com/cloudfoundry/cli/main")
	Expect(err).NotTo(HaveOccurred())

	session, err := Start(exec.Command(path, args...), GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	return session
}

// gexec.Build leaves a compiled binary behind in /tmp.
var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})
