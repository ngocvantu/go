package main

import (
	"os"
	"io"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"fmt"
	"bufio"
)

func main() {
	file, _ := os.Open("D:\\2017B\\WMS\\trunk\\src\\slimks-final\\WebContent\\WEB-INF\\kousu_sql.xml")
	defer file.Close()
	var r io.Reader

	r = transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
	 //butes, _ := ioutil.ReadAll(r)

	 in := bufio.NewScanner(r)
	for in.Scan()  {
		fmt.Println(in.Text())
	}

	 //fmt.Print(string(butes))

}
