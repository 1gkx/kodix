package router

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	store *sql.DB
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {

	rows, err := h.store.QueryContext(r.Context(), `
		SELECT * FROM auto;
	`)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	response(w, http.StatusOK, "")
}

func (h *Handler) addItems(w http.ResponseWriter, r *http.Request) {

	rows, err := h.store.QueryContext(r.Context(), `
		SELECT * FROM auto;
	`)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	response(w, http.StatusOK, "")
}

func (h *Handler) updateItems(w http.ResponseWriter, r *http.Request) {

	rows, err := h.store.QueryContext(r.Context(), `
		SELECT * FROM auto;
	`)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	response(w, http.StatusOK, "")
}

func (h *Handler) deleteItems(w http.ResponseWriter, r *http.Request) {

	rows, err := h.store.QueryContext(r.Context(), `
		SELECT * FROM auto;
	`)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	response(w, http.StatusOK, "")
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
