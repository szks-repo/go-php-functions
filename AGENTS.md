# AGENTS

## Overview
- Repository recreates selected PHP utility functions in Go for developers migrating from PHP ecosystems.
- Current focus is the `phpstring` package that mirrors string helpers like `substr`, `mb_substr`, `nl2br`, `wordwrap`, and `number_format`.
- Project targets Go 1.25 and vendors behaviour via `golang.org/x/exp` and `golang.org/x/text` for numeric formatting.

## Repository Layout
- `string/`: Core Go implementations of PHP-style string helpers plus unit tests.
- `array/`: Placeholder directory; no Go sources yet.
- Root files: `README.md`, `go.mod`, `go.sum`, `LICENSE`, newly added `AGENTS.md`.

## Key Packages & Functions
### `phpstring`
- `NumberFormat[N constraints.Integer | constraints.Float](n N) string`: Wraps `golang.org/x/text/number` with a Japanese locale printer for comma-separated numeric formatting. Tests cover signed/unsigned integers and floats.
- `Wordwrap(s string, opt ...WordwrapOpt) string`: Splits input into lines and wraps them according to optional width/delimiter/cut flags. Internals handle whitespace tokenisation; note that the current helper routing (`wrapLineNoCut` vs `wrapLineCut`) is inverted relative to the option name.
- `WordwrapOpt`: Options struct providing `Width`, `Delim`, and `CutLongWords` flags.
- `MbSubstr(s string, start int, length ...int) string`: Rune-aware substring similar to PHP `mb_substr`, supporting negative indices and lengths.
- `Substr(s string, offset int, length ...int) string`: UTF-8-safe analogue to PHP `substr`, covering negative offsets and lengths.
- `Nl2br[S ~string](s S) string`: Replaces `\n` with `<br />`, generic over types with string underlying representation.
- `Date(layout string) string`: Stub; not yet implemented.

## Tests & Quality Notes
- `go test ./...` exercises the string helpers. Coverage is uneven: `number_format`, `nl2br`, `mb_substr`, and `substr` assertions are meaningful.
- `string/wordwrap_test.go` defines scenarios but does not assert on the returned value, so regressions may slip through.
- `string/date_test.go` is empty, mirroring the unimplemented function.
- Mixed line endings (`\r\n`) appear in several files (e.g., `mb_substr.go`, `substr.go`); standardising could aid diffs.

## External Dependencies
- `golang.org/x/text` (0.28.0) supplies locale-aware number formatting utilities.
- `golang.org/x/exp` (2025-08-08 snapshot) delivers the `constraints` package for generics.

## Recommended Next Steps
1. Implement `phpstring.Date` to match PHP `date` semantics or document intended scope.
2. Fix `string/wordwrap.go` option routing so `CutLongWords` toggles the expected behaviour, and add assertions in tests.
3. Populate `array/` with planned functions or remove the empty directory to clarify roadmap.
4. Standardise line endings and address README typos (`number_fomrat`).

