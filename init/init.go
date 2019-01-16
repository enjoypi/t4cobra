package init

import "github.com/sirupsen/logrus"

func init() {
	logrus.Trace("init.init")
	defer func() {
		logrus.Trace("init.init ok")
	}()

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.TraceLevel)
}
