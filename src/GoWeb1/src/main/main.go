package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type PageVars struct {
	Title              string
	NumberObKnownVocab int
	Resultset          []User
}

var db, errglobal = sql.Open("mysql", "root:@tcp(localhost:3306)/test")
var file, err = os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

type User struct {
	Id   int
	Name string
	Age  int
}

func DeleteFromDB(userID int, w http.ResponseWriter) {
	log.Println("SQL: DELETE FROM users WHERE userid =", userID)
	_, err := db.Exec("DELETE FROM users WHERE userid = ?;", userID)
	if err != nil {
		panic(err)
	}
	log.Println("Xóa thành công")
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
			DeleteFromDB(i, w)
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
	results, err := db.Query("SELECT * FROM users order by userid desc limit 50")

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
	//	resulltSet = resulltSet[1:50]
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

func AddUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Println("add user start")
	// parse form
	r.ParseForm()
	username := r.FormValue("username")
	age, _ := strconv.Atoi(r.FormValue("age"))

	// insert to DB
	log.Println("age: ", age)
	db.Exec("INSERT INTO users (username, age) VALUES (?, ?);", username, age)

	log.Println("INSERT INTO users (username, age) VALUES ('", username, "',", age, ")", "\n")

	// redirect
	http.Redirect(w, r, "/users", http.StatusFound)
	log.Println("add user end")
	excuteTime := time.Since(start)
	fmt.Println("Add user took: ", excuteTime)
	log.Println("Add user took: ", excuteTime)
}

func Knownvocab(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method == "GET" {
		KnownvocabVar := &PageVars{
			Title: "Knownvocab",
		}
		t, err := template.ParseFiles("knownvocab.html")
		if err != nil {
			log.Panicln(err.Error())
		}
		err = t.Execute(w, KnownvocabVar)
		if err != nil { // if there is an error
			log.Print("template executing error: ", err)
		}

	} else {

		r.ParseForm()
		content := r.FormValue("content")
		contentArr := strings.Split(content, "\n")
		var i int
		i = 0
		for _, v := range contentArr {

			contentArrArr := strings.Split(v, " ")
			for _, v1 := range contentArrArr {
				thisIsAKnownVocab := false
				i = i + 1
				listKnowVocab, _ := db.Query("SELECT * FROM known_vocab")
				var processedString string
				for listKnowVocab.Next() {
					var vocab_id string
					var vocab string
					err = listKnowVocab.Scan(&vocab_id, &vocab)
					if err != nil {
						panic(err.Error()) // proper error handling instead of panic in your app
					}

					reg, err := regexp.Compile("[^a-zA-Z0-9]+")
					if err != nil {
						log.Fatal(err)
					}
					processedString = reg.ReplaceAllString(v1, "")
					fmt.Println(processedString)

					if strings.ToUpper(vocab) == strings.ToUpper(strings.Trim(processedString, " ")) {
						thisIsAKnownVocab = true
						break
					}

					fmt.Println(">>>>>>>>>", vocab)
					fmt.Println(">>>>>>>>>", processedString)
					fmt.Println(">>", thisIsAKnownVocab)
				}

				if thisIsAKnownVocab == false {
					db.Exec("insert into known_vocab (known_vocab_word) value (?) ;", strings.Trim(processedString, " "))
				}
			}
		}
		//		numberOfKnowVocab := 0;
		//		Title := "Knownvocab"

		http.Redirect(w, r, "/knownvocab", http.StatusFound)

		elapse := time.Since(start)
		log.Println("Add know vocab take: ", elapse)
		fmt.Println("Add know vocab take: ", elapse)
	}

}

type Task struct {
	Id          int
	Content     string
	DateCreated string
	IdComplete  bool
	StartTime   string
	EndTime     string
	TotalTime   string
}

type TaskPage struct {
	Title    string
	TaskList []Task
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "GET" {
		s := []Task{}
		task := &Task{}
		tasks, _ := db.Query("select tasks.*, TIMEDIFF(tasks.ENDTIME, tasks.STARTTIME) as TOTAL FROM tasks where DATECREATED = CURDATE()")
		for tasks.Next() {
			tasks.Scan(&task.Id, &task.Content, &task.DateCreated, &task.IdComplete, &task.StartTime, &task.EndTime, &task.TotalTime)
			s = append(s, *task)
		}

		t, _ := template.ParseFiles("tasks.html")
		er := t.Execute(w, &TaskPage{"Task today", s})

		if er != nil {
			panic(er)
		}
	} else {
		r.ParseForm()
		taskContent := r.FormValue("task-content")
		if len(taskContent) > 0 {
			db.Exec("INSERT into tasks (CONTENT, DATECREATED, ISCOMPLETE, STARTTIME, ENDTIME) VALUES (?, CURDATE(),0,CURTIME(),null)", taskContent)
		}
		http.Redirect(w, r, "/tasks", http.StatusFound)
	}
}

func Complete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		db.Exec("update tasks set ISCOMPLETE = 1,ENDTIME = CURTIME() WHERE id = ?", string(b))
		w.Write([]byte("ok"))
	}
}

func LamTiep(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		db.Exec("update tasks set ISCOMPLETE = 0, ENDTIME = null WHERE id = ?", string(b))
		w.Write([]byte("ok"))
	}
}

func Xoa(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		db.Exec("delete from tasks WHERE id = ?", string(b))
		w.Write([]byte("ok"))
	}
}

func main() {
	defer file.Close()
	log.SetOutput(file)

	http.HandleFunc("/", root)
	http.HandleFunc("/kinhnghiem", KinhNghiem)
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/users/xoauser", XoaUser)
	http.HandleFunc("/adduser", AddUser)
	http.HandleFunc("/knownvocab", Knownvocab)
	http.HandleFunc("/tasks", Tasks)
	http.HandleFunc("/complete", Complete)
	http.HandleFunc("/lamtiep", LamTiep)
	http.HandleFunc("/xoa", Xoa)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe(":8090", nil))
}
