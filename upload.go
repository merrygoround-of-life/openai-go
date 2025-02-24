// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/openai/openai-go/internal/apijson"
	"github.com/openai/openai-go/internal/requestconfig"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/packages/resp"
	"github.com/openai/openai-go/shared/constant"
)

// UploadService contains methods and other services that help with interacting
// with the openai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUploadService] method instead.
type UploadService struct {
	Options []option.RequestOption
	Parts   UploadPartService
}

// NewUploadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUploadService(opts ...option.RequestOption) (r UploadService) {
	r = UploadService{}
	r.Options = opts
	r.Parts = NewUploadPartService(opts...)
	return
}

// Creates an intermediate
// [Upload](https://platform.openai.com/docs/api-reference/uploads/object) object
// that you can add
// [Parts](https://platform.openai.com/docs/api-reference/uploads/part-object) to.
// Currently, an Upload can accept at most 8 GB in total and expires after an hour
// after you create it.
//
// Once you complete the Upload, we will create a
// [File](https://platform.openai.com/docs/api-reference/files/object) object that
// contains all the parts you uploaded. This File is usable in the rest of our
// platform as a regular File object.
//
// For certain `purpose`s, the correct `mime_type` must be specified. Please refer
// to documentation for the supported MIME types for your use case:
//
// - [Assistants](https://platform.openai.com/docs/assistants/tools/file-search#supported-files)
//
// For guidance on the proper filename extensions for each purpose, please follow
// the documentation on
// [creating a File](https://platform.openai.com/docs/api-reference/files/create).
func (r *UploadService) New(ctx context.Context, body UploadNewParams, opts ...option.RequestOption) (res *Upload, err error) {
	opts = append(r.Options[:], opts...)
	path := "uploads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cancels the Upload. No Parts may be added after an Upload is cancelled.
func (r *UploadService) Cancel(ctx context.Context, uploadID string, opts ...option.RequestOption) (res *Upload, err error) {
	opts = append(r.Options[:], opts...)
	if uploadID == "" {
		err = errors.New("missing required upload_id parameter")
		return
	}
	path := fmt.Sprintf("uploads/%s/cancel", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Completes the
// [Upload](https://platform.openai.com/docs/api-reference/uploads/object).
//
// Within the returned Upload object, there is a nested
// [File](https://platform.openai.com/docs/api-reference/files/object) object that
// is ready to use in the rest of the platform.
//
// You can specify the order of the Parts by passing in an ordered list of the Part
// IDs.
//
// The number of bytes uploaded upon completion must match the number of bytes
// initially specified when creating the Upload object. No Parts may be added after
// an Upload is completed.
func (r *UploadService) Complete(ctx context.Context, uploadID string, body UploadCompleteParams, opts ...option.RequestOption) (res *Upload, err error) {
	opts = append(r.Options[:], opts...)
	if uploadID == "" {
		err = errors.New("missing required upload_id parameter")
		return
	}
	path := fmt.Sprintf("uploads/%s/complete", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The Upload object can accept byte chunks in the form of Parts.
type Upload struct {
	// The Upload unique identifier, which can be referenced in API endpoints.
	ID string `json:"id,omitzero,required"`
	// The intended number of bytes to be uploaded.
	Bytes int64 `json:"bytes,omitzero,required"`
	// The Unix timestamp (in seconds) for when the Upload was created.
	CreatedAt int64 `json:"created_at,omitzero,required"`
	// The Unix timestamp (in seconds) for when the Upload was created.
	ExpiresAt int64 `json:"expires_at,omitzero,required"`
	// The name of the file to be uploaded.
	Filename string `json:"filename,omitzero,required"`
	// The object type, which is always "upload".
	//
	// This field can be elided, and will be automatically set as "upload".
	Object constant.Upload `json:"object,required"`
	// The intended purpose of the file.
	// [Please refer here](https://platform.openai.com/docs/api-reference/files/object#files/object-purpose)
	// for acceptable values.
	Purpose string `json:"purpose,omitzero,required"`
	// The status of the Upload.
	//
	// Any of "pending", "completed", "cancelled", "expired"
	Status string `json:"status,omitzero,required"`
	// The `File` object represents a document that has been uploaded to OpenAI.
	File FileObject `json:"file,omitzero,nullable"`
	JSON struct {
		ID        resp.Field
		Bytes     resp.Field
		CreatedAt resp.Field
		ExpiresAt resp.Field
		Filename  resp.Field
		Object    resp.Field
		Purpose   resp.Field
		Status    resp.Field
		File      resp.Field
		raw       string
	} `json:"-"`
}

func (r Upload) RawJSON() string { return r.JSON.raw }
func (r *Upload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Upload.
type UploadStatus = string

const (
	UploadStatusPending   UploadStatus = "pending"
	UploadStatusCompleted UploadStatus = "completed"
	UploadStatusCancelled UploadStatus = "cancelled"
	UploadStatusExpired   UploadStatus = "expired"
)

type UploadNewParams struct {
	// The number of bytes in the file you are uploading.
	Bytes param.Int `json:"bytes,omitzero,required"`
	// The name of the file to upload.
	Filename param.String `json:"filename,omitzero,required"`
	// The MIME type of the file.
	//
	// This must fall within the supported MIME types for your file purpose. See the
	// supported MIME types for assistants and vision.
	MimeType param.String `json:"mime_type,omitzero,required"`
	// The intended purpose of the uploaded file.
	//
	// See the
	// [documentation on File purposes](https://platform.openai.com/docs/api-reference/files/create#files-create-purpose).
	//
	// Any of "assistants", "batch", "fine-tune", "vision"
	Purpose FilePurpose `json:"purpose,omitzero,required"`
	apiobject
}

func (f UploadNewParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

func (r UploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type UploadCompleteParams struct {
	// The ordered list of Part IDs.
	PartIDs []string `json:"part_ids,omitzero,required"`
	// The optional md5 checksum for the file contents to verify if the bytes uploaded
	// matches what you expect.
	Md5 param.String `json:"md5,omitzero"`
	apiobject
}

func (f UploadCompleteParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

func (r UploadCompleteParams) MarshalJSON() (data []byte, err error) {
	type shadow UploadCompleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
