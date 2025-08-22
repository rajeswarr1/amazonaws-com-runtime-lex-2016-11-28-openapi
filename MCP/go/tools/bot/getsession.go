package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-lex-runtime-service/mcp-server/config"
	"github.com/amazon-lex-runtime-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetsessionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		botNameVal, ok := args["botName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: botName"), nil
		}
		botName, ok := botNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: botName"), nil
		}
		botAliasVal, ok := args["botAlias"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: botAlias"), nil
		}
		botAlias, ok := botAliasVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: botAlias"), nil
		}
		userIdVal, ok := args["userId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: userId"), nil
		}
		userId, ok := userIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: userId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["checkpointLabelFilter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("checkpointLabelFilter=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/bot/%s/alias/%s/user/%s/session/%s", cfg.BaseURL, botName, botAlias, userId, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

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
		var result models.GetSessionResponse
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

func CreateGetsessionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_bot_botName_alias_botAlias_user_userId_session",
		mcp.WithDescription("Returns session information for a specified bot, alias, and user ID."),
		mcp.WithString("botName", mcp.Required(), mcp.Description("The name of the bot that contains the session data.")),
		mcp.WithString("botAlias", mcp.Required(), mcp.Description("The alias in use for the bot that contains the session data.")),
		mcp.WithString("userId", mcp.Required(), mcp.Description("The ID of the client application user. Amazon Lex uses this to identify a user's conversation with your bot. ")),
		mcp.WithString("checkpointLabelFilter", mcp.Description("<p>A string used to filter the intents returned in the <code>recentIntentSummaryView</code> structure. </p> <p>When you specify a filter, only intents with their <code>checkpointLabel</code> field set to that string are returned.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsessionHandler(cfg),
	}
}
