package main

import (
	"bytes"
	"fmt"
)

func Iterasi(N int) string {
	// ar := []interface{}
	var hasil bytes.Buffer

	for i := 1; i <= N; i++ {

		if i%3 == 0 && i%5 == 0 {
			hasil.WriteString("FrontEnd BackEnd")
		}else if i%3 == 0 {
			hasil.WriteString("FrontEnd")
		} else if i%5 == 0 {
			hasil.WriteString("BackEnd")
		}else{
			hasil.WriteString(fmt.Sprint(i))
		}

		if i != N {
			hasil.WriteString(",")
		}
	}

	return hasil.String()
}

func main() {
	fmt.Println(Iterasi(50))

}