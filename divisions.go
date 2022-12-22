package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func divisions(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var id = flags.String("id", "", "division id")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	req := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/divisions/%s", *id)

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
