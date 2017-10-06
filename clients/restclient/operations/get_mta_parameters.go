package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMtaParams creates a new GetMtaParams object
// with the default values initialized.
func NewGetMtaParams() *GetMtaParams {
	var ()
	return &GetMtaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMtaParamsWithTimeout creates a new GetMtaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMtaParamsWithTimeout(timeout time.Duration) *GetMtaParams {
	var ()
	return &GetMtaParams{

		timeout: timeout,
	}
}

// NewGetMtaParamsWithContext creates a new GetMtaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMtaParamsWithContext(ctx context.Context) *GetMtaParams {
	var ()
	return &GetMtaParams{

		Context: ctx,
	}
}

/*GetMtaParams contains all the parameters to send to the API endpoint
for the get mta operation typically these are written to a http.Request
*/
type GetMtaParams struct {

	/*MtaID*/
	MtaID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get mta params
func (o *GetMtaParams) WithTimeout(timeout time.Duration) *GetMtaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mta params
func (o *GetMtaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mta params
func (o *GetMtaParams) WithContext(ctx context.Context) *GetMtaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mta params
func (o *GetMtaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithMtaID adds the mtaID to the get mta params
func (o *GetMtaParams) WithMtaID(mtaID string) *GetMtaParams {
	o.SetMtaID(mtaID)
	return o
}

// SetMtaID adds the mtaId to the get mta params
func (o *GetMtaParams) SetMtaID(mtaID string) {
	o.MtaID = mtaID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMtaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param mtaId
	if err := r.SetPathParam("mtaId", o.MtaID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}