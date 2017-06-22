package mpapi

import (
	"gopkg.in/resty.v0"
	"encoding/json"
	"github.com/burythehammer/golang-api-test-quickstart/model"
	"strings"
	"errors"
)

var baseUrl = "https://www.theyworkforyou.com/api"

func GetMp(personId string, apiKey string) (int, model.Mp, error) {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"id": personId,
		"key": apiKey,
		"output":"js",
	}).Get(baseUrl + "/getMP")

	if (strings.Contains(resp.String(), "error")) {
		return resp.StatusCode(), model.Mp{}, errors.New("Mp not found")
	}

	mps := []model.Mp{}
	err = json.Unmarshal(resp.Body(), &mps)

	return resp.StatusCode(), mps[0], err
}

func GetMpByConstituency(constituency string, apiKey string) (int, model.Mp, error) {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"key": apiKey,
		"constituency": constituency,
		"output":"js",
	}).Get(baseUrl + "/getMP")

	if (strings.Contains(resp.String(), "error")) {
		return resp.StatusCode(), model.Mp{}, errors.New("Mp not found")
	}

	mp := model.Mp{}
	err = json.Unmarshal(resp.Body(), &mp)
	return resp.StatusCode(), mp, err
}

func GetMpByPostcode(postcode string, apiKey string) (int, model.Mp, error) {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"postcode": postcode,
		"key": apiKey,
		"output":"js",
	}).Get(baseUrl + "/getMP")


	if (strings.Contains(resp.String(), "error")) {
		return resp.StatusCode(), model.Mp{}, errors.New("Mp not found")
	}

	mps := model.Mp{}
	err = json.Unmarshal(resp.Body(), &mps)
	return resp.StatusCode(), mps, err
}

func ListMps(apiKey string) (int, []model.MpShort, error) {
	resp, err := resty.R().
		SetQueryParams(map[string]string{
		"key": apiKey,
		"output":"js",
	}).Get(baseUrl + "/getMPs")

	mpList := []model.MpShort{}
	err = json.Unmarshal(resp.Body(), &mpList)
	return resp.StatusCode(), mpList, err
}