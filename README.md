# punchcard

A simple CLI for tracking time.

## Getting started

Create a configuration file in `$HOME/.punchcard.yaml`:

```yaml
# ~/.punchcard.yaml
# configuration for punchcard
storage:
	data: ~/punchcard/data.yaml

# round timestamps to this duration
# for details on the format see https://golang.org/pkg/time/#ParseDuration
roundTimestampsTo: 15m
```

Run the commands for instructions and/or examples:

```
punchcard help

punchcard examples
```

## Developer documentation

Some short information on how to keep this going. (Because I tend to forget these things.)

### Contributing

1.  Create an issue.
1.  Fork the repository.
1.  Run tests (all should be OK): `go test ./...`
1.  Write tests.
1.  Run tests (some should fail).
1.  Implement fix / feature.
1.  Run tests (all should be OK again).
1.  Commit, push and send a pull request, referencing the issue from step 1.

### Releasing a new version

I use [standard-version](https://github.com/conventional-changelog/standard-version#readme) to generate the [CHANGELOG](./CHANGELOG.md).

1.  Run `standard-version --dry-run`. Note the next version number.
1.  Increase the version number in `cmd/root.go` and `git add cmd/root.go`.
1.  Run `standard-version --commit-all`.
1.  Run `git push --follow-tags origin master`.

## LICENSE

GPLv3 (or later). See file [LICENSE](./LICENSE) for the full license text.
