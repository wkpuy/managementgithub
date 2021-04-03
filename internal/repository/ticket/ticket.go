package product

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"management/internal/http/helper"
)

//Tickets struct from table tickets
type Tickets struct {
	Row 		 int       `json:"row"`
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Detail       string    `json:"details"`
	ContactInfo  string    `json:"contact_info"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	TicketStatus  string    `json:"ticket_status"`
	//Status            string               `json:"status"`
}

//InsertTicket is create ticket
func InsertTicket(db *sql.DB, params map[string]interface{}) error {
	fields := []string{
		"name",
	    "detail",
		"contact_info",
	}
	_, err := helper.InsertGetId(db, "tickets", params, fields)

	if err != nil {
		return err
	}
	return err
}

//UpdateTicket with name,detail,contact_info and status_ticket(pending,accepted,resolved,rejected)
func UpdateTicket(db *sql.DB, params map[string]interface{}) error {
	condKey := []string{"id"}
	ignoredKey := []string{
		"id",
		"created_at",
		"updated_at",
		"status",
	}

	if err := helper.UpdateData(db, "tickets", params, ignoredKey, condKey); err != nil {
		return err
	}

	return nil
}

//GetAllTickets is get all tickets filter with status_ticket and duration created_at
func GetAllTickets(db *sql.DB, params map[string]interface{}) ([]Tickets, int, error) {
	items := []Tickets{}
	fields := []string{
		"tickets.id",
		"tickets.name",
		"tickets.detail",
		"tickets.contact_info",
		"tickets.created_at",
		"tickets.updated_at",
		"tickets.ticket_status",
		//"tickets.status",
	}
	str := strings.Join(fields, ",")

	condKey := []string{"t.ticket_status"}
	extracond, condValue := helper.MakeCondStatement(condKey, params["filterObj"].(map[string]interface{}))
	cond := " tickets.status = 'A' "
	if extracond != "" {
		cond += " AND " + extracond
	}
	filterObj := params["filterObj"].(map[string]interface{})
	
	if filterObj["start_at"] != "" {
		cond += " AND tickets.created_at >= ? "
		condValue = append(condValue,helper.GetDateTimeSQL(filterObj["start_at"].(string)))
	}

	if filterObj["end_at"] != "" {

		cond += " AND tickets.created_at <= ? "
		condValue = append(condValue, helper.GetDateTimeSQL(filterObj["end_at"].(string)))

	}

	sql := `SELECT ` + str + ` FROM tickets WHERE `+ cond
	sql += " ORDER BY " + params["orderBy"].(string) + " " + params["orderType"].(string)
	sql += " LIMIT ?, ? "

	

	totalRows := helper.CountRows(db, "tickets", cond, condValue)

	pageSize := params["pageSize"].(float64)
	page := params["page"].(float64)
	offset := (pageSize * page) - pageSize
	

	condValue = append(condValue, offset, pageSize)
	//fmt.Println(sql,condValue)

	rows, err := db.Query(sql, condValue...)
	if err != nil {
		return nil,0, err
	}

	defer rows.Close()
    row :=1;
	for rows.Next() {
		b := Tickets{}
		if err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.Detail,
			&b.ContactInfo,
			&b.CreatedAt,
			&b.UpdatedAt,
			&b.TicketStatus,
		); err != nil {
			return nil, 0, err
		}
		fmt.Println(b.CreatedAt);
		b.Row = row
		items = append(items, b)
		row++;
	}
	return items, totalRows, nil
}


