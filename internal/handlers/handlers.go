package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// для корня "/""
func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// для "/upload"
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// парсим html-форму
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error receiving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	// закрываем файл после работы с ним
	defer file.Close()

	// создаем буфер для чтения данных
	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// читаем данные из буфера в строку
	str := buf.String()
	// передаем строку в функцию определения морзе
	result, err := service.CheckAndConvert(str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// создаем локальный файл предварительно сделав name из текущего времени
	fileName := time.Now().Format("02012006-150405") + ".txt"
	dst, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "failed to create local file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	// копируем данные из строки в локальный файл
	fmt.Fprint(dst, result)
	// возвращаем результат конвертации строки
	w.Write([]byte(result))
}
