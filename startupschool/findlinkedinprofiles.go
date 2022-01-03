package startupschool

import (
	"log"
	"time"

	"github.com/pkg/errors"
	goutilselenium "github.com/spudtrooper/goutil/selenium"
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
	for i := 0; ; i++ {
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

func (b *bot) findLinkedInProfile(num int) (LinkedInURI, error) {
	if err := b.wd.Get("https://www.startupschool.org/cofounder-matching"); err != nil {
		return nilLinkedInURI, err
	}

	buttons, err := goutilselenium.WaitForElements(b.wd, "button", "Respond to invite")
	if err != nil {
		return nilLinkedInURI, err
	}

	if len(buttons) == 0 {
		return nilLinkedInURI, errors.Errorf("no buttons")
	}

	if num >= len(buttons) {
		return nilLinkedInURI, nil
	}

	if err := buttons[num].Click(); err != nil {
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
