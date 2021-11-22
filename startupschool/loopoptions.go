package startupschool

// genopts --opt_type=LoopOption --prefix=Loop --outfile=startupschool/loopoptions.go 'limit:int'

type LoopOption func(*loopOptionImpl)

type LoopOptions interface {
	Limit() int
}

func LoopLimit(limit int) LoopOption {
	return func(opts *loopOptionImpl) {
		opts.limit = limit
	}
}

type loopOptionImpl struct {
	limit int
}

func (l *loopOptionImpl) Limit() int { return l.limit }

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
