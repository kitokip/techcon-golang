package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fibonanceHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	n := params.Get("n")
	nInt, _ := strconv.ParseInt(n, 0, 64)
	result := fibonance(nInt)

	fmt.Fprintf(w, "%v", result)
}

func primethHandler(w http.ResponseWriter, r *http.Request) {
	//https://primes.utm.edu/lists/small/1000.txt
	params := r.URL.Query()
	n := params.Get("n")
	nInt, _ := strconv.ParseInt(n, 0, 64)
	result := getNthPrime(nInt)

	fmt.Fprintf(w, "%d", result)
}

func main() {
	httpPort := 8081

	http.HandleFunc("/fibonance", fibonanceHandler)
	http.HandleFunc("/nthprime", primethHandler)

	fmt.Printf("listening on %v\n", httpPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}

// Calculate Fibonance
func fibonance(n int64) int64 {
	var f1, f2, nth int64 = 1, 1, 0
	if n == 1 || n == 2 {
		return 1
	}
	for i := int64(3); i <= n; i++ {
		nth = f1 + f2
		f1 = f2
		f2 = nth
	}

	return nth
}

// prime check
func isPrime(number int64) bool {

	if number < int64(2) {
		return false
	}

	for i := int64(2); i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func getNthPrime(nth int64) int64 {
	if nth < 1 {
		return -1
	}

	var nthPrime, prime int64 = 0, -1
	for i := int64(2); nthPrime < nth; i++ {
		if isPrime(i) {
			prime = i
			nthPrime++
		}
	}

	return prime
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
