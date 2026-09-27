package model

// DroidModelAlias represents a Factory Droid-native model choice for role assignments.
type DroidModelAlias string

const (
	DroidModelAuto     DroidModelAlias = "auto"
	DroidModelOpus     DroidModelAlias = "opus"
	DroidModelSonnet   DroidModelAlias = "sonnet"
	DroidModelHaiku    DroidModelAlias = "haiku"
	DroidModelO3Mini   DroidModelAlias = "o3-mini"
	DroidModelGPT4o    DroidModelAlias = "gpt-4o"
	DroidModelDeepSeek DroidModelAlias = "deepseek"
	DroidModelGemini   DroidModelAlias = "gemini"
	DroidModelQwen     DroidModelAlias = "qwen"
)

// Valid reports whether the alias is one of the known Droid model options.
func (a DroidModelAlias) Valid() bool {
	switch a {
	case DroidModelAuto, DroidModelOpus, DroidModelSonnet, DroidModelHaiku,
		DroidModelO3Mini, DroidModelGPT4o, DroidModelDeepSeek, DroidModelGemini, DroidModelQwen:
		return true
	default:
		return false
	}
}

// DroidModelID maps a DroidModelAlias to the model identifier Factory Droid expects
// in the `model:` field of a custom agent frontmatter.
func DroidModelID(alias DroidModelAlias) string {
	switch alias {
	case DroidModelAuto:
		return "auto"
	case DroidModelOpus:
		return "claude-opus-4.5"
	case DroidModelSonnet:
		return "claude-3-7-sonnet"
	case DroidModelHaiku:
		return "claude-haiku-4.5"
	case DroidModelO3Mini:
		return "o3-mini"
	case DroidModelGPT4o:
		return "gpt-4o"
	case DroidModelDeepSeek:
		return "deepseek-reasoner"
	case DroidModelGemini:
		return "gemini-2.5-pro"
	case DroidModelQwen:
		return "qwen3-coder-next"
	default:
		return "auto"
	}
}

// DroidModelPresetBalanced lets Factory Droid route all roles automatically with its auto-model.
func DroidModelPresetBalanced() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelAuto,
		"gentle-orchestrator": DroidModelAuto,
		"gentle-ai-init":      DroidModelAuto,
		"gentle-ai-explore":   DroidModelAuto,
		"gentle-ai-propose":   DroidModelAuto,
		"gentle-ai-spec":      DroidModelAuto,
		"gentle-ai-design":    DroidModelAuto,
		"gentle-ai-tasks":     DroidModelAuto,
		"gentle-ai-apply":     DroidModelAuto,
		"gentle-ai-verify":    DroidModelAuto,
		"gentle-ai-archive":   DroidModelAuto,
		"gentle-ai-worker":    DroidModelAuto,
		"jd-judge-a":          DroidModelAuto,
		"jd-judge-b":          DroidModelAuto,
		"jd-fix-agent":        DroidModelAuto,
		"default":             DroidModelAuto,
	}
}

// DroidModelPresetPerformance prioritizes frontier Claude-family models.
func DroidModelPresetPerformance() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelOpus,
		"gentle-orchestrator": DroidModelOpus,
		"gentle-ai-init":      DroidModelSonnet,
		"gentle-ai-explore":   DroidModelSonnet,
		"gentle-ai-propose":   DroidModelOpus,
		"gentle-ai-spec":      DroidModelOpus,
		"gentle-ai-design":    DroidModelOpus,
		"gentle-ai-tasks":     DroidModelSonnet,
		"gentle-ai-apply":     DroidModelSonnet,
		"gentle-ai-verify":    DroidModelOpus,
		"gentle-ai-archive":   DroidModelSonnet,
		"gentle-ai-worker":    DroidModelSonnet,
		"jd-judge-a":          DroidModelOpus,
		"jd-judge-b":          DroidModelOpus,
		"jd-fix-agent":        DroidModelSonnet,
		"default":             DroidModelSonnet,
	}
}

// DroidModelPresetEconomy optimizes token costs with auto/lightweight models for routine work.
func DroidModelPresetEconomy() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelAuto,
		"gentle-orchestrator": DroidModelAuto,
		"gentle-ai-init":      DroidModelAuto,
		"gentle-ai-explore":   DroidModelAuto,
		"gentle-ai-propose":   DroidModelSonnet,
		"gentle-ai-spec":      DroidModelO3Mini,
		"gentle-ai-design":    DroidModelO3Mini,
		"gentle-ai-tasks":     DroidModelAuto,
		"gentle-ai-apply":     DroidModelAuto,
		"gentle-ai-verify":    DroidModelAuto,
		"gentle-ai-archive":   DroidModelAuto,
		"gentle-ai-worker":    DroidModelAuto,
		"jd-judge-a":          DroidModelAuto,
		"jd-judge-b":          DroidModelAuto,
		"jd-fix-agent":        DroidModelAuto,
		"default":             DroidModelAuto,
	}
}

// DroidModelPresetOpenWeight favors open-weight models.
func DroidModelPresetOpenWeight() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelDeepSeek,
		"gentle-orchestrator": DroidModelDeepSeek,
		"gentle-ai-init":      DroidModelQwen,
		"gentle-ai-explore":   DroidModelQwen,
		"gentle-ai-propose":   DroidModelDeepSeek,
		"gentle-ai-spec":      DroidModelDeepSeek,
		"gentle-ai-design":    DroidModelDeepSeek,
		"gentle-ai-tasks":     DroidModelQwen,
		"gentle-ai-apply":     DroidModelQwen,
		"gentle-ai-verify":    DroidModelDeepSeek,
		"gentle-ai-archive":   DroidModelQwen,
		"gentle-ai-worker":    DroidModelQwen,
		"jd-judge-a":          DroidModelDeepSeek,
		"jd-judge-b":          DroidModelDeepSeek,
		"jd-fix-agent":        DroidModelQwen,
		"default":             DroidModelQwen,
	}
}
