package udp

import "net"

const (
	// make configurable?
	writeBuffer = 1024 * 1024
)

func Dial(addr string, options map[string]string) (net.Conn, error) {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return nil, err
	}
	// bump up the packet size for large log lines
	err = conn.SetWriteBuffer(writeBuffer)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
