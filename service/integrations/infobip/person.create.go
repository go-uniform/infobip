package infobip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

/* Create Person
Use this method to create a new person.

POST: /people/2/persons

Reference: https://www.infobip.com/docs/api#customer-engagement/people/create-a-new-person
*/

/*
PersonCreateRequest
*/

type PersonCreateRequest struct {
	Address            string `json:",omitempty"`
	BirthDate          string `json:",omitempty"`
	City               string `json:",omitempty"`
	ContactInformation ContactInformation
	County             string `json:",omitempty"`
	ExternalId         string `json:",omitempty"`
	FirstName          string `json:",omitempty"`
	Gender             string `json:",omitempty"`
	LastName           string `json:",omitempty"`
	MiddleName         string `json:",omitempty"`
	ProfilePicture     string `json:",omitempty"`
	Tags               string `json:",omitempty"`
	Type               string `json:",omitempty"`
}

type ContactInformation struct {
	Email []Email
	Phone []Phone
}

type Email struct {
	Address string
}

type Phone struct {
	Number string
}

/*
PersonCreateResponse
*/

type PersonCreateResponse struct {
	Address            string `json:",omitempty"`
	BirthDate          string `json:",omitempty"`
	City               string `json:",omitempty"`
	ContactInformation ContactInformation
	County             string `json:",omitempty"`
	ExternalId         string `json:",omitempty"`
	FirstName          string `json:",omitempty"`
	Gender             string `json:",omitempty"`
	LastName           string `json:",omitempty"`
	MiddleName         string `json:",omitempty"`
	ProfilePicture     string `json:",omitempty"`
	Tags               string `json:",omitempty"`
	Type               string `json:",omitempty"`
	ModifiedAt         string `json:",omitempty"`
	ModifiedFrom       string `json:",omitempty"`
	Origin             string `json:",omitempty"`
}

/*
PersonCreateError
*/

type PersonCreateErrorRequestErrorServiceException struct {
	MessageId string
	Text      string // details error description
}

type PersonCreateErrorRequestError struct {
	ServiceException PersonRemoveErrorRequestErrorServiceException
}

type PersonCreateError struct {
	RequestError PersonRemoveErrorRequestError
}

/*
Request
*/

func (i *infobip) PersonCreate(request PersonCreateRequest) PersonCreateResponse {

	/* Create Request */
	data, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	uri := fmt.Sprintf("%s/people/2/persons", strings.TrimRight(i.BaseUri, "/"))
	auth := fmt.Sprintf("App %s", i.ApiKey)

	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", auth)

	/* Execute Request */
	body, statusCode, err := executeRequest(client, req)

	/* Handle Response */
	var personCreateResponse PersonCreateResponse
	var responseErr error

	if statusCode != 200 {
		var personCreateError PersonCreateError
		if err := json.Unmarshal(body, &personCreateError); err != nil {
			responseErr = err
		} else {
			responseErr = errors.New(fmt.Sprintf("Infobip Error [%s] %s", personCreateError.RequestError.ServiceException.MessageId, personCreateError.RequestError.ServiceException.Text))
		}
	}

	if responseErr == nil {
		if err := json.Unmarshal(body, &personCreateResponse); err != nil {
			responseErr = err
		}
	}

	return personCreateResponse
}
