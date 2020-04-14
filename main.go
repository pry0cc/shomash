package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// QueryShodan queries a given IP and returns Shodan JSON
func QueryShodan(ips string, key string) string {
	keyFragment := "?key=" + key

	resp, err := http.Get("https://api.shodan.io/shodan/host/" + ips + keyFragment)

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
		return QueryShodan(ips, key)
	}

	return stringBody
}

func main() {
	keys := []string{"fd12FQLsxxz7Rebw3XY8KHG1IywQsvyI"}
	var domain = flag.String("d", "example.com", "The domain (used for storage in output/$domain)")
	flag.Parse()
	outdir := filepath.Join(".", "output", *domain)
	os.MkdirAll(outdir, os.ModePerm)

	ips := make(chan string, 45)
	output := make(chan string)
	var wg sync.WaitGroup

	bulkAmount := 100
	stop := false

	for i := 0; i < len(keys); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			for !stop {
				bulkList := []string{}
				for x := 0; x < bulkAmount; x++ {
					ip, ok := <-ips
					if !ok {
						stop = true
						break
					}
					bulkList = append(bulkList, ip)
				}
				queryString := strings.Join(bulkList, ",")
				output <- QueryShodan(queryString, keys[i])
				time.Sleep(time.Second * 1)
			}
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
		var results []map[string]interface{}
		json.Unmarshal([]byte(item), &results)
		if len(results) == 0 {
			var result map[string]interface{}
			json.Unmarshal([]byte(item), &result)
			results = append(results, result)
		}
		for _, result := range results {
			if result["ip_str"] != nil {
				fmt.Println("Successfully scanned " + result["ip_str"].(string))
				jsonStr, _ := json.Marshal(result)
				path := outdir + "/" + result["ip_str"].(string) + ".json"
				ioutil.WriteFile(path, []byte(string(jsonStr)), 0644)
			}
		}

	}

}
