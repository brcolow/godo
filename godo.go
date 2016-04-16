package main

import "fmt"
import "os"
import "net/http"
import "io/ioutil"
import "time"

func GetExternalIp() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	return string(resp_body)
}

func main() {
	username := os.Args[1]
	password := os.Args[2]

	client := &http.Client{}

	timer := time.NewTicker(time.Second * 10)
	lastIp := ""

	for {
		newIp := GetExternalIp()
		if newIp != lastIp {
			lastIp = newIp
			req, _ := http.NewRequest("GET", "https://domains.google.com/nic/update?hostname=brcolow.com", nil)
			req.SetBasicAuth(username, password)
			fmt.Println(req.URL)

			req.Header.Add("User-Agent", "godo v0.0.1")
			resp, err := client.Do(req)

			if err != nil {
				os.Stderr.WriteString(err.Error() + "\n")
				os.Exit(1)
			}

			defer resp.Body.Close()
			resp_body, _ := ioutil.ReadAll(resp.Body)

			fmt.Println(resp.Status)
			fmt.Println(string(resp_body))
			<-timer.C
		}
	}
}
