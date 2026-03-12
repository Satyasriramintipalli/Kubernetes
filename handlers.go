package main

import (
	"fmt"
	"net/http"
)

func loginPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "templates/login.html")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	row := db.QueryRow("SELECT username FROM users WHERE username=? AND password=?", username, password)

	var user string

	err := row.Scan(&user)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func dashboard(w http.ResponseWriter, r *http.Request) {

	rows, _ := db.Query("SELECT id, username FROM users")

	fmt.Fprintf(w, "<h1>User List</h1>")

	for rows.Next() {

		var id int
		var username string

		rows.Scan(&id, &username)

		fmt.Fprintf(w, "%d - %s <br>", id, username)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	statement, _ := db.Prepare("INSERT INTO users(username,password) VALUES(?,?)")

	statement.Exec(username, password)

	fmt.Fprintf(w, "User Created")
}