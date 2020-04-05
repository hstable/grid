package api

import (
	"fmt"
	"github.com/mzz2017/grip/basestation/common/multipart"
	"github.com/mzz2017/grip/basestation/service/image"
	"net/http"
	"strconv"
)

func UploadImage(file []byte, mark image.Level, deviceID int) (err error) {
	writer := multipart.NewWriter()
	defer writer.Close()
	if err = multipart.AttachFile(writer, "image", "image", file); err != nil {
		return
	}
	if err = multipart.AttachField(writer, "deviceID", strconv.Itoa(deviceID)); err != nil {
		return
	}
	if err = multipart.AttachField(writer, "mark", mark.String()); err != nil {
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/info", writer.Reader())
	if err != nil {
		fmt.Println("req err: ", err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("resp err: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("resp status: %v", resp.StatusCode)
	}
	return nil
}
