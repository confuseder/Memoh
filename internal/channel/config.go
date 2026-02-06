package channel

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NormalizeChannelConfig validates and normalizes a channel configuration map
// using the registered descriptor for the given channel type.
func NormalizeChannelConfig(channelType ChannelType, raw map[string]any) (map[string]any, error) {
	if raw == nil {
		raw = map[string]any{}
	}
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return nil, fmt.Errorf("unsupported channel type: %s", channelType)
	}
	if desc.NormalizeConfig == nil {
		return raw, nil
	}
	return desc.NormalizeConfig(raw)
}

// NormalizeChannelUserConfig validates and normalizes a user-channel binding configuration.
func NormalizeChannelUserConfig(channelType ChannelType, raw map[string]any) (map[string]any, error) {
	if raw == nil {
		raw = map[string]any{}
	}
	desc, ok := GetChannelDescriptor(channelType)
	if !ok {
		return nil, fmt.Errorf("unsupported channel type: %s", channelType)
	}
	if desc.NormalizeUserConfig == nil {
		return raw, nil
	}
	return desc.NormalizeUserConfig(raw)
}

// ResolveTargetFromUserConfig derives a delivery target string from a user-channel binding.
func ResolveTargetFromUserConfig(channelType ChannelType, config map[string]any) (string, error) {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok || desc.ResolveTarget == nil {
		return "", fmt.Errorf("unsupported channel type: %s", channelType)
	}
	return desc.ResolveTarget(config)
}

// MatchUserBinding reports whether the given binding config matches the criteria.
func MatchUserBinding(channelType ChannelType, config map[string]any, criteria BindingCriteria) bool {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok || desc.MatchBinding == nil {
		return false
	}
	return desc.MatchBinding(config, criteria)
}

// BuildUserBindingConfig constructs a user-channel binding config from an Identity.
func BuildUserBindingConfig(channelType ChannelType, identity Identity) map[string]any {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok || desc.BuildUserConfig == nil {
		return map[string]any{}
	}
	return desc.BuildUserConfig(identity)
}

// DecodeConfigMap unmarshals a JSON byte slice into a string-keyed map.
func DecodeConfigMap(raw []byte) (map[string]any, error) {
	if len(raw) == 0 {
		return map[string]any{}, nil
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}
	if payload == nil {
		payload = map[string]any{}
	}
	return payload, nil
}

// ReadString looks up the first matching key in a map and returns its string representation.
// It tries each key in order and converts non-string values using type-safe formatting.
func ReadString(raw map[string]any, keys ...string) string {
	for _, key := range keys {
		if value, ok := raw[key]; ok {
			switch v := value.(type) {
			case string:
				return v
			case float64:
				return strconv.FormatFloat(v, 'f', -1, 64)
			case bool:
				return strconv.FormatBool(v)
			default:
				return fmt.Sprintf("%v", v)
			}
		}
	}
	return ""
}
