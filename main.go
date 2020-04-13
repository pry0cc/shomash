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
	// keys := []string{"R6D6qC9QE0DyFkMW3CeNCurq6IVEYCat", "fd12FQLsxxz7Rebw3XY8KHG1IywQsvyI", "MM72AkzHXdHpC8iP65VVEEVrJjp7zkgd", "VTCNUPjwvgQtAjdUAiwatg00qGeJkuXv", "SLSmbzwIS2oQbAMevr9aXZ6VjPXPozXN", "YbI0bF8iqUJO6Zc58trmFYF0ZUEhdRjK", "EXdwhpQ8y0di0jVGUAx8P9ymtZ5xl0Rg", "j615x8hchMZ8cNVgurPAwePORJz7uVV2", "NaskNg5YlPKZutlMBS7gS0s4nj113um0", "16Za8QtX0kdg4zCcdgGCWp1HoOVVkDyU", "CUB0v414pREWOnciSC6rnSg5qiQP5Ar4", "WqjWlUYXwGlLlhAlyh9JweM9MbLxBN9b", "ZfdCSbmiFXS3yUScyK0Qwn04i4Ns79d7", "m6iE3U8JzekQlteJ9aagLZzop5o6GlyH", "LvVwc2HFpBz80ujOqYQUfyBZliguBxih", "XyvQ904Ts1cL7V6Vt720VwcmazoTfhNT", "qgz3HRM2F2ceDmMzca4l6MtJ2wsNmjMt", "IHWttHEPp25zpNNYfXPbhliqcT8BVcvQ", "eMunoAodm8qSabxv2J6c8KzqYw11OSe7", "9sjTu3AuGdUesGluaQyy48LhhZuiS3Mt", "FRz6i2jqMEy8AL5eOZiJmDk8UkghWmx0", "UFh6iXrpV7fnZejI3wBsxTrZ5yNr1fxw", "Djls4J23qaUnKVrm5KGAx9exhTvSuxqo", "xZGEZR9KYaHV3YDZqJyioNjXpLB0mQ0H", "Kydu5HdKI7vz2dSIQbMTYqHZkYwQzw4D"}
	keys := []string{"fd12FQLsxxz7Rebw3XY8KHG1IywQsvyI"}
	// domain := "example.com"
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
			fmt.Println("Successfully scanned " + result["ip_str"].(string))
			jsonStr, _ := json.Marshal(result)
			path := outdir + "/" + result["ip_str"].(string) + ".json"
			ioutil.WriteFile(path, []byte(string(jsonStr)), 0644)
		}

	}

}
