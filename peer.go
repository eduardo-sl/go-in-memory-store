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
		var cmd Command
		if v.Type() == resp.Array {
			rawCMD := v.Array()[0]
			switch rawCMD.String() {
			case CommandClient:
				cmd = ClientCommand{
					value: v.Array()[1].String(),
				}
			case CommandGET:
				cmd = GetCommand{
					key: v.Array()[1].Bytes(),
				}
			case CommandDEL:
				value := make([]Value, len(v.Array())-1)
				for i := 1; i < len(v.Array()); i++ {
					value[i-1].key = v.Array()[i].Bytes()
				}
				cmd = DelCommand{
					val: value,
				}
			case CommandSET:
				cmd = SetCommand{
					key: v.Array()[1].Bytes(),
					val: v.Array()[2].Bytes(),
				}
			case CommandExists:
				cmd = ExistsCommand{
					key: v.Array()[1].Bytes(),
				}
			case CommandHELLO:
				cmd = HelloCommand{
					value: v.Array()[1].String(),
				}
			default:
				fmt.Println("got this unhandled command", rawCMD)
			}
			p.msgCh <- Message{
				cmd:  cmd,
				peer: p,
			}

		}
	}
	return nil
}
