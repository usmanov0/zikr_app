package http

import (
	"encoding/json"
	"net/http"
	"zikar-app/internal/app"
	"zikar-app/internal/model"
)

type PrayHandler struct {
	prayUseCase app.PrayService
}

func NewPrayHandler(prayService app.PrayService) *PrayHandler {
	return &PrayHandler{prayUseCase: prayService}
}

func (p *PrayHandler) PraySave(w http.ResponseWriter, r *http.Request) {
	var pray model.Pray

	err := json.NewDecoder(r.Body).Decode(&pray)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	err = p.prayUseCase.SavePray(&pray)
	if err != nil {
		http.Error(w, "failed to create a new pray", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Pray created")
}

func (p *PrayHandler) PrayRead(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	prays, err := p.prayUseCase.ReadPray(id)
	if err != nil {
		http.Error(w, "failed to create a new pray", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prays)
}

func (p *PrayHandler) PrayUpdate(w http.ResponseWriter, r *http.Request) {
	var pray model.Pray

	err := json.NewDecoder(r.Body).Decode(&pray)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	err = p.prayUseCase.PutPray(pray.Id, pray.Language, pray.Definition)
	if err != nil {
		http.Error(w, "failed to update pray", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (p *PrayHandler) PrayDelete(w http.ResponseWriter, r *http.Request) {
	var id int

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	err = p.prayUseCase.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete pray", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
