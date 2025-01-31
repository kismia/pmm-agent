// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// StartMySQLShowTableStatusActionReader is a Reader for the StartMySQLShowTableStatusAction structure.
type StartMySQLShowTableStatusActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLShowTableStatusActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartMySQLShowTableStatusActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartMySQLShowTableStatusActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLShowTableStatusActionOK creates a StartMySQLShowTableStatusActionOK with default headers values
func NewStartMySQLShowTableStatusActionOK() *StartMySQLShowTableStatusActionOK {
	return &StartMySQLShowTableStatusActionOK{}
}

/*StartMySQLShowTableStatusActionOK handles this case with default header values.

A successful response.
*/
type StartMySQLShowTableStatusActionOK struct {
	Payload *StartMySQLShowTableStatusActionOKBody
}

func (o *StartMySQLShowTableStatusActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLShowTableStatus][%d] startMySqlShowTableStatusActionOk  %+v", 200, o.Payload)
}

func (o *StartMySQLShowTableStatusActionOK) GetPayload() *StartMySQLShowTableStatusActionOKBody {
	return o.Payload
}

func (o *StartMySQLShowTableStatusActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLShowTableStatusActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLShowTableStatusActionDefault creates a StartMySQLShowTableStatusActionDefault with default headers values
func NewStartMySQLShowTableStatusActionDefault(code int) *StartMySQLShowTableStatusActionDefault {
	return &StartMySQLShowTableStatusActionDefault{
		_statusCode: code,
	}
}

/*StartMySQLShowTableStatusActionDefault handles this case with default header values.

An error response.
*/
type StartMySQLShowTableStatusActionDefault struct {
	_statusCode int

	Payload *StartMySQLShowTableStatusActionDefaultBody
}

// Code gets the status code for the start my SQL show table status action default response
func (o *StartMySQLShowTableStatusActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLShowTableStatusActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/StartMySQLShowTableStatus][%d] StartMySQLShowTableStatusAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMySQLShowTableStatusActionDefault) GetPayload() *StartMySQLShowTableStatusActionDefaultBody {
	return o.Payload
}

func (o *StartMySQLShowTableStatusActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLShowTableStatusActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMySQLShowTableStatusActionBody start my SQL show table status action body
swagger:model StartMySQLShowTableStatusActionBody
*/
type StartMySQLShowTableStatusActionBody struct {

	// Database name. Required if not given in the table_name field.
	Database string `json:"database,omitempty"`

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`

	// Table name. Required. May additionally contain a database name.
	TableName string `json:"table_name,omitempty"`
}

// Validate validates this start my SQL show table status action body
func (o *StartMySQLShowTableStatusActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowTableStatusActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowTableStatusActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowTableStatusActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowTableStatusActionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model StartMySQLShowTableStatusActionDefaultBody
*/
type StartMySQLShowTableStatusActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start my SQL show table status action default body
func (o *StartMySQLShowTableStatusActionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowTableStatusActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowTableStatusActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowTableStatusActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLShowTableStatusActionOKBody start my SQL show table status action OK body
swagger:model StartMySQLShowTableStatusActionOKBody
*/
type StartMySQLShowTableStatusActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL show table status action OK body
func (o *StartMySQLShowTableStatusActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLShowTableStatusActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLShowTableStatusActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLShowTableStatusActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
