package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func schedule(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		date      = flags.String("date", "", "games of day")
		season    = flags.String("season", "", "games of whole season")
		teamId    = flags.String("teamId", "", "pull only data for specified team(s)")
		startDate = flags.String("start", "", "starting from date games")
		endDate   = flags.String("end", "", "end at date games (inclusive)")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	var queryString string
	if *date != "" {
		queryString = fmt.Sprintf("?date=%s", *date)
	} else if *season != "" {
		queryString = fmt.Sprintf("?season=%s", *season)
	} else if *startDate != "" && *endDate != "" {
		queryString = fmt.Sprintf("?startDate=%s&endDate=%s", *startDate, *endDate)
	}

	if *teamId != "" {
		queryString += fmt.Sprintf("&teamId=%s", *teamId)
	}
	req := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/schedule/%s", queryString)

	resp, err := http.Get(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, string(body))

	return nil
}
