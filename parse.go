package main

import (
	"encoding/json"
	"fmt"
)

// Define the structure of your data
type Message struct {
	MsgID   int      `json:"msgid"`
	ErrCode string   `json:"errcode"`
	Field   string   `json:"field"`
	Vals    []string `json:"vals"`
}

type ErrorData struct {
	ClientCode string `json:"client_code"`
	MemberCode string `json:"member_code"`
}

type ErrorResponse struct {
	Status   string    `json:"status"`
	Data     ErrorData `json:"data"`
	Messages []Message `json:"messages"`
}

func main() {
	// Example JSON data
	jsonData := `
	[
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018ANOTPA9001A",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018ANOTPA0002A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018ANOTPA0000A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018AEVPPS5841R",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018ANOTPA9001ANOTPA0002A",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018ANOTPA9001ANOTPA0000A",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018ANOTPA9001AEVPPS5841R",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018ANOTPA0002ANOTPA0000A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018ANOTPA0002ANOTPA9001A",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018ANOTPA0002AEVPPS5841R",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018ANOTPA0000ANOTPA0002A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018ANOTPA0000ANOTPA9001A",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018ANOTPA0000AEVPPS5841R",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "success",
        "data": {
            "client_code": "NOTPA0018AEVPPS5841RNOTPA9001A",
            "member_code": "0103",
            "status": "APPROVED"
        },
        "messages": []
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018AEVPPS5841RNOTPA0002A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    },
    {
        "status": "error",
        "data": {
            "client_code": "NOTPA0018AEVPPS5841RNOTPA0000A",
            "member_code": "0103"
        },
        "messages": [
            {
                "msgid": 539,
                "errcode": "duplicate_record",
                "field": "client and member code",
                "vals": [
                    "client and member code combination already exists in master UCC table"
                ]
            }
        ]
    }
]
	`

	// Parse the JSON
	var errorResponses []ErrorResponse
	err := json.Unmarshal([]byte(jsonData), &errorResponses)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the parsed data
	for _, response := range errorResponses {
		fmt.Println("Status:", response.Status)
		fmt.Println("Client Code:", response.Data.ClientCode)
		fmt.Println("Member Code:", response.Data.MemberCode)
		for _, msg := range response.Messages {
			fmt.Println("Message ID:", msg.MsgID)
			fmt.Println("Error Code:", msg.ErrCode)
			fmt.Println("Field:", msg.Field)
			fmt.Println("Values:", msg.Vals)
		}

	}
}
