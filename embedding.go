// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai

import (
	"context"
	"net/http"

	"github.com/openai/openai-go/internal/apijson"
	"github.com/openai/openai-go/internal/requestconfig"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/packages/resp"
	"github.com/openai/openai-go/shared/constant"
)

// EmbeddingService contains methods and other services that help with interacting
// with the openai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r EmbeddingService) {
	r = EmbeddingService{}
	r.Options = opts
	return
}

// Creates an embedding vector representing the input text.
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *CreateEmbeddingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateEmbeddingResponse struct {
	// The list of embeddings generated by the model.
	Data []Embedding `json:"data,omitzero,required"`
	// The name of the model used to generate the embedding.
	Model string `json:"model,omitzero,required"`
	// The object type, which is always "list".
	//
	// This field can be elided, and will be automatically set as "list".
	Object constant.List `json:"object,required"`
	// The usage information for the request.
	Usage CreateEmbeddingResponseUsage `json:"usage,omitzero,required"`
	JSON  struct {
		Data   resp.Field
		Model  resp.Field
		Object resp.Field
		Usage  resp.Field
		raw    string
	} `json:"-"`
}

func (r CreateEmbeddingResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The usage information for the request.
type CreateEmbeddingResponseUsage struct {
	// The number of tokens used by the prompt.
	PromptTokens int64 `json:"prompt_tokens,omitzero,required"`
	// The total number of tokens used by the request.
	TotalTokens int64 `json:"total_tokens,omitzero,required"`
	JSON        struct {
		PromptTokens resp.Field
		TotalTokens  resp.Field
		raw          string
	} `json:"-"`
}

func (r CreateEmbeddingResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *CreateEmbeddingResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an embedding vector returned by embedding endpoint.
type Embedding struct {
	// The embedding vector, which is a list of floats. The length of vector depends on
	// the model as listed in the
	// [embedding guide](https://platform.openai.com/docs/guides/embeddings).
	Embedding []float64 `json:"embedding,omitzero,required"`
	// The index of the embedding in the list of embeddings.
	Index int64 `json:"index,omitzero,required"`
	// The object type, which is always "embedding".
	//
	// This field can be elided, and will be automatically set as "embedding".
	Object constant.Embedding `json:"object,required"`
	JSON   struct {
		Embedding resp.Field
		Index     resp.Field
		Object    resp.Field
		raw       string
	} `json:"-"`
}

func (r Embedding) RawJSON() string { return r.JSON.raw }
func (r *Embedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingModel = string

const (
	EmbeddingModelTextEmbeddingAda002 EmbeddingModel = "text-embedding-ada-002"
	EmbeddingModelTextEmbedding3Small EmbeddingModel = "text-embedding-3-small"
	EmbeddingModelTextEmbedding3Large EmbeddingModel = "text-embedding-3-large"
)

type EmbeddingNewParams struct {
	// Input text to embed, encoded as a string or array of tokens. To embed multiple
	// inputs in a single request, pass an array of strings or array of token arrays.
	// The input must not exceed the max input tokens for the model (8192 tokens for
	// `text-embedding-ada-002`), cannot be an empty string, and any array must be 2048
	// dimensions or less.
	// [Example Python code](https://cookbook.openai.com/examples/how_to_count_tokens_with_tiktoken)
	// for counting tokens. Some models may also impose a limit on total number of
	// tokens summed across inputs.
	Input EmbeddingNewParamsInputUnion `json:"input,omitzero,required"`
	// ID of the model to use. You can use the
	// [List models](https://platform.openai.com/docs/api-reference/models/list) API to
	// see all of your available models, or see our
	// [Model overview](https://platform.openai.com/docs/models) for descriptions of
	// them.
	Model EmbeddingModel `json:"model,omitzero,required"`
	// The number of dimensions the resulting output embeddings should have. Only
	// supported in `text-embedding-3` and later models.
	Dimensions param.Int `json:"dimensions,omitzero"`
	// The format to return the embeddings in. Can be either `float` or
	// [`base64`](https://pypi.org/project/pybase64/).
	//
	// Any of "float", "base64"
	EncodingFormat EmbeddingNewParamsEncodingFormat `json:"encoding_format,omitzero"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor
	// and detect abuse.
	// [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#end-user-ids).
	User param.String `json:"user,omitzero"`
	apiobject
}

func (f EmbeddingNewParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

func (r EmbeddingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmbeddingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero
type EmbeddingNewParamsInputUnion struct {
	OfString             param.String
	OfArrayOfStrings     []string
	OfArrayOfTokens      []int64
	OfArrayOfTokenArrays [][]int64
	apiunion
}

func (u EmbeddingNewParamsInputUnion) IsMissing() bool { return param.IsOmitted(u) || u.IsNull() }

func (u EmbeddingNewParamsInputUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[EmbeddingNewParamsInputUnion](u.OfString, u.OfArrayOfStrings, u.OfArrayOfTokens, u.OfArrayOfTokenArrays)
}

// The format to return the embeddings in. Can be either `float` or
// [`base64`](https://pypi.org/project/pybase64/).
type EmbeddingNewParamsEncodingFormat string

const (
	EmbeddingNewParamsEncodingFormatFloat  EmbeddingNewParamsEncodingFormat = "float"
	EmbeddingNewParamsEncodingFormatBase64 EmbeddingNewParamsEncodingFormat = "base64"
)
