package receipt

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

//ReceiptDirectory ReceiptDirectory
var ReceiptDirectory string = filepath.Join("uploads")

//Receipt Receipt
type Receipt struct {
	ReceiptName string    `json:"name"`
	UploadDate  time.Time `json:"uploadDate"`
}

//GetReceipts GetReceipts
func GetReceipts() ([]Receipt, error) {
	receipts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptDirectory)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		receipts = append(receipts, Receipt{ReceiptName: f.Name(), UploadDate: f.ModTime()})
	}
	return receipts, nil
}
