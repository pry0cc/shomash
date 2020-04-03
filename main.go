package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	 "os/user"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// QueryShodan queries a given IP and returns Shodan JSON
func QueryShodan(ip string, key string) string {
	keyFragment := "?key=" + key

	resp, err := http.Get("https://api.shodan.io/shodan/host/" + ip + keyFragment)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	stringBody := string(body)
	if strings.Contains(stringBody, "limit reached") {
		fmt.Println("We got limited dude: " + key)
		time.Sleep(2 * time.Second)
		return QueryShodan(ip, key)
	}

	return stringBody
}

func main() {
	usr, _ := user.Current()
	path := usr.HomeDir + "/.shomash"
	keys, _ := readLines(path)

	ips := make(chan string, 10)
	output := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < len(keys); i++ {
		wg.Add(1)

		go func(x int) {
			for ip := range ips {
				// fmt.Println("Querying with " + keys[x])
				resp := QueryShodan(ip, keys[x])
				output <- resp

				time.Sleep(800 * time.Millisecond)
			}

			wg.Done()
		}(i)
	}

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			ips <- sc.Text()
		}
		close(ips)
	}()
	go func() {
		wg.Wait()
		close(output)
	}()
	for item := range output {
		var result map[string]interface{}
		json.Unmarshal([]byte(item), &result)
		if result["ip_str"] != nil {
			// fmt.Println("Successfully scanned " + result["ip_str"].(string))
			fmt.Println(item)
			// path := "./output/" + result["ip_str"].(string) + ".json"
			// ioutil.WriteFile(path, []byte(item), 0644)
		}
	}

}
