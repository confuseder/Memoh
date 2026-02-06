package local

import (
	"context"
	"fmt"
	"strings"

	"github.com/memohai/memoh/internal/channel"
)

// CLIAdapter implements channel.Sender for the local CLI channel.
type CLIAdapter struct {
	hub *channel.SessionHub
}

// NewCLIAdapter creates a CLIAdapter backed by the given session hub.
func NewCLIAdapter(hub *channel.SessionHub) *CLIAdapter {
	return &CLIAdapter{hub: hub}
}

// Type returns the CLI channel type.
func (a *CLIAdapter) Type() channel.ChannelType {
	return CLIType
}

// Send publishes an outbound message to the CLI session hub.
func (a *CLIAdapter) Send(ctx context.Context, cfg channel.ChannelConfig, msg channel.OutboundMessage) error {
	if a.hub == nil {
		return fmt.Errorf("cli hub not configured")
	}
	target := strings.TrimSpace(msg.Target)
	if target == "" {
		return fmt.Errorf("cli target is required")
	}
	if msg.Message.IsEmpty() {
		return fmt.Errorf("message is required")
	}
	a.hub.Publish(target, msg)
	return nil
}
