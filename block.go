package sia

import (
	"bufio"
	"errors"
	"io"
)

type Block struct {
	Function byte
	Data     []byte
}

const kBaseParity byte = 0xff
const kLengthOffset byte = 0x40

var ErrParity = errors.New("sia: parity error")

type Reader struct {
	r *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewReader(r)}
}

func (r *Reader) Read() (*Block, error) {
	b1, err := r.r.ReadByte()
	if err != nil {
		return nil, err
	}

	b2, err := r.r.ReadByte()
	if err != nil {
		return nil, err
	}

	length := b1 - kLengthOffset
	block := Block{
		b2,
		make([]byte, length),
	}

	_, err = io.ReadFull(r.r, block.Data)
	if err != nil {
		return nil, err
	}

	parity, err := r.r.ReadByte()
	if err != nil {
		return nil, err
	}

	// Compute parity
	p := kBaseParity ^ b1 ^ b2
	for _, b := range block.Data {
		p ^= b
	}

	// Compare parity
	if p != parity {
		return nil, ErrParity
	}

	return &block, nil
}

type Writer struct {
	w *bufio.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{bufio.NewWriter(w)}
}

func (w *Writer) Write(block Block) error {
	var err error

	// Write length byte
	length := len(block.Data)
	b1 := byte(length) + kLengthOffset
	err = w.w.WriteByte(b1)
	if err != nil {
		return err
	}

	// Write function code byte
	b2 := block.Function
	err = w.w.WriteByte(b2)
	if err != nil {
		return err
	}

	// Write data
	_, err = w.w.Write(block.Data)
	if err != nil {
		return err
	}

	// Compute parity
	parity := kBaseParity ^ b1 ^ b2
	for _, b := range block.Data {
		parity ^= b
	}

	// Write parity
	err = w.w.WriteByte(parity)
	if err != nil {
		return err
	}

	return w.w.Flush()
}
