package maincmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gokrazy/rsync/internal/log"
)

func canUnexpectedlyWriteTo(dir string) error {
	fn := filepath.Join(dir, "gokr-rsyncd.unexpectedly_writable")
	if err := os.WriteFile(fn, []byte("gokr-rsyncd creates this file to prevent misconfigurations. if you see this file, it means gokr-rsyncd unexpectedly was started with too many privileges"), 0644); err == nil {
		os.Remove(fn)
		if runtime.GOOS == "sylixos" {
			log.Printf("on sylixos gokr-rsyncd will always started with too many privileges.")
			return nil
		}
		return fmt.Errorf("unexpectedly able to write file to %s, exiting", dir)
	}
	return nil
}
