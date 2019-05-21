package main

import (
	canaryboss "github.com/mirzakhany/canaryboss/internal/app/canaryboss"
	"github.com/mirzakhany/pkg/signals"
	"github.com/sirupsen/logrus"
)

const appName = "canaryboss"

func main() {
	defer canaryboss.InitCanaryBoss(appName)()
	sig := signals.WaitExitSignal()
	logrus.Infof("%s received, exiting...", sig.String())
}
