package multipart

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

type Writer struct {
	*multipart.Writer
	buf bytes.Buffer
}

func (w *Writer) Reader() io.Reader {
	return &w.buf
}

func (w *Writer) Close() error {
	return w.Writer.Close()
}

func NewWriter() *Writer {
	t := new(Writer)
	t.buf = bytes.Buffer{}
	t.Writer = multipart.NewWriter(&t.buf)
	return t
}

func AttachField(writer *Writer, keyname, keyvalue string) error {
	if err := writer.WriteField(keyname, keyvalue); err != nil {
		return fmt.Errorf("cannot WriteField: %s, err: %v", keyname, err)
	}
	return nil
}

func AttachFile(writer *Writer, fieldname string, filename string, file []byte) error {
	// file
	part, err := writer.CreateFormFile(fieldname, filename)
	if err != nil {
		return fmt.Errorf("cannot CreateFormFile for: %s , err: %v", filename, err)
	}
	_, err = part.Write(file)
	if err != nil {
		return fmt.Errorf("cannot Copy file: %s , err: %v", filename, err)
	}
	return nil
}
