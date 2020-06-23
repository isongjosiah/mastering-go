package main

import (
	"flag"
	"fmt"
)


minusI := flag.Int("i", 100, "i parameter")
flag.Parse()
i := *minusI
fmt.Println(i)


func main(){
	flag.Int("i", 100, "i parameter")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindFlags(pflag.CommandLine)
	i := viper.GetInt("i")
	fmt.Println(i)
}