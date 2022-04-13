package graphql

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var (
	graphqlCredentialsFile = flag.String("graphql_credentials_file", ".graphql_credentials.json", "file for graphql credentials")
)

type creds struct {
	SSOKey     string `json:"ssoKey"`
	SUSSession string `json:"susSession"`
	XCSRFToken string `json:"xCSRFToken"`
}

func readCredsFromFlags() (*creds, error) {
	b, err := ioutil.ReadFile(*graphqlCredentialsFile)
	if err != nil {
		return nil, err
	}
	creds := &creds{}
	if err := json.Unmarshal(b, creds); err != nil {
		return nil, err
	}
	return creds, nil
}
