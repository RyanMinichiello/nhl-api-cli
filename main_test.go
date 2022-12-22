package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestConferences(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "conferences"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/conferences"))
}

func TestConferencesById(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "conferences", "--id", "6"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/conferences/6"))
}

func TestFranchises(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "franchises"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/franchises"))
}

func TestFranchisesById(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "franchises", "--id", "6"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/franchises/6"))
	is.True(strings.Contains(out, "Boston"))

}

func TestDivisions(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "divisions"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/divisions"))
}

func TestDivisionsById(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "divisions", "--id", "6"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/divisions/6"))
}

func TestTeams(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "teams"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/teams"))
}

func TestTeamsById(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "teams", "--id", "28"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/teams/28"))
}

func TestPeopleErr(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "people"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.True(err != nil)
}

func TestPeopleById(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "people", "-id", "8477474"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "/api/v1/people/8477474"))
}

func TestSchedule(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "schedule"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "totalGames"))
}

func TestScheduleByDate(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "schedule", "-date", "2018-01-09"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, `"date" : "2018-01-09"`))
}

func TestScheduleByStartDateEndDate(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api", "schedule", "-start", "2021-01-17", "-end", "2021-01-19"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)
	out := stdout.String()
	is.True((strings.Contains(out, `"date" : "2021-01-17"`)))
	is.True((strings.Contains(out, `"date" : "2021-01-18"`)))
	is.True((strings.Contains(out, `"date" : "2021-01-19"`)))
}

// Commented as test is expensive
// func TestScheduleBySeason(t *testing.T) {
// 	is := is.New(t)

// 	args := []string{"nhl-api", "schedule", "-season", "20172018"}
// 	var stdout bytes.Buffer

// 	err := run(args, &stdout)
// 	is.NoErr(err)
// 	out := stdout.String()
// 	is.True((strings.Contains(out, `"date" : "2018-01-17"`)))
// }

func TestArgs(t *testing.T) {
	is := is.New(t)

	args := []string{"nhl-api"}
	var stdout bytes.Buffer
	err := run(args, &stdout)
	is.True(err != nil)
}
