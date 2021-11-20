package startupschool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/spudtrooper/goutil/html"
)

func Report() error {
	d, err := makeData()
	if err != nil {
		return err
	}
	dir := d.Dir()

	var files []string
	if err := filepath.Walk(dir, func(f string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(f) == ".json" {
			files = append(files, f)
		}
		return nil
	}); err != nil {
		return err
	}

	var cands []candidate
	for _, f := range files {
		var c candidate
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &c); err != nil {
			return err
		}
		cands = append(cands, c)
	}

	sort.Slice(cands, func(a, b int) bool {
		ca, cb := cands[a], cands[b]
		return ca.Name > cb.Name
	})
	head := html.TableRowData{
		"NAME",
		"INTRO",
		"LINKEDIN",
	}
	var rows []html.TableRowData
	for _, c := range cands {
		row := html.TableRowData{
			fmt.Sprintf(`<a href="%s" target="_">%s</a>`, c.URI, c.Name),
			c.Intro,
			fmt.Sprintf(`<a href="%s" target="_">LinkedIn</a>`, c.LinkedInUri),
		}
		rows = append(rows, row)
	}

	htmlData := html.Data{
		Entities: []html.DataEntity{
			html.MakeDataEntityFromTable(html.TableData{
				Head: head,
				Rows: rows,
			}),
		}}
	html, err := html.Render(htmlData)
	if err != nil {
		return err
	}
	outFile := path.Join(dir, "html", "index.html")
	if err := ioutil.WriteFile(outFile, []byte(html), 0755); err != nil {
		return err
	}

	log.Printf("wrote to %s", outFile)

	return nil
}
