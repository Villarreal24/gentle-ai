package reviewassets

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gentleman-programming/gentle-ai/v3/internal/agents"
	"github.com/gentleman-programming/gentle-ai/v3/internal/assets"
	"github.com/gentleman-programming/gentle-ai/v3/internal/components/mutationjournal"
	"github.com/gentleman-programming/gentle-ai/v3/internal/model"
)

// OwnershipCommandsLedgerFilename identifies the native rendered-command ownership record.
const OwnershipCommandsLedgerFilename = ".gentle-ai-native-command-ownership.json"

// NativeCommandManifest is an explicit allowlist of slash command files managed per runtime.
var NativeCommandManifest = map[model.AgentID][]string{
	model.AgentDroid: {
		"gentle-apply.md",
		"gentle-archive.md",
		"gentle-design.md",
		"gentle-explore.md",
		"gentle-init.md",
		"gentle-judge.md",
		"gentle-new.md",
		"gentle-orchestrator.md",
		"gentle-propose.md",
		"gentle-sdd-apply.md",
		"gentle-sdd-archive.md",
		"gentle-sdd-design.md",
		"gentle-sdd-explore.md",
		"gentle-sdd-init.md",
		"gentle-sdd-new.md",
		"gentle-sdd-propose.md",
		"gentle-sdd-spec.md",
		"gentle-sdd-tasks.md",
		"gentle-sdd-verify.md",
		"gentle-spec.md",
		"gentle-status.md",
		"gentle-tasks.md",
		"gentle-verify.md",
	},
}

// NativeCommandsSupported reports whether the native command installer handles a runtime.
func NativeCommandsSupported(agent model.AgentID) bool {
	_, ok := NativeCommandManifest[agent]
	return ok
}

// NativeCommandFileNames lists every command file the installer may write for a runtime.
func NativeCommandFileNames(agent model.AgentID) []string {
	if names, ok := NativeCommandManifest[agent]; ok {
		return append([]string(nil), names...)
	}
	return nil
}

func commandLedgerPath(dir string) string {
	return filepath.Join(dir, OwnershipCommandsLedgerFilename)
}

// InstallNativeCommands installs native slash commands for runtimes supporting custom commands.
func InstallNativeCommands(home string, adapter agents.Adapter) (InstallResult, error) {
	if !NativeCommandsSupported(adapter.Agent()) {
		return InstallResult{}, nil
	}
	names := NativeCommandManifest[adapter.Agent()]
	dir := adapter.CommandsDir(home)
	if dir == "" {
		return InstallResult{}, fmt.Errorf("empty native commands directory")
	}

	rendered := make(map[string]string, len(names))
	for _, name := range names {
		sourceName := name
		if strings.HasPrefix(name, "gentle-sdd-") {
			sourceName = "gentle-" + strings.TrimPrefix(name, "gentle-sdd-")
		}
		path := fmt.Sprintf("%s/commands/%s", adapter.Agent(), sourceName)
		source, err := assets.Read(path)
		if err != nil {
			return InstallResult{}, fmt.Errorf("read native command %s: %w", path, err)
		}
		rendered[name] = source
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return InstallResult{}, fmt.Errorf("create native commands directory: %w", err)
	}

	ledgerFile := commandLedgerPath(dir)
	journal := mutationjournal.New(dir)
	if err := journal.Capture(ledgerFile); err != nil {
		return InstallResult{}, fmt.Errorf("capture command ownership ledger: %w", err)
	}

	ledger, _, err := readOwnership(ledgerFile, names)
	if err != nil {
		return InstallResult{}, err
	}

	type candidate struct {
		name, path    string
		data          []byte
		exists, owned bool
	}
	candidates := make([]candidate, 0, len(names))
	result := InstallResult{}

	for _, name := range names {
		path := filepath.Join(dir, name)
		if err := journal.Validate(path); err != nil {
			return result, err
		}
		if err := journal.Capture(path); err != nil {
			return result, fmt.Errorf("capture native command %s: %w", name, err)
		}
		data, exists, err := nativeFile(path)
		if err != nil {
			return result, err
		}
		expected, known := ledger.Files[name]
		owned := exists && known && installedHash(data) == expected
		if exists && !owned {
			result.Skipped = append(result.Skipped, path)
		}
		candidates = append(candidates, candidate{name: name, path: path, data: data, exists: exists, owned: owned})
	}

	rollback := func(err error) (InstallResult, error) {
		return InstallResult{Skipped: result.Skipped}, errors.Join(err, journal.Restore())
	}

	ledgerChanged := false
	for _, c := range candidates {
		if c.exists && !c.owned {
			continue
		}
		desired := []byte(rendered[c.name])
		if !c.exists || string(c.data) != string(desired) {
			if _, err := journal.WriteWithMode(c.path, desired, 0o644); err != nil {
				return rollback(fmt.Errorf("write native command %s: %w", c.name, err))
			}
			result.Changed = true
			result.Files = append(result.Files, c.path)
		}
		actual, err := os.ReadFile(c.path)
		if err != nil {
			return rollback(fmt.Errorf("read installed command %s: %w", c.name, err))
		}
		if string(actual) != string(desired) {
			return rollback(fmt.Errorf("native command changed after write: %s", c.path))
		}
		hash := installedHash(actual)
		if ledger.Files[c.name] != hash {
			ledger.Files[c.name] = hash
			ledgerChanged = true
		}
	}

	if ledgerChanged {
		encoded, err := json.MarshalIndent(ledger, "", "  ")
		if err != nil {
			return rollback(fmt.Errorf("encode command ownership ledger: %w", err))
		}
		encoded = append(encoded, '\n')
		if _, err := journal.WriteWithMode(ledgerFile, encoded, 0o644); err != nil {
			return rollback(fmt.Errorf("write command ownership ledger: %w", err))
		}
		result.Changed = true
		result.Files = append(result.Files, ledgerFile)
	}

	return result, nil
}
