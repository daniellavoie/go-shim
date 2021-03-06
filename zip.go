// Code generated by counterfeiter. DO NOT EDIT.
// with command: counterfeiter -p archive/zip
package goshim

import (
	"archive/zip"
	"io"
	"os"
)

//go:generate counterfeiter -o zip_fake/fake_zip.go . Zip
type Zip interface {
	OpenReader(name string) (*zip.ReadCloser, error)
	NewReader(r io.ReaderAt, size int64) (*zip.Reader, error)
	RegisterDecompressor(method uint16, dcomp zip.Decompressor)
	RegisterCompressor(method uint16, comp zip.Compressor)
	FileInfoHeader(fi os.FileInfo) (*zip.FileHeader, error)
	NewWriter(w io.Writer) *zip.Writer
}
