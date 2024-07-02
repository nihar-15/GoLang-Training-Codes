package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string = "Niharika"
	var subStr string = "niharika"

	if strings.Contains(str, subStr) == true {
		fmt.Printf("String %s contains the substring %s ", str, subStr)
	} else {
		fmt.Println("Does not contains")
	}

	println(strings.ToUpper(subStr))
	println(strings.ToLower(subStr))

	// Also Index()
	index := strings.Count(str, "i")
	println(index)

	var str1 string = "My-Name-Is-Niharika"
	var strArr []string = strings.Split(str1, "-")

	for i := 0; i < len(strArr); i++ {
		println(strArr[i])
	}
	str2 := "Niharika"

	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i])
	}
	println("\n")
	var str3 string = "india is a great land"
	var r1 string = strings.Replace(str3, "i", "I", -1)
	var r2 string = strings.Replace(str3, "i", "I", 1)
	println(r1)
	println(r2)

	result := strings.Join(strArr, "_")
	println(result)

	result3 := strings.HasPrefix(str, "Nih")

	if result3 {
		println("Yes")
	} else {
		print("No")
	}

}
