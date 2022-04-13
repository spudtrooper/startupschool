package graphql

//go:generate genopts --prefix=Query --outfile=queryoptions.go "variables:map[string]string"

type QueryOption func(*queryOptionImpl)

type QueryOptions interface {
	Variables() map[string]string
}

func QueryVariables(variables map[string]string) QueryOption {
	return func(opts *queryOptionImpl) {
		opts.variables = variables
	}
}
func QueryVariablesFlag(variables *map[string]string) QueryOption {
	return func(opts *queryOptionImpl) {
		opts.variables = *variables
	}
}

type queryOptionImpl struct {
	variables map[string]string
}

func (q *queryOptionImpl) Variables() map[string]string { return q.variables }

func makeQueryOptionImpl(opts ...QueryOption) *queryOptionImpl {
	res := &queryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeQueryOptions(opts ...QueryOption) QueryOptions {
	return makeQueryOptionImpl(opts...)
}
