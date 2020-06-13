package main

/*
#cgo LDFLAGS: -lpsapi
#include <windows.h>
#include <psapi.h>
#include <time.h>

double get_cpu_time(){
    FILETIME a,b,c,d;
    if (GetProcessTimes(GetCurrentProcess(),&a,&b,&c,&d) != 0){
        //  Returns total user time.
        //  Can be tweaked to include kernel times as well (c).
        return

		(
			(double)(d.dwLowDateTime |
			((unsigned long long)d.dwHighDateTime << 32)) // user time

		) * 0.0000001;
    }else{
        //  Handle error
        return 0;
    }
}
*/
import "C"

import (
	"log"
	"time"
)

var startTime = time.Now()
var startTicksTime = C.get_cpu_time()
var WaitFactor = time.Duration(C.CLOCKS_PER_SEC)

// if samplingRate = 0 then we continue with initial startTime and startTicks
func CpuUsagePercent(samplingRate float64, debug bool) float64 {

	clockSeconds := float64(C.get_cpu_time() - startTicksTime)
	realSeconds := time.Since(startTime).Seconds()

	if debug {
		log.Printf(" current clock  : %v, real seconds: %v, startTicks: %v", clockSeconds, realSeconds, startTicksTime)
	}

	if samplingRate > 0 && realSeconds >= samplingRate {
		startTime = time.Now()
		startTicksTime = C.get_cpu_time()
		if debug {
			log.Printf("Resetting starts !! startTime : %v, startTicks: %v", startTime, startTicksTime)
		}
	}
	return clockSeconds / realSeconds * 100

}
