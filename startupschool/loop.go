package startupschool

import (
	"log"
	"time"
)

func (b *bot) Loop(loopOpts ...LoopOption) error {
	opts := makeLoopOptionImpl(loopOpts...)

	for i := 0; ; i++ {
		if opts.limit > 0 && i == opts.limit {
			log.Printf("hit the limit: %d", opts.limit)
			break
		}
		c, err := b.lookUp("https://www.startupschool.org/cofounder-matching/candidate/next")
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
