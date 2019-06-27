package test

import (
	"fmt"
  
  "github.com/uber-go/zap"
)

func Output() {
	zap.NewExample().Sugar().Info("Hello, playground")
}
