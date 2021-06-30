//WaitGroups
//每次goroutines還沒執行完.main就跳出了
//雖然可以用
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://cn.bing.com",
	"https://www.baidu.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	//通知WaitGroup執行完成了
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

//fetch url
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")

	var wg sync.WaitGroup

	//迴圈抓取url,多一筆就加一的WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	//等待WaitGroup數量執行結束
	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "Responses")
}

//註冊路由
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}

/* package main

import (
	"fmt"
	"log"
	"net/http"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")
	for _, url := range urls {
		go fetch(url)
	}
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "All Responses Received")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
*/
