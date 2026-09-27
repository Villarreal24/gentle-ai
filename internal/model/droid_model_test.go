package model

import "testing"

func TestDroidModelAlias_Valid(t *testing.T) {
	for _, tc := range []struct {
		alias DroidModelAlias
		want  bool
	}{
		{DroidModelAuto, true},
		{DroidModelOpus, true},
		{DroidModelFable, true},
		{DroidModelSonnet, true},
		{DroidModelHaiku, true},
		{DroidModelGPT56, true},
		{"opus", true},
		{"fable", true},
		{"sonnet", true},
		{"o3-mini", false},
		{"gpt-4o", false},
		{"deepseek", false},
		{"gemini", false},
		{"qwen", false},
		{"unknown", false},
		{"", false},
	} {
		if got := tc.alias.Valid(); got != tc.want {
			t.Errorf("Valid(%q) = %v, want %v", tc.alias, got, tc.want)
		}
	}
}

func TestDroidModelID(t *testing.T) {
	for _, tc := range []struct {
		alias DroidModelAlias
		want  string
	}{
		{DroidModelAuto, "auto"},
		{DroidModelOpus, "claude-opus-5.5"},
		{"opus", "claude-opus-5.5"},
		{DroidModelFable, "claude-fable-5.1"},
		{"fable", "claude-fable-5.1"},
		{DroidModelSonnet, "claude-sonnet-5"},
		{"sonnet", "claude-sonnet-5"},
		{DroidModelHaiku, "claude-haiku-4.5"},
		{DroidModelGPT56, "gpt-5.6"},
		{"gpt-5.6", "gpt-5.6"},
		{"custom", "auto"},
	} {
		if got := DroidModelID(tc.alias); got != tc.want {
			t.Errorf("DroidModelID(%q) = %v, want %v", tc.alias, got, tc.want)
		}
	}
}

func TestDroidModelPresets(t *testing.T) {
	presets := []struct {
		name   string
		values map[string]DroidModelAlias
	}{
		{"balanced", DroidModelPresetBalanced()},
		{"performance", DroidModelPresetPerformance()},
		{"economy", DroidModelPresetEconomy()},
		{"openai", DroidModelPresetOpenAI()},
	}

	requiredKeys := []string{
		"orchestrator", "gentle-orchestrator", "gentle-ai-init", "gentle-ai-explore", "gentle-ai-propose",
		"gentle-ai-spec", "gentle-ai-design", "gentle-ai-tasks", "gentle-ai-apply",
		"gentle-ai-verify", "gentle-ai-archive", "gentle-ai-worker",
		"jd-judge-a", "jd-judge-b", "jd-fix-agent", "default",
	}

	for _, p := range presets {
		for _, key := range requiredKeys {
			val, ok := p.values[key]
			if !ok {
				t.Errorf("preset %s missing key %s", p.name, key)
			}
			if !val.Valid() {
				t.Errorf("preset %s has invalid alias for key %s: %s", p.name, key, val)
			}
		}
	}
}
