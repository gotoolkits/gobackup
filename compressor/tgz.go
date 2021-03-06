package compressor

import (
	"github.com/huacnlee/gobackup/config"
	"github.com/huacnlee/gobackup/helper"
)

// Tgz .tar.gz compressor
//
// type: tgz
type Tgz struct {
}

func (ctx *Tgz) perform(model config.ModelConfig) (archivePath string, err error) {
	filePath := archiveFilePath(".tar.gz")

	opts := ctx.options()
	opts = append(opts, filePath)
	opts = append(opts, model.Name)

	_, err = helper.Exec("tar", opts...)
	if err == nil {
		archivePath = filePath
		return
	}
	return
}

func (ctx *Tgz) options() (opts []string) {
	if helper.IsGnuTar {
		opts = append(opts, "--ignore-failed-read")
	}
	opts = append(opts, "-zcf")

	return
}
