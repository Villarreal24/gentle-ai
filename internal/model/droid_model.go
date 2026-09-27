package model

// DroidModelAlias represents a Factory Droid-native model choice for role assignments.
type DroidModelAlias string

const (
	DroidModelAuto   DroidModelAlias = "auto"
	DroidModelOpus   DroidModelAlias = "opus-5.5"
	DroidModelFable  DroidModelAlias = "fable-5.1"
	DroidModelSonnet DroidModelAlias = "sonnet-5"
	DroidModelHaiku  DroidModelAlias = "haiku"
	DroidModelGPT56  DroidModelAlias = "gpt-5.6"
)

// Valid reports whether the alias is one of the known Droid model options.
func (a DroidModelAlias) Valid() bool {
	switch a {
	case DroidModelAuto, DroidModelOpus, DroidModelFable, DroidModelSonnet, DroidModelHaiku, DroidModelGPT56,
		// Backwards-compatible aliases
		"opus", "fable", "sonnet":
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
	case DroidModelOpus, "opus":
		return "claude-opus-5.5"
	case DroidModelFable, "fable":
		return "claude-fable-5.1"
	case DroidModelSonnet, "sonnet":
		return "claude-sonnet-5"
	case DroidModelHaiku:
		return "claude-haiku-4.5"
	case DroidModelGPT56:
		return "gpt-5.6"
	default:
		return "auto"
	}
}

// DroidModelPresetBalanced lets Factory Droid route all roles automatically with its internal Auto Model router.
// Routing is handled natively by Factory Droid, not by Gentle-AI.
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

// DroidModelPresetPerformance prioritizes frontier Anthropic models (Opus 5.5 & Fable 5.1 for architecture/review, Sonnet 5 for execution).
func DroidModelPresetPerformance() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelOpus,
		"gentle-orchestrator": DroidModelOpus,
		"gentle-ai-init":      DroidModelSonnet,
		"gentle-ai-explore":   DroidModelSonnet,
		"gentle-ai-propose":   DroidModelOpus,
		"gentle-ai-spec":      DroidModelFable,
		"gentle-ai-design":    DroidModelFable,
		"gentle-ai-tasks":     DroidModelSonnet,
		"gentle-ai-apply":     DroidModelSonnet,
		"gentle-ai-verify":    DroidModelOpus,
		"gentle-ai-archive":   DroidModelSonnet,
		"gentle-ai-worker":    DroidModelSonnet,
		"jd-judge-a":          DroidModelFable,
		"jd-judge-b":          DroidModelFable,
		"jd-fix-agent":        DroidModelSonnet,
		"default":             DroidModelSonnet,
	}
}

// DroidModelPresetEconomy optimizes token costs with Haiku and Factory Droid Auto Model fallback.
func DroidModelPresetEconomy() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelAuto,
		"gentle-orchestrator": DroidModelAuto,
		"gentle-ai-init":      DroidModelAuto,
		"gentle-ai-explore":   DroidModelAuto,
		"gentle-ai-propose":   DroidModelSonnet,
		"gentle-ai-spec":      DroidModelSonnet,
		"gentle-ai-design":    DroidModelSonnet,
		"gentle-ai-tasks":     DroidModelHaiku,
		"gentle-ai-apply":     DroidModelAuto,
		"gentle-ai-verify":    DroidModelAuto,
		"gentle-ai-archive":   DroidModelHaiku,
		"gentle-ai-worker":    DroidModelAuto,
		"jd-judge-a":          DroidModelSonnet,
		"jd-judge-b":          DroidModelSonnet,
		"jd-fix-agent":        DroidModelAuto,
		"default":             DroidModelAuto,
	}
}

// DroidModelPresetOpenAI favors OpenAI's flagship GPT-5.6 for orchestration and deep reasoning.
func DroidModelPresetOpenAI() map[string]DroidModelAlias {
	return map[string]DroidModelAlias{
		"orchestrator":        DroidModelGPT56,
		"gentle-orchestrator": DroidModelGPT56,
		"gentle-ai-init":      DroidModelAuto,
		"gentle-ai-explore":   DroidModelAuto,
		"gentle-ai-propose":   DroidModelGPT56,
		"gentle-ai-spec":      DroidModelGPT56,
		"gentle-ai-design":    DroidModelGPT56,
		"gentle-ai-tasks":     DroidModelSonnet,
		"gentle-ai-apply":     DroidModelGPT56,
		"gentle-ai-verify":    DroidModelGPT56,
		"gentle-ai-archive":   DroidModelHaiku,
		"gentle-ai-worker":    DroidModelSonnet,
		"jd-judge-a":          DroidModelGPT56,
		"jd-judge-b":          DroidModelGPT56,
		"jd-fix-agent":        DroidModelSonnet,
		"default":             DroidModelSonnet,
	}
}
