package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/Louiservelle/findme/models"
)

func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var err error

	db, err = sql.Open("mysql", "root:root@(127.0.0.1:8889)/findme?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	var missions []models.Mission
	var query, _ = db.Query(`SELECT * FROM Missions`)
	fmt.Println(query)
	for query.Next() {
		var mission models.Mission
		query.Scan(&mission.ID, &mission.Title, &mission.Description, &mission.Skills, &mission.UserId)
		missions = append(missions, mission)
	}
	fmt.Println(missions)
	a, _ := json.Marshal(missions)
	w.Write(a)
}

func OneMission(w http.ResponseWriter, r *http.Request) {
	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	fmt.Println(pathIDint)
}
