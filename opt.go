package dlock

type Options struct {
}

type Option interface {
	apply(*Options)
}

type Endpoints struct {
	endpoints []string
}

func (ep Endpoints) apply(*Options) {
	_opts.Endpoints = ep.endpoints
}

func WithEndpoints(endpoints []string) Option {
	return Endpoints{
		endpoints: endpoints,
	}
}
