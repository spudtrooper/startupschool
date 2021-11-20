package startupschool

import (
	"fmt"
	"log"
	"strings"
	"time"

	goutilselenium "github.com/spudtrooper/goutil/selenium"
	"github.com/tebeka/selenium"
)

func (b *bot) Loop(loopOpts ...LoopOption) error {
	opts := makeLoopOptionImpl(loopOpts...)

	wd, cancel, err := goutilselenium.MakeWebDriver(goutilselenium.MakeWebDriverOptions{
		Verbose:  opts.seleniumVerbose,
		Headless: !opts.seleniumHead,
	})
	if err != nil {
		return err
	}
	defer cancel()

	if err := b.login(wd); err != nil {
		return err
	}

	d, err := makeData()
	if err != nil {
		return err
	}
	for i := 0; ; i++ {
		if opts.limit > 0 && i == opts.limit {
			log.Printf("hit the limit: %d", opts.limit)
			break
		}
		c, err := b.nextCandidate(wd, d)
		if err != nil {
			return err
		}
		if err := d.SaveCandidate(*c); err != nil {
			return err
		}
	}

	return nil
}

func (b *bot) nextCandidate(wd selenium.WebDriver, d *data) (*candidate, error) {
	log.Printf("next candidate...")

	if err := wd.Get("https://www.startupschool.org/cofounder-matching/candidate/next"); err != nil {
		return nil, err
	}

	var name string
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
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

	var linkedInUri string
	as, err := wd.FindElements(selenium.ByTagName, "a")
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

	tables, err := wd.FindElements(selenium.ByTagName, "table")
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

	url, err := wd.CurrentURL()
	if err != nil {
		return nil, err
	}

	c := &candidate{
		Name:        name,
		LinkedInUri: linkedInUri,
		Intro:       intro,
		URI:         url,
	}
	return c, nil
}

func (b *bot) login(wd selenium.WebDriver) error {
	log.Printf("logging in...")

	if err := wd.Get("https://account.ycombinator.com/?continue=https%3A%2F%2Fwww.startupschool.org%2Fusers%2Fsign_in"); err != nil {
		return err
	}

	findCredsInputs := func() (selenium.WebElement, selenium.WebElement, error) {
		usernameInput, err := wd.FindElement(selenium.ByID, "ycid-input")
		if err != nil {
			return nil, nil, err
		}
		if usernameInput == nil {
			return nil, nil, nil
		}
		passwordInput, err := wd.FindElement(selenium.ByID, "password-input")
		if err != nil {
			return nil, nil, err
		}
		if passwordInput == nil {
			return nil, nil, nil
		}
		return usernameInput, passwordInput, nil
	}
	var usernameInput, passwordInput selenium.WebElement
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		log.Printf("waiting for username/password inputs...")
		u, p, err := findCredsInputs()
		if err != nil {
			return false, err
		}
		if u == nil || p == nil {
			return false, nil
		}
		usernameInput = u
		passwordInput = p
		return true, nil
	})
	if usernameInput == nil {
		return fmt.Errorf("no username input")
	}
	if passwordInput == nil {
		return fmt.Errorf("no password input")
	}

	if err := usernameInput.Clear(); err != nil {
		return err
	}
	usernameInput.SendKeys(b.creds.Username)

	if err := passwordInput.Clear(); err != nil {
		return err
	}
	passwordInput.SendKeys(b.creds.Password)

	loginBtn, err := wd.FindElement(selenium.ByClassName, "sign-in-button")
	if err != nil {
		return err
	}
	if err := loginBtn.Click(); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	log.Printf("logged in")
	return nil
}
