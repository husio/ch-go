// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"github.com/go-faster/errors"
)

// ColInt8 represents Int8 column.
type ColInt8 []int8

// Compile-time assertions for ColInt8.
var (
	_ Input  = ColInt8{}
	_ Result = (*ColInt8)(nil)
)

// Type returns ColumnType of Int8.
func (ColInt8) Type() ColumnType {
	return ColumnTypeInt8
}

// Rows returns count of rows in column.
func (c ColInt8) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt8) Reset() {
	*c = (*c)[:0]
}

// EncodeColumn encodes Int8 rows to *Buffer.
func (c ColInt8) EncodeColumn(b *Buffer) {
	start := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, len(c))...)
	for i := range c {
		b.Buf[i+start] = uint8(c[i])
	}
}

// DecodeColumn decodes Int8 rows from *Reader.
func (c *ColInt8) DecodeColumn(r *Reader, rows int) error {
	data, err := r.ReadRaw(rows)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	v = append(v, make([]int8, rows)...)
	for i := range data {
		v[i] = int8(data[i])
	}
	*c = v
	return nil
}