//go:build sylixos && !nonamespacing

package maincmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/gokrazy/rsync/internal/log"
	"github.com/gokrazy/rsync/rsyncd"
)

func namespace(modules []rsyncd.Module, listen string) error {
	if os.Getenv("GOKRAZY_RSYNC_PRIVDROP") != "" {
		log.Printf("pid %d (privileges dropped)", os.Getpid())

		// Expected by the go-systemd package, and hard to set before creating
		// the process in Go.
		os.Setenv("LISTEN_PID", strconv.Itoa(os.Getpid()))

		return nil
	}

	if os.Getuid() != 0 {
		version()
		log.Printf("environment: unprivileged")
		return nil
	}

	version()
	log.Printf("environment: privileged")
	log.Printf("running as root (uid 0), sylixos not drop privileges to nobody (uid/gid 65534), still using (uid 0)")

	return nil
}

var errIsParent = errors.New("re-exec parent process sentinel error")
