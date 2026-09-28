package reviewassets_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gentleman-programming/gentle-ai/v3/internal/agents"
	"github.com/gentleman-programming/gentle-ai/v3/internal/components/reviewassets"
	"github.com/gentleman-programming/gentle-ai/v3/internal/model"
)

func TestNativeCommandsSupported(t *testing.T) {
	if !reviewassets.NativeCommandsSupported(model.AgentDroid) {
		t.Errorf("NativeCommandsSupported(AgentDroid) = false, want true")
	}
	if reviewassets.NativeCommandsSupported(model.AgentClaudeCode) {
		t.Errorf("NativeCommandsSupported(AgentClaudeCode) = true, want false")
	}
}

func TestInstallNativeCommands_Droid(t *testing.T) {
	home := t.TempDir()
	adapter, err := agents.NewAdapter(model.AgentDroid)
	if err != nil {
		t.Fatal(err)
	}

	result, err := reviewassets.InstallNativeCommands(home, adapter)
	if err != nil {
		t.Fatalf("InstallNativeCommands error: %v", err)
	}
	if !result.Changed {
		t.Errorf("expected Changed = true on fresh install")
	}

	commandsDir := adapter.CommandsDir(home)
	ledgerPath := filepath.Join(commandsDir, reviewassets.OwnershipCommandsLedgerFilename)
	if _, err := os.Stat(ledgerPath); err != nil {
		t.Errorf("ownership ledger missing: %v", err)
	}

	manifest := reviewassets.NativeCommandManifest[model.AgentDroid]
	if len(manifest) == 0 {
		t.Fatal("empty Droid command manifest")
	}
	for _, cmd := range manifest {
		path := filepath.Join(commandsDir, cmd)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("command file %s missing: %v", cmd, err)
		}
	}

	// Idempotency check: second run should report no change
	second, err := reviewassets.InstallNativeCommands(home, adapter)
	if err != nil {
		t.Fatalf("second InstallNativeCommands error: %v", err)
	}
	if second.Changed {
		t.Errorf("expected Changed = false on second install, files modified: %v", second.Files)
	}

	// User-edit preservation check:
	customCmd := filepath.Join(commandsDir, "gentle-sdd-new.md")
	userContent := "# User customized command\n"
	if err := os.WriteFile(customCmd, []byte(userContent), 0o644); err != nil {
		t.Fatal(err)
	}

	third, err := reviewassets.InstallNativeCommands(home, adapter)
	if err != nil {
		t.Fatalf("third InstallNativeCommands error: %v", err)
	}
	if containsFile(third.Files, customCmd) {
		t.Errorf("user customized command was overwritten: %s", customCmd)
	}
	data, err := os.ReadFile(customCmd)
	if err != nil || string(data) != userContent {
		t.Errorf("user custom command was not preserved, got %q", string(data))
	}
}
