package infobip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/* Remove Person
Use this method to delete a person.

DELETE: /people/2/persons

Reference: https://www.infobip.com/docs/api#customer-engagement/people/delete-a-person
*/

/*
PersonRemoveRequest
*/

type PersonRemoveQueryRequest struct {
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

/*
PersonRemoveResponse
*/

type PersonRemoveResponse string

/*
PersonRemoveError
*/

type PersonRemoveErrorRequestErrorServiceException struct {
	MessageId string
	Text      string // details error description
}

type PersonRemoveErrorRequestError struct {
	ServiceException PersonRemoveErrorRequestErrorServiceException
}

type PersonRemoveError struct {
	RequestError PersonRemoveErrorRequestError
}

/*
Request
*/

func (i *infobip) PersonRemove(request PersonRemoveQueryRequest) PersonRemoveResponse {

	uri := fmt.Sprintf("%s/people/2/persons?email=%s", strings.TrimRight(i.BaseUri, "/"), url.QueryEscape(request.Email))
	auth := fmt.Sprintf("App %s", i.ApiKey)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", auth)

	body := make([]byte, 0)
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var personRemoveResponse PersonRemoveResponse
	var responseErr error

	if res.StatusCode != 200 {
		var personRemoveError PersonRemoveError
		if err := json.Unmarshal(body, &personRemoveError); err != nil {
			responseErr = err
		} else {
			responseErr = errors.New(fmt.Sprintf("Infobip Error [%s] %s", personRemoveError.RequestError.ServiceException.MessageId, personRemoveError.RequestError.ServiceException.Text))
		}
	}

	if responseErr == nil {
		if err := json.Unmarshal(body, &personRemoveResponse); err != nil {
			responseErr = err
		}
	}

	return personRemoveResponse
}
