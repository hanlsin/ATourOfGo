package exercise

import (
	"io"
)

type MyReader struct {
	stream string
	offset int
}

func NewMyReader(stream string) *MyReader {
	return &MyReader{
		stream: stream,
		offset: 0,
	}
}

func (mr *MyReader) Read(buffer []byte) (int, error) {
	if mr.stream == "" {
		for i := range buffer {
			buffer[i] = 'A'
		}
		return len(buffer), nil
	}

	streamLen := len(mr.stream)
	if streamLen == mr.offset {
		mr.offset = 0
		return 0, io.EOF
	}

	max := cap(buffer)
	if len(mr.stream) < (mr.offset + max) {
		max = len(mr.stream) - mr.offset
	}

	idx := 0
	for ; idx < max; idx++ {
		buffer[idx] = mr.stream[idx + mr.offset]
	}
	mr.offset += idx

	return idx, nil
}
