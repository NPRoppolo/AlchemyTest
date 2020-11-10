package main

import (
    "fmt"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB

func main() {
	fmt.Println("Connecting to local db")

    //connect to imperialfleet db
    db, err := sql.Open("mysql", "root:hacker323@tcp(127.0.0.1:3306)/imperialfleet")

    //handle connection error
    if err != nil {
        panic(err.Error())
    }

	defer db.Close()

	// perform a db.Query insert. Success!
	//insert, err := db.Query("INSERT INTO inventory(id, name, class, armament, crew, image, value, status) VALUES(1, 'Home One', 'Cruiser', 'Turbo Laser', 50000, 'https:\\url.to.image', 1999.99, 'operational')")
	
	//handle insert error
	//if err != nil {
		//panic(err.Error())
	//}

	//defer insert.Close()

	http.HandleFunc("/spaceship", spaceship)

	http.ListenAndServe(":8090", nil)
}

func spaceship(w http.ResponseWriter, req *http.Request) {
	
	fmt.Fprintf(w, "Welcome to the Imperial Fleet\n")
	method := req.Method

	if method == "GET"{
		data := ReadSpaceship()
		fmt.Fprintf(w, data)
	} else if method == "POST" {
		data := UpdateSpaceship()
		fmt.Fprintf(w, data)
	}else if method == "PUT" {
		data := CreateSpaceship()
		fmt.Fprintf(w, data)
	} else if method == "DELETE"{
		data := DeleteSpaceship()
		fmt.Fprintf(w, data)
	}
}

func CreateSpaceship() string{
	return "Create"
}

func ReadSpaceship() string{
	return "Read"
}

func UpdateSpaceship() string{
	return "Update"
}

func DeleteSpaceship() string{
	return "Delete"
}