package bqproto

var defaultOption = option{
	tagKey:         "bigquery",
	fractionalTime: false,
}

type option struct {
	tagKey         string
	fractionalTime bool
}

func newOption(opts ...Option) option {
	opt := defaultOption

	for _, f := range opts {
		f(&opt)
	}

	return opt
}

type Option func(*option)

func WithTagKey(s string) Option {
	return func(o *option) {
		o.tagKey = s
	}
}

func WithFractionalTime() Option {
	return func(o *option) {
		o.fractionalTime = true
	}
}
