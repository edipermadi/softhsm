package slip

import "io"

func NewWriter(w io.Writer) io.Writer {
	return &writer{writer: w}
}

type writer struct {
	writer io.Writer
}

func (w *writer) Write(data []byte) (int, error) {
	mapping := map[byte]byte{
		Esc: EscEsc,
		End: EscEnd,
	}

	length := len(data) + 2
	for _, v := range data {
		if v == Esc || v == End {
			length++
		}
	}

	idx := 0
	buffer := make([]byte, length)

	buffer[idx] = End
	idx++

	for _, v := range data {
		if v == Esc || v == End {
			buffer[idx] = Esc
			idx++
			v = mapping[v]
		}
		buffer[idx] = v
		idx++
	}

	buffer[idx] = End

	return w.writer.Write(buffer)
}
