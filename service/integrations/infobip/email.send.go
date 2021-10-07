package infobip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-diary/diary"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

/* Send Email
This is used to send transactional emails, but in some cases for marketing traffic also. With HTTP API you can set up a
connection quickly and start sending and tracking.

POST: /email/2/send

Reference: https://www.infobip.com/docs/api#channels/email/send-email
*/

/*
EmailSendRequest
*/

type EmailSendRequest struct {
	From               string
	To                 []string
	Cc                 []string `json:",omitempty"`
	Bcc                []string `json:",omitempty"`
	Subject            string
	Text               string       `json:",omitempty"`
	BulkId             string       `json:",omitempty"`
	MessageId          string       `json:",omitempty"`
	TemplateId         int          `json:",omitempty"`
	Attachments        []Attachment `json:",omitempty"`
	InlineImage        []Image      `json:",omitempty"`
	Html               string       `json:",omitempty"`
	ReplyTo            string       `json:",omitempty"`
	DefaultPlaceholder string       `json:",omitempty"`
	PreserveRecipients string       `json:",omitempty"`
	TrackingUrl        string       `json:",omitempty"`
	TrackClicks        bool         `json:",omitempty"`
	TrackOpens         bool         `json:",omitempty"`
	Track              bool         `json:",omitempty"`
	CallbackData       string       `json:",omitempty"`
	IntermediateReport string       `json:",omitempty"`
	NotifyUrl          string       `json:",omitempty"`
	NotifyContentType  string       `json:",omitempty"`
	SendAt             string       `json:",omitempty"`
}

type Attachment struct {
	ContentType string `json:",omitempty"`
	Filename    string `json:",omitempty"`
	Data        string `json:",omitempty"`
}

type Image struct {
	ContentType string `json:",omitempty"`
	Filename    string `json:",omitempty"`
	Data        string `json:",omitempty"`
}

/*
EmailSendResponse
*/

type EmailSendResponse struct {
	BulkId   string
	Messages []EmailSendResponseMessage
}

type EmailSendResponseMessage struct {
	MessageCount int32
	MessageId    string
	Status       EmailSendResponseMessageStatus
	To           string
}

type EmailSendResponseMessageStatus struct {
	Action      string
	Description string
	GroupId     int32
	GroupName   string
	Id          int32
	Name        string
}

/*
EmailSendError
*/

type EmailSendErrorRequestErrorServiceException struct {
	MessageId string
	Text      string // details error description
}

type EmailSendErrorRequestError struct {
	ServiceException EmailSendErrorRequestErrorServiceException
}

type EmailSendError struct {
	RequestError EmailSendErrorRequestError
}

/*
Request
*/

func (i *infobip) SendEmail(request EmailSendRequest) EmailSendResponse {
	/* Create Request */
	bodyMultipart := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyMultipart)
	if err := writer.WriteField("from", request.From); err != nil {
		panic(err)
	}
	if err := writer.WriteField("subject", request.From); err != nil {
		panic(err)
	}
	for _, to := range request.To {
		fw, err := writer.CreateFormField("to")
		if err != nil {
		}
		_, err = io.Copy(fw, strings.NewReader(to))
		if err != nil {
			panic(err)
		}
	}

	uri := fmt.Sprintf("%s/email/2/send", strings.TrimRight(i.BaseUri, "/"))
	auth := fmt.Sprintf("App %s", i.ApiKey)

	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, bodyMultipart)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", auth)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Accept", "application/json")

	i.Page.Debug("infobip.send-email", diary.M{
		"method": "POST",
		"uri":    uri,
		"body":   request,
		"headers": map[string][]string{
			"Authorization": {auth},
		},
	})

	/* Execute Request */
	body, statusCode, err := executeRequest(client, req)

	/* Handle Response */
	var emailSendResponse EmailSendResponse
	var responseErr error

	if statusCode != 200 {
		var emailSendError EmailSendError
		if err := json.Unmarshal(body, &emailSendError); err != nil {
			responseErr = err
		} else {
			responseErr = errors.New(fmt.Sprintf("Infobip Error [%s] %s", emailSendError.RequestError.ServiceException.MessageId, emailSendError.RequestError.ServiceException.Text))
		}
	}

	if responseErr == nil {
		if emailSendResponse.Messages == nil || len(emailSendResponse.Messages) <= 0 {
			responseErr = errors.New("empty response received from Infobip")
		}
	}

	if responseErr == nil {
		if err := json.Unmarshal(body, &emailSendResponse); err != nil {
			responseErr = err
		}
	}

	return emailSendResponse
}
