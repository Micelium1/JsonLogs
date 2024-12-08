package service

import (
	"JsonLogs/model"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// Post-only
func SendLogHandler(w http.ResponseWriter, r *http.Request) {
	logType := r.FormValue("log_type")
	description := r.FormValue("description")
	serviceName := "TestService"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	ip := GetIpAddress(r)
	Log := model.LogStruct{
		Date:    timestamp,
		Level:   logType,
		Message: description,
		Context: model.ContextStruct{
			Ip: ip,
		},
		Logger: serviceName,
	}
	LogJson, _ := json.Marshal(&Log)
	addToLogs("logs.json", LogJson)
	//fmt.Fprintf(w, "Лог сохранен %s - %s", Log.Level, Log.Message)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func GetIpAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	} else {
		ip = strings.Split(ip, ",")[0]
	}
	parsedIp := net.ParseIP(ip)
	if parsedIp != nil && parsedIp.IsLoopback() {
		return "127.0.0.1"
	}
	return ip
}
func addToLogs(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	dataWithNewLine := append(data, '\n')
	_, err = file.Write(dataWithNewLine)
	return err
}
