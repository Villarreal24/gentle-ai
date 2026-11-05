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
- Resolve only straightforward compatibility issues. Stop for user input if
  integration requires redesigning the Factory Droid adapter.
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
    feature head.
  - Commit: task-record work-unit commit pending.

- [ ] **T2. Integrate the official v4.3.1 tag.**
  - Route: delegated implementation by one writer; the six-commit adapter
    footprint spans 76 files. Do not parallelize writers.
  - Refresh `origin` and `upstream`, verify the tag and its official source,
    then merge it into the temporary branch based on the original feature
    branch. Preserve all six fork commits.
  - Apply test-first regression fixes only where a concrete behavior
    incompatibility is observed. Stop and ask if the adapter needs redesign.
  - Commit the release merge and any focused compatibility fix as reviewable
    work units. Record their IDs and authored line counts below.

- [ ] **T3. Verify the merge and build before installation.**
  - Route: writer self-verification plus an independent read-only verifier;
    risk is high because the change affects the CLI, persisted model
    preferences, and local sync behavior.
  - Run the focused Go checks listed under Verification, build a temporary
    binary, and verify its version and command help. Review the exact diff
    against the six fork commits and assess local config effects.
  - Do not install or sync if required checks fail or the independent review
    finds an unresolved high-impact issue.

- [ ] **T4. Install and synchronize within the authorized scope.**
  - Route: bounded local operation after validation; risk is high because sync
    can modify global configuration.
  - Replace the active binary only after its temporary build passes.
  - Run `sync --dry-run` first. The current sync implementation's dry-run
    reports components and step counts, not a per-file diff. A normal sync can
    refresh Persona, Engram, Context7, GGA, Skills, shared skills, and
    persisted CodeGraph wiring even when `--agents droid` is specified.
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
- Focused tests pass; the temporary binary reports the expected fork version
  and exposes the expected commands.
- The installed binary matches the verified temporary build.
- Any sync stays within the approved local assets and preserves user files and
  preferences; it does not update the Engram binary/setup/database or unrelated
  tools.
- The original feature branch is updated and pushed normally only to
  `Villarreal24/gentle-ai`.
- No changes are made to `cars_platform`.

## Verification

Candidate focused checks from the current tree; adjust only if v4.3.1 changes
package boundaries:

```sh
go test ./internal/agents/droid ./internal/model ./internal/components/reviewassets
go test ./internal/cli -run 'TestDroid|TestRunSyncDryRunDoesNotWriteFiles|TestNativeReviewSyncPreservesUnknownAndSnapshotsLedger'
go test ./internal/tui/... -run 'Droid|ModelConfig|Navigation'
go test ./internal/assets ./internal/components/agentguidance ./internal/components/persona ./internal/components/skills ./internal/state ./internal/system ./internal/catalog ./internal/agents
go test ./internal/cli -run 'Doctor|ToolBinaries'
```

Build the temporary binary with the repository's Go version and release
linker settings, then check `--version` and `sync --help`. Before real sync,
review its dry-run output and confirm the resulting config effects are within
the authorization above.

## Delivery Strategy

Use `ask-on-risk` (default). The imported upstream history and the six
pre-existing fork commits are not new authored work for this update. Count new
authored changes before each work-unit commit; if they exceed about 400 lines,
ask once about the chain strategy before committing.

## Progress

- Route: delegated read-only mapping to `gentle-ai-explore` (completed). It
  mapped the six commits to a 76-file footprint and found that `--agents droid`
  is not an asset-only sync fence.
- T1 recovery checkpoint: binary, Git tag, managed-asset backups, private MCP
  backup, and temporary branch verified; the task-record commit is next.
- No remote fetch/push, source changes, tests, builds, binary replacement, or
  sync have been performed yet.

## Next Step

Commit this task record on `update/gentle-ai-v4.3.1`, then fetch and verify the
official release before changing source.
