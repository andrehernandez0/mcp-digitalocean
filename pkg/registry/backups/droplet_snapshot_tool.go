package backups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/mark3labs/mcp-go/mcp"
)

type DropletSnapshotTool struct {
	client func(ctx context.Context) (*godo.Client, error)
}

const (
	defaultDropletSnapshotListPage    = 1
	defaultDropletSnapshotListPerPage = 50
)

func NewDropletSnapshotTool(client func(ctx context.Context) (*godo.Client, error)) *DropletSnapshotTool {
	return &DropletSnapshotTool{client: client}
}

func (ds *DropletSnapshotTool) listDropletSnapshots(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletID, ok := args["dropletID"].(float64)
	if !ok {
		return mcp.NewToolResultError("Droplet ID is required"), nil
	}

	page, ok := args["Page"].(float64)
	if !ok || page < 1 {
		page = defaultDropletSnapshotListPage
	}
	perPage, ok := args["PerPage"].(float64)
	if !ok || perPage < 1 {
		perPage = defaultDropletSnapshotListPerPage
	}

	client, err := ds.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	options := &godo.ListOptions{
		Page:    int(page),
		PerPage: int(perPage),
	}

	snapshots, _, err := client.Droplets.Snapshots(ctx, int(dropletID), options)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	filteredSnapshots := make([]map[string]any, len(snapshots))
	for i, snapshot := range snapshots {
		filteredSnapshots[i] = map[string]any{
			"id":             snapshot.ID,
			"name":           snapshot.Name,
			"created_at":     snapshot.Created,
			"regions":        snapshot.Regions,
			"min_disk_size":  snapshot.MinDiskSize,
			"size_gigabytes": snapshot.SizeGigaBytes,
			"distribution":   snapshot.Distribution,
			"tags":           snapshot.Tags,
			"status":         snapshot.Status,
		}
	}

	jsonSnapshots, err := json.MarshalIndent(filteredSnapshots, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonSnapshots)), nil
}

func (ds *DropletSnapshotTool) createDropletSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dropletID := req.GetArguments()["ID"].(float64)
	name := req.GetArguments()["Name"].(string)

	client, err := ds.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get DigitalOcean client: %w", err)
	}

	action, _, err := client.DropletActions.Snapshot(ctx, int(dropletID), name)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *DropletSnapshotTool) deleteDropletSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}

func (b *DropletSnapshotTool) createDropletFromSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}

func (b *DropletSnapshotTool) restoreDropletFromSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}

func (b *DropletSnapshotTool) addDropletSnapshotToRegion(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}

func (b *DropletSnapshotTool) transferDropletSnapshotTeam(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}
