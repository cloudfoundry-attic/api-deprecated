package app_bits_controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApp_bits_controller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App_bits_controller Suite")
}
