package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spudtrooper/startupschool/startupschool"
)

var (
	credentialsFile    = flag.String("credentials_file", ".credentials.json", "File with credentials")
	seleniumVerbose    = flag.Bool("selenium_verbose", false, "verbose selenium logging")
	seleniumHead       = flag.Bool("selenium_head", false, "Take screenshots withOUT headless chrome")
	loop               = flag.Int("loop", 0, "max number of times to check next candidate")
	uris               = flag.String("uris", "", "comma-delimited list of URIs to search")
	data               = flag.String("data", "data", "directory to store data")
	backfill           = flag.Bool("backfill", false, "Backfill existing entries")
	pause              = flag.Duration("pause", 0, "pause time between requests")
	findLinkedIns      = flag.Bool("find_linkedins", false, "if true we will search for linkedin profiles")
	findLinkedInsStart = flag.Int("find_linkedins_start", 0, "the start index for finding linkedins")
)

func realMain() error {
	creds, err := startupschool.ReadCredentials(*credentialsFile)
	if err != nil {
		return err
	}
	bot := startupschool.MakeBot(*creds)
	cancel, err := bot.Login(*data,
		startupschool.LoginSeleniumVerbose(*seleniumVerbose),
		startupschool.LoginSeleniumHead(*seleniumHead))
	if err != nil {
		return err
	}
	defer cancel()

	if *findLinkedIns {
		uris, err := bot.FindLinkedInProfiles(
			startupschool.FindLinkedInProfilesPause(*pause),
			startupschool.FindLinkedInProfilesStart(*findLinkedInsStart),
		)
		if err != nil {
			return err
		}
		for _, uri := range uris {
			fmt.Println(uri)
		}
	}

	if *uris != "" {
		uris := strings.Split(*uris, ",")
		if err := bot.SearchURIs(uris, startupschool.SearchURIsPause(*pause)); err != nil {
			return err
		}
	}

	if *loop > 0 {
		if err := bot.Loop(
			startupschool.LoopLimit(*loop),
			startupschool.LoopPause(*pause)); err != nil {
			return err
		}
	}

	if *backfill {
		if err := bot.Backfill(startupschool.BackfillPause(*pause)); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if err := realMain(); err != nil {
		panic(err)
	}
}
