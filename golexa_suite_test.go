package golexa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAlexaVoiceUI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golexa Suite")
}
