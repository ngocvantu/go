package main

import (
	"os/exec"
	"log"
	"fmt"
	"os"
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


	//out, err := exec.Command("/usr/bin/yum", "clean", "all").Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", out)

	cmd := exec.Command("/usr/bin/yum", "update", "" )
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println(err1.Error())
		log.Fatal(err1)
	} 
}
