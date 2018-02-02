package sanitizemarathonappid_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSanitizeMarathonAppId(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SanitizeMarathonAppId Suite")
}
