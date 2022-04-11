package deepgram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type invalidResponse struct {
	Error string `json:"error"`
}

var (
	dg Deepgram = Deepgram{
		ApiKey: "123-abc-456-def",
	}
	respSuccess CheckBalanceResponse = CheckBalanceResponse{
		Balance: 123.0,
		UserId:  "123-abc-456-def",
	}
	invalidRes invalidResponse = invalidResponse{
		Error: "invalid contentID / userID",
	}
	goodJson  string = "{\"balance\":123,\"userID\":\"123-abc-456-def\"}"
	errorJson string = "{\"error\":\"invalid contentID / userID\"}"
	badJson   string = "{"
)

func TestMakeRequestSuccess(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(respSuccess)
		fmt.Fprint(w, string(res))
	})
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	payload := checkBalanceRequest{
		Action: "get_balance",
		UserId: dg.ApiKey,
	}
	response, err := makeRequest(testServer.URL, payload)
	if err != nil {
		t.Error("Got error", err, "expected", nil)
	}
	respSuccessBytes, _ := json.Marshal(respSuccess)
	if string(respSuccessBytes) != string(response) {
		t.Error("Expected", string(respSuccessBytes), "Got", string(response))
	}
}

func TestMakeRequestNilResponse(t *testing.T) {

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, nil)
	})
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	payload := checkBalanceRequest{
		Action: "get_balance",
		UserId: dg.ApiKey,
	}
	response, err := makeRequest(testServer.URL, payload)
	if err != nil {
		t.Error("Got error", err, "expected", nil)
	}
	if string(response) == "" {
		t.Error("Expected", nil, "Got", string(response))
	}
}

func TestParseResponseSuccess(t *testing.T) {
	response := new(CheckBalanceResponse)
	err := parseResponse([]byte(goodJson), response)
	if err != nil {
		t.Error("Got error", err, "expected", nil)
	}
	if !reflect.DeepEqual(*response, respSuccess) {
		t.Error("Expected", respSuccess, "Got", *response)
	}
}

func TestParseResponseInvalidJsonResponse(t *testing.T) {
	response := new(CheckBalanceResponse)
	err := parseResponse([]byte(badJson), response)
	if err == nil {
		t.Error("Expecting error", err, "got", nil)
	}
	if *response != *(new(CheckBalanceResponse)) {
		t.Error("Expected", *(new(CheckBalanceResponse)), "Got", *response)
	}
}

func TestParseResponseErrorResponse(t *testing.T) {
	response := new(CheckBalanceResponse)
	err := parseResponse([]byte(errorJson), response)
	if err == nil {
		t.Error("Expecting error", err, "got", nil)
	}
	if *response != *(new(CheckBalanceResponse)) {
		t.Error("Expected", *(new(CheckBalanceResponse)), "Got", *response)
	}
}
