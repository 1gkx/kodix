package router

import (
	"encoding/json"
	"kodix/internal/store"
	"net/http"
)

type Handler struct {
	store *store.Db
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {

	data, err := h.store.GetAuto(r.Context())
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, data)
}

func (h *Handler) addItems(w http.ResponseWriter, r *http.Request) {

	var req *store.Auto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := h.store.AddAuto(r.Context(), req)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, data)
}

func (h *Handler) updateItems(w http.ResponseWriter, r *http.Request) {

	var req *store.Auto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}
	data, err := h.store.UpdateAuto(r.Context(), req)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, data)
}

func (h *Handler) deleteItems(w http.ResponseWriter, r *http.Request) {

	var req map[string]uint32
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.store.DeleteAuto(r.Context(), req["id"]); err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}

	response(w, http.StatusOK, "Ok")
}

func (h *Handler) notFound(w http.ResponseWriter, r *http.Request) {
	response(w, http.StatusNotFound, "Not Found")
}

func response(
	w http.ResponseWriter,
	code int,
	msg interface{},
) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Code":    code,
		"Message": msg,
	})
}
