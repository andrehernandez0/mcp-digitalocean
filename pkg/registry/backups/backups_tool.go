package backups

import (
	"context"
	"encoding/json"

	"github.com/digitalocean/godo"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type BackupsTool struct {
	client func(ctx context.Context) (*godo.Client, error)
}

const (
	defaultBackupListPage    = 1
	defaultBackupListPerPage = 50
)

func NewBackupsTool(client func(ctx context.Context) (*godo.Client, error)) *BackupsTool {
	return &BackupsTool{client: client}
}

func (b *BackupsTool) listDropletBackups(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletID, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("ID is required"), nil
	}
	page, ok := args["Page"].(float64)
	if !ok || page < 1 {
		page = defaultBackupListPage
	}
	perPage, ok := args["PerPage"].(float64)
	if !ok || perPage < 1 {
		perPage = defaultBackupListPerPage
	}

	opt := &godo.ListOptions{
		Page:    int(page),
		PerPage: int(perPage),
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	backups, _, err := client.Droplets.Backups(ctx, int(dropletID), opt)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonBackups, err := json.MarshalIndent(backups, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonBackups)), nil
}

func (b *BackupsTool) enableDropletBackups(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletID, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("ID is required"), nil
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	action, _, err := client.DropletActions.EnableBackups(ctx, int(dropletID))
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *BackupsTool) enableDropletBackupsWithPolicy(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletID, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("ID is required"), nil
	}

	plan, ok := args["Plan"].(string)
	if !ok {
		return mcp.NewToolResultError("Plan is required"), nil
	}
	weekday, ok := args["Weekday"].(string)
	if !ok {
		return mcp.NewToolResultError("Weekday is required"), nil
	}
	hour, ok := args["Hour"].(float64)

	hourInt := int(hour)
	policy := &godo.DropletBackupPolicyRequest{
		Plan:    plan,
		Weekday: weekday,
		Hour:    &hourInt,
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	action, _, err := client.DropletActions.EnableBackupsWithPolicy(ctx, int(dropletID), policy)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *BackupsTool) modifyDropletBackupPolicy(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletID, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("ID is required"), nil
	}

	plan, ok := args["Plan"].(string)
	weekday, ok := args["Weekday"].(string)
	hour, ok := args["Hour"].(float64)

	hourInt := int(hour)
	policy := &godo.DropletBackupPolicyRequest{
		Plan:    plan,
		Weekday: weekday,
		Hour:    &hourInt,
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	action, _, err := client.DropletActions.ChangeBackupPolicy(ctx, int(dropletID), policy)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *BackupsTool) getDropletBackupPolicy(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()

	id, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("ID is required"), nil
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	policy, _, err := client.Droplets.GetBackupPolicy(ctx, int(id))
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonData, err := json.MarshalIndent(policy, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}
	return mcp.NewToolResultText(string(jsonData)), nil
}

func (b *BackupsTool) disableDropletBackups(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletID, ok := args["ID"].(float64)
	if !ok {
		return mcp.NewToolResultError("ID is required"), nil
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	action, _, err := client.DropletActions.DisableBackups(ctx, int(dropletID))
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *BackupsTool) createDropletFromBackup(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	dropletName := args["Name"].(string)
	size := args["Size"].(string)
	backupImageID := args["BackupImageID"].(float64)
	region := args["Region"].(string)
	backup, _ := args["Backup"].(bool)         // Defaults to false
	monitoring, _ := args["Monitoring"].(bool) // Defaults to false

	// Handle SSH keys if provided
	var sshKeys []godo.DropletCreateSSHKey
	if sshKeysRaw, ok := args["SSHKeys"]; ok && sshKeysRaw != nil {
		sshKeysList := sshKeysRaw.([]interface{})
		for _, key := range sshKeysList {
			switch v := key.(type) {
			case float64:
				sshKeys = append(sshKeys, godo.DropletCreateSSHKey{ID: int(v)})
			case string:
				sshKeys = append(sshKeys, godo.DropletCreateSSHKey{Fingerprint: v})
			}
		}
	}

	// Handle tags if provided
	var tags []string
	if tagsRaw, ok := args["Tags"]; ok && tagsRaw != nil {
		tagsList := tagsRaw.([]interface{})
		for _, tag := range tagsList {
			if tagStr, ok := tag.(string); ok {
				tags = append(tags, tagStr)
			}
		}
	}

	// Create the droplet
	dropletCreateRequest := &godo.DropletCreateRequest{
		Name:       dropletName,
		Size:       size,
		Image:      godo.DropletCreateImage{ID: int(backupImageID)},
		Region:     region,
		Backups:    backup,
		Monitoring: monitoring,
		SSHKeys:    sshKeys,
		Tags:       tags,
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	droplet, _, err := client.Droplets.Create(ctx, dropletCreateRequest)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("droplet create", err), nil
	}
	jsonDroplet, err := json.MarshalIndent(droplet, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("json marshal", err), nil
	}
	return mcp.NewToolResultText(string(jsonDroplet)), nil
}

func (b *BackupsTool) restoreDropletFromBackup(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	dropletID := req.GetArguments()["ID"].(float64)
	imageID := req.GetArguments()["ImageID"].(float64)

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	action, _, err := client.DropletActions.Restore(ctx, int(dropletID), int(imageID))
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *BackupsTool) convertBackupToSnapshot(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := req.GetArguments()
	backupImageID, ok := args["BackupImageID"].(float64)
	if !ok {
		return mcp.NewToolResultError("BackupID is required"), nil
	}

	client, err := b.client(ctx)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("Error getting DigitalOcean client", err), nil
	}

	action, _, err := client.ImageActions.Convert(ctx, int(backupImageID))
	if err != nil {
		return mcp.NewToolResultErrorFromErr("api error", err), nil
	}

	jsonAction, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return mcp.NewToolResultErrorFromErr("marshal error", err), nil
	}

	return mcp.NewToolResultText(string(jsonAction)), nil
}

func (b *BackupsTool) Tools() []server.ServerTool {
	tools := []server.ServerTool{
		{
			Handler: b.listDropletBackups,
			Tool: mcp.NewTool(
				"droplet-backup-list",
				mcp.WithDescription("List backups for a droplet. Supports pagination."),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to list backups for")),
				mcp.WithNumber("Page", mcp.DefaultNumber(defaultBackupListPage), mcp.Description("Page number")),
				mcp.WithNumber("PerPage", mcp.DefaultNumber(defaultBackupListPerPage), mcp.Description("Backups per page")),
			),
		},
		{
			Handler: b.enableDropletBackups,
			Tool: mcp.NewTool(
				"droplet-backup-enable",
				mcp.WithDescription("Enable backups for a droplet"),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to enable backups for")),
			),
		},
		{
			Handler: b.enableDropletBackupsWithPolicy,
			Tool: mcp.NewTool(
				"droplet-backup-enable-with-policy",
				mcp.WithDescription("Enable backups for a droplet with a policy"),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to enable backups for")),
				mcp.WithString("Plan", mcp.Required(), mcp.Description("The plan to use for backups."), mcp.Enum("daily", "weekly")),
				mcp.WithString("Weekday", mcp.Required(), mcp.Description("The weekday to backup"), mcp.Enum("mon", "tue", "wed", "thu", "fri", "sat", "sun")),
				mcp.WithNumber("Hour", mcp.Description("The hour to backup")),
			),
		},
		{
			Handler: b.modifyDropletBackupPolicy,
			Tool: mcp.NewTool(
				"droplet-backup-modify-policy",
				mcp.WithDescription("Modify the backup policy for a droplet"),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to modify the backup policy for")),
				mcp.WithString("Plan", mcp.Required(), mcp.Description("The plan to use for backups"), mcp.Enum("daily", "weekly")),
				mcp.WithString("Weekday", mcp.Required(), mcp.Description("The weekday to backup"), mcp.Enum("mon", "tue", "wed", "thu", "fri", "sat", "sun")),
				mcp.WithNumber("Hour", mcp.Description("The hour to backup. Must be between 0 and 23."), mcp.Min(0), mcp.Max(23)),
			),
		},
		{
			Handler: b.getDropletBackupPolicy,
			Tool: mcp.NewTool(
				"droplet-backup-get-policy",
				mcp.WithDescription("Get the backup policy for a droplet"),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to get the backup policy for")),
			),
		},
		{
			Handler: b.disableDropletBackups,
			Tool: mcp.NewTool(
				"droplet-backup-disable",
				mcp.WithDescription("Disable backups for a droplet"),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to disable backups for")),
			),
		},
		{
			Handler: b.createDropletFromBackup,
			Tool: mcp.NewTool("droplet-create",
				mcp.WithDescription("Create a new droplet from a backup image"),
				mcp.WithString("Name", mcp.Required(), mcp.Description("Name of the droplet")),
				mcp.WithString("Size", mcp.Required(), mcp.Description("Slug of the droplet size (e.g., s-1vcpu-1gb)")),
				mcp.WithNumber("BackupImageID", mcp.Required(), mcp.Description("ID of the backup image to create the droplet from")),
				mcp.WithString("Region", mcp.Required(), mcp.Description("Slug of the region (e.g., nyc3)")),
				mcp.WithBoolean("Backup", mcp.DefaultBool(false), mcp.Description("Whether to enable backups")),
				mcp.WithBoolean("Monitoring", mcp.DefaultBool(false), mcp.Description("Whether to enable monitoring")),
				mcp.WithArray("SSHKeys", mcp.Description("Array of SSH key IDs (numbers) or fingerprints (strings) to add to the droplet"), mcp.Items(map[string]any{"type": "string"})),
				mcp.WithArray("Tags", mcp.Description("Array of tag names to apply to the droplet"), mcp.Items(map[string]any{"type": "string"})),
			),
		},
		{
			Handler: b.restoreDropletFromBackup,
			Tool: mcp.NewTool(
				"droplet-backup-restore",
				mcp.WithDescription("Restore a droplet from a backup image"),
				mcp.WithNumber("ID", mcp.Required(), mcp.Description("The ID of the droplet to restore")),
				mcp.WithNumber("ImageID", mcp.Required(), mcp.Description("The ID of the backup image to restore from")),
			),
		},
		{
			Handler: b.convertBackupToSnapshot,
			Tool: mcp.NewTool(
				"backup-to-snapshot",
				mcp.WithDescription("Convert a backup to a snapshot"),
				mcp.WithNumber("BackupImageID", mcp.Required(), mcp.Description("The ID of the backup image to convert")),
			),
		},
	}
	return tools
}