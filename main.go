package main

import(
	"fmt"
	"bufio"
	"os"
)


func main() {
	for {
		fmt.Printf("pokedex > ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n');
		if(err != nil) {
			fmt.Printf("The following error happened:")
			fmt.Printf(err.Error())
			break
		}else if(text == "exit") {
			break
		}else if(text == "help") {
			fmt.Println()
		} else {
		  fmt.Println(text);
		}
		
	}
}

