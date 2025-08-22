package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/issue-tracking-api/mcp-server/config"
	"github.com/issue-tracking-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CollectionticketsdeleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ticket_idVal, ok := args["ticket_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ticket_id"), nil
		}
		ticket_id, ok := ticket_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ticket_id"), nil
		}
		collection_idVal, ok := args["collection_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: collection_id"), nil
		}
		collection_id, ok := collection_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: collection_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["raw"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("raw=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/issue-tracking/collections/%s/tickets/%s%s", cfg.BaseURL, ticket_id, collection_id, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-apideck-consumer-id"]; ok {
			req.Header.Set("x-apideck-consumer-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-app-id"]; ok {
			req.Header.Set("x-apideck-app-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-service-id"]; ok {
			req.Header.Set("x-apideck-service-id", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.DeleteTicketResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCollectionticketsdeleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_issue-tracking_collections_collection_id_tickets_ticket_id",
		mcp.WithDescription("Delete Ticket"),
		mcp.WithString("ticket_id", mcp.Required(), mcp.Description("ID of the ticket you are acting upon.")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("collection_id", mcp.Required(), mcp.Description("The collection ID")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CollectionticketsdeleteHandler(cfg),
	}
}
