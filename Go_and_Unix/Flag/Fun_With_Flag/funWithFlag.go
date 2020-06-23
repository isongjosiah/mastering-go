package main

import (
	"flag"
	"fmt"
	"strings"
)

// NamesFlag contains the name of the flags used
type NamesFlag struct {
	Names []string
}

// String returns the Flag names as a string
func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

// GetNames returns the list of Flags
func (s *NamesFlag) GetNames() []string {
	return s.Names
}

// Set turns a comma-seperated list into a slice of string
func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	flag.Var(&manyNames, "names", "Comma-seperated list")
	flag.Parse()
	fmt.Println("-k", *minusK)
	fmt.Println("-o", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
