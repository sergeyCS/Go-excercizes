package main

import "fmt"

func main() {
	ages := map[string]int{"qwer": 12}
	//ages = make(map[string]int)

	ages["asdf"] = 11
	fmt.Println(ages)
}
