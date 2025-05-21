# Changelog

## [Unreleased] - 2025-05-21

### Added
- Introduced a `Metadata` field of type `json.RawMessage` to both `meta` and `ItemOptions` structs in [`meta.go`](meta.go) and [`options.go`](options.go), allowing storage of arbitrary custom metadata for cache items.
- Added new tests in [`meta_test.go`](meta_test.go) to verify correct handling and round-trip preservation of the `Metadata` field, including cases with non-empty and nil metadata.

### Changed
- Updated `newMeta` and `metaToOptions` functions in [`meta.go`](meta.go) to support copying the new `Metadata` field between `meta` and `ItemOptions`.
- Updated imports in [`options.go`](options.go) to include `encoding/json` for the new metadata functionality.