package handler

import (
	"database/sql"
	"management/internal/http/helper"
	"net/http"
	"os"

	repo "management/internal/repository/ticket"
)

type InsertTicketHandler struct {
	Db *sql.DB
}

type UpdateTicketHandler struct {
	Db *sql.DB
}

type ShowTicketHandler struct {
	Db *sql.DB
}

type TicketHandler struct {
	Db *sql.DB
}


func (h *InsertTicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := helper.DecodeJSONRequest(w, r)
	err := repo.InsertTicket(h.Db, params)
	msg := make(map[string]interface{}, 1)
	statusCode := 200
	msg["message"] = "Insert success"

	if err != nil {
		msg["message"] = err.Error()
		statusCode = 500
	}

	helper.ReturnJsonData(w, statusCode, msg)
}

func (h *UpdateTicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := helper.DecodeJSONRequest(w, r)
	err := repo.UpdateTicket(h.Db, params)
	msg := make(map[string]interface{}, 1)
	statusCode := 200
	msg["message"] = "Edit success"

	if err != nil {
		msg["message"] = err.Error()
		statusCode = 500
	}

	helper.ReturnJsonData(w, statusCode, msg)
}



func (h *TicketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// params := helper.DecodeJSONRequest(w, r)
	// data, totalRows, err := repo.GetAllTickets(h.Db, params)
	testenv:= os.Getenv("NAME_DEVELOPER")
	res := make(map[string]interface{}, 1)
	statusCode := 200
	res["data"] = testenv


	helper.ReturnJsonData(w, statusCode, res)
}


