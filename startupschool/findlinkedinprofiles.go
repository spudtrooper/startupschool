package startupschool

import (
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
	"github.com/tebeka/selenium"
)

type LinkedInURI string

const (
	nilLinkedInURI LinkedInURI = ""
)

func (b *bot) FindLinkedInProfiles(fliOpts ...FindLinkedInProfilesOption) ([]LinkedInURI, error) {
	opts := MakeFindLinkedInProfilesOptions(fliOpts...)

	// TODO: Can't figure out how to dismiss the modal, so request the main page N times
	var linkedInUris []LinkedInURI
	start := or.Int(opts.Start(), 0)
	for i := start; ; i++ {
		uri, err := b.findLinkedInProfile(i)
		if err != nil {
			return nil, err
		}
		if uri == nilLinkedInURI {
			break
		}
		log.Printf("Found %s", uri)
		linkedInUris = append(linkedInUris, uri)
		if opts.Pause() > 0 {
			log.Printf("sleeping for %v", opts.Pause())
			time.Sleep(opts.Pause())
		}
	}

	return linkedInUris, nil
}

func (b *bot) waitForElementsByClassName(className string) ([]selenium.WebElement, error) {
	var res []selenium.WebElement
	var cnt int
	b.wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		log.Printf("waiting for element with className %s  [%d] ...", className, cnt+1)
		cnt++
		els, err := wd.FindElements(selenium.ByCSSSelector, className)
		if err != nil {
			return false, err
		}
		if len(els) > 0 {
			res = els
			return true, nil
		}
		return false, nil
	})
	if res == nil {
		return nil, fmt.Errorf("couldn't find div with className: %s", className)
	}
	return res, nil
}

func (b *bot) findLinkedInProfile(num int) (LinkedInURI, error) {
	if err := b.wd.Get("https://www.startupschool.org/cofounder-matching"); err != nil {
		return nilLinkedInURI, err
	}

	profiles, err := b.waitForElementsByClassName(".css-e8vpit.egpbwzr2")
	if err != nil {
		return nilLinkedInURI, err
	}

	log.Printf("have %d profiles", len(profiles))

	if len(profiles) == 0 {
		return nilLinkedInURI, errors.Errorf("no profiles")
	}

	if num >= len(profiles) {
		return nilLinkedInURI, nil
	}

	if err := profiles[num].Click(); err != nil {
		return nilLinkedInURI, err
	}

	b.wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		log.Printf("waiting for respond to invite text...")
		h2s, err := wd.FindElements(selenium.ByTagName, "h2")
		if err != nil {
			return false, nil
		}
		for _, h2 := range h2s {
			text, err := h2.Text()
			if err != nil {
				return false, nil
			}
			if text == "Respond to Invite" {
				return true, nil
			}
		}
		return false, nil
	})

	as, err := b.wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return nilLinkedInURI, err
	}
	for _, a := range as {
		text, err := a.Text()
		if err != nil {
			return nilLinkedInURI, err
		}
		if text == "LinkedIn" {
			href, err := a.GetAttribute("href")
			if err != nil {
				return nilLinkedInURI, err
			}
			return LinkedInURI(href), nil
		}
	}

	return nilLinkedInURI, nil
}
