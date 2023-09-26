package runtimes_test

import (
	"testing"

	"github.com/nevercatchmeow/penumbra/core/tools/log"
	"github.com/nevercatchmeow/penumbra/core/tools/runtimes"
)

func TestCurrentFuncName(t *testing.T) {
	fnName := runtimes.CurrentFuncName()
	log.Info(fnName)
}

func TestGetWorkingDir(t *testing.T) {
	dir := runtimes.GetWorkingDir()
	log.Info(dir)
}
