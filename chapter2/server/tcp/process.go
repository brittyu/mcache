package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
)

func (this *Server) get(conn net.Conn, r *bufio.Reader) error {
	key, err := this.readKey(r)
	if err != nil {
		return err
	}

	value, err := this.Get(key)
	return sendResponse(value, err, conn)
}

func (this *Server) set(conn net.Conn, r *bufio.Reader) error {
	key, value, err := this.readKeyAndValue(r)
	if err != nil {
		return err
	}

	return sendResponse(nil, this.Set(key, value), conn)
}

func (this *Server) del(conn net.Conn, r *bufio.Reader) error {
	key, err := this.readKey(r)
	if err != nil {
		return err
	}

	return sendResponse(nil, this.Del(key), conn)
}

func (this *Server) process(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		op, e := r.ReadByte()
		if e != nil {
			if e != io.EOF {
				log.Println("close connection due to error:", e)
			}
			return
		}

		if op == 'S' {
			e = this.set(conn, r)
		} else if op == 'G' {
			e = this.get(conn, r)
		} else if op == 'D' {
			e = this.del(conn, r)
		} else {
			log.Println("close connection due to invalid operation:", op)
			return
		}

		if e != nil {
			log.Println("close connection due to error:", e)
			return
		}
	}
}
