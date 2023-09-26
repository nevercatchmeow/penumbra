package network

import "github.com/nevercatchmeow/penumbra/core/tools/log"

type RunMode = log.RunMode

const (
	RunModeDev  RunMode = log.RunModeDev
	RunModeTest RunMode = log.RunModeTest
	RunModeProd RunMode = log.RunModeProd
)
