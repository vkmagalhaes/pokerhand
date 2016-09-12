package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vkmagalhaes/pokerhand/evaluator"
)

const pokerHandGreetings = `
              __
        _..-''--'----_.
      ,''.-''| .---/ _` + "`" + `-._
    ,' \ \  ;| | ,/ / ` + "`" + `-._` + "`" + `-.
  ,' ,',\ \( | |// /,-._  / /
  ;.` + "`" + `. ` + "`" + `,\ \` + "`" + `| |/ / |   )/ /
 / /` + "`" + `_` + "`" + `.\_\ \| /_.-.'-''/ /
/ /_|_:.` + "`" + `. \ |;'` + "`" + `..')  / /
` + "`" + `-._` + "`" + `-._` + "`" + `.` + "`" + `.;` + "`" + `.\  ,'  / /
    ` + "`" + `-._` + "`" + `.` + "`" + `/    ,'-._/ /
      : ` + "`" + `-/     \` + "`" + `-.._/
      |  :      ;._ (
      :  |      \  ` + "`" + ` \
       \         \   |
        :        :   ;
        |           /
        ;         ,'
       /         /
      /         /
               / Welcome To Poker Hand Showdown!
`

func main() {
	fmt.Println(pokerHandGreetings)

	fmt.Println("Please, type the number of players:")
	var nPlayers int
	for nPlayers <= 0 {
		_, err := fmt.Scanf("%d\n", &nPlayers)
		if err != nil {
			fmt.Println("Oops, I wasn't able to parse that number. Please, try again")
			continue
		}

		if nPlayers < 0 {
			fmt.Println("Please, type a number greater that 0")
		}
	}

	e := evaluator.Evaluator{}
	in := bufio.NewReader(os.Stdin) // to read whole lines

	fmt.Printf("Now type the name and the hands of the %d players:\n", nPlayers)
	for i := 0; i < nPlayers; i++ {
		for {
			var line string
			line, err := in.ReadString('\n')
			if err != nil {
				fmt.Println("Oops, I wasn't able to parse that hand. Please, try again")
				continue
			}

			args := strings.Split(line, ",")
			err = e.ParseHand(args[0], args[1:])
			if err != nil {
				fmt.Println(err)
				continue
			}

			// hand parsed and added succesfully
			break
		}
	}

	e.Decide()
	fmt.Println(e.Resolution())
}
