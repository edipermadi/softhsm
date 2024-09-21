package slip

import (
	"bufio"
	"io"
)

func NewReader(r io.Reader) io.Reader {
	return &reader{reader: bufio.NewReader(r)}
}

type reader struct {
	reader io.ByteReader
}

func (r *reader) Read(buffer []byte) (int, error) {
	if cap(buffer) == 0 {
		return 0, nil
	}

	buffer = buffer[:0]
	started := false
	escaping := false
	read := 0
	for remaining := cap(buffer); remaining > 0; {
		read++
		b, err := r.reader.ReadByte()
		if err != nil {
			return read, err
		}

		switch {
		case !started && b == End:
			started = true
		case started && b == End:
			return read, nil
		case started && !escaping && b == Esc:
			escaping = true
		case started && escaping && b == EscEsc:
			buffer = append(buffer, Esc)
			escaping = false
			remaining--
		case started && escaping && b == EscEnd:
			buffer = append(buffer, End)
			escaping = false
			remaining--
		case started && !escaping:
			buffer = append(buffer, b)
			remaining--
		}
	}

	return read + 1, nil
}
