package generator

import (
	"fmt"
	"io"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/saheienko/clab/pkg/util"
)

const (
	ConnTimeout = time.Second * 10
)

type GenNumberFunc func() *big.Int

func Fibonacci() GenNumberFunc {
	var x, y *big.Int
	x, y = big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		defer func() {
			x.Add(x, y)
			x, y = y, x
		}()
		return x
	}
}

type Generator struct {
	period  time.Duration
	out     io.Writer
	genFunc GenNumberFunc
}

func New(endpoint string, freq int, g GenNumberFunc) (*Generator, error) {
	out, err := getWriter(endpoint)
	if err != nil {
		return nil, fmt.Errorf("getWriter: %v", err)
	}

	p, err := util.Period(freq)
	if err != nil {
		return nil, fmt.Errorf("convert %d frequency to period: %v", freq, err)
	}

	return &Generator{
		out:     out,
		period:  p,
		genFunc: g,
	}, nil
}

func (g *Generator) Run() error {
	var err error
	var n *big.Int

	for {
		// generate numbers with the specified speed
		n, err = g.genNumber()
		if err != nil {
			return fmt.Errorf("generate: %v", err)
		}

		// write to logger
		_, err = g.out.Write([]byte(n.String() + string(util.Separator)))
		if err != nil {
			return fmt.Errorf("write %d to writer: %v", n, err)
		}

	}
}

func (g *Generator) genNumber() (*big.Int, error) {
	time.Sleep(g.period)

	n := g.genFunc()
	if n.Cmp(&util.NumberLimit) > 0 {
		return nil, fmt.Errorf("got number of %d bits, limit %d", n.BitLen(), util.NumberLimit.BitLen())
	}

	return n, nil
}

func getWriter(endpoint string) (io.Writer, error) {
	if endpoint == "" {
		return os.Stdout, nil
	}

	conn, err := net.DialTimeout("tcp", endpoint, ConnTimeout)
	if err != nil {
		return nil, fmt.Errorf("endpoint %s: %v", endpoint, err)
	}

	return conn, nil
}
