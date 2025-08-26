// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mdboynton/cortex-cloud-go/api"
	"github.com/stretchr/testify/assert"
)

// setupTest is a helper from appsec/client_test.go
func setupTest(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	server := httptest.NewServer(handler)
	config := &api.Config{
		ApiUrl:     server.URL,
		ApiKey:     "test-key",
		ApiKeyId:   123,
		Transport:  server.Client().Transport.(*http.Transport),
	}
	client, err := NewClient(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	return client, server
}

func TestClient_GetIDPMetadata(t *testing.T) {
	t.Run("should list idp metadata successfully", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/"+GetIDPMetadataEndpoint, r.URL.Path)

			var req GetIDPMetadataRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			assert.NoError(t, err)

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{
				"tenant_id": "my-tenant",
				"sp_entity_id": "sp-entity-id",
				"sp_logout_url": "https://logout.url",
				"sp_url": "https://sp.url"
			}`)
		})
		client, server := setupTest(t, handler)
		defer server.Close()

		resp, err := client.GetIDPMetadata(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, "my-tenant", resp.TenantID)
		assert.Equal(t, "sp-entity-id", resp.SpEntityID)
		assert.Equal(t, "https://logout.url", resp.SpLogoutURL)
		assert.Equal(t, "https://sp.url", resp.SpURL)
	})
}

func TestClient_ListAuthSettings(t *testing.T) {
	t.Run("should list auth settings successfully", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/"+ListAuthSettingsEndpoint, r.URL.Path)

			var req ListAuthSettingsRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			assert.NoError(t, err)

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{
				"reply": [
					{
						"tenant_id": "my-tenant",
						"name": "Test Setting",
						"domain": "example.com"
					}
				]
			}`)
		})
		client, server := setupTest(t, handler)
		defer server.Close()

		resp, err := client.ListAuthSettings(context.Background())
		assert.NoError(t, err)
		assert.Len(t, resp.Reply, 1)
		assert.Equal(t, "Test Setting", resp.Reply[0].Name)
		assert.Equal(t, "example.com", resp.Reply[0].Domain)
	})
}

func TestClient_CreateAuthSettings(t *testing.T) {
	t.Run("should create auth settings successfully", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/"+CreateAuthSettingsEndpoint, r.URL.Path)

			var req CreateAuthSettingsRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			assert.NoError(t, err)
			assert.Equal(t, "New Setting", req.Data.Name)
			assert.Equal(t, "new.example.com", req.Data.Domain)

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"reply": true}`)
		})
		client, server := setupTest(t, handler)
		defer server.Close()

		createReq := CreateAuthSettingsRequest{
			Data: CreateAuthSettingsRequestData{
				Name:   "New Setting",
				Domain: "new.example.com",
			},
		}
		resp, err := client.CreateAuthSettings(context.Background(), createReq)
		assert.NoError(t, err)
		assert.True(t, resp.Reply)
	})
}

func TestClient_UpdateAuthSettings(t *testing.T) {
	t.Run("should update auth settings successfully", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/"+UpdateAuthSettingsEndpoint, r.URL.Path)

			var req UpdateAuthSettingsRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			assert.NoError(t, err)
			assert.Equal(t, "Updated Setting", req.Data.Name)
			assert.Equal(t, "old.example.com", req.Data.CurrentDomain)
			assert.Equal(t, "new.example.com", req.Data.NewDomain)

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"reply": true}`)
		})
		client, server := setupTest(t, handler)
		defer server.Close()

		updateReq := UpdateAuthSettingsRequest{
			Data: UpdateAuthSettingsRequestData{
				Name:          "Updated Setting",
				CurrentDomain: "old.example.com",
				NewDomain:     "new.example.com",
			},
		}
		resp, err := client.UpdateAuthSettings(context.Background(), updateReq)
		assert.NoError(t, err)
		assert.True(t, resp.Reply)
	})
}

func TestClient_DeleteAuthSettings(t *testing.T) {
	t.Run("should delete auth settings successfully", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "/"+DeleteAuthSettingsEndpoint, r.URL.Path)

			var req DeleteAuthSettingsRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			assert.NoError(t, err)
			assert.Equal(t, "delete.example.com", req.Data.Domain)

			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"reply": true}`)
		})
		client, server := setupTest(t, handler)
		defer server.Close()

		resp, err := client.DeleteAuthSettings(context.Background(), "delete.example.com")
		assert.NoError(t, err)
		assert.True(t, resp.Reply)
	})
}
