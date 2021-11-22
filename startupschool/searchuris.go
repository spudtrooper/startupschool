package startupschool

import (
	"log"
	"time"
)

func (b *bot) SearchURIs(uris []string, searchURIsOpts ...SearchURIsOption) error {
	opts := makeSearchURIsOptionImpl(searchURIsOpts...)

	for _, uri := range uris {
		c, err := b.lookUp(uri)
		if err != nil {
			return err
		}
		if err := b.d.SaveCandidate(*c); err != nil {
			return err
		}
		if opts.pause > 0 {
			log.Printf("sleeping for %v", opts.pause)
			time.Sleep(opts.pause)
		}
	}

	return nil
}
