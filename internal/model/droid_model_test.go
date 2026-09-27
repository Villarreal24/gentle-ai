package model

import "testing"

func TestDroidModelAlias_Valid(t *testing.T) {
	for _, tc := range []struct {
		alias DroidModelAlias
		want  bool
	}{
		{DroidModelAuto, true},
		{DroidModelOpus, true},
		{DroidModelSonnet, true},
		{DroidModelHaiku, true},
		{DroidModelO3Mini, true},
		{DroidModelGPT4o, true},
		{DroidModelDeepSeek, true},
		{DroidModelGemini, true},
		{DroidModelQwen, true},
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
		{DroidModelOpus, "claude-opus-4.5"},
		{DroidModelSonnet, "claude-3-7-sonnet"},
		{DroidModelHaiku, "claude-haiku-4.5"},
		{DroidModelO3Mini, "o3-mini"},
		{DroidModelGPT4o, "gpt-4o"},
		{DroidModelDeepSeek, "deepseek-reasoner"},
		{DroidModelGemini, "gemini-2.5-pro"},
		{DroidModelQwen, "qwen3-coder-next"},
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
		{"open-weight", DroidModelPresetOpenWeight()},
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
