---
name: repoguide
description: Generate and interpret `repoguide` repository maps for fast codebase exploration. Use when you need a high-level code map, want to identify important files/symbols/dependencies before deeper reading, need focused symbol or file lookups, or want to cache repo maps for repeated agent runs.
---

# repoguide

Use `repoguide` to get a compact structural map of a repository before doing detailed file reads.

## Run It

Use the repo root as the default target.

```bash
repoguide .
```

Useful variants:

```bash
repoguide . --cache .cache/repoguide.toon
repoguide . --raw
repoguide . --format v1
repoguide . -n 20
repoguide . -l go,typescript
repoguide . --with-tests
repoguide . --symbol BuildGraph
repoguide . --symbol-regex '^(?i)build.*'
repoguide . --file internal/lang
repoguide . --file-regex '(^|/)internal/lang'
repoguide . --symbol Handle --file server
```

## Recommended Workflow

1. Run `repoguide . --cache .cache/repoguide.toon` first for a full cached map.
2. Scan `files` to find the most central files.
3. Scan `defs.*` to find definitions by kind.
4. Scan `deps` to see cross-file relationships.
5. Use `--symbol` / `--file` (substring) or `--symbol-regex` / `--file-regex` (regex) for focused follow-up queries.
6. Read only the small number of files suggested by the map.

## How To Read v2 Output

Default output is `v2`.

- `files[id,rank,path]`: ranked files; lower repetition than v1
- `defs.c`, `defs.f`, `defs.m`, `defs.fld`, `defs.const`: definitions split by kind
- `deps[edge,symbols]`: file-to-file dependencies such as `f3->f1`
- `calls[edge]`: call graph edges like `caller->callee`
- focused queries may also include `callsites` and `sig`

Interpretation tips:

- Start from the highest-rank files first.
- Use `defs.*` as the lookup table for where symbols live.
- Use `deps` to understand subsystem boundaries and imports.
- Use `callsites` from focused queries for exact navigation when available.

## When To Use Focused Queries

Use `--symbol` when the user names a function, method, class, or concept that likely maps to a symbol.

```bash
repoguide . --symbol BuildGraph
```

Use `--symbol-regex` when the user describes a naming pattern.

```bash
repoguide . --symbol-regex '^(?i)build.*graph$'
```

Use `--file` when the user points to a package, directory, or feature area.

```bash
repoguide . --file internal/parse
```

Use `--file-regex` when the user describes a path pattern.

```bash
repoguide . --file-regex '(^|/)parse/'
```

Combine symbol and file filters when the user asks about a symbol inside a subsystem.

```bash
repoguide . --symbol Prompt --file session
```

Focused filters match extracted symbol names and discovered file paths (semantic map data), not arbitrary file content lines. Use `rg` for grep-style full-text searches.

## Choosing Format

- Prefer default `v2` for token efficiency.
- Use `--format v1` only for compatibility with older tooling or examples.

## Caching

Prefer caching for repeated runs in the same repo.

```bash
repoguide . --cache .cache/repoguide.toon
```

- cache is reused when source files have not changed
- focused reads still benefit because the full map is kept warm
- add `.cache/` to `.gitignore`

## Language And Scope Controls

- `-l go,typescript`: restrict to languages relevant to the task
- `-n 20`: trim very large repos to the top-ranked files
- `--with-tests`: include tests only when they matter for the task
- `--raw`: remove the explanatory header when only the TOON payload is needed

## Good Patterns

- Before broad refactors, run a full cached map.
- Before answering architecture questions, inspect `files`, `defs.*`, and `deps`.
- Before reading many files, use `--symbol` / `--file` or regex variants to narrow the search.
- For repeated agent sessions, keep `.cache/repoguide.toon` warm.

## Avoid

- Do not read dozens of files before checking the map.
- Do not default to `--with-tests` unless test structure matters.
- Do not use repoguide focused filters as a line-by-line grep replacement; use `rg` for full-text matching.
- Do not prefer `v1` unless compatibility is required.
