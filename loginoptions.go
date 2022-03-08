package startupschool

//go:generate genopts --opt_type=LoginOption --prefix=Login --outfile=loginoptions.go "seleniumVerbose" "seleniumHead"

type LoginOption func(*loginOptionImpl)

type LoginOptions interface {
	SeleniumVerbose() bool
	SeleniumHead() bool
}

func LoginSeleniumVerbose(seleniumVerbose bool) LoginOption {
	return func(opts *loginOptionImpl) {
		opts.seleniumVerbose = seleniumVerbose
	}
}

func LoginSeleniumHead(seleniumHead bool) LoginOption {
	return func(opts *loginOptionImpl) {
		opts.seleniumHead = seleniumHead
	}
}

type loginOptionImpl struct {
	seleniumVerbose bool
	seleniumHead    bool
}

func (l *loginOptionImpl) SeleniumVerbose() bool { return l.seleniumVerbose }
func (l *loginOptionImpl) SeleniumHead() bool    { return l.seleniumHead }

func makeLoginOptionImpl(opts ...LoginOption) *loginOptionImpl {
	res := &loginOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeLoginOptions(opts ...LoginOption) LoginOptions {
	return makeLoginOptionImpl(opts...)
}
