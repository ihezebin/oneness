package excel

import "github.com/xuri/excelize/v2"

const (
	defaultMaxRow    = 2 << 12
	defaultMaxColumn = 1000
)

type Options = excelize.Options

type ReadOptions struct {
	MaxRow    int
	MaxColumn int
}

type ReadOption func(opt *ReadOptions)

// WithMaxRow set the max row to read when get all rows
func WithMaxRow(maxRow int) ReadOption {
	return func(opt *ReadOptions) {
		opt.MaxRow = maxRow
	}
}

// WithMaxColumn set the max column to read when get all rows
func WithMaxColumn(maxColumn int) ReadOption {
	return func(opt *ReadOptions) {
		opt.MaxColumn = maxColumn
	}
}
