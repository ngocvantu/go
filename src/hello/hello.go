package main

import ( 
	"fmt"  
	"os" 
	"io/ioutil"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	readFileGolang("E:\\GoogleDriver\\khac\\slimks\\src\\jp\\co\\toshiba_sol\\slim\\ks\\ksichiran\\KSKousuPrintToPDFService.java")
}

func readFileGolang(path string) {
	
	
	files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
    }
	
	
	
	
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close( ) 
	
	db, err := sql.Open("mysql", "root:@/test")
	insert, err := db.Query("INSERT INTO chao1 VALUES ( '10j', 'TEST' )")
	
	if err != nil {
        panic(err.Error())
    }
	
	defer insert.Close()
	
	
	
	var input string
	fmt.Scanln(&input)
}
