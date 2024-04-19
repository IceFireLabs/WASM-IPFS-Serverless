package openlog

type Tags map[string]interface{}

type Options struct {
	Tags  Tags
	Depth int
	Err   error
}
type Option func(*Options)

func WithTags(tags Tags) Option {
	return func(o *Options) {
		o.Tags = tags
	}
}
func WithDepth(d int) Option {
	return func(o *Options) {
		o.Depth = d
	}
}
func WithErr(d error) Option {
	return func(o *Options) {
		o.Err = d
	}
}

func ToOptions(opts ...Option) *Options {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

// Logger is a interface for log tool
type Logger interface {
	Debug(message string, opts ...Option)
	Info(message string, opts ...Option)
	Warn(message string, opts ...Option)
	Error(message string, opts ...Option)
	Fatal(message string, opts ...Option)
}
