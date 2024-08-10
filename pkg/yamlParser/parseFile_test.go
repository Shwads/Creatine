package fileParser

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	requests, _, parseFileErr := ParseFile("tests/testFile2.yml")
	if parseFileErr != nil {
		t.Fatal(parseFileErr)
	}

	fmt.Println("Entered test function")
	if requestMap, ok := requests["request-1"].(map[string]interface{}); ok {
		if method, ok := requestMap["method"]; ok {
			fmt.Printf("method: %s\n", method)
		} else {
			t.Fatal("requests['request-1] did not contain method tag")
		}

		if url, ok := requestMap["url"]; ok {
			fmt.Printf("url: %s\n", url)
		} else {
			t.Fatal("requests['request-1] did not contain url tag")
		}

		if headerMap, ok := requestMap["headers"].(map[string][]string); ok {
			if headerList, ok := headerMap["Authorization"]; ok {
				fmt.Printf("Authorization: %s\n", headerList[0])
			} else {
				t.Fatal("requests['request-1']['headers']['Authorization'] does not exist")
			}

			if headerList, ok := headerMap["Names"]; ok {
				fmt.Printf("Names: %s, %s\n", headerList[0], headerList[1])
			} else {
				t.Fatal("requests['request-1']['headers']['Names'] does not exist")
			}
		} else {
			t.Fatal("requests['request-1']['headers'] did not have expected type: map[string][]string")
		}

		if body, ok := requestMap["body"].(string); ok {
			fmt.Printf("body: %s\n", body)
		} else {
			t.Fatal("requests['request-1']['body'] did not have expected type: string")
		}

	} else {
		t.Fatal("requests['request-1'] did not have type map[string]interface{}")
	}

	if requestMap, ok := requests["request-2"].(map[string]interface{}); ok {
		if method, ok := requestMap["method"]; ok {
			fmt.Printf("method: %s\n", method)
		} else {
			t.Fatal("No tag method in request-2")
		}

		if url, ok := requestMap["url"]; ok {
			fmt.Printf("url: %s\n", url)
		} else {
			t.Fatal("No tag url in request-2")
		}

		if headerMap, ok := requestMap["headers"].(map[string][]string); ok {
			if headerList, ok := headerMap["Authorization"]; ok {
				fmt.Printf("Authorization: %s\n", headerList[0])
			} else {
				t.Fatal("requests['request-2']['headers']['Authorization'] does not exist")
			}

			if headerList, ok := headerMap["Whoble"]; ok {
				fmt.Printf("Whoble: %s, %s\n", headerList[0], headerList[1])
			} else {
				t.Fatal("requests['request-2']['headers']['Whoble'] does not exist")
			}
		} else {
			t.Fatalf("requests['request-2']['headers'] did not have expected type: map[string][]string. Found type %T instead\n", requests["request-1"])
		}

		if body, ok := requestMap["body"].(string); ok {
			fmt.Printf("body: %s\n", body)
		} else {
			t.Fatal("No value requests['request-2']['body']\n")
		}

	} else {
		t.Fatalf("requests['request-2'] did not have type map[string]interface{}. Found type %T instead\n", requests["request-2"])
	}
}
