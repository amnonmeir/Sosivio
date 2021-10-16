package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string { //this func handle the random string with the length 'n'
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func main() {

	var local_port, server_port string
	var strings_length int

	local_port = os.Getenv("FRONTEND_PORT")
	server_port = os.Getenv("BACKEND_PORT")

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
			number_of_strings, err := strconv.Atoi(url[2:])
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			str := ""
			for j := 0; j < number_of_strings-1; j++ {
				str = str + randSeq(strings_length) + ","
			}
                        strings_length, err := strconv.Atoi(os.Getenv("STRING_LENGTH"))
			str = str + randSeq(strings_length) //it insert the last string without the ','
			resp, err := http.Get("http://localhost:" + server_port + "/?" + str) //Sent to encrypt
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			fmt.Println("Response status:", resp.Status)
			scanner := bufio.NewScanner(resp.Body)
			scanner.Scan()
			encrypt := scanner.Text() //insrt the respons to encrypt var
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			fmt.Fprintf(w, encrypt)
			fmt.Println(encrypt)
		}
	})
	http.ListenAndServe(":" + local_port, nil)
}
