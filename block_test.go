package sia

import (
	"bytes"
	"testing"
)

func TestReaderOK(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x46,
		0x23,
		0x30,
		0x31,
		0x37,
		0x36,
		0x35,
		0x31,
		0x9e,
	})

	reader := NewReader(buffer)
	block, err := reader.Read()
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}

	if block.Function != 0x23 {
		t.Fatalf("Function mismatch; expected %x, got %x", 0x23, block.Function)
	}

	if len(block.Data) != 6 {
		t.Fatalf("Length mismatch; expected %d, got %d", 6, len(block.Data))
	}
}

func TestReaderParity(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x46,
		0x23,
		0x30,
		0x31,
		0x37,
		0x36,
		0x35,
		0x31,
		0x9f,
	})

	reader := NewReader(buffer)
	block, err := reader.Read()
	if block != nil || err == nil {
		t.Fatal("Expected error")
	}

	if err != ErrParity {
		t.Fatal("Expected parity error")
	}
}

func TestWriteThenRead(t *testing.T) {
	var buffer bytes.Buffer
	var err error

	r := NewReader(&buffer)
	w := NewWriter(&buffer)
	b1 := Block{0x80, []byte{0x01, 0x02}}

	err = w.Write(b1)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}

	b2, err := r.Read()
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}

	if b1.Function != b2.Function {
		t.Fatalf("Block function mismatch")
	}

	if bytes.Compare(b1.Data, b2.Data) != 0 {
		t.Fatalf("Block data mismatch")
	}
}
