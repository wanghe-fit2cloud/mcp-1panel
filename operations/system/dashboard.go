package system

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"mcp-1panel/operations/types"
	"mcp-1panel/utils"
)

const (
	GetDashboardInfo = "get_dashboard_info"
)

var GetDashboardInfoTool = mcp.NewTool(GetDashboardInfo, mcp.WithDescription(
	"show dashboard info"))

func GetDashboardInfoHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := "/dashboard/base/all/all"
	client := utils.NewPanelClient("GET", apiUrl)
	osInfo := &types.DashboardRes{}
	return client.Request(osInfo)
}
