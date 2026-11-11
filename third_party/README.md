# third_party

Patched copies of two dependencies, wired in through `replace` directives in the
root `go.mod`, so that `buildkite-agent` builds and runs on AIX (`aix/ppc64`).
Upstream has no working AIX support in either.

| Directory    | Upstream base                            | AIX changes                                                                                                            |
|--------------|------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| `pty/`       | `github.com/creack/pty` v1.1.19          | Adds `pty_aix.go`, `ioctl_aix.go`, `asm_aix_ppc64.s` (from github.ibm.com/sfertile/pty); drops the AIX stub `ioctl_unsupported.go`; excludes `aix` in `pty_unsupported.go`. |
| `machineid/` | `github.com/denisbrodbeck/machineid` v1.0.1 | Adds `id_aix.go` (from github.ibm.com/daltenty/machineid, branch `daltenty/aix-support`), reading `os_uuid` via `lsattr`; adds a `go.mod`. |

Other changes on this branch: `internal/process/signal_notwindows.go` uses
`unix.Getpgid` because `syscall.Getpgid` does not exist on AIX.

When bumping `creack/pty` or `machineid` in the root `go.mod`, re-apply these
changes to the new upstream version.

## Building for AIX

```
GOOS=aix GOARCH=ppc64 CGO_ENABLED=0 go build \
  -ldflags "-X github.com/buildkite/agent/v4/version.buildNumber=<N>" \
  -o buildkite-agent .
```
