package service

import (
	"JsonLogs/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func LogViewHandler(w http.ResponseWriter, r *http.Request) {
	levelFilter := r.URL.Query().Get("level")
	logger := r.URL.Query().Get("service")
	ip := r.URL.Query().Get("ip")
	from := r.URL.Query().Get("from_date")
	to := r.URL.Query().Get("to_date")

	logs, err := readLogs("logs.json")
	if err != nil {
		http.Error(w, "Ошибка чтения логов", http.StatusInternalServerError)
		return
	}

	logs = filterLogs(&logs, from, to, levelFilter, ip, logger)

	tmpl, err := template.ParseFiles("web/logView.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, logs)
}

func readLogs(filename string) ([]model.LogStruct, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var logs []model.LogStruct
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		var log model.LogStruct
		if err := json.Unmarshal([]byte(line), &log); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}

func filterLogs(logs *[]model.LogStruct, from, to, level, ip, logger string) []model.LogStruct {
	var filteredLogs []model.LogStruct
	for _, log := range *logs {
		if !(from == "" || filterFrom(&log, from)) {
			continue
		}
		if !(to == "" || filterTo(&log, to)) {
			continue
		}
		if !(level == "" || filterByLevel(&log, level)) {
			continue
		}
		if !(ip == "" || filterByIp(&log, ip)) {
			continue
		}
		if !(logger == "" || filterByLogger(&log, logger)) {
			continue
		}

		// if (level == "" || log.Level == level) &&
		// 	(logger == "" || log.Logger == logger) {
		filteredLogs = append(filteredLogs, log)
		// }
	}
	return filteredLogs
}
func filterByLevel(log *model.LogStruct, level string) bool {
	return log.Level == level
}
func filterByLogger(log *model.LogStruct, logger string) bool {
	return strings.HasPrefix(log.Logger, logger)
}
func filterByIp(log *model.LogStruct, ip string) bool {
	return strings.HasPrefix(log.Context.Ip, ip)
}
func filterFrom(log *model.LogStruct, from string) bool {
	logDate, err := time.Parse("2006-01-02 15:04:05", log.Date)
	if err != nil {
		return false
	}
	fromDate, err1 := time.Parse("2006-01-02", from)
	if err1 != nil || logDate.Before(fromDate) {
		return false
	}
	return true
}
func filterTo(log *model.LogStruct, to string) bool {
	logDate, err := time.Parse("2006-01-02 15:04:05", log.Date)
	if err != nil {
		return false
	}

	ToDate, err2 := time.Parse("2006-01-02", to)

	if err2 != nil {
		ToDate = time.Time{}
	} else {
		ToDate = ToDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second) // Конец дня
	}
	if logDate.After(ToDate) {
		return false
	}
	return true
}
