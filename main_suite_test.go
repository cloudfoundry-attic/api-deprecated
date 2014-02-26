package main_test

import (
	"fmt"
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/testhelpers/file"
	testnet "github.com/cloudfoundry-incubator/api/testhelpers/net"
	"github.com/fraenkel/candiedyaml"
	. "github.com/onsi/ginkgo"
	gconfig "github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/tjarratt/mr_t"
	"github.com/vito/cmdtest"
	"io"
	"net"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

var (
	conf           config.Config
	defaultBackend *httptest.Server
	defaultHandler *testnet.TestHandler
	proxyPort      int
	proxyHost      string
	proxyURL       string
)

func TestMain(t *testing.T) {
	proxyPort = 3000 + gconfig.GinkgoConfig.ParallelNode
	proxyHost := "localhost"
	proxyURL = fmt.Sprintf("http://%s:%d", proxyHost, proxyPort)

	defaultBackend, defaultHandler = testnet.NewServer(mr_t.T(), []testnet.TestRequest{})

	conf = config.Config{
		DefaultBackendURL: defaultBackend.URL,
		Port:              proxyPort,
		DB: config.DbConfig{
			URI: "sqlite://:memory:",
		},
		AppPackages: config.BlobstoreConfig{
			Filepath: file.TmpDir(),
		},
	}

	cmd := startApiServer(conf)
	waitForServerStart()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")

	stopApiServer(cmd)
}

func buildApiServer() {
	err := exec.Command("scripts/build").Run()

	if err != nil {
		panic("could not build api: " + err.Error())
	}
}

func Tee(out io.Writer) io.Writer {
	return io.MultiWriter(out, os.Stdout)
}

func startApiServer(c config.Config) *cmdtest.Session {
	filePath := writeTempConfigFile(c)

	session, err := cmdtest.StartWrapped(
		exec.Command("out/api", "-c", filePath),
		Tee,
		Tee,
	)
	if err != nil {
		panic(err)
	}
	return session
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

func stopApiServer(session *cmdtest.Session) {
	session.Cmd.Process.Kill()
	session.Wait(1 * time.Second)

}

func writeTempConfigFile(c config.Config) (filePath string) {
	filePath = filepath.Join(file.TmpDir(), "config.yml")

	confFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic("could not create config file: " + err.Error())
	}
	defer confFile.Close()

	err = candiedyaml.NewEncoder(confFile).Encode(c)
	if err != nil {
		panic("could not marshal config: " + err.Error())
	}

	return
}
