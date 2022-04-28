package go_hello_world

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorld(t *testing.T) {

	Convey("Custom test for Bazel test", t, func() {
		hello := HelloWorld()

		Convey("The value should be equal", func() {
			So(hello, ShouldEqual, "HelloWorld!!")
		})

	})

}
