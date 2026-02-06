package channel

import "strings"

// TargetHint provides a display label and example for a target format.
type TargetHint struct {
	Example string `json:"example,omitempty"`
	Label   string `json:"label,omitempty"`
}

// TargetSpec describes the expected format of a delivery target for a channel type.
type TargetSpec struct {
	Format string       `json:"format"`
	Hints  []TargetHint `json:"hints,omitempty"`
}

// NormalizeTarget applies the channel-specific target normalization function.
// It returns the normalized string and true if a normalizer was found, otherwise the trimmed input and false.
func NormalizeTarget(channelType ChannelType, raw string) (string, bool) {
	desc, ok := GetChannelDescriptor(channelType)
	if !ok || desc.NormalizeTarget == nil {
		return strings.TrimSpace(raw), false
	}
	normalized := strings.TrimSpace(desc.NormalizeTarget(raw))
	if normalized == "" {
		return "", false
	}
	return normalized, true
}
