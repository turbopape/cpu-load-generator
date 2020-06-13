package main

// #include <time.h>
import "C"

import (
	"log"
	"time"
)

var startTime = time.Now()
var startTicks = C.clock()
var WaitFactor = 1

// if samplingRate = 0 then we continue with initial startTime and startTicks
func CpuUsagePercent(samplingRate float64, debug bool) float64 {

	clockSeconds := float64(C.clock()-startTicks) / float64(C.CLOCKS_PER_SEC)
	realSeconds := time.Since(startTime).Seconds()

	if debug {
		log.Printf(" current clock  : %v, real seconds: %v, startTicks: %v", clockSeconds, realSeconds, startTicks)
	}

	if samplingRate > 0 && realSeconds >= samplingRate {
		startTime = time.Now()
		startTicks = C.clock()
		if debug {
			log.Printf("Resetting starts !! startTime : %v, startTicks: %v", startTime, startTicks)
		}
	}
	return clockSeconds / realSeconds * 100

}
