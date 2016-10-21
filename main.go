package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var yamlSeperator = []byte("---\n# Source: ")

func scanYamlSpecs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, yamlSeperator); i >= 0 {
		// We have a full newline-terminated line.
		return i + len(yamlSeperator), data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func splitSpec(token string) (string, string) {
	if i := strings.Index(token, "\n"); i >= 0 {
		return token[0:i], token[i+1:]
	}
	return "", ""
}

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s OUTPUT_DIRECTORY", os.Args[0])
	}

	output_directory := os.Args[1]

	if _, err := os.Stat(output_directory); !os.IsNotExist(err) {
		log.Fatalf(`Output directory "%v" already exists`, output_directory)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(scanYamlSpecs)
	//Allow for tokens (specs) up to 1M when parsing the manifest
	scanner.Buffer(make([]byte, bufio.MaxScanTokenSize), 1048576)
	//Discard the first result, we only care about everything after the first seperator
	scanner.Scan()

	for scanner.Scan() {
		source, content := splitSpec(scanner.Text())
		destinationFile := path.Join(output_directory, source)
		dir := path.Dir(destinationFile)
		if err := os.MkdirAll(dir, 0750); err != nil {
			log.Fatalf("Error creating %s: %s ", dir, err)
		}
		log.Printf("Writing %s", destinationFile)
		if err := ioutil.WriteFile(destinationFile, []byte(content), 0640); err != nil {
			log.Fatalf("Error: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while processing manifest: %s", err)
	}

}
