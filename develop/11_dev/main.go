package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type event struct {
	Day   int
	Month int
	Year  int
	Event string
}
type output struct {
	Result []event `json:"result,omitempty"`
}
type outputDay struct {
	ResultDay event `json:"result,omitempty"`
}
type resultAndError struct {
	Result string `json:"result,omitempty"`
	Err    string `json:"error,omitempty"`
}
type repo struct {
	myMap    map[string]string
	arrayDay []string
}

func makeJSON(w http.ResponseWriter, i interface{}) {
	jSon, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(jSon)
}

type repository interface {
	create(w http.ResponseWriter, evv string, eventT string)
	update(w http.ResponseWriter, evv string, eventT string)
	delete(w http.ResponseWriter, evv string)
	getForDay(w http.ResponseWriter, evv string, day, month, year int)
	getForWeek(w http.ResponseWriter, evv string)
	getForMonth(w http.ResponseWriter, month, year int)
}

type Handler struct {
	r repository
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	goodRequestBool := true
	var evv string
	start := time.Now()
	eventT := req.FormValue("event")
	evv = req.FormValue("year") + "/" + req.FormValue("month") + "/" + req.FormValue("day")
	day, _ := strconv.Atoi(req.FormValue("day"))
	month, _ := strconv.Atoi(req.FormValue("month"))
	year, _ := strconv.Atoi(req.FormValue("year"))
	if _, err := time.Parse("2006/1/2", evv); err != nil && req.URL.Path != "/events_for_month" {
		w.WriteHeader(400)
		return
	}
	switch req.URL.Path {
	case "/create_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		h.r.create(w, evv, eventT)

	case "/update_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		h.r.update(w, evv, eventT)

	case "/delete_event":
		if req.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		h.r.delete(w, evv)
	case "/events_for_day":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		h.r.getForDay(w, evv, day, month, year)

	case "/events_for_month":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		h.r.getForMonth(w, month, year)

	case "/events_for_week":
		if req.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		h.r.getForWeek(w, evv)
	default:
		http.NotFound(w, req)
		goodRequestBool = false
	}
	if goodRequestBool {
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	}
}

func main() {
	port := parseConfig("Port")
	repo := NewRepo()
	handler := &Handler{r: repo}
	fmt.Println("Start server!!!")
	log.Fatal(http.ListenAndServe(port, handler))

}

