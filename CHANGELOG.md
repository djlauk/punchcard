# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [0.4.6](https://github.com/djlauk/punchcard/compare/v0.4.5...v0.4.6) (2020-05-22)


### Features

* add command report daily-totals ([6586a69](https://github.com/djlauk/punchcard/commit/6586a69131961e141849863098a0d41818c5e6b7)), closes [#11](https://github.com/djlauk/punchcard/issues/11)

### [0.4.5](https://github.com/djlauk/punchcard/compare/v0.4.4...v0.4.5) (2020-04-21)


### Features

* add option to list all work entries ([35e7c8a](https://github.com/djlauk/punchcard/commit/35e7c8af5ab125be498fdf032271038705b40a86)), closes [#10](https://github.com/djlauk/punchcard/issues/10)

### [0.4.4](https://github.com/djlauk/punchcard/compare/v0.4.3...v0.4.4) (2020-04-16)

### [0.4.3](https://github.com/djlauk/punchcard/compare/v0.4.2...v0.4.3) (2020-04-16)


### Features

* support resuming arbitrary work log entries ([426c9ab](https://github.com/djlauk/punchcard/commit/426c9ab11aca308074a81fae14499c4820f6d32c)), closes [#9](https://github.com/djlauk/punchcard/issues/9)

### [0.4.2](https://github.com/djlauk/punchcard/compare/v0.4.1...v0.4.2) (2020-04-08)


### Features

* add aliases for primary commands ([e61297a](https://github.com/djlauk/punchcard/commit/e61297af01677ac2ba246277c032b88bee4d8519)), closes [#8](https://github.com/djlauk/punchcard/issues/8)
* add option --finish-current ([2d7da94](https://github.com/djlauk/punchcard/commit/2d7da94ed329b2f0765536c40863f37738792f90)), closes [#7](https://github.com/djlauk/punchcard/issues/7)

### [0.4.1](https://github.com/djlauk/punchcard/compare/v0.4.0...v0.4.1) (2020-04-08)


### Features

* support relative times ([a814f86](https://github.com/djlauk/punchcard/commit/a814f8623cf83fe829e2d62fa897ead64fe2993f)), closes [#5](https://github.com/djlauk/punchcard/issues/5)

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
