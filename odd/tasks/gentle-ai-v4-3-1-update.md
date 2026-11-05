# Feature: Update Gentle-AI Fork to v4.3.1

## Objective

Move `feat/factory-droid-adapter` onto the official Gentle-AI v4.3.1 stable
release while preserving the six Factory Droid commits, then install the
verified build locally and synchronize only the user-authorized assets.

## Problem

The fork is based on `a9e36e9b8a4d7885244466cd9ea6cc3ad330a69b` and has six
Factory Droid commits that must remain intact. Official v4.3.1 is reported to
contain 859 upstream commits not present in the fork. The update must not lose
the adapter or overwrite personal Droid assets and preferences.

## Why

The user wants the fork to follow the latest stable release without abandoning
the Factory Droid integration or the locally installed fork binary.

## Scope

- Preserve the existing feature branch and establish local recovery points.
- Verify the official `v4.3.1` tag, then merge it into a temporary branch based
  on the current feature branch. Do not reset, rebase, or discard fork commits.
- Resolve the v4.3.1 conflicts while preserving both sides. On 2026-11-05, the
  user authorized a minimal v4 Droid migration: update module imports, Droid
  prompt/agent assets, routing hashes and related tests, and the README agent
  badge. Stop if unrelated failures or broader design changes are required.
- Run focused Go tests, build and inspect a temporary binary, and independently
  review the merge and local configuration effects.
- Replace `~/bin/gentle-ai` only after verification.
- Review `sync --dry-run` before any real synchronization. Preserve user-owned
  files and preferences; pause if the effective sync would update tools or
  assets outside the authorized scope.
- Advance and normally push `feat/factory-droid-adapter` only to
  `Villarreal24/gentle-ai`, after confirming the active GitHub identity is
  `Villarreal24`. Do not publish upstream or open a PR.

## Constraints

- Work only in this Gentle-AI checkout and its explicitly authorized local
  installation paths. Do not modify `cars_platform`.
- Do not integrate commits after the official v4.3.1 tag.
- Keep the user-approved Droid migration limited to the v4 module path, Droid
  guidance/assets, routing golden/hash, README badge, and tests directly related
  to those changes. Stop before addressing unrelated CLI failures. The user
  authorized skipping only `TestRenderedGuidanceHasNoRetiredWorkflowReferences/codex`
  because the `codex` executable is unavailable; do not install or configure it.
- Do not run `gentle-ai upgrade`, update the Engram binary, run Engram setup, or
  write directly to Engram's database.
- Regenerating the Engram integration is authorized; do not copy credentials
  into this repository. Do not refresh unrelated tools or shared global
  configuration without stopping for approval.
- Preserve custom files, including
  `~/.factory/droids/scrutiny-feature-reviewer.md` and
  `~/.factory/droids/user-testing-flow-validator.md`.
- Do not expose or commit credential-bearing configuration. If synchronization
  needs to change `~/.factory/mcp.json`, keep any recovery copy outside the
  repository with owner-only permissions and preserve unrelated entries.
- Keep generated local backups outside the repository.

## Tasks

- [x] **T1. Establish recovery points and record this work.**
  - Route: inline for the single task document and bounded local backups;
    read-only mapping delegated to `gentle-ai-explore` because the sync path
    spans more than four files.
  - Risk: medium for local backups; passive for this document.
  - Evidence: the original branch was clean at `29ec8b82`; local tag
    `backup/factory-droid-adapter-pre-v4.3.1-2026-11-05` points to that commit;
    the pre-update binary backup matches SHA-256
    `c7a68e1068424fc10c7f650504e43c97f91784ceceae548f6d8a8a30bf84bde6`.
    Managed local backups were verified under
    `~/Documents/personal/.gentle-ai-backups/2026-11-05-v4.3.1-29ec8b82/managed-pre-sync/`;
    no `~/.factory/system.md` existed. `~/.factory/mcp.json` was copied
    separately outside the repository with mode `0600`; its contents were not
    read. Temporary branch `update/gentle-ai-v4.3.1` is based on the original
    feature head. A pre-activation recovery set was also created outside the
    checkout at
    `~/Documents/personal/.gentle-ai-backups/2026-11-05-v4.3.1-29ec8b82/activation-pre-install-2026-11-05/`.
    Its Git bundle contains both branches at `29ec8b82` and `ed825830` and
    verifies as complete history. It includes a pre-registration copy of the
    pending task document (SHA-256
    `b5423f69b061faabeaeb0b95fa5b1ae73a72f0b63c61744d0b103fcd134ada35`) and
    the pre-activation binary (SHA-256
    `c7a68e1068424fc10c7f650504e43c97f91784ceceae548f6d8a8a30bf84bde6`).
    The untracked `.codegraph/` directory remains in place and untouched.
  - Work-unit commit: `7843d964` (`docs(odd): track Gentle-AI v4.3.1 update`);
    completion-evidence commit: `13b82a4e`.

- [x] **T2. Integrate the official v4.3.1 tag.**
  - Route: conflict mapping and resolution delegated to `gentle-ai-apply`;
    parent performs and commits the Git merge because that writer cannot commit.
    The six-commit adapter footprint spans 76 files, and the upstream release
    changes 1,845 paths. Do not parallelize writers.
  - Refresh `origin` and `upstream`, verify the tag and its official source,
    then merge it into the temporary branch based on the original feature
    branch. Preserve all six fork commits.
  - Tag verification: annotated tag object
    `ee20d2dded25c8e1f1071f92ff73f1e357f36071` points to commit
    `2dba7784114fff57864519c4c0d067e952a8c98b`. Both object IDs match
    `git ls-remote upstream`; the tag is unsigned (verification reported no
    signature found). The original shared base is an ancestor and the tag adds
    exactly 859 commits. Fetched `upstream/main` has 200 commits after the tag;
    none are in scope. `origin/feat/factory-droid-adapter` remains at
    `29ec8b82`.
  - The original feature branch to tag tree diff is 1,845 paths
    (`+163,023/-13,047`), all imported upstream changes for this release.
  - The Git merge found 17 content conflicts; all 17 were resolved and staged
    by `gentle-ai-apply`. No unresolved paths remain. Merge commit
    `ed825830ad0c62833d82304ce992749138281b8e` has parents
    `0ee790351b069a4f2a209af98b44702a708caba3` and
    `2dba7784114fff57864519c4c0d067e952a8c98b`.
  - The authorized migration changed the 19 `/v3` imports in seven Droid files
    to `/v4`, updated 23 Droid assets, the README badge, and the routing digest.
    The Droid digest is
    `0208d789909c092b32c94e97bbbefd2785ca44147d3ee4a77cc6de8640a7a29a`;
    the Pi routing fixture remained byte-identical (SHA-256
    `aa459a16017fc6a507cf4d87783068782ba6250fac05c5eed57a35f47a9caca6`).
  - Focused routing, guidance, embedded-asset, README badge, Droid adapter,
    native command installation, TUI picker, and CLI Droid parity/model
    persistence checks passed. The only authorized skip was
    `TestRenderedGuidanceHasNoRetiredWorkflowReferences/codex`, because
    `codex --version` exits 127 (`codex not found in PATH`).
  - `git diff --check` passed for the new migration changes. The parent
    `git diff --cached --check` found whitespace in imported v4.3.1 content:
    `README.md:46`, `internal/tui/styles/rose_frames.txt`, and
    `odd/tasks/opencode-settings-writer-gaps.md`. The latter two index blobs
    match the release tag exactly, and the README line also has trailing
    whitespace in the tag. No unrelated upstream whitespace was changed.
  - No prompt tests were weakened; the routing golden was regenerated only via
    the existing `-update` path. Merge commit `ed825830` is complete.
  - The user authorized the bounded Droid v4 migration listed above. Use
    applicable test-first checks; stop before fixing unrelated failures.
  - Independent verification passed: all six fork commits remain ancestors,
    only the approved 32 migration paths were changed beyond the imported
    release tree, and no actionable migration defect was found. The verifier
    also confirmed that real global sync may affect additional components;
    inspect dry-run and ask before applying any out-of-scope effect.
  - Work-unit merge commit: `ed825830` (`chore(upstream): merge Gentle-AI v4.3.1`).

- [x] **T3. Verify the merge and build before installation.**
  - Route: writer self-verification plus an independent read-only verifier;
    risk is high because the change affects the CLI, persisted model
    preferences, and local sync behavior.
  - Run the focused Go checks listed under Verification, build a temporary
    binary, and verify its version and command help. Review the exact diff
    against the six fork commits and assess local config effects.
  - Do not install or sync if required checks fail or the independent review
    finds an unresolved high-impact issue.
  - Evidence: all T2-focused commands passed except the user-authorized Codex
    subtest; the parent reran the routing golden test successfully. The
    independent verifier found no in-scope defect and confirmed all six fork
    commits remain ancestors.
  - Temporary build command:

    ```sh
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=4.3.1-villarreal24.1" -o /tmp/gentle-ai-v4.3.1-factory-droid-merge-check-0ee790351 ./cmd/gentle-ai
    ```

    It passed with Go `1.26.6` on `darwin/arm64` (the repo declares Go
    `1.25.10`).
    The binary reports `gentle-ai 4.3.1-villarreal24.1`; `--help` and
    `sync --help` both passed. `MINISIGN_PUBLIC_KEYS_CANONICAL` was unavailable,
    so this temporary binary keeps the updater's fail-closed default. It has
    not been installed.

- [ ] **T4. Install and synchronize within the authorized scope.**
  - Route: bounded local operation after validation; risk is high because sync
    can modify global configuration.
  - Replace the active binary only after its temporary build passes.
  - Run `sync --dry-run` first. The current sync implementation's dry-run
    reports components and step counts, not a per-file diff. A normal sync can
    refresh Persona, Engram, Context7, GGA, Skills, shared skills, and
    persisted CodeGraph wiring even when `--agents droid` is specified.
  - Independent v4 source review confirmed that selecting `droid` does not
    fence those global stages. `--scope workspace` skips global `~/.factory`
    assets and cannot satisfy this task. Do not run real sync until the dry-run
    shows an authorized component set or the user approves additional effects.
  - Proceed only if the effective actions stay within the approved Droid
    assets and Engram integration. If unrelated tools/configuration would be
    refreshed, stop and ask before applying sync.
  - Preserve native ownership semantics, user-edited files, custom droids,
    model assignments, and other preferences. Do not run setup or update
    another tool.

- [ ] **T5. Advance and publish the verified feature branch.**
  - Route: parent-controlled Git operation after all checks; push risk is high.
  - Fast-forward the original `feat/factory-droid-adapter` branch to the
    verified temporary branch. Confirm the authenticated GitHub account is
    `Villarreal24`; then use a normal push only to `origin`.
  - Stop if the intended account is unavailable. Never force-push, push to
    `upstream`, or create a PR.

## Acceptance Criteria

- The verified official `v4.3.1` tag is in the updated branch history, and all
  six Factory Droid commits remain represented without reset or rebase.
- Focused tests pass except the user-authorized, reported Codex-dependent
  subtest; the temporary binary reports the expected fork version and exposes
  the expected commands.
- The Droid integration's module imports, prompts/assets, routing hashes, and
  related tests comply with the v4.3.1 contracts.
- The installed binary matches the verified temporary build.
- Any sync stays within the approved local assets and preserves user files and
  preferences; it does not update the Engram binary/setup/database or unrelated
  tools.
- The original feature branch is updated and pushed normally only to
  `Villarreal24/gentle-ai`.
- No changes are made to `cars_platform`.

## Verification

T2 focused migration checks passed, except the user-authorized Codex subtest:

```sh
go test ./internal/components/agentguidance -run '^TestRenderRoutingMatchesGoldens$' -update
go test ./internal/components/agentguidance -run '^TestRenderRoutingMatchesGoldens$' -count=1
go test ./internal/components/agentguidance -run '^(TestInstalledEvidenceBudgetGuidance|TestInstalledOrchestratorDelegatesForReason|TestVerifierRunsOnlyForANamedHighRiskItem|TestInstalledPromptCarriesRDDOnlyOnRDDRuntimes|TestRenderedGuidanceHasNoRetiredWorkflowReferences)$' -skip 'TestRenderedGuidanceHasNoRetiredWorkflowReferences/codex$' -count=1
go test ./internal/assets -run '^(TestEmbeddedAssetsHaveNoRetiredWorkflowReferences|TestOrchestratorsDelegateOnlyWhenATriggerTrulyHolds|TestOrchestratorsProjectOrganicRouting|TestOrchestratorSectionDriftDoesNotGrow)$' -count=1
go test ./internal/catalog -run '^TestReadmeAgentsBadgeMatchesCatalog$' -count=1
go test ./internal/agents/droid -count=1
go test ./internal/components/reviewassets -run '^(TestNativeCommandsSupported|TestInstallNativeCommands_Droid)$' -count=1
go test ./internal/tui/screens -run Droid -count=1
go test ./internal/cli -run '^TestDroid(OrchestratorInstallAndSyncParity|ModelAssignmentsSubstitutionAndPersistence)$' -count=1 -timeout=3m
```

`git diff --cached --check` is not clean because the imported v4.3.1 tag
contains whitespace in the files recorded under T2. No unrelated upstream
whitespace was changed. The T3 temporary build, command checks, and independent
review passed; the recorded Codex-dependent subtest remains the only authorized
skip.

## Deferred Checks

- `go test ./...` and the broad `internal/cli` suite were not run for the final
  migration. The broad CLI suite timed out after 10 minutes in an earlier
  disposable pre-migration copy with failures not attributed to this change;
  focused Droid CLI tests passed.
- The repository-declared Go version is `1.25.10`; the available compiler was
  Go `1.26.6`. The temporary build passed on that newer compiler.
- Real sync remains pending review of its dry-run. The v4 sync can touch global
  components beyond Droid and Engram; do not run it until the effective scope
  is authorized.

`TestRenderedGuidanceHasNoRetiredWorkflowReferences/codex` is an authorized
environmental skip because `codex` is not available in `PATH`. Do not install
or configure another tool to satisfy it.

## Delivery Strategy

Use `ask-on-risk` (default). The imported upstream history and the six
pre-existing fork commits are not new authored work for this update. Count new
authored changes before each work-unit commit; if they exceed about 400 lines,
ask once about the chain strategy before committing.

## Progress

- Route: delegated read-only mapping to `gentle-ai-explore` (completed). It
  mapped the six commits to a 76-file footprint and found that `--agents droid`
  is not an asset-only sync fence.
- T1 closed by `7843d964` and completion evidence recorded in `13b82a4e`:
  binary, Git tag, managed-asset backups, private MCP backup, temporary branch,
  and task record are verified.
- `origin` and `upstream` refs were fetched and the official v4.3.1 tag was
  verified. The merge is committed as `ed825830`; the original feature branch
  and `origin/feat/factory-droid-adapter` remain at `29ec8b82`.
- User approved a minimal Droid v4 migration: update the seven stale `/v3`
  imports, Droid guidance/assets, routing hash/test expectations, README agent
  badge, and directly related tests. Unrelated CLI failures remain out of scope.
- The migration changed the 19 imports, 23 Droid assets, README badge,
  parity expectations, and routing digest; those changes are part of merge
  commit `ed825830`.
  Focused checks passed except the user-authorized Codex-dependent subtest; the
  independent review and temporary build passed.
- The temporary build used Go 1.26.6, not the repository-declared Go 1.25.10;
  no release trust-anchor variable was available, so the temp binary's updater
  remains fail-closed. It has not replaced the active binary.
- An untracked `.codegraph/.gitignore` remains outside the commit and was not
  inspected or changed.

## Next Step

Install the verified temporary binary and inspect `sync --dry-run`; pause before
real sync if it would refresh out-of-scope components.
