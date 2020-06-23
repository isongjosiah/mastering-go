package main

import "fmt"

func main() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./myJSONConfig.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("item1.key1") {

	}
}
