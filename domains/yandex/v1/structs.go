// Yandex Disk File Pusher
// Copyright (c) 2019 Vladimir "fat0troll" Hodakov

package yandexv1

import (
	"fmt"
	"github.com/schollz/progressbar/v2"
	"io"
)

type progressReader struct {
	r           io.Reader
	atEOF       bool
	progressbar *progressbar.ProgressBar
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.r.Read(p)
	if err == io.EOF {
		pr.atEOF = true
	}
	pr.report(int64(n))
	return n, err
}

func (pr *progressReader) report(n int64) {
	_ = pr.progressbar.Add64(n)
	if pr.atEOF {
		fmt.Print("\n\n")
	}
}

type TokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type UploadError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type UploadInfo struct {
	URL            string `json:"href"`
	Method         string `json:"method"`
	URLIsTemplated bool   `json:"templated"`
}
