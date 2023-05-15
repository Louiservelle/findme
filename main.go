package main

import (
	"net/http"
	"os"

	"github.com/Louiservelle/findme/api"
	"github.com/Louiservelle/findme/handlers"
)

func main() {

	//filesystem server for images and stylesheet files.
	rootdir, _ := os.Getwd()
	cssfileserver := http.FileServer(http.Dir(rootdir + "/assets/css/"))
	imgfileserver := http.FileServer(http.Dir(rootdir + "/assets/img/"))
	jsfileserver := http.FileServer(http.Dir(rootdir + "/assets/js/"))

	http.Handle("/static/", http.StripPrefix("/static/", cssfileserver))
	http.Handle("/staticimg/", http.StripPrefix("/staticimg/", imgfileserver))
	http.Handle("/staticjs/", http.StripPrefix("/staticjs/", jsfileserver))

	// All Routes
	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/registerHandler", handlers.RegisterHandler)
	http.HandleFunc("/register", handlers.Register)

	http.HandleFunc("/loginHandler", handlers.LoginHandler)
	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/about", handlers.About)

	http.HandleFunc("/createMission", handlers.MissionsHandler)
	http.HandleFunc("/missions", handlers.Missions)
	http.HandleFunc("/viewMissionHandler", handlers.ViewMissionsHandler)
	http.HandleFunc("/viewMissions", api.MissionsHandler)

	// All Api Routes
	http.HandleFunc("/api/users", api.UsersHandler)
	http.HandleFunc("/api/Missions", api.MissionsHandler)
	http.HandleFunc("/api/Mission/", api.OneMission)

	// Listen on http://localhost:80/
	http.ListenAndServe(":80", nil)
}
