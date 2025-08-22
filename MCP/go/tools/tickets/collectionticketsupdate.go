package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/issue-tracking-api/mcp-server/config"
	"github.com/issue-tracking-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CollectionticketsupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.Ticket
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/issue-tracking/collections/%s/tickets/%s%s", cfg.BaseURL, ticket_id, collection_id, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.UpdateTicketResponse
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

func CreateCollectionticketsupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_issue-tracking_collections_collection_id_tickets_ticket_id",
		mcp.WithDescription("Update Ticket"),
		mcp.WithString("ticket_id", mcp.Required(), mcp.Description("ID of the ticket you are acting upon.")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("collection_id", mcp.Required(), mcp.Description("The collection ID")),
		mcp.WithArray("assignees", mcp.Description("")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Input parameter: A unique identifier for an object.")),
		mcp.WithString("collection_id", mcp.Description("Input parameter: The ticket's collection ID")),
		mcp.WithString("parent_id", mcp.Description("Input parameter: The ticket's parent ID")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("due_date", mcp.Description("Input parameter: Due date of the ticket")),
		mcp.WithString("priority", mcp.Description("Input parameter: Priority of the ticket")),
		mcp.WithString("type", mcp.Description("Input parameter: The ticket's type")),
		mcp.WithString("completed_at", mcp.Description("Input parameter: When the ticket was completed")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithArray("tags", mcp.Description("")),
		mcp.WithString("subject", mcp.Description("Input parameter: Subject of the ticket")),
		mcp.WithString("description", mcp.Description("Input parameter: The ticket's description. HTML version of description is mapped if supported by the third-party platform")),
		mcp.WithString("status", mcp.Description("Input parameter: The current status of the ticket. Possible values include: open, in_progress, closed, or - in cases where there is no clear mapping - the original value passed through.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CollectionticketsupdateHandler(cfg),
	}
}
