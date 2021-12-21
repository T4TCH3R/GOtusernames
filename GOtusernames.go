package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	//"log"
)

func createUser(file string, output string) {

	read, err := os.Open(file)
	if err != nil {
		fmt.Println("Sorry couldnt open given file")
	} else {

		if output != "" {
			scanner := bufio.NewScanner(read)

			outputfile, err := os.Create(output)
			if err != nil {
				fmt.Println("File cannot be created")
				os.Exit(1)
			}
			defer outputfile.Close()

			for scanner.Scan() {
				line := scanner.Text()
				first_last := strings.Split(line, " ")

				first_name := first_last[0]
				last_name := first_last[1]

				fmt.Fprintf(outputfile, "%s.%s\n", first_name[0:1], last_name)
				fmt.Fprintf(outputfile, "%s"+"%s\n", first_name, last_name)
				fmt.Fprintf(outputfile, "%s"+"%s\n", first_name[0:1], last_name)
				fmt.Fprintf(outputfile, "%s.%s\n", first_name, last_name)
				fmt.Fprintf(outputfile, "%s.%s\n", first_name, last_name[0:1])
				fmt.Fprintf(outputfile, "%s\n", first_name)
				fmt.Fprintf(outputfile, "%s\n", last_name)
			}
		} else {
			scanner := bufio.NewScanner(read)

			for scanner.Scan() {
				line := scanner.Text()
				first_last := strings.Split(line, " ")

				first_name := first_last[0]
				last_name := first_last[1]

				fmt.Printf("%s.%s\n", first_name[0:1], last_name)
				fmt.Println(first_name + last_name)
				fmt.Println(first_name[0:1] + last_name)
				fmt.Printf("%s.%s\n", first_name, last_name)
				fmt.Printf("%s.%s\n", first_name, last_name[0:1])
				fmt.Println(first_name)
				fmt.Println(last_name)
			}
		}

	}
}

func main() {
	var f string
	var o string
	flag.StringVar(&f, "f", "", "Please select the file with first and lastnames like: John Doe")
	flag.StringVar(&o, "o", "", "Output the new domain usernames to a file.")
	flag.Parse()

	userlist := f
	outputfile := o

	if userlist != "" {
		createUser(userlist, outputfile)
	} else {
		flag.Args()
	}
}
