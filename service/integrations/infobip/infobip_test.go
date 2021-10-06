package infobip

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInfobip_SendEmailSuccess(t *testing.T) {
	/* Expected */
	injectedResponseBody := `
{
	"messages": [
		{
			"to": "joan.doe0@example.com",
			"messageCount": 1,
			"messageId": "somexternalMessageId0",
			"status": {
			"groupId": 1,
			"groupName": "PENDING",
			"id": 7,
			"name": "PENDING_ENROUTE",
			"description": "Message sent to next instance"
		}
	]
}
`

	/* Prepare */
	defer func() {
		// restore orig
	}()
	// inject scenario
	var requestBody []byte
	var err error
	executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
		// extract request body content for assertion
		defer req.Body.Close()
		requestBody, err = ioutil.ReadAll(req.Body)
		if err != nil {
			t.Error(err)
		}

		return []byte(injectedResponseBody), 200, nil
	}

	/* Execute */
	instance := NewInfobipConnector("", "")
	response := instance.SendEmail(EmailSendRequest{
		To:      []string{"joan.doe0@example.com"},
		From:    "noreply@uniform.co.za",
		Subject: "Test",
	})

	/* Assert */
	if response.BulkId != "" {
		// todo: error
	}
	if response.Messages == nil || len(response.Messages) != 1 {
		// todo: error
	}
}

func TestInfobip_SendEmailFail(t *testing.T) {
	//todo: write test
}

func TestInfobip_SmsTextAdvancedSuccess(t *testing.T) {
	//todo: write test
}

func TestInfobip_PersonCreateSuccess(t *testing.T) {
	//todo: write test
}

func TestInfobip_PersonCreateFail(t *testing.T) {
	//todo: write test
}

func TestInfobip_PersonUpdateSuccess(t *testing.T) {
	//todo: write test
}

func TestInfobip_PersonUpdateFail(t *testing.T) {
	//todo: write test
}

func TestInfobip_PersonRemoveSuccess(t *testing.T) {
	//todo: write test
}

func TestInfobip_PersonRemoveFail(t *testing.T) {
	//todo: write test
}
