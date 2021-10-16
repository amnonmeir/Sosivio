package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func encrypt(split []string, number_of_threads int) string { //Change sting array to sting array encrypted in SHA256
	//NOTE only COUNTER times of threads can run at the same time
	var wg sync.WaitGroup
	var guard = make(chan int, number_of_threads) //This channel let us control the number of threads we want to open as max
	for i := range split {
		wg.Add(1)
		guard <- 1
		go func(i int) {
			enc := sha256.New()
			enc.Write([]byte(split[i]))
			result := enc.Sum(nil)
			split[i] = string(result)
			<-guard
			wg.Done()
		}(i)
	}
	wg.Wait()    // Wait for the threads to finish
	close(guard) // This tells the goroutines there's nothing else to do
	result := ""
	for j := 0; j < len(split); j++ {
		result = result + split[j]
	}
	return base64.StdEncoding.EncodeToString(([]byte(result)))
}
func main() {
	var local_port string

	local_port = os.Getenv("BACKEND_PORT")
	if local_port == "" {
		fmt.Println("NO PORT DEFINED")
		os.Exit(2)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}
		switch r.Method {
		case "GET":
			var url string = r.URL.String()
			if len(url) < 2 {
				return
			}
			str := url[2:]
			split := strings.Split(str, ",")
			number_of_threads, err := strconv.Atoi(os.Getenv("THREADS"))
			if err != nil {
				// handle error
				number_of_threads = 4
			}
			encrypt := encrypt(split, number_of_threads) //send the splited string to encryption
			fmt.Fprintf(w, encrypt)                      //return the encrypted string to the user
		}
	})
	http.ListenAndServe(":"+local_port, nil)

}

