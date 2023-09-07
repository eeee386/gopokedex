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
		splitText := strings.Split(text, " ")
		if err != nil {
			fmt.Printf("The following error happened:")
			fmt.Printf(err.Error())
			break
		} else {
			fmt.Println(splitText[0])
			command ,e := cliMap[splitText[0]]
			if e {
				errc := command.Callback(splitText[1:]...)
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
