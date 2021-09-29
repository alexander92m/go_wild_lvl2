package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/joho/godotenv"
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
	evv = req.FormValue("date")
	day, _ := strconv.Atoi(req.FormValue("date")[8:10])
	month, _ := strconv.Atoi(req.FormValue("date")[5:7])
	year, _ := strconv.Atoi(req.FormValue("date")[:4])
	if _, err := time.Parse("2006-1-2", evv); err != nil && req.URL.Path != "/events_for_month" {
		w.WriteHeader(400)
		return
	}
	strings.ReplaceAll(evv, "-", "/")
	fmt.Println(req.URL.Path)
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

func parseConfig(value string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error", err)
	}
	return os.Getenv(value)
}
func NewRepo() *repo {
	return &repo{
		myMap: make(map[string]string),
		arrayDay: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
			"15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"},
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
		if len(vvv) < 2 {
			vvv = "0" + vvv
		}
		m := strconv.Itoa(month)
		if len(m) < 2 {
			m = "0" + m
		}
		fmt.Printf("%d-%d-%s\n", year, month, vvv)
		value, ok := r.myMap[fmt.Sprintf("%d-%v-%v", year, m, vvv)]
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

//month_day() фикс для приведения к нормальному формату
func month_day(time1 time.Time) (string, string){
	month	:= ""
	day 	:= ""
	if int(time1.Month()) < 10 {
		month = "0"+ strconv.Itoa(int(time1.Month()))
	} else {
		month = strconv.Itoa(int(time1.Month()))
	}
	if int(time1.Day()) < 10 {
		day = "0"+ strconv.Itoa(int(time1.Day()))
	} else {
		day = strconv.Itoa(int(time1.Day()))
	}
	return month, day
}

func (r *repo) getForWeek(w http.ResponseWriter, evv string) {
	var events []event
	layout := "2006-1-2"
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
		y := time1.Year()
		m, d := month_day(time1)
		value, ok := r.myMap[fmt.Sprintf("%d-%v-%v", y, m, d)]
		if ok {
			newEvent := event{Day: time1.Day(), Month: int(time1.Month()), Year: time1.Year(), Event: value}
			events = append(events, newEvent)
		}
	}
	if len(events) == 0 {
		result := resultAndError{Err: "Значение не найдено!"}
		makeJSON(w, result)
	} else {
		NewOutput := output{Result: events}
		makeJSON(w, NewOutput)

	}
}

func main() {
	port := parseConfig("Port")
	repo := NewRepo()
	handler := &Handler{r: repo}
	fmt.Println("Start server!!!")
	log.Fatal(http.ListenAndServe(port, handler))

}