package readdata

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id     int
	x      float64
	y      float64
	chosen bool
}

type Nodes []Node

func ReadData(filename string) Nodes {
	data := Nodes{}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error")
	}

	// when reading (mission of the function) finished, close the file:
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	bf := bufio.NewReader(f)
	//n := 0

	for {
		l, _, err := bf.ReadLine()
		if err != nil {
			return nil
		}
		line := string(l)
		if strings.HasPrefix(line, "NODE") {
			for {
				l, _, err := bf.ReadLine()
				if err != nil {
					if err != io.EOF {
						return nil
					}
					break
				}
				line := string(l)

				if !strings.HasPrefix(line, "EOF") {

					fb := strings.Index(line, " ")
					nid := string(line[:fb])
					sb := strings.Index(line[fb+1:], " ")
					x := line[fb+1 : fb+sb+1]
					y := line[fb+sb+2:]

					var n Node
					n.id, _ = strconv.Atoi(nid)
					n.x, _ = strconv.ParseFloat(x, 64)
					n.y, _ = strconv.ParseFloat(y, 64)

					data = append(data, n)

				}

				if strings.HasPrefix(line, "EOF") {
					return data
				}

			}
		}

	}

}
