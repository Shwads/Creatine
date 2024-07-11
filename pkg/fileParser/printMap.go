package fileParser

import "fmt"

var multiRequest bool = false

func printMap(requests map[string]interface{}, indent int) {
	line := ""

	for x := 0; x < indent; x++ {
		line = fmt.Sprintf(" %s", line)
	}

	for key, _ := range requests {
		fmt.Printf("%s%s: \n", line, key)

		if tagString, ok := requests[key].(string); ok {
			fmt.Printf("%s %s\n",line, tagString)
		} else if tagList, ok := requests[key].([]string); ok {
			for _, val := range tagList {
				fmt.Printf("%s %s", line, val)
			}
		} else if requestMap, ok := requests[key].(map[string][]string); ok {
			for key, list := range requestMap {
				fmt.Printf("%s %s: \n", line, key)

				for _, item := range list {
					fmt.Printf("%s %s\n", line, item)
				}
			}
		} else if requestMap, ok := requests[key].(map[string]interface{}); ok {
			printMap(requestMap, indent+1)
		}
		fmt.Print("\n")

	}
}
