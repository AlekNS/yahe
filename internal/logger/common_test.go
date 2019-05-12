package logger

import (
	"testing"

	"github.com/alekns/yahe/internal/config"
	"github.com/sirupsen/logrus"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInitCommon(t *testing.T) {
	Convey("Given initiated common logger with error level", t, func(c C) {
		initCommon(&config.LoggerSettings{
			ConsoleLevel: "error",
		})

		c.Convey("Then logging level should be error", func(c C) {
			c.So(GetRootLogger().Level, ShouldEqual, logrus.ErrorLevel)
		})
	})
}
