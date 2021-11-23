package startupschool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func (b *bot) lookUp(uri string) (*candidate, error) {
	log.Printf("looking up %s", uri)

	if err := b.wd.Get(uri); err != nil {
		return nil, err
	}

	var name string
	b.wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		log.Printf("waiting for name...")
		h1, err := wd.FindElement(selenium.ByTagName, "h1")
		if err != nil {
			return false, nil
		}
		text, err := h1.Text()
		if err != nil || text == "" {
			return false, nil
		}
		name = text
		return true, nil
	})

	ps, err := b.wd.FindElements(selenium.ByTagName, "p")
	if err != nil {
		return nil, err
	}
	locationP := ps[0]
	locationText, err := locationP.Text()
	if err != nil {
		return nil, err
	}
	location := strings.Split(locationText, "|")[0]

	var linkedInUri string
	as, err := b.wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return nil, err
	}
	for _, a := range as {
		text, err := a.Text()
		if err != nil {
			return nil, err
		}
		if text == "LinkedIn" {
			href, err := a.GetAttribute("href")
			if err != nil {
				return nil, err
			}
			linkedInUri = href
		}
	}

	tables, err := b.wd.FindElements(selenium.ByTagName, "table")
	if err != nil {
		return nil, err
	}
	table := tables[0]
	trs, err := table.FindElements(selenium.ByTagName, "tr")
	if err != nil {
		return nil, err
	}
	introTR := trs[0]
	introText, err := introTR.Text()
	if err != nil {
		return nil, err
	}
	intro := strings.Replace(introText, "Intro", "", 1)

	var companyLinks []link
	var companyText string
	for _, tr := range trs {
		tds, err := tr.FindElements(selenium.ByTagName, "td")
		if err != nil {
			return nil, err
		}
		if len(tds) != 1 {
			continue
		}
		td := tds[0]
		tdText, err := td.Text()
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(tdText, "Company") {
			companyText = tdText
			companyText = strings.Replace(companyText, "Company\n", "", 1)
			as, err := tr.FindElements(selenium.ByTagName, "a")
			if err != nil {
				return nil, err
			}
			for _, a := range as {
				href, err := a.GetAttribute("href")
				if err != nil {
					return nil, err
				}
				text, err := a.Text()
				if err != nil {
					return nil, err
				}
				companyLinks = append(companyLinks, link{
					URI:  href,
					Text: text,
				})
			}
		}
	}

	url, err := b.wd.CurrentURL()
	if err != nil {
		return nil, err
	}

	var profileURI string
	images, err := b.wd.FindElements(selenium.ByTagName, "img")
	if err != nil {
		return nil, err
	}
	for _, img := range images {
		alt, err := img.GetAttribute("alt")
		if err == nil && alt == "candidate avatar" {
			src, err := img.GetAttribute("src")
			if err != nil {
				return nil, err
			}
			profileURI = src
		}
	}

	updatedMillis := time.Now().UnixMilli()
	c := &candidate{
		Name:          name,
		Location:      location,
		LinkedInUri:   linkedInUri,
		Intro:         intro,
		URI:           url,
		ProfileURI:    profileURI,
		CompanyLinks:  companyLinks,
		CompanyText:   companyText,
		UpdatedMillis: updatedMillis,
	}
	return c, nil
}

func findExistingCandidates(dir string) ([]candidate, error) {
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
		return nil, err
	}

	var cands []candidate
	for _, f := range files {
		var c candidate
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, &c); err != nil {
			return nil, err
		}
		cands = append(cands, c)
	}

	return cands, nil
}
