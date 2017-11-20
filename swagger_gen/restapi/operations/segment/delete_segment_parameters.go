// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSegmentParams creates a new DeleteSegmentParams object
// with the default values initialized.
func NewDeleteSegmentParams() DeleteSegmentParams {
	var ()
	return DeleteSegmentParams{}
}

// DeleteSegmentParams contains all the bound params for the delete segment operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteSegment
type DeleteSegmentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*numeric ID of the flag
	  Required: true
	  Minimum: 1
	  In: path
	*/
	FlagID int64
	/*numeric ID of the segment
	  Required: true
	  Minimum: 1
	  In: path
	*/
	SegmentID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteSegmentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rFlagID, rhkFlagID, _ := route.Params.GetOK("flagID")
	if err := o.bindFlagID(rFlagID, rhkFlagID, route.Formats); err != nil {
		res = append(res, err)
	}

	rSegmentID, rhkSegmentID, _ := route.Params.GetOK("segmentID")
	if err := o.bindSegmentID(rSegmentID, rhkSegmentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteSegmentParams) bindFlagID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("flagID", "path", "int64", raw)
	}
	o.FlagID = value

	if err := o.validateFlagID(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteSegmentParams) validateFlagID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("flagID", "path", int64(o.FlagID), 1, false); err != nil {
		return err
	}

	return nil
}

func (o *DeleteSegmentParams) bindSegmentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("segmentID", "path", "int64", raw)
	}
	o.SegmentID = value

	if err := o.validateSegmentID(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteSegmentParams) validateSegmentID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("segmentID", "path", int64(o.SegmentID), 1, false); err != nil {
		return err
	}

	return nil
}
