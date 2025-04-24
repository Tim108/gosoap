package gosoap

import (
	"encoding/xml"
	"fmt"
)

// Response Soap Response
type Response struct {
	Body    []byte
	Header  []byte
	Payload []byte
}

// Unmarshal get the body and unmarshal into the interface
func (r *Response) Unmarshal(v interface{}) error {
	if len(r.Body) == 0 {
		return fmt.Errorf("response body is empty")
	}

	var fault Fault
	err := xml.Unmarshal(r.Body, &fault)
	if err != nil {
		return fmt.Errorf("error unmarshalling the body to Fault: %v", err.Error())
	}

	if fault.Code != "" {
		return fault
	}

	return xml.Unmarshal(r.Body, v)
}
