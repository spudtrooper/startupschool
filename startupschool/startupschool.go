package startupschool

import (
	"encoding/json"
	"io/ioutil"
)

type candidate struct {
	URI         string
	Name        string
	LinkedInUri string
	Intro       string
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
}

func MakeBot(creds Credentials) *bot {
	return &bot{creds}
}
