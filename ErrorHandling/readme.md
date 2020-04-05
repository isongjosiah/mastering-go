Error handling and printing to error output are different,

Error handling has to do with Go code that handles error conditions
the latter has to do with writing something to the standard error file descriptor

===========================================================
The following code prints error to the screen(terminal)
if err != nil {
    fmt.Println(err)
    os.Exit(10) 
}
note that you should use return to exit your program inside a function, calling os.Exit in a function other than 
the main function is considered bad practice.

The following code sends the error to the logging service instead.

if err != nil {
    log.Println(err)
    os.Exit
}

Another variation of the preceeding code is used when something really bad has happend and termination of the program is necessary

if err != nil {
    panic(err)
    os.Exit(10)
}