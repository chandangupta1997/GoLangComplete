package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// var wg sync.WaitGroup // it is a pointer usually
// // now wg has functionality of wait group

// func main() {
// 	// go greeter("hello")
// 	// greeter("world")

// 	websiteList := []string{
// 		//"https.youtube.com",
// 		//"https.facebook.com",
// 		"https.github.com",
// 		"https.lco.dev",
// 	}

// 	for _, website := range websiteList {
// 		fmt.Println(website)
// 		GetStatusCode(website)

// 		wg.Add(1)

// 	}

// 	wg.Wait()
// }

// // func greeter(s string) {

// // 	// whcih can print 5 times

// // 	for i := 0; i < 6; i++ {

// // 		time.Sleep(1 * time.Second)
// // 		fmt.Println(s)

// // 	}

// // }

// func GetStatusCode(endPoint string) {

// 	defer wg.Done()
// 	result, err := http.Get(endPoint)

// 	if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Printf("%d statuse code for %s\n", result.StatusCode, endPoint)
// 	}

// }

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer

var mut sync.Mutex // usually a pointer

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("world")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {

		signals = append(signals, endpoint)

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}
