package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)


type SharkAttack struct {
	Id       int    `json:"id"`
	Date     string `json:"date"`
	Country  string `json:"country"`
	Activity string `json:"activity"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	FatalYN  string `json:"fatal_y_n"`
}

var (
	sharkAttacksMutex sync.Mutex
	sharkAttacks      = make(map[int]SharkAttack)
	nextID            = 0
)

func main() {
	var jsonSharkAttacks []SharkAttack
	postsToLoad := 10
	byteValue, err := os.ReadFile("C:\\Users\\michalpc\\Desktop\\Studia\\ug-golang\\zadanie_3\\global-shark-attack.json")
	if err != nil {
		println("Error reading file")
		return
	}
	sharkAttacksMutex.Lock()

	json.Unmarshal(byteValue, &jsonSharkAttacks)

	for i := 0; i < len(jsonSharkAttacks); i++ {
		jsonSharkAttacks[i].Id = i
	}
	
	postsToLoad = min(postsToLoad, len(jsonSharkAttacks))
	for i := 0; i < postsToLoad; i++ {
		sharkAttacks[nextID] = jsonSharkAttacks[i]
		nextID++
	}
	
	
	sharkAttacksMutex.Unlock()
	fmt.Printf("Loaded number of shark attacks: %d\n", len(sharkAttacks))
	http.HandleFunc("/attacks", attacksHandler)
	http.HandleFunc("/attacks/", attackHandler)
	fmt.Println("Server is listening on https://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func attackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/attacks/"):])
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	switch r.Method {
		case http.MethodGet:
			handleGetAttack(w, r, id)
		case http.MethodDelete:
			handleDeleteAttack(w, r, id)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func attacksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
				handleGetAttacks(w, r)
		case http.MethodPost:
			handlePostAttack(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleGetAttacks(w http.ResponseWriter, r *http.Request) {
	sharkAttacksMutex.Lock()
	defer sharkAttacksMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sharkAttacks)
}

func handleGetAttack(w http.ResponseWriter, r *http.Request, id int) {
	sharkAttacksMutex.Lock()
	defer sharkAttacksMutex.Unlock()

	_, ok := sharkAttacks[id]
	if !ok {
		http.Error(w, "Shark attack not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sharkAttacks[id])
}

func handlePostAttack(w http.ResponseWriter, r *http.Request) {
	var sharkAttack SharkAttack
	err := json.NewDecoder(r.Body).Decode(&sharkAttack)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sharkAttacksMutex.Lock()
	defer sharkAttacksMutex.Unlock()
	sharkAttack.Id = nextID
	sharkAttacks[nextID] = sharkAttack
	nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sharkAttack)
}

func handleDeleteAttack(w http.ResponseWriter, r *http.Request, id int) {
	sharkAttacksMutex.Lock()
	defer sharkAttacksMutex.Unlock()

	_, ok := sharkAttacks[id]
	if !ok {
		http.Error(w, "Shark attack not found", http.StatusNotFound)
		return
	}

	delete(sharkAttacks, id)
	w.WriteHeader(http.StatusOK)
}