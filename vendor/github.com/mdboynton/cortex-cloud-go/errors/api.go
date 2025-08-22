package errors

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	"strings"

	"github.com/mdboynton/cortex-cloud-go/internal/types"
)

type CortexCloudAPIError struct {
	Reply *CortexCloudAPIErrorReply `json:"reply,omitempty"`
	Code *string `json:"errorCode,omitempty"`
	Message *string `json:"message,omitempty"`
	//Details map[string]any `json:"details,omitempty"`
	Details *CortexCloudAPIErrorDetails `json:"details"`
}

type CortexCloudAPIErrorReply struct {
	Code    int                        `json:"err_code"`
	Message string                     `json:"err_msg"`
	Extra   []CortexCloudAPIErrorExtra `json:"err_extra"`
}

type CortexCloudAPIErrorExtra struct {
	Type     string                     `json:"type"`
	Location []any                      `json:"loc"`
	Message  string                     `json:"msg"`
	Input    any                        `json:"input"`
	Context  CortexCloudAPIErrorContext `json:"ctx"`
}

type CortexCloudAPIErrorContext struct {
	Expected  string `json:"expected,omitempty"`
	MinLength int    `json:"min_length,omitempty"`
}

type CortexCloudAPIErrorDetails struct {
	Params CortexCloudAPIErrorParams `json:"params"`
}

type CortexCloudAPIErrorParams struct {
	Message string `json:"message"`
}

type CortexCloudAPIErrorDetail struct {
	Type     string                     `json:"type"`
	Location []any                      `json:"loc"`
	Message  string                     `json:"msg"`
	Input    any                        `json:"input"`
	Context  CortexCloudAPIErrorContext `json:"ctx"`
}

func (e CortexCloudAPIErrorExtra) locationAsStringSlice() []string {
	result := []string{}
	for _, elem := range e.Location {
		stringElem, err := types.ConvertInterfaceToString(elem)
		if err != nil {
			stringElem = "UNKNOWN_TYPE"
		}

		result = append(result, stringElem)
	}

	return result
}

func (e CortexCloudAPIErrorExtra) inputAsString() string {
	stringInput, err := types.ConvertInterfaceToString(e.Input)
	if err != nil {
		return "UNKNOWN_TYPE"
	}

	return stringInput
}

func NewCortexCloudAPIError(code string, message string, details CortexCloudAPIErrorDetails) CortexCloudAPIError {
	return CortexCloudAPIError{
		Code: &code,
		Message: &message,
		Details: &details,
	}
	//var e CortexCloudAPIError
	//err := json.Unmarshal(body, &e)
	//if err != nil {
	//	return CortexCloudAPIError{
	//		// TODO: should we populate any of the non-Reply fields here?
	//		Reply: CortexCloudAPIErrorReply{
	//			Code:    code,
	//			Message: string(body),
	//			Extra:   []CortexCloudAPIErrorExtra{},
	//		},
	//	}
	//}

	//return e
}

func (e CortexCloudAPIError) Error() string {
	var sb strings.Builder

	//if (e.Reply != (CortexCloudAPIErrorReply{})) {
	if e.Reply != nil {
		sb.WriteString(fmt.Sprintf("Error Code: %d\n", e.Reply.Code))
		sb.WriteString(fmt.Sprintf("Error Message: %s\n", e.Reply.Message))
		sb.WriteString("Error Details:\n")
		for _, extra := range e.Reply.Extra {
			sb.WriteString(fmt.Sprintf("  - Type: \"%s\"\n", extra.Type))
			sb.WriteString(fmt.Sprintf("    Location: [\"%s\"]\n", strings.Join(extra.locationAsStringSlice(), "\", \"")))
			sb.WriteString(fmt.Sprintf("    Message: \"%s\"\n", extra.Message))
			sb.WriteString(fmt.Sprintf("    Input: \"%s\"\n", extra.inputAsString()))

			if extra.Context.Expected != "" {
				sb.WriteString(fmt.Sprintf("    Expected: \"%s\"\n", extra.Context.Expected))
			}
			if extra.Context.MinLength != 0 {
				sb.WriteString(fmt.Sprintf("    MinLength: \"%d\"\n", extra.Context.MinLength))
			}
		}
	} else {
		var (
			code string
			message string
			details CortexCloudAPIErrorDetails
		)

		if e.Code != nil {
			code = *e.Code
		} else {
			code = ""
		}

		if e.Message != nil {
			message = *e.Message
		} else {
			message = ""
		}

		if e.Details != nil {
			details = *e.Details
		} else {
			details = CortexCloudAPIErrorDetails{}
		}

		sb.WriteString(fmt.Sprintf("Error Code: %s\n", code))
		sb.WriteString(fmt.Sprintf("Error Message: %s\n", message))
		sb.WriteString(fmt.Sprintf("Error Details: %s\n", details))
		//if e.Details != nil {
		//	details := *e.Details
		//	sb.WriteString("Error Details:\n")
		//}
	}

	return sb.String()
}

func (e CortexCloudAPIError) ToBuiltin() error {
	return fmt.Errorf("%+v", e.Error())
}

//func NewResponse(code int, body []byte) Response {
//	ans := Response{ErrorCode: code}
//	if bytes.HasPrefix(bytes.TrimSpace(body), []byte("{")) {
//		json.Unmarshal(body, &ans)
//	}
//	return ans
//}
//
//type Response struct {
//	ErrorCode int            `json:"errorCode"`
//	Message   string         `json:"message"`
//	Details   map[string]any `json:"details"`
//}
//
//func (e Response) Failed() bool {
//	return len(e.Details) > 0
//}
//
//func (e Response) Error() string {
//	var buf strings.Builder
//
//	buf.WriteString(fmt.Sprintf("[HTTP %d] %s", e.ErrorCode, e.Message))
//
//	detailStrs := []string{}
//	for _, detail := range e.Details {
//		if detailStr, ok := detail.(string); ok {
//			detailStrs = append(detailStrs, detailStr)
//		}
//
//	}
//
//	if len(detailStrs) > 0 {
//		buf.WriteString("\n\n{")
//		buf.WriteString(strings.Join(detailStrs, "\n\t"))
//		buf.WriteString("}")
//	}
//
//	return buf.String()
//}

