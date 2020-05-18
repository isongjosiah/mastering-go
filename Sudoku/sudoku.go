package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// this function takes in a file and creates a sudoku matrix.
func importFile(file string) ([][]int, error) {
	var err error
	var mySlice = make([][]int, 0)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		fields := strings.Fields(line)
		temp := make([]int, 0)
		for _, v := range fields {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			temp = append(temp, n)
		}
		if len(temp) != 0 {
			mySlice = append(mySlice, temp)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(temp) != len(mySlice[0]) {
			return nil, errors.New("Wrong number of elements")
		}
	}

	return mySlice, nil
}
