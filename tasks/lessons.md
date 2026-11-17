# Lessons

- **No comments in code.** Do not add comments, not even package doc comments or
  `ponytail:` markers. Name things so the code reads without them. (2026-08-06)
- **Settle the concept before writing code.** For spec-shaped requests ("extend
  the scheme so X"), present the grammar, its ambiguities and the trade-offs
  first; write the file only after the rules are agreed. (2026-08-06)
- **Never lowercase or otherwise normalise user data.** Env var names and values
  are passed through exactly as written. (2026-08-06)
- **Never commit unless explicitly asked.** "Copy to main" means put the files on
  the main branch's working tree, not commit. Leave changes uncommitted and let
  the user review and commit. (2026-11-16)
- **Build the mechanism the user names.** When the user asks for a specific
  mechanism ("a min version"), implement that instead of substituting a simpler
  alternative (an explicit branch list) and arguing for it. (2026-11-17)
