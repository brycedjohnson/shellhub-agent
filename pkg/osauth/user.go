package osauth

import (
	"os"
	"os/user"
	"strconv"
)

type User struct {
	UID      uint32
	GID      uint32
	Username string
	Password string
	Name     string
	HomeDir  string
	Shell    string
}

func singleUser() *User {
	var uid, gid int
	var username, name, homeDir, shell string
	u, err := user.Current()
	uid, _ = strconv.Atoi(os.Getenv("UID"))
	homeDir = os.Getenv("HOME")
	shell = os.Getenv("SHELL")
	if err == nil {
		uid, _ = strconv.Atoi(u.Uid)
		gid, _ = strconv.Atoi(u.Gid)
		username = u.Username
		name = u.Name
		homeDir = u.HomeDir
	}

	return &User{
		UID:      uint32(uid),
		GID:      uint32(gid),
		Username: username,
		Name:     name,
		HomeDir:  homeDir,
		Shell:    shell,
	}
}

func LookupUser(username string) *User {
	return singleUser()
}
