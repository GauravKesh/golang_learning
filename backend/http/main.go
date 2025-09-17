package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func apiControllerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is working  and you've requested %s\n", r.URL.Path)

}

func getNoOfUser(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"name": "gaurav",
		"age":  "20",
	}
	// Also, data is a map, so you can’t directly print it without formatting.
	fmt.Printf("%s", r.URL.Host)
	// fmt.Fprintf(w, "%v", data)

	/* to set content-type header  */
	w.Header().Set("content/type", "application/json")

	/* encoding the data into json*/
	json.NewEncoder(w).Encode(data)
}

func main() {

	fmt.Println("/////////////Starting your server wait for a while\\\\\\\\\\\\\\")
	/* function to handle new request */
	http.HandleFunc("/api", apiControllerTest)
	http.HandleFunc("/api/user/count", getNoOfUser)

	//  running the server at certain port address
	http.ListenAndServe(":8080", nil)

}
