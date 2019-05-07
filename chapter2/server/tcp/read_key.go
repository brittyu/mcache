package tcp

import (
	"bufio"
	"io"
)

func (this *Server) readKey(r *bufio.Reader) (string, error) {
	klen, err := readLen(r)
	if err != nil {
		return "", err
	}

	key := make([]byte, klen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

func (this *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	vlen, err := readLen(r)
	if err != nil {
		return "", nil, err
	}
	key := make([]byte, klen)
	_, err = io.ReadFull(r, key)
	if err != nil {
		return "", nil, err
	}
	value := make([]byte, vlen)
	_, err = io.ReadFull(r, value)
	if err != nil {
		return "", nil, err
	}
	return string(key), value, nil
}
