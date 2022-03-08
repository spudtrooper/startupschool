package startupschool

import "time"

//go:generate genopts --opt_type=LoopOption --prefix=Loop --outfile=loopoptions.go "limit:int" "pause:time.Duration"

type LoopOption func(*loopOptionImpl)

type LoopOptions interface {
	Limit() int
	Pause() time.Duration
}

func LoopLimit(limit int) LoopOption {
	return func(opts *loopOptionImpl) {
		opts.limit = limit
	}
}

func LoopPause(pause time.Duration) LoopOption {
	return func(opts *loopOptionImpl) {
		opts.pause = pause
	}
}

type loopOptionImpl struct {
	limit int
	pause time.Duration
}

func (l *loopOptionImpl) Limit() int           { return l.limit }
func (l *loopOptionImpl) Pause() time.Duration { return l.pause }

func makeLoopOptionImpl(opts ...LoopOption) *loopOptionImpl {
	res := &loopOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeLoopOptions(opts ...LoopOption) LoopOptions {
	return makeLoopOptionImpl(opts...)
}
