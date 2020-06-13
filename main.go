package main

// #include <time.h>
import "C"

import (
	. "log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	logger := New(os.Stderr, "eip-great-ptretender ", Llongfile|Ldate|Ltime)

	app.Flags = []cli.Flag{

		cli.Float64Flag{
			Name:  "targetCPU, t",
			Value: 30.0,
			Usage: "The load to be generated.",
		},
		cli.Float64Flag{
			Name:  "samplingRate, s",
			Value: 0,
			Usage: "A rate, in second, to sample CPU usage to determine percentage(used ticks / total ticks since start). If 0, starts are not reset.",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "If Set, outputs metrics about system load.",
		},
	}

	app.Action = func(c *cli.Context) error {
		var wg sync.WaitGroup
		defer func() {
			if errP := recover(); errP != nil {
				logger.Printf("Couldn't recover from panic - reason: %e", errP)
			}
		}()
		wg.Add(1)
		go func(wg *sync.WaitGroup) {

			targetCPU := c.Float64("targetCPU")
			samplingRate := c.Float64("samplingRate")
			debug := c.Bool("debug")
			//done := make(chan int)
			//for CPU
			for i := 0; i < runtime.NumCPU(); i++ {
				go func() {

					for {
						cpuLoad := CpuUsagePercent(samplingRate, debug)

						if debug {
							logger.Printf(" Load : %f, Target : %f", cpuLoad, targetCPU)
						}

						if cpuLoad >= targetCPU {
							waitTime := time.Second / time.Duration(WaitFactor)
							time.Sleep(waitTime)

						}

					}
					defer wg.Done()
				}()
			}

		}(&wg)
		wg.Wait()
		return nil
	}

	errRun := app.Run(os.Args)

	if errRun != nil {
		logger.Panic(errRun)
	}
}
