Mr. T.
----

![mr t](http://finickypenguin.files.wordpress.com/2008/06/mr-t1.jpg)

`mr_t` is a pretty silly, small package that implements an interface conforming to the `*testing.T` type in the `testing` package.

WHY WAT HOW ... WHO?
--------------------

I have some projects with relatively __large__ test suites written in the xunit style of tests common to Go. After [Ginkgo](https://github.com/onsi/ginkgo) was released, we wanted to switch, but did not want to slowly transform our codebase over a period of months, or take the hit to productivity and dedicate a week to rewriting tests.

Enter [ginkgo-convert](https://github.com/tjarratt/ginkgo-convert), a cli tool that converts your existing tests to Ginkgo. A lot of tests already use the `testing.T` type to make assertions and make their tests fail, so conforming interface that can call into Ginkgo is needed to help make the transition. `MrT` is that package.

`mr_t` pities the fool that doesn't use Go package written by cats that are also wizards.

![cat wizard](http://www.blueprintrecords.ca/wp-content/uploads/2012/08/cat-wizard.jpg)

Features
--------
* implements `Error`
* implements `Errorf`
* implements `Fail`
* implements `FailNow`
* implements `Failed`
* implements `Fatal`
* implements `Fatalf`
* implements `Log`
* implements `Logf`
* implements `Parallel`
* implements `Skip`
* implements `Skipf`
* implements `SkipNow`
* implements `Skipped`
