package env_differ_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEnvDiffer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EnvDiffer Suite")
}
