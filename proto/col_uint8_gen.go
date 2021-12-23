// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColUInt8 represents UInt8 column.
type ColUInt8 []uint8

// Compile-time assertions for ColUInt8.
var (
	_ ColInput  = ColUInt8{}
	_ ColResult = (*ColUInt8)(nil)
	_ Column    = (*ColUInt8)(nil)
)

// Type returns ColumnType of UInt8.
func (ColUInt8) Type() ColumnType {
	return ColumnTypeUInt8
}

// Rows returns count of rows in column.
func (c ColUInt8) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt8) Reset() {
	*c = (*c)[:0]
}

// NewArrUInt8 returns new Array(UInt8).
func NewArrUInt8() *ColArr {
	return &ColArr{
		Data: new(ColUInt8),
	}
}

// AppendUInt8 appends slice of uint8 to Array(UInt8).
func (c *ColArr) AppendUInt8(data []uint8) {
	d := c.Data.(*ColUInt8)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes UInt8 rows to *Buffer.
func (c ColUInt8) EncodeColumn(b *Buffer) {
	b.Buf = append(b.Buf, c...)
}

// DecodeColumn decodes UInt8 rows from *Reader.
func (c *ColUInt8) DecodeColumn(r *Reader, rows int) error {
	data, err := r.ReadRaw(rows)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	*c = append(*c, data...)
	return nil
}
