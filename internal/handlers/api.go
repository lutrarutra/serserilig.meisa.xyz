package handlers

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/iMeisa/serserilig.meisa.xyz/internal/models"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/network"
	"github.com/iMeisa/serserilig.meisa.xyz/internal/tools"
)

func (m *Repository) AddRace(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
		return
	}

	gp_id_raw, ok := r.URL.Query()["gp_id"]
	if !ok || len(gp_id_raw) < 1 {
		log.Println("GP id not ok")
		w.Write([]byte("GP id not ok"))
		return
	}

	date_raw, ok := r.URL.Query()["date"]
	if !ok || len(date_raw) < 1 {
		log.Println("Date not ok")
		w.Write([]byte("Date not ok"))
		return
	}

	time_raw, ok := r.URL.Query()["time"]
	if !ok || len(date_raw) < 1 {
		log.Println("Time not ok")
		w.Write([]byte("Time not ok"))
		return
	}

	var gps []models.GrandPrix

	gp_file, err := ioutil.ReadFile("./static/json/gp_reference.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(gp_file, &gps)
	if err != nil {
		return
	}

	gp_id, _ := strconv.Atoi(gp_id_raw[0])
	gp := gps[gp_id-1]

	newRace := models.Race{
		Date: date_raw[0],
		Time: time_raw[0],
		Gp:   gp,
	}

	err = m.DB.InsertRace(newRace)

	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	m.updateCalendarJSON()
	w.Write([]byte("Added Race"))
	log.Println("Added Race")
}

func (m *Repository) DeleteRace(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
		return
	}

	raceId, ok := r.URL.Query()["id"]
	if !ok || len(raceId) < 1 {
		w.Write([]byte("id not ok"))
		return
	}

	if _, err := strconv.Atoi(raceId[0]); err != nil {
		w.Write([]byte("Invalid ID param"))
		return
	}

	err := m.DB.DeleteRace(raceId[0])
	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	m.updateCalendarJSON()
	w.Write([]byte("Deleted Race"))
	log.Println("Deleted Race")
}

func (m *Repository) UpdateRace(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
	}

	id_raw, ok := r.URL.Query()["id"]
	if !ok || len(id_raw) < 1 {
		log.Println("ID not ok")
		w.Write([]byte("ID not ok"))
		return
	}

	date_raw, ok := r.URL.Query()["date"]
	if !ok || len(date_raw) < 1 {
		log.Println("Date not ok")
		w.Write([]byte("Date not ok"))
		return
	}

	time_raw, ok := r.URL.Query()["time"]
	if !ok || len(date_raw) < 1 {
		log.Println("Time not ok")
		w.Write([]byte("Time not ok"))
		return
	}
	log.Println(date_raw[0])
	log.Println(time_raw[0])

	err := m.DB.UpdateRace(id_raw[0], date_raw[0], time_raw[0])

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}

	m.updateCalendarJSON()
	log.Println("Race Updated!")
}

func (m *Repository) GetAllRaces(w http.ResponseWriter, r *http.Request) {
	//m.DB.CreateCalendarTable()

	races, err := m.DB.QueryAllRaces()
	if err != nil {
		log.Fatal(err)
	}

	racesJSON, err := json.Marshal(races)
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(racesJSON)
}

func (m *Repository) GetAllStaff(w http.ResponseWriter, r *http.Request) {
	//m.DB.CreateCalendarTable()

	races, err := m.DB.QueryAllStaff()
	if err != nil {
		log.Fatal(err)
	}

	racesJSON, err := json.Marshal(races)
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(racesJSON)
}

func (m *Repository) AddStaff(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name_field")
	title := r.Form.Get("title_field")
	role := r.Form.Get("role_field")

	crop_x, _ := strconv.Atoi(r.Form.Get("crop_x"))
	crop_y, _ := strconv.Atoi(r.Form.Get("crop_y"))
	crop_w, _ := strconv.Atoi(r.Form.Get("crop_w"))
	crop_h, _ := strconv.Atoi(r.Form.Get("crop_h"))
	img_w, _ := strconv.ParseFloat(r.Form.Get("img_w"), 32)
	img_h, _ := strconv.ParseFloat(r.Form.Get("img_h"), 32)

	filename := "/static/images/default.svg"

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("picture")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Redirect(w, r, r.Header.Get("Referer"), 302)

	} else {
		defer file.Close()

		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		file_extension := ""
		contentType := http.DetectContentType(fileBytes)
		switch contentType {
		case "image/jpeg":
			file_extension = ".jpg"
		case "image/png":
			file_extension = ".png"
		}

		tempFile, err := ioutil.TempFile("static/images/temp", "*"+file_extension)

		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		tempFile.Write(fileBytes)

		log.Println(tempFile.Name())
		filename, err = tools.SaveProfilePicture(crop_x, crop_y, crop_w, crop_h, img_w, img_h, tempFile.Name(), file_extension)
		if err != nil {
			log.Println(err)
		}
		os.Remove(tempFile.Name())
	}

	var roles []models.Role

	role_file, err := ioutil.ReadFile("./static/json/role_reference.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(role_file, &roles)
	if err != nil {
		return
	}

	role_id, _ := strconv.Atoi(role)
	role_type := roles[role_id-1]

	newStaff := models.Staff{
		Name:      name,
		Role:      role_type,
		Title:     title,
		ImagePath: filename,
	}

	err = m.DB.InsertStaff(newStaff)

	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	m.updateStaffJSON()
	w.Write([]byte("Added Staff Member"))
	log.Println("Added Staff Member")

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func (m *Repository) DeleteStaff(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
		return
	}

	id_raw, ok := r.URL.Query()["id"]
	if !ok || len(id_raw) < 1 {
		log.Println("Id not ok")
		w.Write([]byte("Id not ok"))
		return
	}

	if _, err := strconv.Atoi(id_raw[0]); err != nil {
		w.Write([]byte("Invalid ID param"))
		return
	}

	staff, err := m.DB.QueryStaff(id_raw[0])
	if err != nil {
		log.Println(err)
	}

	if staff.ImagePath != "/static/images/default.svg" {
		os.Remove(staff.ImagePath[1:])
	}

	err = m.DB.DeleteStaff(id_raw[0])
	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	m.updateStaffJSON()
	w.Write([]byte("Deleted Staff Member"))
	log.Println("Deleted Staff Member")
}

func (m *Repository) UpdateStaff(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
		return
	}

	id_raw, ok := r.URL.Query()["id"]
	if !ok || len(id_raw) < 1 {
		log.Println("Id not ok")
		w.Write([]byte("Id not ok"))
		return
	}

	name_raw, ok := r.URL.Query()["name"]
	if !ok || len(name_raw) < 1 {
		log.Println("Name not ok")
		w.Write([]byte("Name not ok"))
		return
	}

	title_raw, ok := r.URL.Query()["title"]
	if !ok || len(title_raw) < 1 {
		log.Println("Title not ok")
		w.Write([]byte("Title not ok"))
		return
	}

	role_raw, ok := r.URL.Query()["role"]
	if !ok || len(role_raw) < 1 {
		log.Println("Role not ok")
		w.Write([]byte("Role not ok"))
		return
	}

	err := m.DB.UpdateStaff(id_raw[0], name_raw[0], title_raw[0], role_raw[0])

	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
	}

	m.updateStaffJSON()
	log.Println("Staff Updated!")
}

func (m *Repository) AddDriver(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
		return
	}

	name, ok := r.URL.Query()["name"]
	if !ok || len(name) < 1 {
		w.Write([]byte("name not ok"))
		return
	}

	newDriver := models.Driver{
		Name: name[0],
	}

	err := m.DB.InsertDriver(newDriver)
	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	m.updateDriverJSON()
	w.Write([]byte("Added Driver"))
}

func (m *Repository) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
		return
	}

	driverId, ok := r.URL.Query()["id"]
	if !ok || len(driverId) < 1 {
		w.Write([]byte("name not ok"))
		return
	}

	if _, err := strconv.Atoi(driverId[0]); err != nil {
		w.Write([]byte("Invalid ID param"))
		return
	}

	err := m.DB.DeleteDriver(driverId[0])
	if err != nil {
		log.Println(err)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	m.updateDriverJSON()
	m.updateTeamJSON()
	w.Write([]byte("Deleted Driver"))
}

func (m *Repository) GetDriver(w http.ResponseWriter, r *http.Request) {
	//m.DB.CreateDriverTable()

	var driverJSON []byte

	query := r.URL.Query()
	if len(query) < 1 {
		w.Write([]byte("Invalid request"))
		return
	} else if val, ok := query["id"]; ok {
		driver, err := m.DB.QueryDriver("id", val[0])
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		driverJSON, err = json.Marshal(driver)
		if err != nil {
			log.Fatal("Could not convert to JSON:", err)
		}
	} else if val, ok = query["name"]; ok {
		driver, err := m.DB.QueryDriver("name", val[0])
		if err != nil {
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		driverJSON, err = json.Marshal(driver)
		if err != nil {
			log.Fatal("Could not convert to JSON:", err)
		}
	} else {
		w.Write([]byte("Invalid request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(driverJSON)
}

func (m *Repository) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
	}

	validIdCols := []string{"id", "name"}

	var idCol string
	var idVal string
	for _, col := range validIdCols {
		if val, ok := r.URL.Query()[col]; ok && len(val) > 0 {

			idCol = col
			idVal = val[0]
			break
		}
	}

	validCols := []string{"team_id", "points", "penalty_points"}

	for _, col := range validCols {

		if val, ok := r.URL.Query()[col]; ok && len(val) > 0 {

			if _, err := strconv.Atoi(val[0]); err == nil {

				err = m.DB.UpdateDriver(idCol, idVal, col, val[0])
				if err != nil {
					w.Write([]byte(fmt.Sprint(err)))
				}
				m.updateDriverJSON()
				return
			}
			w.Write([]byte(fmt.Sprintf("Invalid %v", col)))
		}
	}

	w.Write([]byte("Missing Update Column name"))
}

func (m *Repository) GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	//m.DB.CreateDriverTable()

	drivers, err := m.DB.QueryAllDrivers()
	if err != nil {
		log.Fatal(err)
	}

	driversJSON, err := json.Marshal(drivers)
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(driversJSON)
}

func (m *Repository) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	if !checkIP(r) {
		w.Write([]byte("Invalid IP"))
	}

	validIdCols := []string{"id", "name"}

	var idCol string
	var idVal string
	for _, col := range validIdCols {
		if val, ok := r.URL.Query()[col]; ok && len(val) > 0 {

			idCol = col
			idVal = val[0]
			break
		}
	}

	validCols := []string{"driver1", "driver2", "points"}

	for _, col := range validCols {

		if val, ok := r.URL.Query()[col]; ok && len(val) > 0 {

			if _, err := strconv.Atoi(val[0]); err == nil {

				err = m.DB.UpdateTeam(idCol, idVal, col, val[0])
				if err != nil {
					w.Write([]byte(fmt.Sprint(err)))
				}
				m.updateTeamJSON()
				return
			}
			w.Write([]byte(fmt.Sprintf("Invalid %v", col)))
		}
	}

	w.Write([]byte("Missing Update Column name"))
}

func (m *Repository) GetAllTeams(w http.ResponseWriter, r *http.Request) {
	//m.DB.CreateTeamTable()

	teams, err := m.DB.QueryAllTeams()
	if err != nil {
		log.Fatal(err)
	}

	teamsJSON, err := json.Marshal(teams)
	if err != nil {
		log.Fatal("Could not convert to JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(teamsJSON)
}

func checkIP(r *http.Request) bool {
	remoteIP, ok := r.URL.Query()["ip"]
	if !ok || len(remoteIP) < 1 {
		return false
	}

	userIP := strings.Split(remoteIP[0], ":")[0]
	userIP2 := strings.Split(network.GetRealIP(r), ":")[0]

	if userIP != userIP2 {
		return false
	}
	return true
}
