package logger

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/saheienko/clab/pkg/util"
)

type Logger struct {
	addr string
	out  *bufio.Writer

	dataBuffer int

	file       string
	fileBuffer int

	period time.Duration
}

func New(endpoint, file string, speed, buffSize int32) (*Logger, error) {
	p, err := util.Period(int(speed))
	if err != nil {
		return nil, fmt.Errorf("convert %d speed to period: %v", speed, err)
	}

	return &Logger{
		addr:       endpoint,
		file:       file,
		fileBuffer: int(buffSize),
		dataBuffer: 16,
		period:     p,
	}, nil
}

func (l *Logger) Run() error {
	f, err := os.OpenFile(l.file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("open %s file: %v", l.file, err)
	}
	defer f.Close()

	l.out = bufio.NewWriterSize(f, l.fileBuffer)

	listener, err := net.Listen("tcp", l.addr)
	if err != nil {
		return fmt.Errorf("endpoint %s: %v", l.addr, err)
	}
	log.Info("Listening on", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Errorf("Logger: accepting connection: %v", err)
			continue
		}

		l.handle(conn)
	}
}

func (l *Logger) handle(conn net.Conn) {
	start := time.Now()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Errorf("Failed to close %s connection: %v",
				conn.RemoteAddr(), err)
		}

		log.Infof("Close %s connection, worked %s", conn.RemoteAddr(), time.Now().Sub(start))
	}()

	log.Infof("Accept %s connection", conn.RemoteAddr())
	l.out.WriteString(l.separator(conn.RemoteAddr().String()))

	dataCh := make(chan string, l.dataBuffer)
	go l.numReader(conn, dataCh)

	for {
		d, ok := <-dataCh
		if !ok {
			break
		}

		l.out.WriteString(d)

		// write with the specified speed
		time.Sleep(l.period)
	}

	l.out.Flush()
}

func (l *Logger) numReader(r io.Reader, dataCh chan string) {
	b := bufio.NewReader(r)
	for {
		d, err := b.ReadString(util.Separator)
		if err != nil {
			if err == io.EOF {
				close(dataCh)
				return
			}

			log.Errorf("numReader: read error: %v", err)
			continue
		}

		dataCh <- string(d)
	}
}

func (l *Logger) separator(addr string) string {
	return fmt.Sprintf("\n---Numbers from %s\n", addr)
}
