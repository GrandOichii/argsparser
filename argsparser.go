package argsparser

import (
	"fmt"
	"log"
)

func init() {
	log.Print("argsparser - initializing package")
}

func Parse(args []string, correlationMap map[string]string) (map[string]string, error) {
	// example:
	/*
		map[string]string{
			"--out": "out path",
			"--outpath": "out path",
			"--set": "set name",
			"--setname": "set name"
		}
	*/
	// initialize the result
	result := map[string]string{}
	for _, value := range correlationMap {
		result[value] = ""
	}
	for i, arg := range args {
		// skip the last element
		if i == len(args)-1 {
			break
		}
		for key, value := range correlationMap {
			if arg == key {
				// check if the value is already set
				if result[value] != "" {
					return nil, fmt.Errorf("%v is already set", value)
				}
				// set the value
				result[value] = args[i+1]
			}
		}
	}
	return result, nil
}
