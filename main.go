package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var template = `package main

import (
	"log"
	"os"
)

var _ = func() interface{} {
	var envValues = map[string]string{%v}

	for key, value := range envValues {
		if err := os.Setenv(key, value); err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}()
`

func main() {
	values, err := readFile(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	envValues := strings.Join(values, ",")
	fmt.Print(fmt.Sprintf(template, envValues))
}

func readFile(f *os.File) ([]string, error) {
	var values []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()

		// Skip empty line.
		if len(text) == 0 {
			continue
		}

		// Skip comment line.
		if strings.HasPrefix(text, "#") {
			continue
		}

		text = strings.Replace(text, "=", "\":\"", 1)
		values = append(values, "\""+text+"\"")
	}

	if err := scanner.Err(); err != nil {
		return values, err
	}

	return values, nil
}
