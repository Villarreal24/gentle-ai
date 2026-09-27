package screens

import (
	"fmt"
	"maps"
	"strings"

	"github.com/gentleman-programming/gentle-ai/v3/internal/model"
	"github.com/gentleman-programming/gentle-ai/v3/internal/tui/styles"
)

type DroidModelPreset string

const (
	DroidPresetBalanced    DroidModelPreset = "balanced"
	DroidPresetPerformance DroidModelPreset = "performance"
	DroidPresetEconomy     DroidModelPreset = "economy"
	DroidPresetOpenWeight  DroidModelPreset = "open-weight"
	DroidPresetCustom      DroidModelPreset = "custom"
)

var droidPresetDescriptions = map[DroidModelPreset]string{
	DroidPresetBalanced:    "Factory Droid Auto for orchestrator and all sub-agents",
	DroidPresetPerformance: "Opus for orchestrator & high-cognition phases, Sonnet for execution",
	DroidPresetEconomy:     "Cost-optimized: Sonnet/o3-mini for reasoning, Auto for routine execution",
	DroidPresetOpenWeight:  "DeepSeek Reasoner for planning & judging, Qwen Coder for execution",
	DroidPresetCustom:      "Pick the model for orchestrator, SDD phases, and Judgment Day individually",
}

var droidPresetOrder = []DroidModelPreset{
	DroidPresetBalanced,
	DroidPresetPerformance,
	DroidPresetEconomy,
	DroidPresetOpenWeight,
	DroidPresetCustom,
}

var droidPresetConstructors = map[DroidModelPreset]func() map[string]model.DroidModelAlias{
	DroidPresetBalanced:    model.DroidModelPresetBalanced,
	DroidPresetPerformance: model.DroidModelPresetPerformance,
	DroidPresetEconomy:     model.DroidModelPresetEconomy,
	DroidPresetOpenWeight:  model.DroidModelPresetOpenWeight,
}

var droidAliasOrder = []model.DroidModelAlias{
	model.DroidModelAuto,
	model.DroidModelOpus,
	model.DroidModelSonnet,
	model.DroidModelHaiku,
	model.DroidModelO3Mini,
	model.DroidModelGPT4o,
	model.DroidModelDeepSeek,
	model.DroidModelGemini,
	model.DroidModelQwen,
}

var droidPhases = []string{
	"orchestrator",
	"gentle-ai-init",
	"gentle-ai-explore",
	"gentle-ai-propose",
	"gentle-ai-spec",
	"gentle-ai-design",
	"gentle-ai-tasks",
	"gentle-ai-apply",
	"gentle-ai-verify",
	"gentle-ai-archive",
	"gentle-ai-worker",
	"jd-judge-a",
	"jd-judge-b",
	"jd-fix-agent",
}

var droidPhaseLabels = map[string]string{
	"orchestrator":      "Gentleman Orchestrator",
	"gentle-ai-init":     "Phase 1: Init",
	"gentle-ai-explore":  "Phase 2: Explore",
	"gentle-ai-propose":  "Phase 3: Propose",
	"gentle-ai-spec":     "Phase 4: Spec",
	"gentle-ai-design":   "Phase 5: Design",
	"gentle-ai-tasks":    "Phase 6: Tasks",
	"gentle-ai-apply":    "Phase 7: Apply",
	"gentle-ai-verify":   "Phase 8: Verify",
	"gentle-ai-archive":  "Phase 9: Archive",
	"gentle-ai-worker":   "Worker (Sub-tasks)",
	"jd-judge-a":         "Judgment Day Judge A",
	"jd-judge-b":         "Judgment Day Judge B",
	"jd-fix-agent":       "Judgment Day Fix Agent",
}

// DroidModelPickerState holds navigation state for the Factory Droid model picker screen.
type DroidModelPickerState struct {
	Preset            DroidModelPreset
	CustomAssignments map[string]model.DroidModelAlias
	InCustomMode      bool
}

func NewDroidModelPickerState() DroidModelPickerState {
	return DroidModelPickerState{
		Preset:            DroidPresetBalanced,
		CustomAssignments: model.DroidModelPresetBalanced(),
		InCustomMode:      false,
	}
}

func NewDroidModelPickerStateFromAssignments(assignments map[string]model.DroidModelAlias) DroidModelPickerState {
	if len(assignments) == 0 {
		return NewDroidModelPickerState()
	}
	for preset, constructor := range droidPresetConstructors {
		if droidAssignmentsEqual(constructor(), assignments) {
			return DroidModelPickerState{
				Preset:            preset,
				CustomAssignments: maps.Clone(assignments),
				InCustomMode:      false,
			}
		}
	}
	return DroidModelPickerState{
		Preset:            DroidPresetCustom,
		CustomAssignments: maps.Clone(assignments),
		InCustomMode:      false,
	}
}

func droidAssignmentsEqual(a, b map[string]model.DroidModelAlias) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func HandleDroidModelPickerNav(
	key string,
	state *DroidModelPickerState,
	cursor int,
) (handled bool, assignments map[string]model.DroidModelAlias) {
	if !state.InCustomMode {
		return handleDroidPresetNav(key, state, cursor)
	}
	return handleDroidCustomPhaseNav(key, state, cursor)
}

func handleDroidPresetNav(
	key string,
	state *DroidModelPickerState,
	cursor int,
) (bool, map[string]model.DroidModelAlias) {
	if key != "enter" {
		return false, nil
	}
	if cursor >= len(droidPresetOrder) {
		return false, nil
	}

	selected := droidPresetOrder[cursor]
	state.Preset = selected
	if selected == DroidPresetCustom {
		state.InCustomMode = true
		if state.CustomAssignments == nil {
			state.CustomAssignments = model.DroidModelPresetBalanced()
		}
		return true, nil
	}

	assignments := droidPresetConstructors[selected]()
	for key, alias := range state.CustomAssignments {
		if _, active := assignments[key]; !active {
			assignments[key] = alias
		}
	}
	if orch, ok := assignments["orchestrator"]; ok {
		assignments["gentle-orchestrator"] = orch
	}
	state.CustomAssignments = maps.Clone(assignments)
	return true, assignments
}

func handleDroidCustomPhaseNav(
	key string,
	state *DroidModelPickerState,
	cursor int,
) (bool, map[string]model.DroidModelAlias) {
	switch key {
	case "esc":
		state.InCustomMode = false
		return true, nil
	case "enter":
		if cursor < len(droidPhases) {
			phase := droidPhases[cursor]
			next := nextDroidAlias(state.CustomAssignments[phase])
			state.CustomAssignments[phase] = next
			if phase == "orchestrator" {
				state.CustomAssignments["gentle-orchestrator"] = next
			}
			return true, nil
		}
		if cursor == len(droidPhases) {
			cloned := maps.Clone(state.CustomAssignments)
			if orch, ok := cloned["orchestrator"]; ok {
				cloned["gentle-orchestrator"] = orch
			}
			return true, cloned
		}
		state.InCustomMode = false
		return true, nil
	}
	return false, nil
}

func nextDroidAlias(current model.DroidModelAlias) model.DroidModelAlias {
	for i, alias := range droidAliasOrder {
		if alias == current {
			return droidAliasOrder[(i+1)%len(droidAliasOrder)]
		}
	}
	return model.DroidModelAuto
}

func DroidModelPickerOptionCount(state DroidModelPickerState) int {
	if state.InCustomMode {
		return len(droidPhases) + 2 // phase rows + Confirm + Back
	}
	return len(droidPresetOrder) + 1 // presets + Back
}

func RenderDroidModelPicker(state DroidModelPickerState, cursor int) string {
	if state.InCustomMode {
		return renderDroidCustomPhaseList(state, cursor)
	}
	return renderDroidPresetList(state, cursor)
}

func renderDroidPresetList(state DroidModelPickerState, cursor int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("Factory Droid Model Assignments"))
	b.WriteString("\n\n")
	b.WriteString(styles.SubtextStyle.Render("Choose how models are assigned to Factory Droid SDD phases and roles:"))
	b.WriteString("\n\n")

	for idx, preset := range droidPresetOrder {
		isSelected := preset == state.Preset
		focused := idx == cursor
		b.WriteString(renderRadio(string(preset), isSelected, focused))
		b.WriteString(styles.SubtextStyle.Render("    "+droidPresetDescriptions[preset]) + "\n")
	}

	b.WriteString("\n")
	b.WriteString(renderOptions([]string{"← Back"}, cursor-len(droidPresetOrder)))
	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate • enter: select • esc: back"))

	return b.String()
}

func renderDroidCustomPhaseList(state DroidModelPickerState, cursor int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("Custom Factory Droid Model Assignments"))
	b.WriteString("\n\n")
	b.WriteString(styles.SubtextStyle.Render("Press enter on a role to cycle: auto → opus → sonnet → haiku → o3-mini → gpt-4o → deepseek → gemini → qwen"))
	b.WriteString("\n\n")

	for idx, phase := range droidPhases {
		focused := idx == cursor
		alias := state.CustomAssignments[phase]
		if alias == "" {
			alias = model.DroidModelAuto
		}

		label := fmt.Sprintf("%-26s %s", droidPhaseLabels[phase], droidAliasTag(alias))

		if focused {
			b.WriteString(styles.SelectedStyle.Render(styles.Cursor+label) + "\n")
		} else {
			b.WriteString(styles.UnselectedStyle.Render("  "+label) + "\n")
		}
	}

	b.WriteString("\n")
	actionCursor := cursor - len(droidPhases)
	b.WriteString(renderOptions([]string{"Confirm", "← Back"}, actionCursor))
	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate • enter: cycle/select • esc: back"))

	return b.String()
}

func droidAliasTag(alias model.DroidModelAlias) string {
	switch alias {
	case model.DroidModelAuto:
		return styles.SuccessStyle.Render("[auto]")
	case model.DroidModelOpus:
		return styles.WarningStyle.Render("[opus]")
	case model.DroidModelSonnet:
		return styles.SuccessStyle.Render("[sonnet]")
	case model.DroidModelHaiku:
		return styles.SubtextStyle.Render("[haiku]")
	case model.DroidModelO3Mini:
		return styles.WarningStyle.Render("[o3-mini]")
	case model.DroidModelGPT4o:
		return styles.SuccessStyle.Render("[gpt-4o]")
	case model.DroidModelDeepSeek:
		return styles.WarningStyle.Render("[deepseek]")
	case model.DroidModelGemini:
		return styles.SuccessStyle.Render("[gemini]")
	case model.DroidModelQwen:
		return styles.SubtextStyle.Render("[qwen]")
	default:
		return styles.SuccessStyle.Render("[auto]")
	}
}
