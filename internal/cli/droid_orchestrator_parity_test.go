package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gentleman-programming/gentle-ai/v3/internal/agents"
	"github.com/gentleman-programming/gentle-ai/v3/internal/model"
	"github.com/gentleman-programming/gentle-ai/v3/internal/state"
	"github.com/gentleman-programming/gentle-ai/v3/internal/system"
)

func TestDroidOrchestratorInstallAndSyncParity(t *testing.T) {
	home := installTestHome(t)

	_, err := RunInstall([]string{"--agent", "droid", "--preset", "full-gentleman"}, system.DetectionResult{})
	if err != nil {
		t.Fatalf("RunInstall error = %v", err)
	}

	adapter, err := agents.NewAdapter(model.AgentDroid)
	if err != nil {
		t.Fatal(err)
	}
	droidsDir := adapter.SubAgentsDir(home)

	// 1. Verify gentle-orchestrator.md sub-droid exists and has the coordinator definition
	orchPath := filepath.Join(droidsDir, "gentle-orchestrator.md")
	data, err := os.ReadFile(orchPath)
	if err != nil {
		t.Fatalf("gentle-orchestrator.md missing: %v", err)
	}
	orchContent := string(data)

	for _, expectedText := range []string{
		"Master SDD / ODD Coordinator droid for Gentle-AI",
		"## Core SDD Pipeline",
		"**Init** (`@gentle-ai-init`)",
		"**Apply** (`@gentle-ai-apply`)",
		"**Verify** (`@gentle-ai-verify`)",
		"Adversarial Review",
		"Judgment Day",
	} {
		if !strings.Contains(orchContent, expectedText) {
			t.Errorf("gentle-orchestrator.md missing expected text: %q", expectedText)
		}
	}

	// 2. Verify ~/.factory/AGENTS.md carries the full Gentleman Orchestrator prompt with SDD pipeline
	agentsMdPath := filepath.Join(home, ".factory", "AGENTS.md")
	agentsData, err := os.ReadFile(agentsMdPath)
	if err != nil {
		t.Fatalf("~/.factory/AGENTS.md missing: %v", err)
	}
	agentsContent := string(agentsData)

	for _, expectedText := range []string{
		"Gentle AI — ODD Orchestrator Instructions (Factory Droid)",
		"Agent Teams Orchestrator",
		"Route work through Factory Droid's native sub-droids:",
		"@gentle-ai-init",
		"@gentle-ai-apply",
		"@gentle-ai-verify",
		"Judgment Day adversarial judges",
		"⏳ Delegating",
		"completed",
		"Senior Architect, 15+ years experience, GDE & MVP",
	} {
		if !strings.Contains(agentsContent, expectedText) {
			t.Errorf("AGENTS.md missing expected text: %q", expectedText)
		}
	}

	// 3. Verify all 13 sub-agents exist
	subAgents := []string{
		"gentle-ai-init.md",
		"gentle-ai-explore.md",
		"gentle-ai-propose.md",
		"gentle-ai-spec.md",
		"gentle-ai-design.md",
		"gentle-ai-tasks.md",
		"gentle-ai-apply.md",
		"gentle-ai-verify.md",
		"gentle-ai-archive.md",
		"gentle-ai-worker.md",
		"jd-judge-a.md",
		"jd-judge-b.md",
		"jd-fix-agent.md",
	}
	for _, sub := range subAgents {
		p := filepath.Join(droidsDir, sub)
		if _, err := os.Stat(p); err != nil {
			t.Errorf("sub-droid %s missing: %v", sub, err)
		}
	}
}

func TestDroidModelAssignmentsSubstitutionAndPersistence(t *testing.T) {
	home := installTestHome(t)

	// Persist DroidModelAssignments in state.json
	err := state.Write(home, state.InstallState{
		InstalledAgents:     []string{"droid"},
		SelectionConfigured: true,
		Components:          []model.ComponentID{model.ComponentPersona, model.ComponentSDD},
		Preset:              model.PresetFullGentleman,
		Persona:             string(model.PersonaGentleman),
		DroidModelAssignments: map[string]string{
			"orchestrator":    "opus",
			"gentle-ai-apply": "deepseek",
			"gentle-ai-spec":  "o3-mini",
			"default":         "sonnet",
		},
	})
	if err != nil {
		t.Fatalf("state.Write error = %v", err)
	}

	// Run sync and check that persisted models are loaded and rendered
	result, err := RunSync([]string{"--agents", "droid"})
	if err != nil {
		t.Fatalf("RunSync error = %v", err)
	}

	// Verify selection received the assignments from state
	if got := result.Selection.DroidModelAssignments["orchestrator"]; got != model.DroidModelOpus {
		t.Errorf("sync selection orchestrator = %q, want opus", got)
	}
	if got := result.Selection.DroidModelAssignments["gentle-ai-apply"]; got != model.DroidModelDeepSeek {
		t.Errorf("sync selection gentle-ai-apply = %q, want deepseek", got)
	}

	adapter, err := agents.NewAdapter(model.AgentDroid)
	if err != nil {
		t.Fatal(err)
	}
	droidsDir := adapter.SubAgentsDir(home)

	// Check gentle-orchestrator model substitution
	orchData, err := os.ReadFile(filepath.Join(droidsDir, "gentle-orchestrator.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(orchData), "model: claude-opus-4.5") {
		t.Errorf("gentle-orchestrator model not substituted with claude-opus-4.5, got:\n%s", string(orchData)[:200])
	}

	// Check gentle-ai-apply model substitution
	applyData, err := os.ReadFile(filepath.Join(droidsDir, "gentle-ai-apply.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(applyData), "model: deepseek-reasoner") {
		t.Errorf("gentle-ai-apply model not substituted with deepseek-reasoner, got:\n%s", string(applyData)[:200])
	}

	// Check gentle-ai-spec model substitution
	specData, err := os.ReadFile(filepath.Join(droidsDir, "gentle-ai-spec.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(specData), "model: o3-mini") {
		t.Errorf("gentle-ai-spec model not substituted with o3-mini, got:\n%s", string(specData)[:200])
	}

	// Check default fallback model substitution (gentle-ai-explore)
	exploreData, err := os.ReadFile(filepath.Join(droidsDir, "gentle-ai-explore.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(exploreData), "model: claude-3-7-sonnet") {
		t.Errorf("gentle-ai-explore model not substituted with default claude-3-7-sonnet, got:\n%s", string(exploreData)[:200])
	}
}
