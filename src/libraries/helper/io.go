package helper

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	},
}

func IoCopy(r io.Reader) ([]byte, error) {
	dst := bufferPool.Get().(*bytes.Buffer)
	dst.Reset()
	defer func() {
		if dst != nil {
			bufferPool.Put(dst)
			dst = nil
		}
	}()

	if _, err := io.Copy(dst, r); err != nil {
		return nil, err
	}
	bodyB := dst.Bytes()

	bufferPool.Put(dst)
	dst = nil

	return bodyB, nil
}

func DrainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
