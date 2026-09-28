package droid

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gentleman-programming/gentle-ai/v3/internal/model"
	"github.com/gentleman-programming/gentle-ai/v3/internal/system"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		name            string
		lookPathPath    string
		lookPathErr     error
		stat            statResult
		wantInstalled   bool
		wantBinaryPath  string
		wantConfigFound bool
		wantErr         bool
	}{
		{
			name:            "binary and config directory found",
			lookPathPath:    "/usr/local/bin/droid",
			stat:            statResult{isDir: true},
			wantInstalled:   true,
			wantBinaryPath:  "/usr/local/bin/droid",
			wantConfigFound: true,
		},
		{
			name:            "binary missing config missing",
			lookPathErr:     errors.New("missing"),
			stat:            statResult{err: os.ErrNotExist},
			wantInstalled:   false,
			wantBinaryPath:  "",
			wantConfigFound: false,
		},
		{
			name:            "binary found config missing",
			lookPathPath:    "/usr/local/bin/droid",
			stat:            statResult{err: os.ErrNotExist},
			wantInstalled:   true,
			wantBinaryPath:  "/usr/local/bin/droid",
			wantConfigFound: false,
		},
		{
			name:            "binary missing config found",
			lookPathErr:     errors.New("missing"),
			stat:            statResult{isDir: true},
			wantInstalled:   false,
			wantBinaryPath:  "",
			wantConfigFound: true,
		},
		{
			name:    "stat error propagates",
			stat:    statResult{err: errors.New("permission denied")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				lookPath: func(string) (string, error) {
					return tt.lookPathPath, tt.lookPathErr
				},
				statPath: func(string) statResult {
					return tt.stat
				},
			}
			homeDir := filepath.Join(string(filepath.Separator), "home", "test")

			installed, binaryPath, configPath, configFound, err := a.Detect(context.Background(), homeDir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Detect() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if installed != tt.wantInstalled {
				t.Fatalf("Detect() installed = %v, want %v", installed, tt.wantInstalled)
			}

			if binaryPath != tt.wantBinaryPath {
				t.Fatalf("Detect() binaryPath = %q, want %q", binaryPath, tt.wantBinaryPath)
			}

			wantConfigPath := filepath.Join(homeDir, ".factory")
			if configPath != wantConfigPath {
				t.Fatalf("Detect() configPath = %q, want %q", configPath, wantConfigPath)
			}

			if configFound != tt.wantConfigFound {
				t.Fatalf("Detect() configFound = %v, want %v", configFound, tt.wantConfigFound)
			}
		})
	}
}

func TestInstallCommand(t *testing.T) {
	a := NewAdapter()

	commands, err := a.InstallCommand(system.PlatformProfile{})
	if err == nil {
		t.Fatalf("InstallCommand() error = nil, want non-installable error")
	}
	if commands != nil {
		t.Fatalf("InstallCommand() commands = %v, want nil", commands)
	}

	var notInstallable AgentNotInstallableError
	if !errors.As(err, &notInstallable) {
		t.Fatalf("InstallCommand() error type = %T, want AgentNotInstallableError", err)
	}
	if got := err.Error(); !strings.Contains(got, "must be installed manually") {
		t.Fatalf("InstallCommand() error = %q, want message containing 'must be installed manually'", got)
	}
}

func TestConfigPaths(t *testing.T) {
	a := NewAdapter()
	homeDir := filepath.Join(string(filepath.Separator), "home", "test")
	configDir := filepath.Join(homeDir, ".factory")
	settingsJSON := filepath.Join(configDir, "settings.json")
	mcpJSON := filepath.Join(configDir, "mcp.json")
	agentsMD := filepath.Join(configDir, "AGENTS.md")
	skillsDir := filepath.Join(configDir, "skills")
	droidsDir := filepath.Join(configDir, "droids")

	tests := []struct {
		name string
		got  string
		want string
	}{
		{"GlobalConfigDir", a.GlobalConfigDir(homeDir), configDir},
		{"SystemPromptDir", a.SystemPromptDir(homeDir), configDir},
		{"SystemPromptFile", a.SystemPromptFile(homeDir), agentsMD},
		{"SkillsDir", a.SkillsDir(homeDir), skillsDir},
		{"SubAgentsDir", a.SubAgentsDir(homeDir), droidsDir},
		{"SettingsPath", a.SettingsPath(homeDir), settingsJSON},
		{"MCPConfigPath (engram)", a.MCPConfigPath(homeDir, "engram"), mcpJSON},
		{"MCPConfigPath (context7)", a.MCPConfigPath(homeDir, "context7"), mcpJSON},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("%s = %q, want %q", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestCapabilities(t *testing.T) {
	a := NewAdapter()

	if got := a.Agent(); got != model.AgentDroid {
		t.Fatalf("Agent() = %q, want %q", got, model.AgentDroid)
	}

	if got := a.Tier(); got != model.TierFull {
		t.Fatalf("Tier() = %v, want TierFull", got)
	}

	if a.SupportsOutputStyles() {
		t.Fatalf("SupportsOutputStyles() = true, want false")
	}

	if got := a.OutputStyleDir("/home/test"); got != "" {
		t.Fatalf("OutputStyleDir() = %q, want empty", got)
	}

	if !a.SupportsSlashCommands() {
		t.Fatalf("SupportsSlashCommands() = false, want true")
	}

	if got := a.CommandsDir("/home/test"); got != filepath.Join("/home/test", ".factory", "commands") {
		t.Fatalf("CommandsDir() = %q, want ~/.factory/commands", got)
	}

	if !a.SupportsSubAgents() {
		t.Fatalf("SupportsSubAgents() = false, want true")
	}

	if got := a.SubAgentsDir("/home/test"); got != filepath.Join("/home/test", ".factory", "droids") {
		t.Fatalf("SubAgentsDir() = %q, want ~/.factory/droids", got)
	}

	if got := a.EmbeddedSubAgentsDir(); got != "droid/agents" {
		t.Fatalf("EmbeddedSubAgentsDir() = %q, want 'droid/agents'", got)
	}

	if !a.SupportsSkills() {
		t.Fatalf("SupportsSkills() = false, want true")
	}

	if !a.SupportsSystemPrompt() {
		t.Fatalf("SupportsSystemPrompt() = false, want true")
	}

	if !a.SupportsMCP() {
		t.Fatalf("SupportsMCP() = false, want true")
	}

	if got := a.SystemPromptStrategy(); got != model.StrategyMarkdownSections {
		t.Fatalf("SystemPromptStrategy() = %v, want StrategyMarkdownSections", got)
	}

	if got := a.MCPStrategy(); got != model.StrategyMCPConfigFile {
		t.Fatalf("MCPStrategy() = %v, want StrategyMCPConfigFile", got)
	}
}

func TestSystemPromptFileResolution(t *testing.T) {
	a := NewAdapter()

	t.Run("defaults to factory AGENTS.md when no files exist", func(t *testing.T) {
		tempDir := t.TempDir()
		got := a.SystemPromptFile(tempDir)
		want := filepath.Join(tempDir, ".factory", "AGENTS.md")
		if got != want {
			t.Fatalf("SystemPromptFile() = %q, want %q", got, want)
		}
	})

	t.Run("prefers factory system.md if present", func(t *testing.T) {
		tempDir := t.TempDir()
		factoryDir := filepath.Join(tempDir, ".factory")
		if err := os.MkdirAll(factoryDir, 0o755); err != nil {
			t.Fatal(err)
		}
		systemMD := filepath.Join(factoryDir, "system.md")
		if err := os.WriteFile(systemMD, []byte("# System"), 0o644); err != nil {
			t.Fatal(err)
		}

		got := a.SystemPromptFile(tempDir)
		if got != systemMD {
			t.Fatalf("SystemPromptFile() = %q, want %q", got, systemMD)
		}
	})

	t.Run("prefers workspace AGENTS.md if present", func(t *testing.T) {
		tempDir := t.TempDir()
		agentsMD := filepath.Join(tempDir, "AGENTS.md")
		if err := os.WriteFile(agentsMD, []byte("# Agents"), 0o644); err != nil {
			t.Fatal(err)
		}

		got := a.SystemPromptFile(tempDir)
		if got != agentsMD {
			t.Fatalf("SystemPromptFile() = %q, want %q", got, agentsMD)
		}
	})
}

