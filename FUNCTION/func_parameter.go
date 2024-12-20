// package main

// import "fmt"

// func sayHelloTo(firstName string, lastName string){
// 	//nama parameter tdk boleh sama
// 	fmt.Println("Hello", firstName, lastName)
// }
// func main(){
// 	sayHelloTo("Roma", "Yana")
// 	sayHelloTo("sim", "jaehyun")
// }
package main

import "fmt"
import "strings"

func filter(data []string, callback func(string)bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
func main(){
	var data = []string{"wick", "jason", "ethan"}
	var dataContains0 = filter(data, func(each string)bool {
		return strings.Contains(each, "o")
	})
	var dataLenght5 = filter(data, func(each string)bool {
		return len(each) == 5
	})
	fmt.Println("data asli \t\t:", data)
	fmt.Println("filter ada huruf \"i\"\t:", dataContains0)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
}