package graphql

type api struct {
	creds creds
}

func MakeAPIFromFlags() (*api, error) {
	creds, err := readCredsFromFlags()
	if err != nil {
		return nil, err
	}
	res := &api{
		creds: *creds,
	}
	return res, nil
}
