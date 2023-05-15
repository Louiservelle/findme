package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Louiservelle/findme/models"
)

func ViewMissionsHandler(w http.ResponseWriter, r *http.Request) {
	getApiMissions, _ := http.Get("http://localhost:80/api/Missions")
	var apiObject []models.Mission
	apiBody, _ := ioutil.ReadAll(getApiMissions.Body)
	json.Unmarshal(apiBody, &apiObject)
	models.RenderTemplate(w, "viewMissionHandler")

}

func MissionsHandler(w http.ResponseWriter, r *http.Request) {

	models.RenderTemplate(w, "createMission")
}

func Missions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var err error

	db, err = sql.Open("mysql", "root:root@(127.0.0.1:8889)/findme?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	r.ParseForm()

	title := r.FormValue("title")
	fmt.Println(title)
	description := r.FormValue("description")
	fmt.Println(description)
	skills := r.FormValue("skills")
	user_id := r.FormValue("user_id")

	insert := `INSERT INTO Missions (title, description, skills, user_id) VALUES ("` + title + `","` + description + `","` + skills + `","` + user_id + `")`
	db.Query(insert)
	http.Redirect(w, r, "/api/Missions", 302)
}
