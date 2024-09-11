package main

import (
	"fmt"
	resp "go-in-memory-store/resp"
	"io"
	"log"
	"net"
)

type Peer struct {
	conn  net.Conn
	msgCh chan Message
	delCh chan *Peer
}

func NewPeer(conn net.Conn, msgCh chan Message, delCh chan *Peer) *Peer {
	return &Peer{
		conn:  conn,
		msgCh: msgCh,
		delCh: delCh,
	}
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func (p *Peer) readLoop() error {
	rd := resp.NewReader(p.conn)

	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			p.delCh <- p
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if v.Type() == resp.Array {
			cmd := v.Array()[0]
			switch cmd.String() {
			case CommandGET:
			case CommandSET:
			case CommandHELLO:
				cmd := HelloCommand{
					value: v.Array()[1].String(),
				}
				p.msgCh <- Message{
					cmd:  cmd,
					peer: p,
				}
			}
			fmt.Println("this should be the cmd", v.Array()[0])
		}
	}
	return nil
}
