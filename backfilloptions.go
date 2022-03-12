package startupschool

import "time"

//go:generate genopts --prefix=Backfill --outfile=backfilloptions.go "pause:time.Duration"

type BackfillOption func(*backfillOptionImpl)

type BackfillOptions interface {
	Pause() time.Duration
}

func BackfillPause(pause time.Duration) BackfillOption {
	return func(opts *backfillOptionImpl) {
		opts.pause = pause
	}
}
func BackfillPauseFlag(pause *time.Duration) BackfillOption {
	return func(opts *backfillOptionImpl) {
		opts.pause = *pause
	}
}

type backfillOptionImpl struct {
	pause time.Duration
}

func (b *backfillOptionImpl) Pause() time.Duration { return b.pause }

func makeBackfillOptionImpl(opts ...BackfillOption) *backfillOptionImpl {
	res := &backfillOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeBackfillOptions(opts ...BackfillOption) BackfillOptions {
	return makeBackfillOptionImpl(opts...)
}
