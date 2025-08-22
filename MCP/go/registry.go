package main

import (
	"github.com/issue-tracking-api/mcp-server/config"
	"github.com/issue-tracking-api/mcp-server/models"
	tools_collections "github.com/issue-tracking-api/mcp-server/tools/collections"
	tools_tickets "github.com/issue-tracking-api/mcp-server/tools/tickets"
	tools_tags "github.com/issue-tracking-api/mcp-server/tools/tags"
	tools_comments "github.com/issue-tracking-api/mcp-server/tools/comments"
	tools_users "github.com/issue-tracking-api/mcp-server/tools/users"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_collections.CreateCollectionsallTool(cfg),
		tools_collections.CreateCollectionsoneTool(cfg),
		tools_tickets.CreateCollectionticketsaddTool(cfg),
		tools_tickets.CreateCollectionticketsallTool(cfg),
		tools_tags.CreateCollectiontagsallTool(cfg),
		tools_tickets.CreateCollectionticketsdeleteTool(cfg),
		tools_tickets.CreateCollectionticketsoneTool(cfg),
		tools_tickets.CreateCollectionticketsupdateTool(cfg),
		tools_comments.CreateCollectionticketcommentsallTool(cfg),
		tools_comments.CreateCollectionticketcommentsaddTool(cfg),
		tools_comments.CreateCollectionticketcommentsupdateTool(cfg),
		tools_comments.CreateCollectionticketcommentsdeleteTool(cfg),
		tools_comments.CreateCollectionticketcommentsoneTool(cfg),
		tools_users.CreateCollectionusersallTool(cfg),
		tools_users.CreateCollectionusersoneTool(cfg),
	}
}
