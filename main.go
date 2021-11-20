package main

import (
	"flag"

	"github.com/spudtrooper/startupschool/startupschool"
)

var (
	credentialsFile = flag.String("credentials_file", ".credentials.json", "File with credentials")
	seleniumVerbose = flag.Bool("selenium_verbose", false, "verbose selenium logging")
	seleniumHead    = flag.Bool("selenium_head", false, "Take screenshots withOUT headless chrome")
	limit           = flag.Int("limit", 0, "max number of times to check next candidate")
)

func realMain() error {
	creds, err := startupschool.ReadCredentials(*credentialsFile)
	if err != nil {
		return err
	}
	bot := startupschool.MakeBot(*creds)
	if err := bot.Loop(
		startupschool.LoopSeleniumVerbose(*seleniumVerbose),
		startupschool.LoopSeleniumHead(*seleniumHead),
		startupschool.LoopLimit(*limit)); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	if err := realMain(); err != nil {
		panic(err)
	}
}
