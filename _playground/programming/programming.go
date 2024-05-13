package programming

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ProgrammingHandler struct {
	Db *sql.DB
}

func NewProgrammingHandler(db *sql.DB) *ProgrammingHandler {
	return &ProgrammingHandler{
		Db: db,
	}
}

func (p ProgrammingHandler) Languages(w http.ResponseWriter, r *http.Request) {
	type Language struct {
		Name     string `json:"name"`
		ImageURL string `json:"imageUrl"`
	}

	var languages []Language
	rows, err := p.Db.Query("select name, imageUrl from languages")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var imageUrl string
		if err := rows.Scan(&name, &imageUrl); err != nil {
			slog.Error(err.Error())
			return
		}

		languages = append(languages, Language{
			Name:     name,
			ImageURL: imageUrl,
		})
	}

	if err := json.NewEncoder(w).Encode(languages); err != nil {
		slog.Error(err.Error())
		return
	}
}
