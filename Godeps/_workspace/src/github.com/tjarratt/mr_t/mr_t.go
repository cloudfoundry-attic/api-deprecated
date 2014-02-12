package mr_t

import (
	"fmt"
	"github.com/onsi/ginkgo"
)

type TestingT interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Parallel()
	Skip(args ...interface{})
	Skipf(format string, args ...interface{})
	SkipNow(args ...interface{})
	Skipped() bool
}

type mrT struct{}

func T() TestingT {
	return mrT{}
}

func (m mrT) Error(args ...interface{}) {
	m.Log(args)
	ginkgo.Fail("failed")
}

func (m mrT) Errorf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args...))
}

func (m mrT) Fail() {
	ginkgo.Fail("failed")
}

func (m mrT) FailNow() {
	ginkgo.Fail("failed")
}

func (m mrT) Failed() bool {
	return false
}

func (m mrT) Fatal(args ...interface{}) {
	m.Log(args)
	m.FailNow()
}

func (m mrT) Fatalf(format string, args ...interface{}) {
	m.Logf(format, args...)
	m.FailNow()
}

func (m mrT) Log(args ...interface{}) {
	for _, log := range args {
		println(log)
	}
}

func (m mrT) Logf(format string, args ...interface{}) {
	println(fmt.Sprintf(format, args...))
}

func (m mrT) Parallel() {
	return
}

func (m mrT) Skip(args ...interface{}) {
	m.Log(args...)
}

func (m mrT) Skipf(format string, args ...interface{}) {
	m.Logf(format, args...)
}

func (m mrT) SkipNow(args ...interface{}) {
	m.Log(args...)
	return
}

func (m mrT) Skipped() bool {
	return false
}
