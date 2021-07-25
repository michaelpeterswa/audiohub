package main

import (
	"go.uber.org/zap"
)

func runner(sugar zap.SugaredLogger) {

	sugar.Info("runner is initializing...")
	server(sugar)
}
