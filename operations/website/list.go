package website

import (
	"context"
	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListWebsites = "list_websites"
)

var ListWebsitesTool = mcp.NewTool(
	ListWebsites,
	mcp.WithDescription("list websites"),
	mcp.WithString("name", mcp.Description("search by website name")),
)

func ListWebsiteHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req := &types.ListWebsiteRequest{
		Order:   "null",
		OrderBy: "created_at",
		PageRequest: types.PageRequest{
			Page:     1,
			PageSize: 500,
			Name:     "",
		},
	}
	client := utils.NewPanelClient("POST", "/websites/search", utils.WithPayload(req))
	listWebsiteRes := &types.ListWebsiteRes{}
	return client.Request(listWebsiteRes)
}
