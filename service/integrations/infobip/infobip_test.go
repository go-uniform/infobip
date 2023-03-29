package infobip

import (
    "github.com/go-diary/diary"
    "net/http"
    "testing"
)

var d diary.IDiary = diary.Dear("client", "project", "appNmae", nil, "repo", "commit", []string{"version"}, nil, diary.LevelTrace, func(log diary.Log) {
})
var p diary.IPage

func init() {
    d.Page(-1, 1000, true, "appNmae", nil, "", "", nil, func(subPage diary.IPage) {
        p = subPage
    })
}
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
    executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
        return []byte(injectedResponseBody), 200, nil
    }

    /* Execute */
    instance := NewInfobipConnector(p, "", "", false)
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

func TestInfobip_SmsTextAdvancedSuccess(t *testing.T) {
    /* Expected */
    injectedResponseBody := `
{
	"messages": 
		[
			{
				"to": "1234",
			}
		]
}
`

    /* Prepare */
    defer func() {
        // restore orig
    }()
    // inject scenario
    executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
        return []byte(injectedResponseBody), 200, nil
    }

    /* Execute */
    instance := NewInfobipConnector(p, "", "", false)
    response := instance.SmsTextAdvanced(SmsTextAdvanceRequest{
        Messages: []SmsTextAdvanceRequestMessage{
            {
                Destinations: []SmsTextAdvanceRequestMessageDestination{
                    {
                        To: "1234",
                    },
                },
            },
        },
    })

    /* Assert */
    if response.BulkId != "" {
        // todo: error
    }
    if response.Messages == nil || len(response.Messages) != 1 {
        // todo: error
    }
}

func TestInfobip_PersonCreateSuccess(t *testing.T) {
    /* Expected */
    injectedResponseBody := `
{
	"contactInformation": {
		"email": [
			{
				"address": "janewilliams@acme.com"
			}
		]
	}
}
`

    /* Prepare */
    defer func() {
        // restore orig
    }()
    // inject scenario
    executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
        return []byte(injectedResponseBody), 200, nil
    }

    /* Execute */
    instance := NewInfobipConnector(p, "", "", false)
    response := instance.PersonCreate(PersonCreateRequest{
        ContactInformation: ContactInformation{
            Email: []Email{
                {
                    Address: "janewilliams@acme.com",
                },
            },
        },
    })

    /* Assert */
    if response.ContactInformation.Email == nil {
        // todo: error
    }
}

func TestInfobip_PersonUpdateSuccess(t *testing.T) {
    //todo: write test
}

func TestInfobip_PersonRemoveSuccess(t *testing.T) {
    /* Expected */
    injectedResponseBody := `
{
	"phone": "1234"
}
`
    /* Prepare */
    defer func() {
        // restore orig
    }()
    // inject scenario
    executeRequest = func(client *http.Client, req *http.Request) ([]byte, int, error) {
        return []byte(injectedResponseBody), 200, nil
    }

    /* Execute */
    instance := NewInfobipConnector(p, "", "", false)
    response := instance.PersonRemove(PersonRemoveQueryRequest{
        Phone: "1234",
    })

    /* Assert */
    if response != "" {
        // todo: error
    }
}
