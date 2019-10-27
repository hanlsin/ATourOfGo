package exercise

import (
	"io"
	"strings"
)

/*
ROT13: "rotate by 13 places"
https://en.wikipedia.org/wiki/ROT13

ex:
Index:  0123456789012 3456789012345 6789012345678 9012345678901
Input:  ABCDEFGHIJKLM NOPQRSTUVWXYZ abcdefghijklm nopqrstuvwxyz
Output: NOPQRSTUVWXYZ ABCDEFGHIJKLM nopqrstuvwxyz abcdefghijklm
 */

type Rot13Reader struct {
	stream *strings.Reader
	offset int
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const rot13bet = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

func NewReader(stream *strings.Reader) *Rot13Reader {
	return &Rot13Reader{
		stream: stream,
		offset: 0,
	}
}

func (rot13 *Rot13Reader) Read(buffer []byte) (int, error) {
	if rot13 == nil || rot13.stream == nil {
		return 0, io.EOF
	}

	streamLen := rot13.stream.Len()
	if streamLen == rot13.offset {
		return 0, io.EOF
	}

	max := cap(buffer)
	if streamLen < max {
		max = streamLen
	}

	for i := 0; i < max; i++ {
		idx, err := rot13.stream.ReadByte()
		if idx >= 65 && idx <= 90 {
			idx -= 65
			buffer[i] = rot13bet[idx]
		} else if idx >= 97 && idx <= 122 {
			idx = idx - 97 + 26
			buffer[i] = rot13bet[idx]
		} else {
			buffer[i] = idx
		}

		if err != nil {
			return len(buffer), err
		}
	}

	return len(buffer), nil
}
