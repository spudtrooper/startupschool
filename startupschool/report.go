package startupschool

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sort"

	"github.com/spudtrooper/goutil/html"
)

func Report(dataDir string) error {
	cands, err := findExistingCandidates(dataDir)
	if err != nil {
		return err
	}

	sort.Slice(cands, func(a, b int) bool {
		ca, cb := cands[a], cands[b]
		return ca.Name > cb.Name
	})
	head := html.TableRowData{
		"IMAGE",
		"NAME",
		"LOCATION",
		"COMPANY",
		"INTRO",
		"LINKEDIN",
	}
	var rows []html.TableRowData
	for _, c := range cands {
		image := fmt.Sprintf(`<a href="%s" target="_"><img style="max-width:100px" src="%s"/></a>`, c.ProfileURI, c.ProfileURI)
		name := fmt.Sprintf(`<a href="%s" target="_">%s</a>`, c.URI, c.Name)
		location := c.Location
		companyHTML := html.Linkify(c.CompanyText)
		intro := c.Intro
		linkedIn := fmt.Sprintf(`<a href="%s" target="_">LinkedIn</a>`, c.LinkedInUri)
		row := html.TableRowData{
			image,
			name,
			location,
			companyHTML,
			intro,
			linkedIn,
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
	outFile := path.Join(dataDir, "html", "index.html")
	if err := ioutil.WriteFile(outFile, []byte(html), 0755); err != nil {
		return err
	}

	log.Printf("wrote to %s", outFile)

	return nil
}
