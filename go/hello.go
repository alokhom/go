package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
)

const url = "https://services.explorecalifornia.org/json/tours.php"

func main() {

	t := time.Now()
	fmt.Printf("Current time: %s\n", t.Format(time.RFC1123))
	fmt.Printf("Current time: %s\n", t.Format(time.ANSIC))

	tomorrow := t.Add(24 * time.Hour)
	fmt.Printf("Current time: %s\n", tomorrow)

	tFormat := t.Format("Mon 2006-01-02 15:04:05")
	fmt.Printf("Current time: %s\n", tFormat)

	var colours [3]string
	colours[0] = "red"
	colours[1] = "green"
	colours[2] = "blue"
	fmt.Printf("Colours: %v\n", colours)

	var numbers = [3]int{1, 2, 3}
	fmt.Printf("Numbers: %v\n", numbers)

	//slice
	var colours1 = make([]string, 0, 3)
	colours1 = append(colours1, "red")
	colours1 = append(colours1, "green")
	fmt.Printf("Colours1: %v\n", colours1)

	colours1 = remove(colours1, 0)
	fmt.Printf("Colours1: %v\n", colours1)

	sort.Strings(colours1)
	fmt.Println(colours1)

	// map
	states := make(map[string]string)
	states["CA"] = "California"
	states["NY"] = "New York"
	states["TX"] = "Texas"
	fmt.Printf("States: %v\n", states)
	for k, v := range states {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	//slice
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

	//http
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("User-Agent", "")
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()
	fmt.Printf("Response type: %T\n", resp)
	bytes, err := io.ReadAll(resp.Body)
	checkErr(err)
	content := string(bytes)
	fmt.Print(content)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

// func swap(a *int, b *int) {
// 	*a, *b = *b, *a
// }
