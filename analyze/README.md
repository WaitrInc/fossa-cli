Should something be a _supported tool_ (package) or a _resolution strategy_ (subpackage)?

Criteria:
- Is the grouping intuitive or idiomatic?
- Is a tool cross-language? (E.g. grouping Maven and Ivy under Java might not make sense if they support all JVM languages.)
- Favour grouping by _tool_ over grouping by _language_ or _registry_.
- Are the underlying dependency resolution strategies very similar?
- Do they share package manifests or underlying implementations?
- Is one tool a wrapper over another?
- Do they share package registries?
- Do they share package manifest formats?
- Does it make sense to fall back to one tool when the other does not succeed?
- Are the tools used in a mutually exclusive way?
- Only include tools where _code_ or _manifests_ can be scanned (e.g. `package.json` or `Dockerfile`).

Examples:
- Would I say "this is a Go project?" or "this is a GDM project?" (Go)
- Would I say "this is a Java project?" or "this is a Maven project?" (Maven)
- Would I say "this is a Ruby project?" or "this is a Bundler project?" (Ruby)
- Both `npm` and `yarn` use `node_modules` and `package.json`.
- `stack` is a wrapper over `cabal`.
- `go` has a language-specified import mechanism, but all package managers use different manifests to specify revisions.