package main

import (
	"flag"

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant"
	"github.com/tiborm/barefoot-bear/internal/utils/config"
)

var (
	fetchSleepTime float64
	forceFetch     bool
)

// FISME sleep time cli arg heavn't applied
func init() {
	sleepTimeInt := config.GetEnvAsFloat64("FETCH_SLEEP_TIME", constants.FetchSleepTime)
	forceFetchBool := config.GetEnvAsBool("FORCE_FETCH", constants.ForceFetch)

	flag.Float64Var(&fetchSleepTime, "fetchSleepTime", float64(sleepTimeInt), "The sleep time between fetches. Environment variable: FETCH_SLEEP_TIME")
	flag.BoolVar(&forceFetch, "forceFetch", forceFetchBool, "Whether to force fetch or not. Environment variable: FORCE_FETCH")
}

func main() {
	transplant.Transplant(fetchSleepTime, forceFetch)
}
