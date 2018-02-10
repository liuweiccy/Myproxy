package structure

import "net"

type Channel struct {
	SrcConn net.Conn
	DstConn net.Conn
}