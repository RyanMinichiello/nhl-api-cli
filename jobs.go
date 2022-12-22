package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func jobs(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New("no job specified")
	}

	switch job := (args[1]); job {
	case "generate-teams-schedule":
		return fetchSeasonByTeam(args[1:], stdout)
	default:
		return errors.New("no job found")
	}
}

func fetchSeasonByTeam(args []string, stdout io.Writer) error {
	for i, team := range teamIdAndNames {
		if team == "" {
			continue
		}
		req := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/schedule?teamId=%s&season=20212022&gameType=R", fmt.Sprint(i))
		resp, err := http.Get(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		ioutil.WriteFile(fmt.Sprintf("%s-%s.json", fmt.Sprint(i), team), body, os.ModePerm)
	}
	return nil
}
