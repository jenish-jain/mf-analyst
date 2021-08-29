package mfapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchMFDataBytes(schemeCode int) []byte {
	baseApi := fmt.Sprintf("%s/mf/%d", "https://api.mfapi.in", schemeCode) //TODO: move to constants or config
	request, err := http.NewRequest(http.MethodGet, baseApi, nil)
	if err != nil {
		log.Printf("error fetching mf data - %v\n", err)
	}
	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("error making request to mf api server - %v\n", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error reading mf data - %v\n", err)
	}

	return responseBytes
}

func GetMfData(schemeCode int) MFData {
	mfData := MFData{}
	mfDataBytes := fetchMFDataBytes(schemeCode)
	err := json.Unmarshal(mfDataBytes, &mfData)

	if err != nil {
		log.Printf("error unmarshaling mf data bytes -  %v\n", err)
	}

	return mfData
}
