package channel

import (
	"fmt"
	"strings"
	"sync"
)

// ChannelDescriptor holds all metadata and hooks for a registered channel type.
type ChannelDescriptor struct {
	Type                ChannelType
	DisplayName         string
	NormalizeConfig     func(map[string]any) (map[string]any, error)
	NormalizeUserConfig func(map[string]any) (map[string]any, error)
	ResolveTarget       func(map[string]any) (string, error)
	MatchBinding        func(map[string]any, BindingCriteria) bool
	BuildUserConfig     func(Identity) map[string]any
	Configless          bool
	Capabilities        ChannelCapabilities
	OutboundPolicy      OutboundPolicy
	ConfigSchema        ConfigSchema
	UserConfigSchema    ConfigSchema
	TargetSpec          TargetSpec
	NormalizeTarget     func(string) string
}

type channelRegistry struct {
	mu    sync.RWMutex
	items map[ChannelType]ChannelDescriptor
}

var registry = &channelRegistry{
	items: map[ChannelType]ChannelDescriptor{},
}

// RegisterChannel adds a channel descriptor to the global registry.
func RegisterChannel(desc ChannelDescriptor) error {
	normalized := normalizeChannelType(string(desc.Type))
	if normalized == "" {
		return fmt.Errorf("channel type is required")
	}
	desc.Type = normalized
	if strings.TrimSpace(desc.DisplayName) == "" {
		desc.DisplayName = normalized.String()
	}
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if _, exists := registry.items[desc.Type]; exists {
		return fmt.Errorf("channel type already registered: %s", desc.Type)
	}
	registry.items[desc.Type] = desc
	return nil
}

// MustRegisterChannel calls RegisterChannel and panics on error.
func MustRegisterChannel(desc ChannelDescriptor) {
	if err := RegisterChannel(desc); err != nil {
		panic(err)
	}
}

// UnregisterChannel removes a channel type from the global registry.
func UnregisterChannel(channelType ChannelType) bool {
	normalized := normalizeChannelType(channelType.String())
	if normalized == "" {
		return false
	}
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if _, exists := registry.items[normalized]; !exists {
		return false
	}
	delete(registry.items, normalized)
	return true
}

// GetChannelDescriptor returns the descriptor for the given channel type.
func GetChannelDescriptor(channelType ChannelType) (ChannelDescriptor, bool) {
	normalized := normalizeChannelType(channelType.String())
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	desc, ok := registry.items[normalized]
	return desc, ok
}

// ListChannelDescriptors returns all registered channel descriptors.
func ListChannelDescriptors() []ChannelDescriptor {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	items := make([]ChannelDescriptor, 0, len(registry.items))
	for _, item := range registry.items {
		items = append(items, item)
	}
	return items
}

// GetChannelCapabilities returns the capability matrix for the given channel type.
func GetChannelCapabilities(channelType ChannelType) (ChannelCapabilities, bool) {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return ChannelCapabilities{}, false
	}
	return desc.Capabilities, true
}

// GetChannelOutboundPolicy returns the outbound policy for the given channel type.
func GetChannelOutboundPolicy(channelType ChannelType) (OutboundPolicy, bool) {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return OutboundPolicy{}, false
	}
	return desc.OutboundPolicy, true
}

// GetChannelConfigSchema returns the configuration schema for the given channel type.
func GetChannelConfigSchema(channelType ChannelType) (ConfigSchema, bool) {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return ConfigSchema{}, false
	}
	return desc.ConfigSchema, true
}

// GetChannelUserConfigSchema returns the user-binding configuration schema for the given channel type.
func GetChannelUserConfigSchema(channelType ChannelType) (ConfigSchema, bool) {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return ConfigSchema{}, false
	}
	return desc.UserConfigSchema, true
}

// IsConfigless reports whether the channel type operates without per-bot configuration.
func IsConfigless(channelType ChannelType) bool {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return false
	}
	return desc.Configless
}

func normalizeChannelType(raw string) ChannelType {
	normalized := strings.TrimSpace(strings.ToLower(raw))
	if normalized == "" {
		return ""
	}
	return ChannelType(normalized)
}
