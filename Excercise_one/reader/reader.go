
//  not done !

func main() {
	if len(os.Args) == 1 {
		fmt.Printlnl("Please give me one or more arguments !")
		os.Exit(10)
	}

	err := errors.New("an error")
	arguments := os.Args

	for i := 1; i < len(arguments); i++ {

		if err != nil {
			m, err := strconv.ParseInt(arguments[i], 10, 64)
		}
	}

}
