package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	kvdb "github.com/kvdb/src"
)

func exector(option string, args string) bool {

	return false
}

func main() {
	var queue []string
	inpReader := bufio.NewReader(os.Stdin)

	fmt.Println("Type `exit` to quit.")
	for {
		input, err := inpReader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if strings.Trim(input, "\n") == "exit" {
			break
		}

		args := strings.Split(strings.Trim(input, "\n"), " ")

		switch strings.Trim(args[0], " ") {

		case "SET":
			res := kvdb.Set(args[1], args[2])
			if res != true {
				fmt.Println("Error: Unable to add key to kvdb")
			}

		case "GET":
			value, found := kvdb.Get(args[1])
			if found == true {
				fmt.Println(value)
			} else {
				fmt.Println("Key not found!")
			}

		case "DEL":
			res := kvdb.Del(args[1])
			if res == false {
				fmt.Println("Key not found!")
			}

		case "INCR":
			res := kvdb.Incr(args[1])
			if res != true {
				fmt.Println("Error: Unable to increment value")
			}

		case "INCRBY":
			if len(args) < 3 {
				fmt.Println("Too less arguments!")
			} else {
				incrValue, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

				res := kvdb.IncrBy(args[1], incrValue)
				if res != true {
					fmt.Println("Error: Unable to increment value")
				}
			}

		case "COMPACT":
			db := kvdb.Compact()
			for k, v := range db {
				fmt.Println("SET", k, v)
			}

		case "MULTI":
			for {
				input, _ := inpReader.ReadString('\n')
				queue = append(queue, input)
				if strings.Trim(input, "\n") == "DISCARD" {
					queue = nil
					break
				}

				if strings.Trim(input, "\n") == "EXEC" {
					fmt.Println(kvdb.Multi(queue[:len(queue)-1]))
					break
				}
			}

		default:
			fmt.Println("Invalid argument")
		}
	}

}
