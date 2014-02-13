package main_test

import (
	"fmt"
	"github.com/cloudfoundry-incubator/api/config"
	testnet "github.com/cloudfoundry-incubator/api/testhelpers/net"
	. "github.com/onsi/ginkgo"
	gconfig "github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/tjarratt/mr_t"
	"io/ioutil"
	"launchpad.net/goyaml"
	"net"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

var (
	defaultBackend *httptest.Server
	defaultHandler *testnet.TestHandler
	proxyPort      int
	proxyHost      string
	proxyURL       string
)

func buildApiServer() {
	err := exec.Command("scripts/build").Run()

	if err != nil {
		panic("could not build api: " + err.Error())
	}
}

func startApiServer() (cmd *exec.Cmd) {
	defaultBackend, defaultHandler = testnet.NewServer(mr_t.T(), []testnet.TestRequest{})
	c := config.Config{
		DefaultBackendURL: defaultBackend.URL,
		Port:              proxyPort,
	}

	filePath := writeTempConfigFile(c)
	cmd = exec.Command("out/api", "-c", filePath)
	go func() {
		out, err := cmd.CombinedOutput()
		println("API OUT", string(out))
		if err != nil {
			panic("could not run api server: " + err.Error())
		}
	}()

	time.Sleep(30 * time.Millisecond)
	return
}

func waitForServerStart() {
	uri := fmt.Sprintf("%s:%d", proxyHost, proxyPort)

	for {
		select {
		case <-time.NewTimer(3 * time.Second).C:
			panic("server failed to start before timeout")
		case <-time.NewTicker(30 * time.Millisecond).C:
			_, err := net.Dial("tcp", uri)
			if err == nil {
				return
			}
		}
	}
}

func stopApiServer(cmd *exec.Cmd) {
	cmd.Process.Kill()
	cmd.Wait()
}

func writeTempConfigFile(c config.Config) (filePath string) {
	dir, err := ioutil.TempDir("", "api_test")
	if err != nil {
		panic("could not create temp dir: " + err.Error())
	}

	filePath = filepath.Join(dir, "config.yml")

	data, err := goyaml.Marshal(c)
	if err != nil {
		panic("could not marshal config: " + err.Error())
	}

	err = ioutil.WriteFile(filePath, data, os.FileMode(0600))
	if err != nil {
		panic("could not write config file: " + err.Error())
	}

	return
}

func TestMain(t *testing.T) {
	proxyPort = 3000 + gconfig.GinkgoConfig.ParallelNode
	proxyHost := "localhost"
	proxyURL = fmt.Sprintf("http://%s:%d", proxyHost, proxyPort)

	cmd := startApiServer()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")

	stopApiServer(cmd)
}
