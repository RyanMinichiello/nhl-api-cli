package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New("no subroutine called")
	}
	switch subroutine := (args[1]); subroutine {
	case "conferences":
		return conferences(args[1:], stdout)
	case "divisions":
		return divisions(args[1:], stdout)
	case "franchises":
		return franchises(args[1:], stdout)
	case "teams":
		return teams(args[1:], stdout)
	case "people":
		return people(args[1:], stdout)
	case "schedule":
		return schedule(args[1:], stdout)
	case "jobs":
		return jobs(args[1:], stdout)
	default:
		return errors.New("unable to route request")
	}
}

var teamIdAndNames = []string{
	1:  "NJD",
	2:  "NYI",
	3:  "NYR",
	4:  "PHI",
	5:  "PIT",
	6:  "BOS",
	7:  "BUF",
	8:  "MTL",
	9:  "OTT",
	10: "TOR",
	12: "CAR",
	13: "FLA",
	14: "TBL",
	15: "WSH",
	16: "CHI",
	17: "DET",
	18: "NSH",
	19: "STL",
	20: "CGY",
	21: "COL",
	22: "EDM",
	23: "VAN",
	24: "ANA",
	25: "DAL",
	26: "LAK",
	28: "SJS",
	29: "CBJ",
	30: "MIN",
	52: "WPG",
	53: "ARI",
	54: "VGK",
	55: "SEA",
}
