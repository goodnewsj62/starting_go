package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("type","pe"))
	fmt.Println(strings.HasPrefix("type","t"))
	fmt.Println(strings.HasSuffix("type","e"))
	fmt.Println(strings.Count("libido","i"))
	fmt.Println(strings.Title("hey you"))
	fmt.Println(strings.Join([]string{"hey","yo","there"}, ""))
	fmt.Println(strings.Index("whatever string defined here","o"))// -1 not found
	fmt.Println(strings.Replace("some text of choice","e","f",3))


	x :=  []byte("abcd")
	y := string(x)

	fmt.Println("\n\n", x , y)

	var buf bytes.Buffer
	buf.Write([]byte("apple"))

	fmt.Println(buf)

	// os 

	file, err := os.Open("test.txt")
	if err != nil{
		return
	}

	defer file.Close()

	stat,err := file.Stat()
	if err != nil{
		return
	}

	
	bs := make([]byte, stat.Size())

	_ , err = file.Read(bs)

	if err != nil{
		return
	}


	fmt.Println(string(bs))
}

