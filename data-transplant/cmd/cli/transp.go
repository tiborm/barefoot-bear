package main

import (
	"flag"
	"log"

	"github.com/tiborm/barefoot-bear/constants"
	"github.com/tiborm/barefoot-bear/internal/data/transplant"
	"github.com/tiborm/barefoot-bear/internal/utils/config"
)

// fetchSleepTime and forceFetch are configurable parameters for the transplant operation.
var (
	fetchSleepTime float64
	forceFetch     bool
)

// init initializes the fetchSleepTime and forceFetch variables from environment variables or command-line flags.
func init() {
	sleepTimeInt := config.GetEnvAsFloat64("FETCH_SLEEP_TIME", constants.FetchSleepTime)
	forceFetchBool := config.GetEnvAsBool("FORCE_FETCH", constants.ForceFetch)

	flag.Float64Var(&fetchSleepTime, "fetchSleepTime", float64(sleepTimeInt), "The sleep time between fetches. Environment variable: FETCH_SLEEP_TIME")
	flag.BoolVar(&forceFetch, "forceFetch", forceFetchBool, "Whether to force fetch or not. Environment variable: FORCE_FETCH")
}
// FIXME: read API information from ENV variables
func main() {
	flag.Parse()

	err := transplant.StartDataTransplant(forceFetch, fetchSleepTime)
	if err != nil {
		log.Fatalf("Error during transplant operation: %v", err)
	}

	log.Println("Transplant operation completed successfully.")
}
