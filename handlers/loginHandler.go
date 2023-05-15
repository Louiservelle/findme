package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Louiservelle/findme/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	models.RenderTemplate(w, "login")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var erreur error

	db, erreur = sql.Open("mysql", "root:root@(127.0.0.1:8889)/findme?parseTime=true")
	if erreur != nil {
		panic(erreur.Error())
	}

	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println(email)
	fmt.Println(password)

	var hash string
	var user models.User

	emailVar := `SELECT email, password, FROM Users WHERE email="` + email + `" AND password="` + password + `"`

	row := db.QueryRow(emailVar, email)
	fmt.Println(row)
	var getRaw = db.QueryRow(emailVar)
	fmt.Println(getRaw)
	getRaw.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Organization, &user.Phone)
	fmt.Println(getRaw)
	db.Query(emailVar)

	err := row.Scan(&user.Password)
	fmt.Println("hash from db:", hash)
	if err != nil {
		fmt.Println("error selecting Hash in db by Username")
		models.RenderTemplate(w, "login")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// returns nill on succcess
	if err == nil {
		fmt.Fprint(w, "You have successfully logged in :)")
		return
	}
}
