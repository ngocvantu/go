package main

import (
	"os/exec"
	"log"
	"fmt"
)

func main(){
	//con,_ := net.Dial("tcp", "45.32.118.97:8070")
	//defer con.Close()
	//io.Copy(os.Stdout, con)
	//res, _ := http.Get("http://dantri.com.vn/")
	//defer res.Body.Close()
	//
	//doc, _ := html.Parse(res.Body)
	//fmt.Println(string(doc.))

	out, err := exec.Command("/usr/bin/yum", "clean", "all").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
	out1, err1 := exec.Command("/usr/bin/yum", "update" ).Output()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("%s\n", out1)
}
