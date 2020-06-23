package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	var configFile *string = flag.String("c", "myconfig", "Setting the configuration file")
	flag.Parse()

	_, err := os.Stat(*configFile)

	if err == nil {
		fmt.Println("Using User Specified Configuration file!")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	if viper.IsSet("item1.k1") {
		fmt.Println("item1.val1:", viper.Get("item1.k1"))
	} else {
		fmt.Println("item1.k1 not set!")
	}

	if viper.IsSet("item1.k2") {
		fmt.Println("item1.val2:", viper.Get("item.k2"))
	} else {
		fmt.Println("item1.k2 not set!")
	}

	if !viper.IsSet("Item3.k1") {
		fmt.Println("item3.k1 is not set!")
	}
}
