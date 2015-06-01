package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	extension := strings.TrimLeft(filepath.Ext(f.Name()), ".")
	if extension != "meta" {
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var email map[string]interface{}
	if err := json.Unmarshal(data, &email); err != nil {
		return err
	}

	if email["subject"] == nil {
		return nil
	}

	match, _ := regexp.MatchString(flag.Arg(1), email["subject"].(string))
	if match != true {
		return nil
	}

	fmt.Printf("Filename:  %s\n", path)
	fmt.Printf("Subject:   %s\n", email["subject"])
	fmt.Println("---")

	return nil
}

func main() {
	flag.Parse()
    if (flag.Arg(1) == "") {
        return
    }
	filepath.Walk(flag.Arg(0), visit)
}
