package internal

import (
	"net/http"
	"os"
)

type FileHandler struct {
  http.Handler
}

func NewFileHandler() *FileHandler {
  return &FileHandler{}
}

func (h *FileHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  var err error 
  filename := req.URL.Path 
  file, err := os.ReadFile(filename) // Naughty
  if err != nil {
    res.WriteHeader(http.StatusBadRequest)
    res.Write([]byte(""))
  }

  res.Write(file)
}
