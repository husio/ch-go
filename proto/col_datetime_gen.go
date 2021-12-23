// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColDateTime represents DateTime column.
type ColDateTime []DateTime

// Compile-time assertions for ColDateTime.
var (
	_ ColInput  = ColDateTime{}
	_ ColResult = (*ColDateTime)(nil)
	_ Column    = (*ColDateTime)(nil)
)

// Type returns ColumnType of DateTime.
func (ColDateTime) Type() ColumnType {
	return ColumnTypeDateTime
}

// Rows returns count of rows in column.
func (c ColDateTime) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDateTime) Reset() {
	*c = (*c)[:0]
}

// NewArrDateTime returns new Array(DateTime).
func NewArrDateTime() *ColArr {
	return &ColArr{
		Data: new(ColDateTime),
	}
}

// AppendDateTime appends slice of DateTime to Array(DateTime).
func (c *ColArr) AppendDateTime(data []DateTime) {
	d := c.Data.(*ColDateTime)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes DateTime rows to *Buffer.
func (c ColDateTime) EncodeColumn(b *Buffer) {
	const size = 32 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint32(
			b.Buf[offset:offset+size],
			uint32(v),
		)
		offset += size
	}
}

// DecodeColumn decodes DateTime rows from *Reader.
func (c *ColDateTime) DecodeColumn(r *Reader, rows int) error {
	const size = 32 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			DateTime(binary.LittleEndian.Uint32(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
