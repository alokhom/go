package main

import (
	"fmt"
	"net/http"
)

func main() {
	firstName := "Bob"
	familyName := "Smith"
	age := 34
	peanutAllergy := false

	fmt.Println(firstName)
	fmt.Println(familyName)
	fmt.Println(age)
	fmt.Println(peanutAllergy)

	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)

	count := 5
	var message string
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)

	count1 := 0
	if count1 < 5 {
		count1 = 10
		count1++
	}
	fmt.Println(count1 == 11)

	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
