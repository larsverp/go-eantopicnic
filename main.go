package eantopicnic

import (
	"net/http"
	"strings"
)

func GetPicnicID(ean string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://picnic.app/qr/gtin/"+ean, nil)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	splitedUrl := strings.Split(resp.Request.URL.String(), "id=")
	id := splitedUrl[len(splitedUrl)-1]

	return id, nil
}
