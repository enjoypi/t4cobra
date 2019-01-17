package init

import "github.com/sirupsen/logrus"

func init() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}
