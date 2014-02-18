package hood_example_test

import (
	"github.com/eaigner/hood"
	_ "github.com/ziutek/mymysql/godrv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var hd *hood.Hood

func TestHood_example(t *testing.T) {
	RegisterFailHandler(Fail)
	var err error
	hd, err = hood.Open("mymysql", "test/pivotal/")
	hd.Log = true
	if err != nil {
		panic("Error when trying to connect to db: " + err.Error())
	}

	RunSpecs(t, "Hood_example Suite")
}
