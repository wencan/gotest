package gotest


import (
	"go.uber.org/zap"
)

func Output() {
	zap.NewExample().Sugar().Info("Hello, playground")
}
