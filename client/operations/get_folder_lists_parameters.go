// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFolderListsParams creates a new GetFolderListsParams object
// with the default values initialized.
func NewGetFolderListsParams() *GetFolderListsParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetFolderListsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFolderListsParamsWithTimeout creates a new GetFolderListsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFolderListsParamsWithTimeout(timeout time.Duration) *GetFolderListsParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetFolderListsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetFolderListsParamsWithContext creates a new GetFolderListsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFolderListsParamsWithContext(ctx context.Context) *GetFolderListsParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetFolderListsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetFolderListsParamsWithHTTPClient creates a new GetFolderListsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFolderListsParamsWithHTTPClient(client *http.Client) *GetFolderListsParams {
	var (
		limitDefault  = int64(10)
		offsetDefault = int64(0)
	)
	return &GetFolderListsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetFolderListsParams contains all the parameters to send to the API endpoint
for the get folder lists operation typically these are written to a http.Request
*/
type GetFolderListsParams struct {

	/*FolderID
	  Id of the folder

	*/
	FolderID int64
	/*Limit
	  Number of documents per page

	*/
	Limit *int64
	/*Offset
	  Index of the first document of the page

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get folder lists params
func (o *GetFolderListsParams) WithTimeout(timeout time.Duration) *GetFolderListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get folder lists params
func (o *GetFolderListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get folder lists params
func (o *GetFolderListsParams) WithContext(ctx context.Context) *GetFolderListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get folder lists params
func (o *GetFolderListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get folder lists params
func (o *GetFolderListsParams) WithHTTPClient(client *http.Client) *GetFolderListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get folder lists params
func (o *GetFolderListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFolderID adds the folderID to the get folder lists params
func (o *GetFolderListsParams) WithFolderID(folderID int64) *GetFolderListsParams {
	o.SetFolderID(folderID)
	return o
}

// SetFolderID adds the folderId to the get folder lists params
func (o *GetFolderListsParams) SetFolderID(folderID int64) {
	o.FolderID = folderID
}

// WithLimit adds the limit to the get folder lists params
func (o *GetFolderListsParams) WithLimit(limit *int64) *GetFolderListsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get folder lists params
func (o *GetFolderListsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get folder lists params
func (o *GetFolderListsParams) WithOffset(offset *int64) *GetFolderListsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get folder lists params
func (o *GetFolderListsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetFolderListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param folderId
	if err := r.SetPathParam("folderId", swag.FormatInt64(o.FolderID)); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}