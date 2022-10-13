package main

import "go.uber.org/zap"

var (
	logger *zap.SugaredLogger
)

func init() {
	l, _ := zap.NewDevelopment()
	logger = l.Sugar()
}
