package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/darlio88/go-orm/internals"
	chi "github.com/go-chi/chi/v5"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := internals.DatabaseInstance()
	defer db.Close()
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal("Row querying err", err)
		w.WriteHeader(500)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("Server error"))
		return
	}
	log.Println(rows)
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Fatal("row scanning error", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	fmt.Println(users)
	res, err := json.Marshal(users)
	if err != nil {
		log.Fatal("JSON marshaling error:", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)

		return
	}
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := internals.DatabaseInstance()
	defer db.Close()
	// Get the param which is the id
	paramId := chi.URLParam(r, "id")
	log.Println(paramId)
	// User
	var SelectedUser User

	// Query database
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", paramId).Scan(&SelectedUser.Id, &SelectedUser.Name, &SelectedUser.Age)
	if err != nil {
		log.Println("failed to prepare statement", err)
	}
	//
	log.Println(SelectedUser)
	// Return the SelectedUser to the client
	res, err := json.Marshal(SelectedUser)
	if err != nil {
		log.Println("JSON marshaling error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//get the id of the user to be updated
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Failed to read body data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	// get the user from the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body data", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	var userUpdate User
	err = json.Unmarshal(body, &userUpdate)
	if err != nil {
		log.Println("Json unmarshal error", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	log.Println(userId)
	log.Println(userUpdate)
	db := internals.DatabaseInstance()
	_, err = db.Query("UPDATE users SET name=$1, age=$2 WHERE id=$3", userUpdate.Name, userUpdate.Age, int16(userId))
	if err != nil {
		log.Println("update user query", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	userUpdate.Id = userId
	//send a response back with status code OK and the result of the operation
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(userUpdate)
	w.Write([]byte(res))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//get the id from the params and convert it to an integer
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatal("Error reading the user-id param", err)
		http.Error(w, "Server Error", http.StatusBadRequest)
		return
	}
	//database instance
	db := internals.DatabaseInstance()
	//delete the user
	result, err := db.Exec("DELETE FROM users where id=$1", userId)
	if err != nil {
		log.Fatal("Error deleting user", err)
		http.Error(w, "Server Error", http.StatusBadRequest)
	}
	log.Println(result)
	//return ok to the user
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	res, _ := json.Marshal(&map[string]string{"msg": "success"})
	w.Write([]byte(res))
}
