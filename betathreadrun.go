// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/openai/openai-go/internal/apijson"
	"github.com/openai/openai-go/internal/apiquery"
	"github.com/openai/openai-go/internal/requestconfig"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/pagination"
	"github.com/openai/openai-go/packages/param"
	"github.com/openai/openai-go/packages/resp"
	"github.com/openai/openai-go/packages/ssestream"
	"github.com/openai/openai-go/shared"
	"github.com/openai/openai-go/shared/constant"
)

// BetaThreadRunService contains methods and other services that help with
// interacting with the openai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaThreadRunService] method instead.
type BetaThreadRunService struct {
	Options []option.RequestOption
	Steps   BetaThreadRunStepService
}

// NewBetaThreadRunService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaThreadRunService(opts ...option.RequestOption) (r BetaThreadRunService) {
	r = BetaThreadRunService{}
	r.Options = opts
	r.Steps = NewBetaThreadRunStepService(opts...)
	return
}

// Create a run.
func (r *BetaThreadRunService) New(ctx context.Context, threadID string, params BetaThreadRunNewParams, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a run.
func (r *BetaThreadRunService) NewStreaming(ctx context.Context, threadID string, params BetaThreadRunNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AssistantStreamEventUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2"), option.WithJSONSet("stream", true)}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &raw, opts...)
	return ssestream.NewStream[AssistantStreamEventUnion](ssestream.NewDecoder(raw), err)
}

// Retrieves a run.
func (r *BetaThreadRunService) Get(ctx context.Context, threadID string, runID string, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs/%s", threadID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modifies a run.
func (r *BetaThreadRunService) Update(ctx context.Context, threadID string, runID string, body BetaThreadRunUpdateParams, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs/%s", threadID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of runs belonging to a thread.
func (r *BetaThreadRunService) List(ctx context.Context, threadID string, query BetaThreadRunListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Run], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2"), option.WithResponseInto(&raw)}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs", threadID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of runs belonging to a thread.
func (r *BetaThreadRunService) ListAutoPaging(ctx context.Context, threadID string, query BetaThreadRunListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Run] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, threadID, query, opts...))
}

// Cancels a run that is `in_progress`.
func (r *BetaThreadRunService) Cancel(ctx context.Context, threadID string, runID string, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs/%s/cancel", threadID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// When a run has the `status: "requires_action"` and `required_action.type` is
// `submit_tool_outputs`, this endpoint can be used to submit the outputs from the
// tool calls once they're all completed. All outputs must be submitted in a single
// request.
func (r *BetaThreadRunService) SubmitToolOutputs(ctx context.Context, threadID string, runID string, body BetaThreadRunSubmitToolOutputsParams, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2")}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs/%s/submit_tool_outputs", threadID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// When a run has the `status: "requires_action"` and `required_action.type` is
// `submit_tool_outputs`, this endpoint can be used to submit the outputs from the
// tool calls once they're all completed. All outputs must be submitted in a single
// request.
func (r *BetaThreadRunService) SubmitToolOutputsStreaming(ctx context.Context, threadID string, runID string, body BetaThreadRunSubmitToolOutputsParams, opts ...option.RequestOption) (stream *ssestream.Stream[AssistantStreamEventUnion]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("OpenAI-Beta", "assistants=v2"), option.WithJSONSet("stream", true)}, opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("threads/%s/runs/%s/submit_tool_outputs", threadID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[AssistantStreamEventUnion](ssestream.NewDecoder(raw), err)
}

// Tool call objects
type RequiredActionFunctionToolCall struct {
	// The ID of the tool call. This ID must be referenced when you submit the tool
	// outputs in using the
	// [Submit tool outputs to run](https://platform.openai.com/docs/api-reference/runs/submitToolOutputs)
	// endpoint.
	ID string `json:"id,omitzero,required"`
	// The function definition.
	Function RequiredActionFunctionToolCallFunction `json:"function,omitzero,required"`
	// The type of tool call the output is required for. For now, this is always
	// `function`.
	//
	// This field can be elided, and will be automatically set as "function".
	Type constant.Function `json:"type,required"`
	JSON struct {
		ID       resp.Field
		Function resp.Field
		Type     resp.Field
		raw      string
	} `json:"-"`
}

func (r RequiredActionFunctionToolCall) RawJSON() string { return r.JSON.raw }
func (r *RequiredActionFunctionToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The function definition.
type RequiredActionFunctionToolCallFunction struct {
	// The arguments that the model expects you to pass to the function.
	Arguments string `json:"arguments,omitzero,required"`
	// The name of the function.
	Name string `json:"name,omitzero,required"`
	JSON struct {
		Arguments resp.Field
		Name      resp.Field
		raw       string
	} `json:"-"`
}

func (r RequiredActionFunctionToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *RequiredActionFunctionToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an execution run on a
// [thread](https://platform.openai.com/docs/api-reference/threads).
type Run struct {
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id,omitzero,required"`
	// The ID of the
	// [assistant](https://platform.openai.com/docs/api-reference/assistants) used for
	// execution of this run.
	AssistantID string `json:"assistant_id,omitzero,required"`
	// The Unix timestamp (in seconds) for when the run was cancelled.
	CancelledAt int64 `json:"cancelled_at,omitzero,required,nullable"`
	// The Unix timestamp (in seconds) for when the run was completed.
	CompletedAt int64 `json:"completed_at,omitzero,required,nullable"`
	// The Unix timestamp (in seconds) for when the run was created.
	CreatedAt int64 `json:"created_at,omitzero,required"`
	// The Unix timestamp (in seconds) for when the run will expire.
	ExpiresAt int64 `json:"expires_at,omitzero,required,nullable"`
	// The Unix timestamp (in seconds) for when the run failed.
	FailedAt int64 `json:"failed_at,omitzero,required,nullable"`
	// Details on why the run is incomplete. Will be `null` if the run is not
	// incomplete.
	IncompleteDetails RunIncompleteDetails `json:"incomplete_details,omitzero,required,nullable"`
	// The instructions that the
	// [assistant](https://platform.openai.com/docs/api-reference/assistants) used for
	// this run.
	Instructions string `json:"instructions,omitzero,required"`
	// The last error associated with this run. Will be `null` if there are no errors.
	LastError RunLastError `json:"last_error,omitzero,required,nullable"`
	// The maximum number of completion tokens specified to have been used over the
	// course of the run.
	MaxCompletionTokens int64 `json:"max_completion_tokens,omitzero,required,nullable"`
	// The maximum number of prompt tokens specified to have been used over the course
	// of the run.
	MaxPromptTokens int64 `json:"max_prompt_tokens,omitzero,required,nullable"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard.
	//
	// Keys are strings with a maximum length of 64 characters. Values are strings with
	// a maximum length of 512 characters.
	Metadata shared.Metadata `json:"metadata,omitzero,required,nullable"`
	// The model that the
	// [assistant](https://platform.openai.com/docs/api-reference/assistants) used for
	// this run.
	Model string `json:"model,omitzero,required"`
	// The object type, which is always `thread.run`.
	//
	// This field can be elided, and will be automatically set as "thread.run".
	Object constant.ThreadRun `json:"object,required"`
	// Whether to enable
	// [parallel function calling](https://platform.openai.com/docs/guides/function-calling#configuring-parallel-function-calling)
	// during tool use.
	ParallelToolCalls bool `json:"parallel_tool_calls,omitzero,required"`
	// Details on the action required to continue the run. Will be `null` if no action
	// is required.
	RequiredAction RunRequiredAction `json:"required_action,omitzero,required,nullable"`
	// The Unix timestamp (in seconds) for when the run was started.
	StartedAt int64 `json:"started_at,omitzero,required,nullable"`
	// The status of the run, which can be either `queued`, `in_progress`,
	// `requires_action`, `cancelling`, `cancelled`, `failed`, `completed`,
	// `incomplete`, or `expired`.
	//
	// Any of "queued", "in_progress", "requires_action", "cancelling", "cancelled",
	// "failed", "completed", "incomplete", "expired"
	Status RunStatus `json:"status,omitzero,required"`
	// The ID of the [thread](https://platform.openai.com/docs/api-reference/threads)
	// that was executed on as a part of this run.
	ThreadID string `json:"thread_id,omitzero,required"`
	// Controls which (if any) tool is called by the model. `none` means the model will
	// not call any tools and instead generates a message. `auto` is the default value
	// and means the model can pick between generating a message or calling one or more
	// tools. `required` means the model must call one or more tools before responding
	// to the user. Specifying a particular tool like `{"type": "file_search"}` or
	// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
	// call that tool.
	ToolChoice AssistantToolChoiceOptionUnion `json:"tool_choice,omitzero,required,nullable"`
	// The list of tools that the
	// [assistant](https://platform.openai.com/docs/api-reference/assistants) used for
	// this run.
	Tools []AssistantToolUnion `json:"tools,omitzero,required"`
	// Controls for how a thread will be truncated prior to the run. Use this to
	// control the intial context window of the run.
	TruncationStrategy RunTruncationStrategy `json:"truncation_strategy,omitzero,required,nullable"`
	// Usage statistics related to the run. This value will be `null` if the run is not
	// in a terminal state (i.e. `in_progress`, `queued`, etc.).
	Usage RunUsage `json:"usage,omitzero,required,nullable"`
	// The sampling temperature used for this run. If not set, defaults to 1.
	Temperature float64 `json:"temperature,omitzero,nullable"`
	// The nucleus sampling value used for this run. If not set, defaults to 1.
	TopP float64 `json:"top_p,omitzero,nullable"`
	JSON struct {
		ID                  resp.Field
		AssistantID         resp.Field
		CancelledAt         resp.Field
		CompletedAt         resp.Field
		CreatedAt           resp.Field
		ExpiresAt           resp.Field
		FailedAt            resp.Field
		IncompleteDetails   resp.Field
		Instructions        resp.Field
		LastError           resp.Field
		MaxCompletionTokens resp.Field
		MaxPromptTokens     resp.Field
		Metadata            resp.Field
		Model               resp.Field
		Object              resp.Field
		ParallelToolCalls   resp.Field
		RequiredAction      resp.Field
		StartedAt           resp.Field
		Status              resp.Field
		ThreadID            resp.Field
		ToolChoice          resp.Field
		Tools               resp.Field
		TruncationStrategy  resp.Field
		Usage               resp.Field
		Temperature         resp.Field
		TopP                resp.Field
		raw                 string
	} `json:"-"`
}

func (r Run) RawJSON() string { return r.JSON.raw }
func (r *Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details on why the run is incomplete. Will be `null` if the run is not
// incomplete.
type RunIncompleteDetails struct {
	// The reason why the run is incomplete. This will point to which specific token
	// limit was reached over the course of the run.
	//
	// Any of "max_completion_tokens", "max_prompt_tokens"
	Reason string `json:"reason,omitzero"`
	JSON   struct {
		Reason resp.Field
		raw    string
	} `json:"-"`
}

func (r RunIncompleteDetails) RawJSON() string { return r.JSON.raw }
func (r *RunIncompleteDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why the run is incomplete. This will point to which specific token
// limit was reached over the course of the run.
type RunIncompleteDetailsReason = string

const (
	RunIncompleteDetailsReasonMaxCompletionTokens RunIncompleteDetailsReason = "max_completion_tokens"
	RunIncompleteDetailsReasonMaxPromptTokens     RunIncompleteDetailsReason = "max_prompt_tokens"
)

// The last error associated with this run. Will be `null` if there are no errors.
type RunLastError struct {
	// One of `server_error`, `rate_limit_exceeded`, or `invalid_prompt`.
	//
	// Any of "server_error", "rate_limit_exceeded", "invalid_prompt"
	Code string `json:"code,omitzero,required"`
	// A human-readable description of the error.
	Message string `json:"message,omitzero,required"`
	JSON    struct {
		Code    resp.Field
		Message resp.Field
		raw     string
	} `json:"-"`
}

func (r RunLastError) RawJSON() string { return r.JSON.raw }
func (r *RunLastError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One of `server_error`, `rate_limit_exceeded`, or `invalid_prompt`.
type RunLastErrorCode = string

const (
	RunLastErrorCodeServerError       RunLastErrorCode = "server_error"
	RunLastErrorCodeRateLimitExceeded RunLastErrorCode = "rate_limit_exceeded"
	RunLastErrorCodeInvalidPrompt     RunLastErrorCode = "invalid_prompt"
)

// Details on the action required to continue the run. Will be `null` if no action
// is required.
type RunRequiredAction struct {
	// Details on the tool outputs needed for this run to continue.
	SubmitToolOutputs RunRequiredActionSubmitToolOutputs `json:"submit_tool_outputs,omitzero,required"`
	// For now, this is always `submit_tool_outputs`.
	//
	// This field can be elided, and will be automatically set as
	// "submit_tool_outputs".
	Type constant.SubmitToolOutputs `json:"type,required"`
	JSON struct {
		SubmitToolOutputs resp.Field
		Type              resp.Field
		raw               string
	} `json:"-"`
}

func (r RunRequiredAction) RawJSON() string { return r.JSON.raw }
func (r *RunRequiredAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details on the tool outputs needed for this run to continue.
type RunRequiredActionSubmitToolOutputs struct {
	// A list of the relevant tool calls.
	ToolCalls []RequiredActionFunctionToolCall `json:"tool_calls,omitzero,required"`
	JSON      struct {
		ToolCalls resp.Field
		raw       string
	} `json:"-"`
}

func (r RunRequiredActionSubmitToolOutputs) RawJSON() string { return r.JSON.raw }
func (r *RunRequiredActionSubmitToolOutputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls for how a thread will be truncated prior to the run. Use this to
// control the intial context window of the run.
type RunTruncationStrategy struct {
	// The truncation strategy to use for the thread. The default is `auto`. If set to
	// `last_messages`, the thread will be truncated to the n most recent messages in
	// the thread. When set to `auto`, messages in the middle of the thread will be
	// dropped to fit the context length of the model, `max_prompt_tokens`.
	//
	// Any of "auto", "last_messages"
	Type string `json:"type,omitzero,required"`
	// The number of most recent messages from the thread when constructing the context
	// for the run.
	LastMessages int64 `json:"last_messages,omitzero,nullable"`
	JSON         struct {
		Type         resp.Field
		LastMessages resp.Field
		raw          string
	} `json:"-"`
}

func (r RunTruncationStrategy) RawJSON() string { return r.JSON.raw }
func (r *RunTruncationStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The truncation strategy to use for the thread. The default is `auto`. If set to
// `last_messages`, the thread will be truncated to the n most recent messages in
// the thread. When set to `auto`, messages in the middle of the thread will be
// dropped to fit the context length of the model, `max_prompt_tokens`.
type RunTruncationStrategyType = string

const (
	RunTruncationStrategyTypeAuto         RunTruncationStrategyType = "auto"
	RunTruncationStrategyTypeLastMessages RunTruncationStrategyType = "last_messages"
)

// Usage statistics related to the run. This value will be `null` if the run is not
// in a terminal state (i.e. `in_progress`, `queued`, etc.).
type RunUsage struct {
	// Number of completion tokens used over the course of the run.
	CompletionTokens int64 `json:"completion_tokens,omitzero,required"`
	// Number of prompt tokens used over the course of the run.
	PromptTokens int64 `json:"prompt_tokens,omitzero,required"`
	// Total number of tokens used (prompt + completion).
	TotalTokens int64 `json:"total_tokens,omitzero,required"`
	JSON        struct {
		CompletionTokens resp.Field
		PromptTokens     resp.Field
		TotalTokens      resp.Field
		raw              string
	} `json:"-"`
}

func (r RunUsage) RawJSON() string { return r.JSON.raw }
func (r *RunUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the run, which can be either `queued`, `in_progress`,
// `requires_action`, `cancelling`, `cancelled`, `failed`, `completed`,
// `incomplete`, or `expired`.
type RunStatus string

const (
	RunStatusQueued         RunStatus = "queued"
	RunStatusInProgress     RunStatus = "in_progress"
	RunStatusRequiresAction RunStatus = "requires_action"
	RunStatusCancelling     RunStatus = "cancelling"
	RunStatusCancelled      RunStatus = "cancelled"
	RunStatusFailed         RunStatus = "failed"
	RunStatusCompleted      RunStatus = "completed"
	RunStatusIncomplete     RunStatus = "incomplete"
	RunStatusExpired        RunStatus = "expired"
)

type BetaThreadRunNewParams struct {
	// The ID of the
	// [assistant](https://platform.openai.com/docs/api-reference/assistants) to use to
	// execute this run.
	AssistantID param.String `json:"assistant_id,omitzero,required"`
	// A list of additional fields to include in the response. Currently the only
	// supported value is `step_details.tool_calls[*].file_search.results[*].content`
	// to fetch the file search result content.
	//
	// See the
	// [file search tool documentation](https://platform.openai.com/docs/assistants/tools/file-search#customizing-file-search-settings)
	// for more information.
	Include []RunStepInclude `query:"include,omitzero"`
	// Appends additional instructions at the end of the instructions for the run. This
	// is useful for modifying the behavior on a per-run basis without overriding other
	// instructions.
	AdditionalInstructions param.String `json:"additional_instructions,omitzero"`
	// Adds additional messages to the thread before creating the run.
	AdditionalMessages []BetaThreadRunNewParamsAdditionalMessage `json:"additional_messages,omitzero"`
	// Overrides the
	// [instructions](https://platform.openai.com/docs/api-reference/assistants/createAssistant)
	// of the assistant. This is useful for modifying the behavior on a per-run basis.
	Instructions param.String `json:"instructions,omitzero"`
	// The maximum number of completion tokens that may be used over the course of the
	// run. The run will make a best effort to use only the number of completion tokens
	// specified, across multiple turns of the run. If the run exceeds the number of
	// completion tokens specified, the run will end with status `incomplete`. See
	// `incomplete_details` for more info.
	MaxCompletionTokens param.Int `json:"max_completion_tokens,omitzero"`
	// The maximum number of prompt tokens that may be used over the course of the run.
	// The run will make a best effort to use only the number of prompt tokens
	// specified, across multiple turns of the run. If the run exceeds the number of
	// prompt tokens specified, the run will end with status `incomplete`. See
	// `incomplete_details` for more info.
	MaxPromptTokens param.Int `json:"max_prompt_tokens,omitzero"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard.
	//
	// Keys are strings with a maximum length of 64 characters. Values are strings with
	// a maximum length of 512 characters.
	Metadata shared.MetadataParam `json:"metadata,omitzero"`
	// The ID of the [Model](https://platform.openai.com/docs/api-reference/models) to
	// be used to execute this run. If a value is provided here, it will override the
	// model associated with the assistant. If not, the model associated with the
	// assistant will be used.
	Model ChatModel `json:"model,omitzero"`
	// Whether to enable
	// [parallel function calling](https://platform.openai.com/docs/guides/function-calling#configuring-parallel-function-calling)
	// during tool use.
	ParallelToolCalls param.Bool `json:"parallel_tool_calls,omitzero"`
	// **o1 and o3-mini models only**
	//
	// Constrains effort on reasoning for
	// [reasoning models](https://platform.openai.com/docs/guides/reasoning). Currently
	// supported values are `low`, `medium`, and `high`. Reducing reasoning effort can
	// result in faster responses and fewer tokens used on reasoning in a response.
	//
	// Any of "low", "medium", "high"
	ReasoningEffort BetaThreadRunNewParamsReasoningEffort `json:"reasoning_effort,omitzero"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will
	// make the output more random, while lower values like 0.2 will make it more
	// focused and deterministic.
	Temperature param.Float `json:"temperature,omitzero"`
	// Controls which (if any) tool is called by the model. `none` means the model will
	// not call any tools and instead generates a message. `auto` is the default value
	// and means the model can pick between generating a message or calling one or more
	// tools. `required` means the model must call one or more tools before responding
	// to the user. Specifying a particular tool like `{"type": "file_search"}` or
	// `{"type": "function", "function": {"name": "my_function"}}` forces the model to
	// call that tool.
	ToolChoice AssistantToolChoiceOptionUnionParam `json:"tool_choice,omitzero"`
	// Override the tools the assistant can use for this run. This is useful for
	// modifying the behavior on a per-run basis.
	Tools []AssistantToolUnionParam `json:"tools,omitzero"`
	// An alternative to sampling with temperature, called nucleus sampling, where the
	// model considers the results of the tokens with top_p probability mass. So 0.1
	// means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or temperature but not both.
	TopP param.Float `json:"top_p,omitzero"`
	// Controls for how a thread will be truncated prior to the run. Use this to
	// control the intial context window of the run.
	TruncationStrategy BetaThreadRunNewParamsTruncationStrategy `json:"truncation_strategy,omitzero"`
	apiobject
}

func (f BetaThreadRunNewParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

func (r BetaThreadRunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [BetaThreadRunNewParams]'s query parameters as `url.Values`.
func (r BetaThreadRunNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaThreadRunNewParamsAdditionalMessage struct {
	// An array of content parts with a defined type, each can be of type `text` or
	// images can be passed with `image_url` or `image_file`. Image types are only
	// supported on
	// [Vision-compatible models](https://platform.openai.com/docs/models).
	Content []MessageContentPartParamUnion `json:"content,omitzero,required"`
	// The role of the entity that is creating the message. Allowed values include:
	//
	//   - `user`: Indicates the message is sent by an actual user and should be used in
	//     most cases to represent user-generated messages.
	//   - `assistant`: Indicates the message is generated by the assistant. Use this
	//     value to insert messages from the assistant into the conversation.
	//
	// Any of "user", "assistant"
	Role string `json:"role,omitzero,required"`
	// A list of files attached to the message, and the tools they should be added to.
	Attachments []BetaThreadRunNewParamsAdditionalMessagesAttachment `json:"attachments,omitzero"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard.
	//
	// Keys are strings with a maximum length of 64 characters. Values are strings with
	// a maximum length of 512 characters.
	Metadata shared.MetadataParam `json:"metadata,omitzero"`
	apiobject
}

func (f BetaThreadRunNewParamsAdditionalMessage) IsMissing() bool {
	return param.IsOmitted(f) || f.IsNull()
}

func (r BetaThreadRunNewParamsAdditionalMessage) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunNewParamsAdditionalMessage
	return param.MarshalObject(r, (*shadow)(&r))
}

// The role of the entity that is creating the message. Allowed values include:
//
//   - `user`: Indicates the message is sent by an actual user and should be used in
//     most cases to represent user-generated messages.
//   - `assistant`: Indicates the message is generated by the assistant. Use this
//     value to insert messages from the assistant into the conversation.
type BetaThreadRunNewParamsAdditionalMessagesRole = string

const (
	BetaThreadRunNewParamsAdditionalMessagesRoleUser      BetaThreadRunNewParamsAdditionalMessagesRole = "user"
	BetaThreadRunNewParamsAdditionalMessagesRoleAssistant BetaThreadRunNewParamsAdditionalMessagesRole = "assistant"
)

type BetaThreadRunNewParamsAdditionalMessagesAttachment struct {
	// The ID of the file to attach to the message.
	FileID param.String `json:"file_id,omitzero"`
	// The tools to add this file to.
	Tools []BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolUnion `json:"tools,omitzero"`
	apiobject
}

func (f BetaThreadRunNewParamsAdditionalMessagesAttachment) IsMissing() bool {
	return param.IsOmitted(f) || f.IsNull()
}

func (r BetaThreadRunNewParamsAdditionalMessagesAttachment) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunNewParamsAdditionalMessagesAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero
type BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolUnion struct {
	OfCodeInterpreter *CodeInterpreterToolParam
	OfFileSearch      *BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolsFileSearch
	apiunion
}

func (u BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolUnion) IsMissing() bool {
	return param.IsOmitted(u) || u.IsNull()
}

func (u BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolUnion](u.OfCodeInterpreter, u.OfFileSearch)
}

func (u BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolUnion) GetType() *string {
	if vt := u.OfCodeInterpreter; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFileSearch; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolsFileSearch struct {
	// The type of tool being defined: `file_search`
	//
	// This field can be elided, and will be automatically set as "file_search".
	Type constant.FileSearch `json:"type,required"`
	apiobject
}

func (f BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolsFileSearch) IsMissing() bool {
	return param.IsOmitted(f) || f.IsNull()
}

func (r BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolsFileSearch) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunNewParamsAdditionalMessagesAttachmentsToolsFileSearch
	return param.MarshalObject(r, (*shadow)(&r))
}

// **o1 and o3-mini models only**
//
// Constrains effort on reasoning for
// [reasoning models](https://platform.openai.com/docs/guides/reasoning). Currently
// supported values are `low`, `medium`, and `high`. Reducing reasoning effort can
// result in faster responses and fewer tokens used on reasoning in a response.
type BetaThreadRunNewParamsReasoningEffort string

const (
	BetaThreadRunNewParamsReasoningEffortLow    BetaThreadRunNewParamsReasoningEffort = "low"
	BetaThreadRunNewParamsReasoningEffortMedium BetaThreadRunNewParamsReasoningEffort = "medium"
	BetaThreadRunNewParamsReasoningEffortHigh   BetaThreadRunNewParamsReasoningEffort = "high"
)

// Controls for how a thread will be truncated prior to the run. Use this to
// control the intial context window of the run.
type BetaThreadRunNewParamsTruncationStrategy struct {
	// The truncation strategy to use for the thread. The default is `auto`. If set to
	// `last_messages`, the thread will be truncated to the n most recent messages in
	// the thread. When set to `auto`, messages in the middle of the thread will be
	// dropped to fit the context length of the model, `max_prompt_tokens`.
	//
	// Any of "auto", "last_messages"
	Type string `json:"type,omitzero,required"`
	// The number of most recent messages from the thread when constructing the context
	// for the run.
	LastMessages param.Int `json:"last_messages,omitzero"`
	apiobject
}

func (f BetaThreadRunNewParamsTruncationStrategy) IsMissing() bool {
	return param.IsOmitted(f) || f.IsNull()
}

func (r BetaThreadRunNewParamsTruncationStrategy) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunNewParamsTruncationStrategy
	return param.MarshalObject(r, (*shadow)(&r))
}

// The truncation strategy to use for the thread. The default is `auto`. If set to
// `last_messages`, the thread will be truncated to the n most recent messages in
// the thread. When set to `auto`, messages in the middle of the thread will be
// dropped to fit the context length of the model, `max_prompt_tokens`.
type BetaThreadRunNewParamsTruncationStrategyType = string

const (
	BetaThreadRunNewParamsTruncationStrategyTypeAuto         BetaThreadRunNewParamsTruncationStrategyType = "auto"
	BetaThreadRunNewParamsTruncationStrategyTypeLastMessages BetaThreadRunNewParamsTruncationStrategyType = "last_messages"
)

type BetaThreadRunUpdateParams struct {
	// Set of 16 key-value pairs that can be attached to an object. This can be useful
	// for storing additional information about the object in a structured format, and
	// querying for objects via API or the dashboard.
	//
	// Keys are strings with a maximum length of 64 characters. Values are strings with
	// a maximum length of 512 characters.
	Metadata shared.MetadataParam `json:"metadata,omitzero"`
	apiobject
}

func (f BetaThreadRunUpdateParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

func (r BetaThreadRunUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type BetaThreadRunListParams struct {
	// A cursor for use in pagination. `after` is an object ID that defines your place
	// in the list. For instance, if you make a list request and receive 100 objects,
	// ending with obj_foo, your subsequent call can include after=obj_foo in order to
	// fetch the next page of the list.
	After param.String `query:"after,omitzero"`
	// A cursor for use in pagination. `before` is an object ID that defines your place
	// in the list. For instance, if you make a list request and receive 100 objects,
	// starting with obj_foo, your subsequent call can include before=obj_foo in order
	// to fetch the previous page of the list.
	Before param.String `query:"before,omitzero"`
	// A limit on the number of objects to be returned. Limit can range between 1 and
	// 100, and the default is 20.
	Limit param.Int `query:"limit,omitzero"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
	// order and `desc` for descending order.
	//
	// Any of "asc", "desc"
	Order BetaThreadRunListParamsOrder `query:"order,omitzero"`
	apiobject
}

func (f BetaThreadRunListParams) IsMissing() bool { return param.IsOmitted(f) || f.IsNull() }

// URLQuery serializes [BetaThreadRunListParams]'s query parameters as
// `url.Values`.
func (r BetaThreadRunListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order by the `created_at` timestamp of the objects. `asc` for ascending
// order and `desc` for descending order.
type BetaThreadRunListParamsOrder string

const (
	BetaThreadRunListParamsOrderAsc  BetaThreadRunListParamsOrder = "asc"
	BetaThreadRunListParamsOrderDesc BetaThreadRunListParamsOrder = "desc"
)

type BetaThreadRunSubmitToolOutputsParams struct {
	// A list of tools for which the outputs are being submitted.
	ToolOutputs []BetaThreadRunSubmitToolOutputsParamsToolOutput `json:"tool_outputs,omitzero,required"`
	apiobject
}

func (f BetaThreadRunSubmitToolOutputsParams) IsMissing() bool {
	return param.IsOmitted(f) || f.IsNull()
}

func (r BetaThreadRunSubmitToolOutputsParams) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunSubmitToolOutputsParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type BetaThreadRunSubmitToolOutputsParamsToolOutput struct {
	// The output of the tool call to be submitted to continue the run.
	Output param.String `json:"output,omitzero"`
	// The ID of the tool call in the `required_action` object within the run object
	// the output is being submitted for.
	ToolCallID param.String `json:"tool_call_id,omitzero"`
	apiobject
}

func (f BetaThreadRunSubmitToolOutputsParamsToolOutput) IsMissing() bool {
	return param.IsOmitted(f) || f.IsNull()
}

func (r BetaThreadRunSubmitToolOutputsParamsToolOutput) MarshalJSON() (data []byte, err error) {
	type shadow BetaThreadRunSubmitToolOutputsParamsToolOutput
	return param.MarshalObject(r, (*shadow)(&r))
}
