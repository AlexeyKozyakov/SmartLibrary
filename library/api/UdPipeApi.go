package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const apiUrl = "https://lindat.mff.cuni.cz/services/udpipe/api/process"

type ProcessResponse struct {
	model            string
	acknowledgements []string
	Result           string
}

func Process(data string) (ProcessResponse, error) {
	requestUrl, _ := url.Parse(apiUrl)
	requestUrl.RawQuery = url.Values{
		"tokenizer": {""},
		"tagger":    {""},
		"parser":    {""},
		"data":      {data},
	}.Encode()
	res, err := http.Get(requestUrl.String())
	if err != nil {
		return ProcessResponse{}, err
	}
	var response ProcessResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ProcessResponse{}, err
	}
	return response, nil
}
