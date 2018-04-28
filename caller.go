package telebot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Error represents error from Telegram Bot API.
type Error struct {
	HTTPCode    int    `json:"http_code"`
	Description string `json:"description"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.Description
}

// Caller is a low-level interface to wrap REST API request/response with http
// client into an RPC method. You should not directly use Caller from your
// application.
type Caller struct {
	prefix     string
	client     *http.Client
	pollClient *http.Client
}

// outerResponse sets a standard message formatting for result data on request
// success or error message on request fail.
type outerResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	Description string          `json:"description"`
}

func (c *Caller) do(client *http.Client, name string, request, response interface{}) error {
	// Store HTTP method and optional input reader
	var method string
	var input io.Reader
	// Check if method call has request body
	if request != nil {
		// Set HTTP method to POST
		method = "POST"
		// Marshal request into JSON and set body reader
		buf, err := json.Marshal(request)
		if err != nil {
			return err
		}
		input = bytes.NewBuffer(buf)
	} else {
		// Set HTTP method to GET
		method = "GET"
	}
	// Create HTTP request data
	req, err := http.NewRequest(method, c.prefix+name, input)
	if err != nil {
		return err
	}
	// Add content-type to header
	if request != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// Do HTTP transaction
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	// Parse outer HTTP response
	var outer outerResponse
	if err = json.NewDecoder(res.Body).Decode(&outer); err != nil {
		res.Body.Close()
		return err
	}
	res.Body.Close()
	// Check if telegram API return an error instead
	if !outer.Ok {
		if len(outer.Description) == 0 {
			return &Error{
				HTTPCode:    res.StatusCode,
				Description: "undefined error",
			}
		}
		return &Error{
			HTTPCode:    res.StatusCode,
			Description: outer.Description,
		}
	}
	// Unmarshal inner response if asked
	if response != nil {
		if err := json.Unmarshal(outer.Result, response); err != nil {
			return err
		}
	}
	// Return with no error
	return nil
}

// Call a method name to Telegram API. Request should be a struct that can be
// encoded to JSON and response should be a pointer to struct that can receive
// decoded result from JSON response. If method takes no request parameter
// and/or response data, leave them with nil value.
func (c *Caller) Call(name string, request, response interface{}) error {
	return c.do(c.client, name, request, response)
}

// Poll a method name to Telegram API. This method yields a long timeout value
// that only meant to be used for long-polling operations. Request should be a
// struct that can be encoded to JSON and response should be a pointer to struct
// that can receive decoded result from JSON response. If method takes no
// request parameter and/or response data, leave them with nil value.
func (c *Caller) Poll(name string, request, response interface{}) error {
	return c.do(c.pollClient, name, request, response)
}

// NewCaller creates new caller wraper given telegram bot API endpoint and
// token. You should not directly call NewCaller from your application.
func NewCaller(endpoint, token string) *Caller {
	return &Caller{
		prefix: endpoint + token + "/",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		pollClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}
