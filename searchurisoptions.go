package startupschool

import "time"

//go:generate genopts --opt_type=SearchURIsOption --prefix=SearchURIs --outfile=searchurisoptions.go "pause:time.Duration"

type SearchURIsOption func(*searchURIsOptionImpl)

type SearchURIsOptions interface {
	Pause() time.Duration
}

func SearchURIsPause(pause time.Duration) SearchURIsOption {
	return func(opts *searchURIsOptionImpl) {
		opts.pause = pause
	}
}

type searchURIsOptionImpl struct {
	pause time.Duration
}

func (s *searchURIsOptionImpl) Pause() time.Duration { return s.pause }

func makeSearchURIsOptionImpl(opts ...SearchURIsOption) *searchURIsOptionImpl {
	res := &searchURIsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSearchURIsOptions(opts ...SearchURIsOption) SearchURIsOptions {
	return makeSearchURIsOptionImpl(opts...)
}
