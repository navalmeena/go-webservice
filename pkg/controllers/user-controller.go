package controllers

import (
	"encoding/json"
	"fmt"
	"go-curd-ops/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"go-curd-ops/pkg/models"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b:= CreateUser.CreateUser()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers:= models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["UserId"]
	ID, err:= strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	UserDetails, _:= models.GetUserById(ID)
	res, _ := json.Marshal(UserDetails)
	utils.SendResponse(r,w, res);
	
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	var updateUser = &models.User{}
// 	utils.ParseBody(r, updateUser)
// 	vars := mux.Vars(r)
// 	UserId := vars["UserId"]
// 	ID, err:= strconv.ParseInt(UserId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}
// 	UserDetails, db:= models.GetUserById(ID)
// 	if updateUser.Name != "" {
// 		UserDetails.Name = updateUser.Name
// 	}
// 	if updateUser.Author != "" {
// 		UserDetails.Author = updateUser.Author
// 	}
// 	if updateUser.Publication != "" {
// 		UserDetails.Publication = updateUser.Publication
// 	}
// 	db.Save(&UserDetails)
// 	res, _ := json.Marshal(UserDetails)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	UserId := vars["UserId"]
// 	ID, err:= strconv.ParseInt(UserId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}
// 	User:= models.DeleteUser(ID)
// 	res, _ := json.Marshal(User)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }