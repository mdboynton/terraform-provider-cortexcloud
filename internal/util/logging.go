package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func LogHttpRequest(ctx context.Context, method, url, xdrAuthId, authorization, payload string) string {
	requestId := uuid.New().String()

	tflog.Debug(ctx, fmt.Sprintf("Sending HTTP Request: request_uid=%s method=\"%s\" url=\"%s\" headers.x-xdr-auth-id=\"%s\" headers.Authorization=\"%s\" body=\n%s\n", requestId, method, url, xdrAuthId, authorization, payload))

	return requestId
}

func LogHttpResponse(ctx context.Context, requestId string, statusCode int, body interface{}) error {
	jsonBody, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		return err
	}

	tflog.Debug(ctx, fmt.Sprintf("Recieved HTTP Response: request_uid=%s status_code=%d body=\n%s\n", requestId, statusCode, string(jsonBody)))

	return nil
}
