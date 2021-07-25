package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	prog := "AuDiOhUb"
	ver := "v0.0.1"

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	sugar.Infof("%s (%s) is initializing...", prog, ver)

	runner(*sugar)
}
