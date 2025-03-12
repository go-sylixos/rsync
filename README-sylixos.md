# gokrazy rsync for sylixos

## Privilege dropping on sylixos

Privilege will not be dorpped on sylixos! Be careful when using it.

## Which is not working on sylixos

- rsync client with SSH is not working due to sylixos /usr/bin/ssh behavior.

## How do I set up a config for SSH deamon

```bash
touch ~/.config/gokr-rsyncd.toml
```

```toml
[[listener]]
  [listener.authorized_ssh]
    address = ":22873"
    authorized_keys = "/root/.ssh/authorized_keys"

[[module]]
name = "pwd"
path = "/apps/rsync-test"
comment = "Backup storage"
writable = true

dont_namespace = true
```

### Be aware of:

- Donot set any thing like rsyncd = ":873" under [[listener]]
- SSH authorized_keys only support ed25519.
- When gokr-rsyncd.toml exist cmdline will not work anymore.

## How to build this for sylixos

```bash
cd $GOPATH/src
git clone https://github.com/go-sylixos/rsync
cd rsync

git checkout main-add-sylixos-support
GOOS=sylixos GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-w -s"  ./cmd/gokr-rsync
```
