package Services

import (
	"IntegracionDeDatos/Domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetBusLines(busStopID int) ([]Domain.BusLine, error) {
	busToken, err := getBusToken()
	if err != nil {
		return nil, err
	}
	url := "https://api.montevideo.gub.uy/api/transportepublico/buses/busstops/" + strconv.Itoa(busStopID) + "/lines"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Printf("el error al hacer el request es: %v \n", err)
		return nil, err
	}
	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Authorization", "Bearer "+busToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("el error del response es: %v \n", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))
	var busLines []Domain.BusLine
	errUnmarshal := json.Unmarshal(body, &busLines)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	return busLines, nil
}

func getBusToken() (string, error) {

	type ResponseToken struct {
		AccessToken string `json:"access_token"`
	}
	//Get bus token
	url := "https://mvdapi-auth.montevideo.gub.uy/auth/realms/pci/protocol/openid-connect/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials&client_id=629152ca&client_secret=a36ce14498c2667b9da4918e74fa2d4b")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var resp ResponseToken
	errUnmarshal := json.Unmarshal(body, &resp)

	if errUnmarshal != nil {
		fmt.Println(errUnmarshal)
		return "", errUnmarshal
	}

	return resp.AccessToken, nil
}
