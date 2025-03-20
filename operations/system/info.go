package system

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"mcp-1panel/operations/types"
	"mcp-1panel/utils"
)

const (
	GetSystemInfo = "get_system_info"
)

var GetSystemInfoTool = mcp.NewTool(GetSystemInfo, mcp.WithDescription(
	"show host system information, The unit of diskSize is bytes"))

func GetSystemInfoHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	apiUrl := "/dashboard/base/os"
	client := utils.NewPanelClient("GET", apiUrl)
	osInfo := &types.OsInfoRes{}
	return client.Request(osInfo)
}
