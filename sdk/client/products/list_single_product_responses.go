// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cesarvspr/golang-modules/sdk/models"
)

// ListSingleProductReader is a Reader for the ListSingleProduct structure.
type ListSingleProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSingleProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSingleProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListSingleProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSingleProductOK creates a ListSingleProductOK with default headers values
func NewListSingleProductOK() *ListSingleProductOK {
	return &ListSingleProductOK{}
}

/* ListSingleProductOK describes a response with status code 200, with default header values.

Data structure representing a single product
*/
type ListSingleProductOK struct {
	Payload *models.Product
}

func (o *ListSingleProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductOK  %+v", 200, o.Payload)
}
func (o *ListSingleProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *ListSingleProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSingleProductNotFound creates a ListSingleProductNotFound with default headers values
func NewListSingleProductNotFound() *ListSingleProductNotFound {
	return &ListSingleProductNotFound{}
}

/* ListSingleProductNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type ListSingleProductNotFound struct {
	Payload *ListSingleProductNotFoundBody
}

func (o *ListSingleProductNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductNotFound  %+v", 404, o.Payload)
}
func (o *ListSingleProductNotFound) GetPayload() *ListSingleProductNotFoundBody {
	return o.Payload
}

func (o *ListSingleProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListSingleProductNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListSingleProductNotFoundBody list single product not found body
swagger:model ListSingleProductNotFoundBody
*/
type ListSingleProductNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this list single product not found body
func (o *ListSingleProductNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list single product not found body based on context it is used
func (o *ListSingleProductNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSingleProductNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSingleProductNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ListSingleProductNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
