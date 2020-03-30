/* package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World2")
	time.Sleep(1 * time.Second)
}

func coolFunc(a func()) {
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
	coolFunc(myFunc)
}
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Checking to see if Authorized header set...")

		if val, ok := r.Header["Authorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set! We can serve content!")
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {

	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
