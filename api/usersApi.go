package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Louiservelle/findme/models"
)

var db *sql.DB

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var err error

	db, err = sql.Open("mysql", "root:root@(127.0.0.1:8889)/findme?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	var users []models.User
	var query, _ = db.Query(`SELECT * FROM Users`)
	fmt.Println(query)
	for query.Next() {
		var user models.User
		query.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Organization, &user.Phone)
		users = append(users, user)
	}
	fmt.Println(users)
	a, _ := json.Marshal(users)
	w.Write(a)
}
