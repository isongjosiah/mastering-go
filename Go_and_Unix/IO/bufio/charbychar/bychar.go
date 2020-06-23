package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charBychar(filename string) error {
	var err error
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("cannot read the file, encounter err %s", err)
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}
