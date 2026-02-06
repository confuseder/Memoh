package local

import (
	"context"
	"fmt"
	"strings"

	"github.com/memohai/memoh/internal/channel"
)

// WebAdapter implements channel.Sender for the local Web channel.
type WebAdapter struct {
	hub *channel.SessionHub
}

// NewWebAdapter creates a WebAdapter backed by the given session hub.
func NewWebAdapter(hub *channel.SessionHub) *WebAdapter {
	return &WebAdapter{hub: hub}
}

// Type returns the Web channel type.
func (a *WebAdapter) Type() channel.ChannelType {
	return WebType
}

// Send publishes an outbound message to the Web session hub.
func (a *WebAdapter) Send(ctx context.Context, cfg channel.ChannelConfig, msg channel.OutboundMessage) error {
	if a.hub == nil {
		return fmt.Errorf("web hub not configured")
	}
	target := strings.TrimSpace(msg.Target)
	if target == "" {
		return fmt.Errorf("web target is required")
	}
	if msg.Message.IsEmpty() {
		return fmt.Errorf("message is required")
	}
	a.hub.Publish(target, msg)
	return nil
}
