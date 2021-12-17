package logging

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PembuatanKoneksi() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db, nil
}

//------ utils -----//

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error nii om", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}

func ResponseJSONSEARCH(w http.ResponseWriter, p interface{}, q interface{}, r interface{}, status int) {
	ubahkeByte, err := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "error nii om", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(ubahkeByte))
}
