package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aleroxac/alura-golang/database"
	"github.com/aleroxac/alura-golang/models"
	"github.com/gorilla/mux"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func List(w http.ResponseWriter, r *http.Request) {
	var skills []models.Skill
	database.DB.Find(&skills)
	json.NewEncoder(w).Encode(skills)
}

func GetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	skill_name := vars["name"]

	var skills []models.Skill
	database.DB.Where("sk_name = ?", skill_name).Find(&skills)
	json.NewEncoder(w).Encode(skills)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var new_skill models.Skill
	err := json.NewDecoder(r.Body).Decode(&new_skill)
	if err != nil {
		log.Fatal(err)
	}

	database.DB.Create(&new_skill)
	json.NewEncoder(w).Encode(new_skill)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	skill_name := vars["name"]

	var skill models.Skill
	database.DB.Where("sk_name = ?", skill_name).Find(&skill)

	var updated_skill models.Skill
	err := json.NewDecoder(r.Body).Decode(&updated_skill)
	if err != nil {
		log.Fatal(err)
	}
	database.DB.Where("sk_name = ?", skill_name).Updates(&updated_skill)

	GetByName(w, r)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	skill_name := vars["name"]

	var skill models.Skill
	database.DB.Where("sk_name = ?", skill_name).Delete(&skill)
}
