package hub

import (
	"context"
	"fmt"
	"log"
	"time"
)

// version params
var (
	LongHash   string
	CommitDate string
	BuildDate  string
)

// InitCanaryBoss init canaryboss service
func InitCanaryBoss(appName string) func() {
	_, cnl := context.WithCancel(context.Background())
	err := initModules(appName)
	if err != nil {
		log.Fatalf("init modules failed with error: %v", err)
	}
	return func() {
		cnl()
		<-time.After(1 * time.Second)
	}
}

func printVersion(appName string) {
	fmt.Printf("===\nApp:%s\nCommitDate:%s\nBuildDate:%s\nLongHash:%s\n===\n",
		appName, CommitDate, BuildDate, LongHash)
}

func initModules(appName string) error {

	// print app name and version params to stdout
	printVersion(appName)

	return nil
}
