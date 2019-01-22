package dlog

import (
	"log"
	"os"
	"testing"
)

func TestCheckSum(t *testing.T) {
	logger, err := NewLogger(os.Stdout, LevelDebug)
	if err != nil {
		log.Println(err)
	}

	logger.Debug("anode", "hello")
	SwitchLevel(&logger, LevelInfo)
	logger.Debug("anode", "hello")
	SwitchLevel(&logger, LevelDebug)
	logger.Debug("anode", "world")
}
