package ssl

import (
	"context"
	"github.com/1Panel-dev/mcp-1panel/operations/types"
	"github.com/1Panel-dev/mcp-1panel/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	ListSSLs = "list_ssls"
)

var ListSSLsTool = mcp.NewTool(
	ListSSLs,
	mcp.WithDescription("list ssls"),
)

func ListSSLHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req := &types.PageRequest{
		Page:     1,
		PageSize: 500,
	}
	client := utils.NewPanelClient("POST", "/websites/ssl/search", utils.WithPayload(req))
	listWebsiteSSLRes := &types.ListWebsiteSSLRes{}
	return client.Request(listWebsiteSSLRes)
}
