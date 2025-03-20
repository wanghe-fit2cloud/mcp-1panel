package website

import (
	"context"
	"errors"
	"mcp-1panel/operations/types"
	"mcp-1panel/utils"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	CreateWebsite = "create_website"
)

var CreateWebsiteTool = mcp.NewTool(CreateWebsite,
	mcp.WithDescription("create website"),
	mcp.WithString("domain", mcp.Description("domain"), mcp.Required()),
	mcp.WithString("website_type", mcp.Description("website type,only support static and proxy"), mcp.Required()),
	mcp.WithString("proxy_address", mcp.Description("proxy address,only support for proxy website"), mcp.Required()),
)

func CreateWebsiteHandle(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if request.Params.Arguments["domain"] == nil {
		return nil, errors.New("domain is required")
	}
	domain := request.Params.Arguments["domain"].(string)
	alias := domain
	var proxyAddress string
	if request.Params.Arguments["website_type"] == "proxy" {
		if request.Params.Arguments["proxy_address"] == nil {
			return nil, errors.New("proxy_address is required")
		}
		proxyAddress = request.Params.Arguments["proxy_address"].(string)
	}

	groupUrl := "/groups/search"
	groupReq := &types.GroupRequest{
		Type: "website",
	}
	groupRes := &types.GroupRes{}
	client := utils.NewPanelClient("POST", groupUrl, utils.WithPayload(groupReq))
	_, err := client.Request(groupRes)
	if err != nil {
		return nil, err
	}
	var groupID uint
	for _, group := range groupRes.Data {
		if group.IsDefault {
			groupID = group.ID
			break
		}
	}

	createUrl := "/websites"
	req := &types.CreateWebsiteRequest{
		PrimaryDomain:  domain,
		Alias:          alias,
		Type:           request.Params.Arguments["website_type"].(string),
		WebsiteGroupID: groupID,
		Proxy:          proxyAddress,
	}
	createCli := utils.NewPanelClient("POST", createUrl, utils.WithPayload(req))
	res := &types.Response{}
	return createCli.Request(res)
}
