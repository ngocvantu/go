package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"io/ioutil"
)

var (
	fileInfo os.FileInfo
	fileInfoErr error
)

func main() {
	start := time.Now()
	// file info
	fileInfo, fileInfoErr = os.Stat("E:\\Lap Trinh\\Udemy\\Mastering Go Programming\\Watching Video.txt")

	if fileInfoErr != nil {
		log.Fatal(fileInfoErr)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

	// open a file with the default associated program
	//open.Start("http://dantri.com.vn/")
	//open.Start("https://google.com")

	//open.Start("E:\\Music\\Nhac hieu chuong trinh\\A kiss of goodbye - Ocarina.mp3")

	/* Scan for drive in window computer */
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ"{
		_, err := os.Open(string(drive)+":\\")
		if err == nil {
			fmt.Println("*************************************************************",string(drive))
			files, _ := ioutil.ReadDir(string(drive) + ":\\")
			for _, f := range files {
				fmt.Println(f.Name())
			}
		}
	}


	fmt.Println("time execute main: ", time.Since(start))

}

