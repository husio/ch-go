// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"github.com/go-faster/errors"
)

// ColDecimal256 represents Decimal256 column.
type ColDecimal256 []Decimal256

// Compile-time assertions for ColDecimal256.
var (
	_ ColInput  = ColDecimal256{}
	_ ColResult = (*ColDecimal256)(nil)
	_ Column    = (*ColDecimal256)(nil)
)

// Type returns ColumnType of Decimal256.
func (ColDecimal256) Type() ColumnType {
	return ColumnTypeDecimal256
}

// Rows returns count of rows in column.
func (c ColDecimal256) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDecimal256) Reset() {
	*c = (*c)[:0]
}

// NewArrDecimal256 returns new Array(Decimal256).
func NewArrDecimal256() *ColArr {
	return &ColArr{
		Data: new(ColDecimal256),
	}
}

// AppendDecimal256 appends slice of Decimal256 to Array(Decimal256).
func (c *ColArr) AppendDecimal256(data []Decimal256) {
	d := c.Data.(*ColDecimal256)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Decimal256 rows to *Buffer.
func (c ColDecimal256) EncodeColumn(b *Buffer) {
	const size = 256 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binPutUInt256(
			b.Buf[offset:offset+size],
			UInt256(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Decimal256 rows from *Reader.
func (c *ColDecimal256) DecodeColumn(r *Reader, rows int) error {
	const size = 256 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	for i := 0; i < len(data); i += size {
		v = append(v,
			Decimal256(binUInt256(data[i:i+size])),
		)
	}
	*c = v
	return nil
}
