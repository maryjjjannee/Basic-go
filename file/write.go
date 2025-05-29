package main

import "os"

func check(err error) {
	if err != nil {
		panic(err) // panic will stop the program and print the error
	}
}

func main() {
	data1 := []byte("Hi\n Jane")
	err := os.WriteFile("data.txt", data1, 0644) // 0644 is the permission for the file
	check(err) // check if there is an error
	f, err := os.Create("employeeName")
	check(err) // check if there is an error
	defer f.Close() // defer will close the file at the end of the function

	data2 := []byte("Onw\nSak")
	os.WriteFile("employeeName", data2, 0644) // 0644

}