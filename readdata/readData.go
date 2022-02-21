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
	id      int
	x_coord float64
	y_coord float64
	choosen bool
}

type Nodes []Node

func Read_data(filename string) Nodes {
	data := Nodes{}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error")
		os.Exit(1)
	}

	// when reading (mission of the function) finished, close the file:
	defer f.Close()

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
					x_c := line[fb+1 : fb+sb+1]
					y_c := line[fb+sb+2:]

					var n Node
					n.id, _ = strconv.Atoi(nid)
					n.x_coord, _ = strconv.ParseFloat(x_c, 64)
					n.y_coord, _ = strconv.ParseFloat(y_c, 64)

					data = append(data, n)

				}

				if strings.HasPrefix(line, "EOF") {
					return data
				}

			}
		}

	}

}
