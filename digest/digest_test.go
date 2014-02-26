package digest_test

import (
	"github.com/cloudfoundry-incubator/api/digest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Digest", func() {
	Context("Hex", func() {
		var (
			expectedDigest = "0a4d55a8d778e5022fab701977c5d840bbc486d0"
			hexdigest      string
			err            error
		)
		BeforeEach(func() {
			r := strings.NewReader("Hello World")
			hexdigest, err = digest.Hex(r)
		})

		It("produces the same digest twice", func() {
			r := strings.NewReader("Hello World")
			hexdigestAgain, err := digest.Hex(r)
			Expect(err).NotTo(HaveOccurred())
			Expect(hexdigest).To(Equal(hexdigestAgain))
		})

		It("produces a different digest for different inputs", func() {
			r := strings.NewReader("Guttentag")
			differentDigest, err := digest.Hex(r)
			Expect(err).NotTo(HaveOccurred())
			Expect(hexdigest).ToNot(Equal(differentDigest))
		})

		It("matches expected sha1 hex digest from ruby", func() {
			Expect(hexdigest).To(Equal(expectedDigest))
		})

		It("doesn't error out", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
