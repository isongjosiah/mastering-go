package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByline(filename string) error {
	var err error
	var output []string

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
			fmt.Printf("error reading file %s", err)
			break
		}
		output = append(output, line)
	}
	fmt.Println(output)
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, filename := range flag.Args() {
		err := lineByline(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}
