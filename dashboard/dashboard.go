package dashboard

import (
	"database/sql"
	"encoding/json"
	"github.com/kinxyo/Servebix.git/types"
	"log"
	"net/http"
)

func ShowAll(w http.ResponseWriter, db *sql.DB) {
	var patientList []types.User

	rows, err := db.Query("SELECT * FROM patient_creds")
	if err != nil {
		log.Println("Query error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var patient types.User
		if err := rows.Scan(&patient.ID, &patient.Name, &patient.Email, &patient.Phone, &patient.Password, &patient.CreatedAt); err != nil {
			log.Println("Scan error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		patientList = append(patientList, patient)
	}

	if err := rows.Err(); err != nil {
		log.Println("Rows error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(patientList); err != nil {
		log.Println("Encode error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
