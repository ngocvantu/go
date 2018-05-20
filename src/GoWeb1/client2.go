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
	cmds := []*exec.Cmd{}
	cmds = append(cmds,exec.Command("/usr/bin/yum", "clean", "all" ))
	cmds = append(cmds,exec.Command("/usr/bin/yum", "update", "" ))
	executeCommand(cmds)
}

func executeCommand(cmds []*exec.Cmd){
	for _, cmd := range cmds {
		fmt.Println("Executing:",cmd.Args, "...")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		err1 := cmd.Run()
		if err1 != nil {
			log.Fatal(err1)
			return
		}
		fmt.Println("End success:", cmd.Args)
	}
}
