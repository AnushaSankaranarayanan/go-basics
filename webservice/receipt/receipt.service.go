package receipt

import (
	"github.com/AnushaSankaranarayanan/inventoryservice/cors"
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const receiptPath = "receipts"

func handleReceipts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		receiptList, err := GetReceipts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		j, err := json.Marshal(receiptList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		r.ParseMultipartForm(5 << 20) //5MB
		file, handler, err := r.FormFile("receipt")
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		f, err := os.OpenFile(filepath.Join(ReceiptDirectory, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		defer f.Close()
		io.Copy(f, file)
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}
}

// SetupRoutes  SetupRoutes
func SetupRoutes(apiBasePath string){
	receiptHandler := http.HandlerFunc(handleReceipts)
	http.Handle(fmt.Sprintf("%s/%s",apiBasePath,receiptPath),cors.MiddlewareHandler(receiptHandler))
}