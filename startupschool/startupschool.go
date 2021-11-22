package startupschool

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tebeka/selenium"
)

type candidate struct {
	URI          string
	ProfileURI   string
	Name         string
	LinkedInUri  string
	Intro        string
	CompanyLinks []link
	CompanyText  string
}

type link struct {
	URI, Text string
}

type Credentials struct {
	Username string `json:"startupschoolUsername"`
	Password string `json:"startupschoolPassword"`
}

func ReadCredentials(f string) (*Credentials, error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	creds := &Credentials{}
	if err := json.Unmarshal(b, creds); err != nil {
		return nil, err
	}
	return creds, nil
}

type bot struct {
	creds Credentials
	wd    selenium.WebDriver
	d     *data
}

func MakeBot(creds Credentials) *bot {
	return &bot{creds: creds}
}
