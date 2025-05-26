package main
import "fmt"
func main() {
	country := map[string]string{"TH": "Thailand", "JP": "Japan", "EN": "England"}
	fmt.Println("The country is", country["JP"])

	Province := map[string]string{}
	Province["ANK"] = "Akihabara" //key
	Province["TKY"] = "Tokyo"
	Province["BKK"] = "Bangkok"

	value, check := Province["ANK"] // check if the key exists
	if check {
		fmt.Println("The Province is", value)
	} else {
		fmt.Println("The key does not exist")
	}

	
}
