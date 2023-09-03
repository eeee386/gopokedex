package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
	"GoPokedex/commands"
)


func main() {

	var cliMap = commands.GetCLICommands()

	for {
		fmt.Printf("pokedex > ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.Trim(text, " \n")
		if err != nil {
			fmt.Printf("The following error happened:")
			fmt.Printf(err.Error())
			break
		} else {
			command ,e := cliMap[text]
			if e {
				errc := command.Callback()
				if errc != nil {
			    fmt.Printf("The following error happened:")
			    fmt.Println(errc.Error())
				}
				if text == "exit" {
					break
				}
			} else {
			  fmt.Printf("%s is not a valid command\n", text)
			}
		}
	}
}
