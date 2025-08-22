package main

import (
	"github.com/amazon-lex-runtime-service/mcp-server/config"
	"github.com/amazon-lex-runtime-service/mcp-server/models"
	tools_bot "github.com/amazon-lex-runtime-service/mcp-server/tools/bot"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_bot.CreatePostcontentTool(cfg),
		tools_bot.CreateDeletesessionTool(cfg),
		tools_bot.CreatePutsessionTool(cfg),
		tools_bot.CreateGetsessionTool(cfg),
		tools_bot.CreatePosttextTool(cfg),
	}
}
