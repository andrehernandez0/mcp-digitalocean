package backups

import (
	"context"
	"encoding/json"

	"github.com/digitalocean/godo"
	"github.com/mark3labs/mcp-go/mcp"
)

type SnapshotsTool struct {
	client func(ctx context.Context) (*godo.Client, error)
}

const (
	defaultSnapshotListPage    = 1
	defaultSnapshotListPerPage = 50
)

func NewSnapshotsTool(client func(ctx context.Context) (*godo.Client, error)) *DropletSnapshotTool {
	return &DropletSnapshotTool{client: client}
}

func (s *SnapshotsTool) listSnapshots(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	page, ok := args["Page"].(float64)
	if !ok || page < 1 {
		page = defaultSnapshotListPage
	}
	perPage, ok := args["PerPage"].(float64)
	if !ok || perPage < 1 {
		perPage = defaultSnapshotListPerPage
	}

	options := &godo.ListOptions{
		Page:    int(page),
		PerPage: int(perPage),
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	snapshots, _, err := client.Snapshots.List(ctx, options)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	filteredSnapshots := make([]map[string]any, len(snapshots))
	for i, snapshot := range snapshots {
		filteredSnapshots[i] = map[string]any{
			"id":             snapshot.ID,
			"name":           snapshot.Name,
			"resource_id":    snapshot.ResourceID,
			"resource_type":  snapshot.ResourceType,
			"regions":        snapshot.Regions,
			"min_disk_size":  snapshot.MinDiskSize,
			"size_gigabytes": snapshot.SizeGigaBytes,
			"tags":           snapshot.Tags,
			"created_at":     snapshot.Created,
		}
	}

	jsonSnapshots, err := json.MarshalIndent(filteredSnapshots, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonSnapshots)), nil
}

func (s *SnapshotsTool) listDropletSnapshots(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	page, ok := args["Page"].(float64)
	if !ok || page < 1 {
		page = defaultSnapshotListPage
	}
	perPage, ok := args["PerPage"].(float64)
	if !ok || perPage < 1 {
		perPage = defaultSnapshotListPerPage
	}

	options := &godo.ListOptions{
		Page:    int(page),
		PerPage: int(perPage),
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	snapshots, _, err := client.Snapshots.ListDroplet(ctx, options)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	filteredSnapshots := make([]map[string]any, len(snapshots))
	for i, snapshot := range snapshots {
		filteredSnapshots[i] = map[string]any{
			"id":             snapshot.ID,
			"name":           snapshot.Name,
			"resource_id":    snapshot.ResourceID,
			"resource_type":  snapshot.ResourceType,
			"regions":        snapshot.Regions,
			"min_disk_size":  snapshot.MinDiskSize,
			"size_gigabytes": snapshot.SizeGigaBytes,
			"tags":           snapshot.Tags,
			"created_at":     snapshot.Created,
		}
	}

	jsonSnapshots, err := json.MarshalIndent(filteredSnapshots, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonSnapshots)), nil
}

func (s *SnapshotsTool) listVolumeSnapshots(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	page, ok := args["Page"].(float64)
	if !ok || page < 1 {
		page = defaultSnapshotListPage
	}
	perPage, ok := args["PerPage"].(float64)
	if !ok || perPage < 1 {
		perPage = defaultSnapshotListPerPage
	}

	options := &godo.ListOptions{
		Page:    int(page),
		PerPage: int(perPage),
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	snapshots, _, err := client.Snapshots.ListVolume(ctx, options)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	filteredSnapshots := make([]map[string]any, len(snapshots))
	for i, snapshot := range snapshots {
		filteredSnapshots[i] = map[string]any{
			"id":             snapshot.ID,
			"name":           snapshot.Name,
			"resource_id":    snapshot.ResourceID,
			"resource_type":  snapshot.ResourceType,
			"regions":        snapshot.Regions,
			"min_disk_size":  snapshot.MinDiskSize,
			"size_gigabytes": snapshot.SizeGigaBytes,
			"tags":           snapshot.Tags,
			"created_at":     snapshot.Created,
		}
	}

	jsonSnapshots, err := json.MarshalIndent(filteredSnapshots, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonSnapshots)), nil
}

func (s *SnapshotsTool) getSnapshotByID(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	snapshotID, ok := args["ID"].(string)
	if !ok || snapshotID == "" {
		return mcp.NewToolResultError("Snapshot ID is required"), nil
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	snapshot, _, err := client.Snapshots.Get(ctx, snapshotID)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonSnapshot, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonSnapshot)), nil
}

func (s *SnapshotsTool) deleteSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	snapshotID, ok := args["ID"].(string)
	if !ok || snapshotID == "" {
		return mcp.NewToolResultError("Snapshot ID is required"), nil
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	_, err = client.Snapshots.Delete(ctx, snapshotID)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}
	return mcp.NewToolResultText("Snapshot deleted successfully"), nil
}

func (s *SnapshotsTool) transferDropletSnapshotRegion(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	snapshotID, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("Snapshot ID is required"), nil
	}
	region, ok := args["Region"].(string)
	if !ok || region == "" {
		return mcp.NewToolResultError("Region is required"), nil
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	actionRequest := &godo.ActionRequest{
		"type": "transfer",
		"region": region,
	}

	action, _, err := client.ImageActions.Transfer(ctx, int(snapshotID), actionRequest)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (s *SnapshotsTool) createVolumeSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	volumeID, ok := args["VolumeID"].(string)
	if !ok || volumeID == "" {
		return mcp.NewToolResultError("Volume ID is required"), nil
	}

	snapshotName, ok := args["Name"].(string)
	if !ok || snapshotName == "" {
		return mcp.NewToolResultError("Snapshot name is required"), nil
	}

	tagsArg, _ := args["Tags"].([]any)

	var tags []string
	for _, t := range tagsArg {
		if s, ok := t.(string); ok {
			tags = append(tags, s)
		}
	}

	request := &godo.SnapshotCreateRequest{
		VolumeID: volumeID,
		Name:     snapshotName,
		Tags:     tags,
	}

	client, err := s.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	snapshot, _, err := client.Storage.CreateSnapshot(ctx, request)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonSnapshot, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonSnapshot)), nil
}



