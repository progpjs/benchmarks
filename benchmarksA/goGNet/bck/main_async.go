package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/evanphx/wildcat"
	"github.com/panjf2000/gnet/v2"
)

// https://gnet.host/docs/quickstart/
// https://github.com/gnet-io/gnet-examples

var (
	errMsg      = "Internal Server Error"
	errMsgBytes = []byte(errMsg)
)

type httpServer struct {
	gnet.BuiltinEventEngine

	addr      string
	multicore bool
	eng       gnet.Engine
}

type httpCodec struct {
	parser *wildcat.HTTPParser
	buf    []byte
}

var gCounter = 0

func (hc *httpCodec) appendResponse() {
	gCounter++
	txt := fmt.Sprintf("Hello from GNet ! (%d)", gCounter)
	size := strconv.Itoa(len(txt))

	hc.buf = append(hc.buf, "HTTP/1.1 200 OK\r\nServer: gnet\r\nContent-Type: text/plain\r\nDate: "...)
	hc.buf = time.Now().AppendFormat(hc.buf, "Mon, 02 Jan 2006 15:04:05 GMT")
	hc.buf = append(hc.buf, "\r\nContent-Length: "+size+"\r\n\r\n"+txt...)
}

func (hs *httpServer) OnBoot(eng gnet.Engine) gnet.Action {
	hs.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", hs.multicore, hs.addr)
	return gnet.None
}

func (hs *httpServer) OnOpen(c gnet.Conn) ([]byte, gnet.Action) {
	// Le context va permettre à OnTraffic de retrouver l'instance associée à la requête.
	c.SetContext(&httpCodec{parser: wildcat.NewHTTPParser()})
	return nil, gnet.None
}

func (hs *httpServer) OnTraffic(c gnet.Conn) gnet.Action {
	hc := c.Context().(*httpCodec)
	buf, _ := c.Next(-1)

	//region Read the headers

pipeline:
	headerOffset, err := hc.parser.Parse(buf)
	if err != nil {
		c.Write(errMsgBytes)
		return gnet.Close
	}

	hc.appendResponse()
	bodyLen := int(hc.parser.ContentLength())
	if bodyLen == -1 {
		bodyLen = 0
	}

	buf = buf[headerOffset+bodyLen:]
	if len(buf) > 0 {
		// Cas de requête concaténées ?
		goto pipeline
	}

	//endregion

	//c.Write(hc.buf)
	//hc.buf = hc.buf[:0]

	myCount := gCounter

	go func() {
		PauseMs(1000)

		err := c.AsyncWrite(hc.buf, func(c gnet.Conn, err error) error {
			println("Responsed to ", myCount)
			hc.buf = hc.buf[:0]

			if err != nil {
				println(err.Error())
			}

			return nil
		})

		if err != nil {
			println(err.Error())
		}
	}()

	return gnet.None
}

func PauseMs(timeInMs int) {
	duration := time.Millisecond * time.Duration(timeInMs)
	time.Sleep(duration)
}

func main() {
	var port int
	var multicore bool

	flag.IntVar(&port, "port", 8000, "server port")
	flag.BoolVar(&multicore, "multicore", true, "multicore")
	flag.Parse()

	hs := &httpServer{addr: fmt.Sprintf("tcp://127.0.0.1:%d", port), multicore: multicore}

	// Start serving!
	log.Println("server exits:", gnet.Run(hs, hs.addr, gnet.WithMulticore(multicore)))
}
