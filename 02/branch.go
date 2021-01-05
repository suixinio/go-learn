package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func test_if() {
	const filename = "./02/abc.txt"

	//contents ,err := ioutil.ReadFile(filename)
	//if err!=nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s\n",contents)
	//}
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s \n ", contents)
	}
}
func test_switch() {
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		//grade(101),
	)
}
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	default:
		panic(fmt.Sprintf("Wrong score : %d", score))
	}
	return g
}
func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for {

	}
}
func test_for() {
	fmt.Println(converToBin(5), converToBin(13))
	printfile("./02/abc.txt")
}
func main() {
	test_if()
	test_switch()
	test_for()
}
