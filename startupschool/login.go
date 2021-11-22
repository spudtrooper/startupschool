package startupschool

import (
	"fmt"
	"log"
	"time"

	goutilselenium "github.com/spudtrooper/goutil/selenium"
	"github.com/tebeka/selenium"
)

func (b *bot) Login(loginOpts ...LoginOption) (func(), error) {
	opts := makeLoginOptionImpl(loginOpts...)

	wd, cancel, err := goutilselenium.MakeWebDriver(goutilselenium.MakeWebDriverOptions{
		Verbose:  opts.seleniumVerbose,
		Headless: !opts.seleniumHead,
	})
	if err != nil {
		return cancel, err
	}

	if err := b.login(wd); err != nil {
		return cancel, err
	}

	return cancel, nil
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

	b.wd = wd

	d, err := makeData()
	if err != nil {
		return err
	}
	b.d = d

	log.Printf("logged in")

	return nil
}
