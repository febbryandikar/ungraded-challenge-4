package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"ungraded-challenge-4/entity"

	"github.com/julienschmidt/httprouter"
)

type NewRecordHandler struct {
	*sql.DB
}

func (h *NewRecordHandler) GetCriminalRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var records []entity.Record

	rows, err := h.Query("SELECT ce.event_id, h.hero_name, v.villain_name, ce.description, ce.event_time FROM criminal_event ce JOIN hero h ON ce.hero_id = h.hero_id JOIN villain v ON ce.villain_id = v.villain_id")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record entity.Record
		err = rows.Scan(&record.ID, &record.HeroName, &record.VillainName, &record.Description, &record.EventTime)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.Message{
				Status:  "failed",
				Message: err.Error(),
			})
			return
		}
		records = append(records, record)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status:  "success",
		Message: "Criminal Records retrieved successfully",
		Data:    records,
	})
}

func (h *NewRecordHandler) GetCriminalRecordById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var record entity.Record
	paramsId := p.ByName("id")

	query := "SELECT ce.event_id, h.hero_name, v.villain_name, ce.description, ce.event_time FROM criminal_event ce JOIN hero h ON ce.hero_id = h.hero_id JOIN villain v ON ce.villain_id = v.villain_id WHERE ce.event_id = ?"
	if err := h.QueryRow(query, paramsId).Scan(&record.ID, &record.HeroName, &record.VillainName, &record.Description, &record.EventTime); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status:  "success",
		Message: "Criminal Record retrieved successfully",
		Data:    record,
	})
}

func (h *NewRecordHandler) AddNewCriminalRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newRecord entity.Record
	var heroID int
	var VillainID int

	err := json.NewDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	queryHero := "SELECT hero_id FROM hero WHERE hero_name = ?"
	if err = h.QueryRow(queryHero, newRecord.HeroName).Scan(&heroID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	queryVillain := "SELECT villain_id FROM villain WHERE villain_name = ?"
	if err = h.QueryRow(queryVillain, newRecord.VillainName).Scan(&VillainID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	query := "INSERT INTO criminal_event (hero_id, villain_id, description, event_time) VALUES (?, ?, ?, ?)"
	result, err := h.Exec(query, heroID, VillainID, newRecord.Description, newRecord.EventTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}
	newRecord.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status:  "success",
		Message: "New Criminal Record added successfully",
		Data:    newRecord,
	})
}

func (h *NewRecordHandler) UpdateCriminalRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var record entity.Record
	var heroID int
	var VillainID int

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	paramsId := p.ByName("id")
	record.ID, err = strconv.Atoi(paramsId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	queryHero := "SELECT hero_id FROM hero WHERE hero_name = ?"
	if err = h.QueryRow(queryHero, record.HeroName).Scan(&heroID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	queryVillain := "SELECT villain_id FROM villain WHERE villain_name = ?"
	if err = h.QueryRow(queryVillain, record.VillainName).Scan(&VillainID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	query := "UPDATE criminal_event SET hero_id = ?, villain_id = ?, description = ?, event_time = ? WHERE event_id =?"
	_, err = h.Exec(query, heroID, VillainID, record.Description, record.EventTime, paramsId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status:  "success",
		Message: "Criminal Record updated successfully",
		Data:    record,
	})
}

func (h *NewRecordHandler) DeleteCriminalRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	paramsId := p.ByName("id")

	query := "DELETE FROM criminal_event WHERE event_id = ?"
	_, err := h.Exec(query, paramsId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity.Message{
		Status:  "success",
		Message: "Criminal Record deleted successfully",
	})
}
