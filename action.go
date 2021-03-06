package lumada

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var loginEndpoint = "https://%v/v1/security/oauth/token"

//Data Sink
var createDataSinkEndpoint = "https://%v/v1/asset-data/sinks"

//Asset
var assetViewEndpoint = "https://%v/v1/asset-management/assets/%v"
var assetViewEventEndpoint = "https://%v/v1/asset-data/assets/%v/events?startTime=%v&endTime=%v"
var assetGetAccessTokenEndpoint = "https://%v/v1/asset-management/assets/%v/token"
var assetAddNewAvatarEndpoint = "https://%v/v1/asset-management/assets"

var trans *http.Transport
var Debug bool

func init() {
	trans = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func createEndpoint(u string, h string) string {
	apiUrl := fmt.Sprintf(u, h)
	if Debug {
		fmt.Println("api url: " + apiUrl)
	}
	return apiUrl
}

//Login to Lumada
func Login(req LoginRequest, host string) (*LoginResponse, error) {
	apiUrl := createEndpoint(loginEndpoint, host)

	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("client_id", "lumada-ui")
	form.Add("scope", "all")
	form.Add("username", req.Username)
	form.Add("password", req.Password)
	form.Add("Realm", "local")

	r, _ := http.NewRequest("POST", apiUrl, strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//r.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	client := &http.Client{Transport: trans}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err)
	}

	var loginRsep LoginResponse
	if err := json.Unmarshal(body, &loginRsep); err != nil {
		return nil, err
	}

	return &loginRsep, nil
}

func CreateDataSink(req CreateDataSinkRequest, host string, token string) (*CreateDataSinkResponse, error) {
	apiUrl := createEndpoint(createDataSinkEndpoint, host)

	body, _ := json.Marshal(req)
	r, _ := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	r.Header.Add("Content-Type", "application/json; charset=utf-8")
	r.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{Transport: trans}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err)
	}
	var createDataSinkRsep CreateDataSinkResponse
	if err := json.Unmarshal(body, &createDataSinkRsep); err != nil {
		return nil, err
	}

	return &createDataSinkRsep, nil
}
