package startupschool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/spudtrooper/goutil/io"
)

const (
	dataDir = "data"
)

type data struct {
	dir string
}

func makeData() (*data, error) {
	dir, err := io.MkdirAll(dataDir)
	if err != nil {
		return nil, err
	}
	return &data{dir}, nil
}

func (d *data) Dir() string {
	return d.dir
}

func (d *data) SaveCandidate(cand candidate) error {
	b, err := json.Marshal(&cand)
	if err != nil {
		return err
	}
	f := path.Join(d.dir, "candidates", strings.ReplaceAll(cand.Name, " ", "-")+".json")
	log.Printf("saving to %s: %+v", f, cand)
	if ioutil.WriteFile(f, b, 0755); err != nil {
		return err
	}
	return nil
}
