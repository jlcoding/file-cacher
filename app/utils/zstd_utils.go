package utils

import (
	"github.com/datadog/zstd"
	"io"
	"os"
)

type ZstdEncoder struct {
	File *os.File
	ZstdWriter *zstd.Writer
}

//level是压缩级别：1~19，越大越慢。默认为3

func NewZstdEncoder(filename string, level int) (*ZstdEncoder,error) {
	f,err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	w := zstd.NewWriterLevel(f, level)
	return &ZstdEncoder{File:f, ZstdWriter:w},nil
}

func (s *ZstdEncoder) Close() {

	s.ZstdWriter.Close()
	s.File.Close()
}

func ZstdCompressFile(dst,src string, compress_level int) error {
	//log.Println("Zstd ", src, "->", dst)
	s,err := NewZstdEncoder(dst,compress_level)
	defer s.Close()
	if err != nil {
		return err
	}
	src_file,err := os.Open(src)
	if err != nil {
		return err
	}
	defer src_file.Close()
	io.Copy(s.ZstdWriter, src_file)

	return nil
}

func ZstdDecompressFile(dst, src string) error {
	//log.Println("Zstd ", src, "->", dst)
	in_file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in_file.Close()
	out_file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out_file.Close()

	zst_reader := zstd.NewReader(in_file)
	io.Copy(out_file, zst_reader)

	return nil
}
