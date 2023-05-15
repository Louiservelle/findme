package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"unicode"

	"github.com/Louiservelle/findme/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	models.RenderTemplate(w, "register")
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var err error

	db, err = sql.Open("mysql", "root:root@(127.0.0.1:8889)/findme?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	r.ParseForm()
	name := r.FormValue("name")
	fmt.Println(name)
	password := r.FormValue("Password")
	email := r.FormValue("email")
	organisation := r.FormValue("organisation")
	phone := r.FormValue("Phone")

	// check name for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range name {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	// check name length
	var nameLength bool
	if 5 <= len(name) && len(name) <= 50 {
		nameLength = true
	}

	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true

	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}

	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
		models.RenderTemplate(w, "register")
	}

	// create hash from password
	var hash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		models.RenderTemplate(w, "register")
	}
	fmt.Println("hash:", hash)
	fmt.Println("string(hash):", string(hash))
	// execute query.
	insert := `INSERT INTO Users (name, password, email, organisation, Phone) VALUES ("` + name + `","` + string(hash) + `","` + email + `","` + organisation + `","` + phone + `")`
	db.Query(insert)
	http.Redirect(w, r, "/login", 302)
}
