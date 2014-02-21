package router_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApi_proxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api_proxy Suite")
}
