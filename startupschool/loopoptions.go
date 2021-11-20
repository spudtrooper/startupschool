package startupschool

// ~/go/bin/genopts -opt_type=LoopOption --prefix=Loop seleniumVerbose seleniumHead limit:int

type LoopOption func(*loopOptionImpl)

type LoopOptions interface {
	SeleniumVerbose() bool
	SeleniumHead() bool
	Limit() int
}

func LoopSeleniumVerbose(seleniumVerbose bool) LoopOption {
	return func(opts *loopOptionImpl) {
		opts.seleniumVerbose = seleniumVerbose
	}
}

func LoopSeleniumHead(seleniumHead bool) LoopOption {
	return func(opts *loopOptionImpl) {
		opts.seleniumHead = seleniumHead
	}
}

func LoopLimit(limit int) LoopOption {
	return func(opts *loopOptionImpl) {
		opts.limit = limit
	}
}

type loopOptionImpl struct {
	seleniumVerbose bool
	seleniumHead    bool
	limit           int
}

func (l *loopOptionImpl) SeleniumVerbose() bool { return l.seleniumVerbose }
func (l *loopOptionImpl) SeleniumHead() bool    { return l.seleniumHead }
func (l *loopOptionImpl) Limit() int            { return l.limit }

func makeLoopOptionImpl(opts ...LoopOption) loopOptionImpl {
	var res loopOptionImpl
	for _, opt := range opts {
		opt(&res)
	}
	return res
}
