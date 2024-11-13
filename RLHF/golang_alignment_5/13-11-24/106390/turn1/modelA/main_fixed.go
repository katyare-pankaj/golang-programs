package main

import (
	"bytes"
	"fmt"
	"io"
)

const chunkSize = 10 // Adjusted chunk size for demonstration purposes

// ChunkReader reads data from r and returns a new reader that returns data in chunks of at most size bytes.
func ChunkReader(r io.Reader, size int) io.Reader {
	return &chunkReader{r: r, size: size}
}

type chunkReader struct {
	r    io.Reader
	size int
	buf  bytes.Buffer
	eof  bool
}

func (cr *chunkReader) Read(p []byte) (int, error) {
	// Return any remaining data in the buffer first
	if cr.buf.Len() > 0 {
		return cr.buf.Read(p)
	}

	// Check if EOF has been reached previously
	if cr.eof {
		return 0, io.EOF
	}

	// Read data in chunks from the underlying reader
	temp := make([]byte, cr.size)
	n, err := cr.r.Read(temp)
	if err != nil {
		if err == io.EOF {
			cr.eof = true
			if n == 0 {
				return 0, io.EOF
			}
		} else {
			return 0, err
		}
	}

	// Write the data into the buffer
	cr.buf.Write(temp[:n])

	// Read from buffer into p
	return cr.buf.Read(p)
}

func main() {
	// Example data to be sent in chunks
	data := []byte("This is some sample data that needs to be sent in chunks.")

	// Wrap the data reader with the ChunkReader
	chunkedReader := ChunkReader(bytes.NewReader(data), chunkSize)

	// Receive and process the chunks
	buf := make([]byte, chunkSize)
	for {
		n, err := chunkedReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading:", err)
			return
		}

		// Process the received chunk here
		fmt.Printf("Received chunk: %s\n", string(buf[:n]))
	}
}
