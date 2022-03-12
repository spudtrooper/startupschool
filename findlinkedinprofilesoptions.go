package startupschool

import "time"

//go:generate genopts --prefix=FindLinkedInProfiles --outfile=findlinkedinprofilesoptions.go "pause:time.Duration" "start:int"

type FindLinkedInProfilesOption func(*findLinkedInProfilesOptionImpl)

type FindLinkedInProfilesOptions interface {
	Pause() time.Duration
	Start() int
}

func FindLinkedInProfilesPause(pause time.Duration) FindLinkedInProfilesOption {
	return func(opts *findLinkedInProfilesOptionImpl) {
		opts.pause = pause
	}
}
func FindLinkedInProfilesPauseFlag(pause *time.Duration) FindLinkedInProfilesOption {
	return func(opts *findLinkedInProfilesOptionImpl) {
		opts.pause = *pause
	}
}

func FindLinkedInProfilesStart(start int) FindLinkedInProfilesOption {
	return func(opts *findLinkedInProfilesOptionImpl) {
		opts.start = start
	}
}
func FindLinkedInProfilesStartFlag(start *int) FindLinkedInProfilesOption {
	return func(opts *findLinkedInProfilesOptionImpl) {
		opts.start = *start
	}
}

type findLinkedInProfilesOptionImpl struct {
	pause time.Duration
	start int
}

func (f *findLinkedInProfilesOptionImpl) Pause() time.Duration { return f.pause }
func (f *findLinkedInProfilesOptionImpl) Start() int           { return f.start }

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
