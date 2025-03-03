//go:build !docker
// +build !docker

package command

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	"github.com/brycedjohnson/shellhub-agent/pkg/osauth"
)

func NewCmd(u *osauth.User, shell, term, host string, command ...string) *exec.Cmd {
	user, _ := user.Lookup(u.Username)
	userGroups, _ := user.GroupIds()

	// Supplementary groups for the user
	groups := make([]uint32, 0)
	for _, sgid := range userGroups {
		igid, _ := strconv.Atoi(sgid)
		groups = append(groups, uint32(igid))
	}
	if len(groups) == 0 {
		groups = append(groups, u.GID)
	}

	cmd := exec.Command(command[0], command[1:]...) //nolint:gosec
	cmd.Env = []string{
		"TERM=" + term,
		"HOME=" + u.HomeDir,
		"SHELL=" + shell,
		"SHELLHUB_HOST=" + host,
	}
	cmd.Dir = u.HomeDir

	if os.Geteuid() == 0 {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: u.UID, Gid: u.GID, Groups: groups}
	}

	return cmd
}
