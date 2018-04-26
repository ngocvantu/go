package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time" 
)

type PageVars struct {
	Title     string
	Resultset []User
}

var db, errglobal = sql.Open("mysql", "root:@tcp(localhost:3306)/test")

type User struct {
	Id   int
	Name string
	Age  int
}

func DeleteFromDB(userID int, w http.ResponseWriter) {
	_, err := db.Exec("DELETE FROM users WHERE userid = ?;", userID)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("Xóa thành công"))
}

func root(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Host %s!", r.Host)
	//	fmt.Fprintf(w, "\n")
	//	fmt.Fprintf(w, "URL %s!", r.URL)
	//	fmt.Fprintf(w, "\n")
	//	fmt.Fprintf(w, "RemoteAddr %s!", r.RemoteAddr) 
	if r.Method == "POST" {

		if r.URL.Path[0:15] == "/users/xoauser/" {
			start := time.Now()
			fmt.Println(r.URL.Path[0:15])
			var userId string
			userId = r.URL.Path[15:]

			i, err := strconv.Atoi(userId)
			fmt.Println(i)
			go DeleteFromDB(i, w)
			if err != nil {
				// handle error
				fmt.Println(err)
			} 
			elapsed := time.Since(start)
		    fmt.Println("Delete took: ", elapsed)
		} 
	} else {
		HomePageVars := &PageVars{
			Title: "Xin chao",
		}

		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		err = t.Execute(w, HomePageVars)
		if err != nil { // if there is an error
			log.Print("template executing error: ", err) //log it
		}
	}
}

func KinhNghiem(w http.ResponseWriter, r *http.Request) {

	HomePageVars := PageVars{
		Title: "Kinh nghiem",
	}

	t, err := template.ParseFiles("kinhnghiem.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("list user")
	start := time.Now()
	HomePageVars := PageVars{
		Title: "Danh sách user",
	}

	//	insert, err := db.Query("INSERT INTO users (username, age) VALUES ( 'tunguyen', '25' )")
	results, err := db.Query("SELECT * FROM users")

	if err != nil {
		//		fmt.Println("error")
		panic(err.Error())
	}

	defer results.Close()
	resulltSet := make([]User, 0)
	for results.Next() {
		var user User
		err = results.Scan(&user.Id, &user.Name, &user.Age)
		resulltSet = append(resulltSet, user)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

	}

	m := template.New("users.html").Funcs(template.FuncMap{
		"inc": func(n int) int {
			return n + 1
		},
	})

	t, err := m.ParseFiles("users.html")
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	resulltSet = resulltSet[1:50]
	HomePageVars.Resultset = resulltSet
//	fmt.Println("result set: ", resulltSet)
	err = t.Execute(w, HomePageVars)
	elapsed := time.Since(start)
    fmt.Println("Listuser took: ", elapsed)
}

func XoaUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("xoa")
	userId := r.URL.Path[3:]
	fmt.Println(userId)
	if r.Method == "POST" {
		_, err := db.Exec("DELETE FROM users WHERE userid = ?;", 1)
		if err != nil {
			panic(err)
		}
		w.Write([]byte("success"))
	}
}

func main() {
	log.Println("chao")

	http.HandleFunc("/", root)
	http.HandleFunc("/kinhnghiem", KinhNghiem)
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/users/xoauser", XoaUser)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe(":8090", nil))
}
