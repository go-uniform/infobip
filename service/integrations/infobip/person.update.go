package infobip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-diary/diary"
	"net/http"
	"net/url"
	"strings"
)

/* Update Person
Use this method to update person data.

PUT: /people/2/persons

Reference: https://www.infobip.com/docs/api#customer-engagement/people/update-a-person
*/

/*
PersonUpdateRequest
*/

type PersonUpdateQueryRequest struct {
	Phone              string `json:",omitempty"`
	Email              string `json:",omitempty"`
	ExternalId         string `json:",omitempty"`
	PushRegistrationId string `json:",omitempty"`
	Id                 int64  `json:",omitempty"`
	LineId             string `json:",omitempty"`
	LineSenderId       int64  `json:",omitempty"`
	FacebookId         string `json:",omitempty"`
	FacebookSenderId   string `json:",omitempty"`
	TelegramUserId     string `json:",omitempty"`
	TelegramBotId      int64  `json:",omitempty"`
	ViberBotUserId     string `json:",omitempty"`
	ViberBotId         int64  `json:",omitempty"`
	InstagramId        string `json:",omitempty"`
	InstagramSenderId  int64  `json:",omitempty"`
	TwitterId          string `json:",omitempty"`
	TwitterSenderId    int64  `json:",omitempty"`
}

type PersonUpdateRequest struct {
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

/*
PersonUpdateResponse
*/

type PersonUpdateResponse struct {
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
PersonUpdateError
*/

type PersonUpdateErrorRequestErrorServiceException struct {
	MessageId string
	Text      string // details error description
}

type PersonUpdateErrorRequestError struct {
	ServiceException PersonRemoveErrorRequestErrorServiceException
}

type PersonUpdateError struct {
	RequestError PersonRemoveErrorRequestError
}

/*
Request
*/

func (i *infobip) PersonUpdate(request PersonUpdateRequest, queryRequest PersonUpdateQueryRequest) PersonUpdateResponse {

	/* Create Request */
	uri := fmt.Sprintf("%s/people/2/persons?email=%s", strings.TrimRight(i.BaseUri, "/"), url.QueryEscape(queryRequest.Email))
	auth := fmt.Sprintf("App %s", i.ApiKey)

	data, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", uri, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", auth)

	i.Page.Debug("infobip.person-update", diary.M{
		"method": "POST",
		"uri":    uri,
		"body":   request,
		"headers": map[string][]string{
			"Authorization": {auth},
		},
	})

	/* Handle Response */
	var personUpdateResponse PersonUpdateResponse
	var responseErr error

	if i.Virtual {
		printRequest(client, req)
		return personUpdateResponse
	}

	/* Execute Request */
	body, statusCode, err := executeRequest(client, req)

	if statusCode != 200 {
		var personUpdateError PersonUpdateError
		if err := json.Unmarshal(body, &personUpdateError); err != nil {
			responseErr = err
		} else {
			responseErr = errors.New(fmt.Sprintf("Infobip Error [%s] %s", personUpdateError.RequestError.ServiceException.MessageId, personUpdateError.RequestError.ServiceException.Text))
		}
	}

	if responseErr == nil {
		if err := json.Unmarshal(body, &personUpdateResponse); err != nil {
			responseErr = err
		}
	}

	return personUpdateResponse
}
