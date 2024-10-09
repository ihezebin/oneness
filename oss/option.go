package oss

type GetObjectsOptions struct {
	Limit int
}

type GetObjectsOption func(*GetObjectsOptions)

func newGetObjectsOptions(opts ...GetObjectsOption) *GetObjectsOptions {
	options := &GetObjectsOptions{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

func WithLimit() GetObjectsOption {
	return func(options *GetObjectsOptions) {
		options.Limit = 1000
	}

}

type PutOptions struct {
	Size        int64
	ContentType string
}

type PutOption func(*PutOptions)

func WithSize(size int64) PutOption {
	return func(opt *PutOptions) {
		opt.Size = size
	}
}

func WithContentType(contentType string) PutOption {
	return func(opt *PutOptions) {
		opt.ContentType = contentType
	}
}

func newPutOptions(opts ...PutOption) *PutOptions {
	opt := &PutOptions{
		Size: -1,
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	return opt
}
