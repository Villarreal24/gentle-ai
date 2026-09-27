package screens

import (
	"strings"
	"testing"

	"github.com/gentleman-programming/gentle-ai/v3/internal/model"
)

func TestRenderDroidModelPicker_ShowsRequestedCopy(t *testing.T) {
	state := NewDroidModelPickerState()
	out := RenderDroidModelPicker(state, 0)

	if !strings.Contains(out, "Factory Droid Model Assignments") {
		t.Fatalf("expected title 'Factory Droid Model Assignments' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Choose how models are assigned to Factory Droid SDD phases and roles") {
		t.Fatalf("expected Droid subtitle in output, got:\n%s", out)
	}
	for _, want := range []string{"balanced", "performance", "economy", "openai", "custom"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected preset %q in output, got:\n%s", want, out)
		}
	}
	if !strings.Contains(out, "Auto") {
		t.Fatalf("expected Auto copy in output, got:\n%s", out)
	}
}

func TestHandleDroidModelPickerNav_SelectsDroidPreset(t *testing.T) {
	state := NewDroidModelPickerState()

	handled, assignments := HandleDroidModelPickerNav("enter", &state, 0)

	if !handled {
		t.Fatal("expected enter on preset to be handled")
	}
	if assignments == nil {
		t.Fatal("expected preset selection to return assignments")
	}
	if got := assignments["default"]; got != model.DroidModelAuto {
		t.Fatalf("default assignment = %q, want %q", got, model.DroidModelAuto)
	}
	if got := assignments["orchestrator"]; got != model.DroidModelAuto {
		t.Fatalf("orchestrator assignment = %q, want auto", got)
	}
	if got := assignments["gentle-orchestrator"]; got != model.DroidModelAuto {
		t.Fatalf("gentle-orchestrator assignment = %q, want auto", got)
	}
}

func TestDroidCustomRowsContainAll14Roles(t *testing.T) {
	state := NewDroidModelPickerState()
	HandleDroidModelPickerNav("enter", &state, 4) // select custom
	out := RenderDroidModelPicker(state, 0)

	for _, want := range []string{
		"Gentleman Orchestrator",
		"Phase 1: Init",
		"Phase 2: Explore",
		"Phase 3: Propose",
		"Phase 4: Spec",
		"Phase 5: Design",
		"Phase 6: Tasks",
		"Phase 7: Apply",
		"Phase 8: Verify",
		"Phase 9: Archive",
		"Worker (Sub-tasks)",
		"Judgment Day Judge A",
		"Judgment Day Judge B",
		"Judgment Day Fix Agent",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("missing active role %q in custom picker output", want)
		}
	}
}

func TestHandleDroidModelPickerNav_CustomCyclesAcrossDroidOptions(t *testing.T) {
	state := NewDroidModelPickerState()

	handled, assignments := HandleDroidModelPickerNav("enter", &state, 4)
	if !handled || assignments != nil || !state.InCustomMode {
		t.Fatalf("expected custom preset to enter custom mode, handled=%v assignments=%v inCustom=%v", handled, assignments, state.InCustomMode)
	}

	handled, assignments = HandleDroidModelPickerNav("enter", &state, 0)
	if !handled || assignments != nil {
		t.Fatalf("expected phase cycle to be handled without confirming, handled=%v assignments=%v", handled, assignments)
	}
	if got := state.CustomAssignments["orchestrator"]; got != model.DroidModelOpus {
		t.Fatalf("first cycle from auto should become opus, got %q", got)
	}

	for _, want := range []model.DroidModelAlias{
		model.DroidModelFable,
		model.DroidModelSonnet,
		model.DroidModelHaiku,
		model.DroidModelGPT56,
		model.DroidModelAuto,
	} {
		handled, _ = HandleDroidModelPickerNav("enter", &state, 0)
		if !handled {
			t.Fatal("expected cycle to be handled")
		}
		if got := state.CustomAssignments["orchestrator"]; got != want {
			t.Fatalf("cycled assignment = %q, want %q", got, want)
		}
	}

	// Confirm action is at cursor len(droidPhases)
	confirmIdx := len(droidPhases)
	handled, saved := HandleDroidModelPickerNav("enter", &state, confirmIdx)
	if !handled || saved == nil {
		t.Fatal("expected confirm to return saved assignments")
	}
	if saved["orchestrator"] != model.DroidModelAuto {
		t.Fatalf("saved orchestrator = %q, want auto", saved["orchestrator"])
	}
	if saved["gentle-orchestrator"] != model.DroidModelAuto {
		t.Fatalf("saved gentle-orchestrator = %q, want auto", saved["gentle-orchestrator"])
	}
}

func TestDroidNamedPresetPreservesCustomKeys(t *testing.T) {
	state := NewDroidModelPickerStateFromAssignments(map[string]model.DroidModelAlias{
		"my-custom-phase": model.DroidModelGPT56,
		"default":         model.DroidModelHaiku,
	})
	_, saved := HandleDroidModelPickerNav("enter", &state, 1) // performance preset
	if saved["my-custom-phase"] != model.DroidModelGPT56 {
		t.Fatalf("custom assignment lost when selecting preset: %v", saved)
	}
	if saved["orchestrator"] != model.DroidModelOpus {
		t.Fatalf("performance preset lost: %v", saved)
	}
}

func TestNewDroidModelPickerStateFromAssignments_MatchesPresetsAndCustom(t *testing.T) {
	balancedState := NewDroidModelPickerStateFromAssignments(model.DroidModelPresetBalanced())
	if balancedState.Preset != DroidPresetBalanced {
		t.Fatalf("balanced preset not detected: got %s", balancedState.Preset)
	}

	customState := NewDroidModelPickerStateFromAssignments(map[string]model.DroidModelAlias{
		"orchestrator": model.DroidModelGPT56,
		"default":      model.DroidModelAuto,
	})
	if customState.Preset != DroidPresetCustom {
		t.Fatalf("custom preset not detected: got %s", customState.Preset)
	}
}
