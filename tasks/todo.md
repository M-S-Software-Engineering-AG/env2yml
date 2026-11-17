# Release branches

- [x] auto-release: tag only on `release/**`, main never releases
- [x] release: notes start at previous tag of the same line, "Latest" only for the highest version
- [x] renovate: unchanged, main + all `release/*` (EOL branches get deleted, tags stay)
- [x] renovate: release branches stay security-only (existing packageRule)
- [x] daily-scan: matrix over main + all `release/*` branches
- [ ] user: commit + merge to main
- [ ] rulesets: `releases/**/*` -> `release/**/*` (protect-release-branches, ci-required)
- [ ] create `release/1.2` from main

## Review
Verified locally: branch listing against the real repo, previous-tag
and highest-tag resolution against existing tags, JSON/YAML parse.
