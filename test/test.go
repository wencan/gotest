package test

import (
	"go.uber.org/zap"
)

func Output() {
	zap.NewExample().Sugar().Info("Hello, playground")
}
