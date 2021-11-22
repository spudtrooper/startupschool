package startupschool

import "log"

func (b *bot) Backfill() error {
	cands, err := findExistingCandidates(b.d.Dir())
	if err != nil {
		return err
	}
	for _, c := range cands {
		c, err := b.lookUp(c.URI)
		if err != nil {
			return err
		}
		if err := b.d.SaveCandidate(*c); err != nil {
			return err
		}
	}

	return nil
}

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
	}

	return nil
}

func (b *bot) SearchURIs(uris []string) error {
	for _, uri := range uris {
		c, err := b.lookUp(uri)
		if err != nil {
			return err
		}
		if err := b.d.SaveCandidate(*c); err != nil {
			return err
		}
	}

	return nil
}
