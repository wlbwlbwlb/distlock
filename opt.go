package distlock

type Options struct {
}

type Option interface {
	apply(*Options)
}

type Endpoints struct {
	endpoints []string
}

func (o Endpoints) apply(*Options) {
	_opts.Endpoints = o.endpoints
}

func WithEndpoints(endpoints []string) Option {
	return Endpoints{
		endpoints: endpoints,
	}
}
