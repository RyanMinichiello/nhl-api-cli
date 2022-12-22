package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func people(args []string, stdout io.Writer) error {
	if len(args) < 3 {
		return errors.New("no id supplied to people service")
	}
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var id = flags.String("id", "", "people id")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	req := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/people/%s", *id)

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
