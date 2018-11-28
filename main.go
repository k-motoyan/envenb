package main

import (
	"bufio"
	"fmt"
	"io"
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

var usage = `USAGE:
	echo FOO=foo\nBAR=bar | envenb > [output].go
	go build main.go [output].go`

func main() {
	ok, err := Usage(usage, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	if ok {
		return
	}

	values, err := ReadFile(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	DumpSource(template, MapKeyValueText(values))
}

func Usage(usage string, f *os.File) (ok bool, err error) {
	stat, err := f.Stat()
	if err != nil {
		return false, err
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println(usage)
		return true, nil
	}

	return false, nil
}

func ReadFile(r io.Reader) ([]string, error) {
	var values []string

	scanner := bufio.NewScanner(r)

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

		values = append(values, text)
	}

	if err := scanner.Err(); err != nil {
		return values, err
	}

	return values, nil
}

func MapKeyValueText(values []string) []string {
	var mappedValues []string

	for _, value := range values {
		value = strings.Replace(value, "=", "\":\"", 1)
		mappedValues = append(mappedValues, "\""+value+"\"")
	}

	return mappedValues
}

func DumpSource(template string, values []string) {
	envValues := strings.Join(values, ",")
	fmt.Print(fmt.Sprintf(template, envValues))
}
