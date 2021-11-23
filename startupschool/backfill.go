package startupschool

import (
	"log"
	"time"
)

func (b *bot) Backfill(backfillOpts ...BackfillOption) error {
	opts := makeBackfillOptionImpl(backfillOpts...)

	cands, err := findExistingCandidates(b.d.Dir())
	if err != nil {
		return err
	}
	for _, c := range cands {
		if !c.NeedsBackfill() {
			log.Printf("skipping %v", c)
			continue
		}
		c, err := b.lookUp(c.URI)
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
