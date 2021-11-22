package startupschool

// genopts --opt_type=LoginOption --prefix=Login --outfile=startupschool/loginoptions.go 'seleniumVerbose' 'seleniumHead' 'data:string'

type LoginOption func(*loginOptionImpl)

type LoginOptions interface {
	SeleniumVerbose() bool
	SeleniumHead() bool
	Data() string
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

func LoginData(data string) LoginOption {
	return func(opts *loginOptionImpl) {
		opts.data = data
	}
}

type loginOptionImpl struct {
	seleniumVerbose bool
	seleniumHead    bool
	data            string
}

func (l *loginOptionImpl) SeleniumVerbose() bool { return l.seleniumVerbose }
func (l *loginOptionImpl) SeleniumHead() bool    { return l.seleniumHead }
func (l *loginOptionImpl) Data() string          { return l.data }

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
