# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.0]

### Added

- Add parsing relative dates for `--start` and `--end` parameters

### Changed

- Fix parsing times for `--time` parameter
- Default time range for `report list` and `report summary` is the current day

## [0.3.0]

### Added

- Command `punchcard project list`
- Command `punchcard project rename`

### Changed

- Fixed output of command `punchcard report summary`

## [0.1.0] - 2020-03-19

Initial release

### Added

- Command `punchcard examples`
- Command `punchcard project add`
- Command `punchcard project close`
- Command `punchcard project reopen`
- Command `punchcard report list`
- Command `punchcard report summary`
- Command `punchcard work start`
- Command `punchcard work status`
- Command `punchcard work stop`
- Command `punchcard work resume`

[unreleased]: https://github.com/djlauk/punchcard/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/djlauk/punchcard/releases/v0.4.0
[0.3.0]: https://github.com/djlauk/punchcard/releases/v0.3.0
[0.1.0]: https://github.com/djlauk/punchcard/releases/v0.1.0
