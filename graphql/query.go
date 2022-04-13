package graphql

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/request"
)

func (a *api) Query(operationName, query string, qOpts ...QueryOption) error {
	opts := MakeQueryOptions(qOpts...)

	query = strings.ReplaceAll(query, "\\n", "\n")
	log.Printf("operationName: %s", operationName)
	log.Printf("Query >>>\n%s\n<<< Query\n", query)

	uri := request.MakeURL("https://www.startupschool.org/graphql")
	cookie := [][2]string{
		{"_sso.key", a.creds.SSOKey},
		{"_sus_session", a.creds.SUSSession},
	}
	headers := map[string]string{
		"X-CSRF-Token": a.creds.XCSRFToken,
		"accept":       `*/*`,
		"content-type": `application/json`,
		"cookie":       request.CreateCookie(cookie),
	}

	type Body struct {
		OperationName string            `json:"operationName"`
		Variables     map[string]string `json:"variables"`
		Query         string            `json:"query"`
	}

	createdQuery := fmt.Sprintf("query %s {\n%s\n}", operationName, query)
	bodyObject := Body{
		OperationName: operationName,
		Variables:     opts.Variables(),
		Query:         createdQuery,
	}
	body, err := request.JSONMarshal(bodyObject)
	if err != nil {
		return errors.Errorf("JSONMarshal: %v", err)
	}

	var payload interface{}
	res, err := request.Post(uri, &payload, strings.NewReader(string(body)), request.RequestExtraHeaders(headers))
	if err != nil {
		return errors.Errorf("Post: %v", err)
	}

	// TODO
	log.Printf("res: %v", res)
	log.Printf("payload: %+v", payload)

	return nil
}
