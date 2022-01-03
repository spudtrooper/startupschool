package startupschool

import "time"

// genopts --opt_type=FindLinkedInProfilesOption --prefix=FindLinkedInProfiles --outfile=startupschool/findlinkedinprofilesoptions.go 'pause:time.Duration'

type FindLinkedInProfilesOption func(*findLinkedInProfilesOptionImpl)

type FindLinkedInProfilesOptions interface {
	Pause() time.Duration
}

func FindLinkedInProfilesPause(pause time.Duration) FindLinkedInProfilesOption {
	return func(opts *findLinkedInProfilesOptionImpl) {
		opts.pause = pause
	}
}

type findLinkedInProfilesOptionImpl struct {
	pause time.Duration
}

func (f *findLinkedInProfilesOptionImpl) Pause() time.Duration { return f.pause }

func makeFindLinkedInProfilesOptionImpl(opts ...FindLinkedInProfilesOption) *findLinkedInProfilesOptionImpl {
	res := &findLinkedInProfilesOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeFindLinkedInProfilesOptions(opts ...FindLinkedInProfilesOption) FindLinkedInProfilesOptions {
	return makeFindLinkedInProfilesOptionImpl(opts...)
}
