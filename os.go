// Code generated by counterfeiter. DO NOT EDIT.
// with command: counterfeiter -p os -fake-name fake_os
package goshim

import (
	"os"
	"time"
)

//go:generate counterfeiter -o os_fake/fake_os.go . Os
type Os interface {
	Expand(s string, mapping func(string) string) string
	ExpandEnv(s string) string
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
	Setenv(key, value string) error
	Unsetenv(key string) error
	Clearenv()
	Environ() []string
	NewSyscallError(syscall string, err error) error
	IsExist(err error) bool
	IsNotExist(err error) bool
	IsPermission(err error) bool
	Getpid() int
	Getppid() int
	FindProcess(pid int) (*os.Process, error)
	StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)
	Executable() (string, error)
	Mkdir(name string, perm os.FileMode) error
	Chdir(dir string) error
	Open(name string) (*os.File, error)
	Create(name string) (*os.File, error)
	Rename(oldpath, newpath string) error
	TempDir() string
	Chmod(name string, mode os.FileMode) error
	NewFile(fd uintptr, name string) *os.File
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	Truncate(name string, size int64) error
	Remove(name string) error
	Chtimes(name string, atime time.Time, mtime time.Time) error
	Pipe() (r *os.File, w *os.File, err error)
	Link(oldname, newname string) error
	Symlink(oldname, newname string) error
	Readlink(name string) (string, error)
	Chown(name string, uid, gid int) error
	Lchown(name string, uid, gid int) error
	Getwd() (dir string, err error)
	MkdirAll(path string, perm os.FileMode) error
	RemoveAll(path string) error
	IsPathSeparator(c uint8) bool
	Getuid() int
	Geteuid() int
	Getgid() int
	Getegid() int
	Getgroups() ([]int, error)
	Exit(code int)
	Stat(name string) (os.FileInfo, error)
	Lstat(name string) (os.FileInfo, error)
	Hostname() (name string, err error)
	Getpagesize() int
	SameFile(fi1, fi2 os.FileInfo) bool
}
