package sysoptimize

import (
	"runtime/debug"
	"time"
)

func ForceGC(t time.Duration) {
	for {
		tc := time.After(time.Millisecond * t)
		//logrus.Info("Optimized GC.....")
		debug.FreeOSMemory()
		<-tc
	}

}
