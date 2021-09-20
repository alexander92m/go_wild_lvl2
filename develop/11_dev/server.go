package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/joho/godotenv"
	"net/http"
	"os"
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

func main() {
	port := parseConfig("Port")
	repo := NewRepo()
	handler := &Handler{r: repo}
	fmt.Println("Start server!!!")
	log.Fatal(http.ListenAndServe(port, handler))

}
//аргументы из конфига
func parseConfig(value string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error", err)
	}
	return os.Getenv(value)
}
//создается указатель на хранилище
func NewRepo() *repo {
	return &repo{
		myMap: make(map[string]string),
		arrayDay: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
			"15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
	}

}

func makeJSON(w http.ResponseWriter, i interface{}) {
	jSon, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(jSon)
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
	fmt.Println(req.URL.Path, "PRIVET|", req.URL.ForceQuery, "|", req.URL.RawQuery[:len(req.URL.RawQuery) - 1])
	if _, err := time.Parse("2006/1/2", evv); err != nil && req.URL.Path != "/events_for_month" {
		w.WriteHeader(400)
		return
	}
	fmt.Println(req.URL.Path, "PRIVET2|", req.URL.ForceQuery, "|", req.URL.RawQuery[:len(req.URL.RawQuery) - 1])
	switch req.URL.Path {
	case "/create_event":
		fmt.Println(req.URL.Path, "PRIVET3|", req.URL.ForceQuery, "|", req.URL.RawQuery[:len(req.URL.RawQuery) - 1])
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




func (r *repo) create(w http.ResponseWriter, evv string, eventT string) {
	r.myMap[evv] = eventT
	result := resultAndError{Result: "Событие создано успешно!"}
	makeJSON(w, result)
}
func (r *repo) update(w http.ResponseWriter, evv string, eventT string) {
	var result resultAndError
	_, ok := r.myMap[evv]
	if ok {
		r.myMap[evv] = eventT
		result = resultAndError{Result: "Событие обновлено успешно!"}
	} else {
		result = resultAndError{Err: "Значение не найдено!"}
	}
	makeJSON(w, result)
}
func (r *repo) delete(w http.ResponseWriter, evv string) {
	var result resultAndError
	_, ok := r.myMap[evv]
	if ok {
		delete(r.myMap, evv)
		result = resultAndError{Result: "Событие удалено успешно!"}
	} else {
		result = resultAndError{Err: "Значение не найдено!"}
	}
	makeJSON(w, result)
}
func (r *repo) getForDay(w http.ResponseWriter, evv string, day, month, year int) {
	value, ok := r.myMap[evv]
	if ok {
		newEvent := event{Day: day, Month: month, Year: year, Event: value}
		newOutput := outputDay{ResultDay: newEvent}
		makeJSON(w, newOutput)
	} else {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
	}
}
func (r *repo) getForMonth(w http.ResponseWriter, month, year int) {
	var events []event
	for _, vvv := range r.arrayDay {
		value, ok := r.myMap[fmt.Sprintf("%d/%d/%s", year, month, vvv)]
		vv, _ := strconv.Atoi(vvv)
		if ok {
			newEvent := event{Day: vv, Month: month, Year: year, Event: value}
			events = append(events, newEvent)
		}
	}
	if len(events) == 0 {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
		return
	}
	NewOutput := output{Result: events}
	makeJSON(w, NewOutput)
}
func (r *repo) getForWeek(w http.ResponseWriter, evv string) {
	var events []event
	layout := "2006/1/2"
	t, err := time.Parse(layout, evv)
	if err != nil {
		fmt.Printf("%v", err)
	}
	nDay := int(t.Weekday())
	if nDay == 0 {
		nDay = 7
	}
	for i := 1 - nDay; i <= 7-nDay; i++ {
		time1 := t.AddDate(0, 0, i)
		value, ok := r.myMap[fmt.Sprintf("%d/%d/%d", time1.Year(), time1.Month(), time1.Day())]
		if ok {
			newEvent := event{Day: time1.Day(), Month: int(time1.Month()), Year: time1.Year(), Event: value}
			events = append(events, newEvent)
		}
	}
	if len(events) == 0 {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
		return
	}
	NewOutput := output{Result: events}
	makeJSON(w, NewOutput)

}
