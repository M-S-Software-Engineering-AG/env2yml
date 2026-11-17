# Release branches

- [x] auto-release: tag only on `release/**`, main never releases
- [x] release: notes start at previous tag of the same line, "Latest" only for the highest version
- [x] renovate: unchanged, main + all `release/*` (EOL branches get deleted, tags stay)
- [x] renovate: release branches stay security-only (existing packageRule)
- [x] daily-scan: matrix over main + all `release/*` branches
- [x] user: commit + merge to main
- [x] rulesets: `release/**/*` added to protect-release-branches (org, by user) and ci-required (repo)
- [x] create `release/1.2` from main (5e9ed9a)

## Review
Verified locally: branch listing against the real repo, previous-tag
and highest-tag resolution against existing tags, JSON/YAML parse.

## Auto-release on every release-branch merge
- [x] tag on every push to `release/**`, not only Renovate merges (Renovate merges still batch)
- [x] version line from branch name: new `release/X.Y` gets `vX.Y.0`, existing line gets next patch
- [x] verified: gate cases with stub gh; tag logic for 1.2, 1.3, 1.1, 1.10 in scratch worktree
- [x] release only when `*.go` (not tests), go.mod or go.sum changed; verified per file in scratch worktree
- [ ] user: commit + push to main, then bring the change to release/1.2
