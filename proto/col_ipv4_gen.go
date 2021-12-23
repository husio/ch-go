// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColIPv4 represents IPv4 column.
type ColIPv4 []IPv4

// Compile-time assertions for ColIPv4.
var (
	_ ColInput  = ColIPv4{}
	_ ColResult = (*ColIPv4)(nil)
	_ Column    = (*ColIPv4)(nil)
)

// Type returns ColumnType of IPv4.
func (ColIPv4) Type() ColumnType {
	return ColumnTypeIPv4
}

// Rows returns count of rows in column.
func (c ColIPv4) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColIPv4) Reset() {
	*c = (*c)[:0]
}

// NewArrIPv4 returns new Array(IPv4).
func NewArrIPv4() *ColArr {
	return &ColArr{
		Data: new(ColIPv4),
	}
}

// AppendIPv4 appends slice of IPv4 to Array(IPv4).
func (c *ColArr) AppendIPv4(data []IPv4) {
	d := c.Data.(*ColIPv4)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes IPv4 rows to *Buffer.
func (c ColIPv4) EncodeColumn(b *Buffer) {
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

// DecodeColumn decodes IPv4 rows from *Reader.
func (c *ColIPv4) DecodeColumn(r *Reader, rows int) error {
	const size = 32 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			IPv4(binary.LittleEndian.Uint32(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
