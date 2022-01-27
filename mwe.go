package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
)

var (
	errorImplementMe = errors.New("implement me")
)

func init() {
	fsi := &fs.RegInfo{
		Name:        "dummy",
		Description: "dummy rclone backend",
		NewFs:       NewFs,
	}
	fs.Register(fsi)
}

// NewFs creates a new Fs object from the name and root. It connects to
// the host specified in the config file.
func NewFs(ctx context.Context, name, root string, m configmap.Mapper) (fs.Fs, error) {
	return nil, errorImplementMe
}
