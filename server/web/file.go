package web

import (
	"fmt"
	"github.com/duruyao/gochat/server/util"
	"strings"
)

func HtmlDir() string {
	return fmt.Sprintf("%s/templates", util.GoChatDir())
}

func HtmlPath(name string) string {
	if strings.HasSuffix(name, ".html") || strings.HasSuffix(name, ".HTML") {
		return fmt.Sprintf("%s/%s", HtmlDir(), name)
	}
	return fmt.Sprintf("%s/%s.html", HtmlDir(), name)
}

func ResourceDir() string {
	return fmt.Sprintf("%s/public", util.GoChatDir())
}

func TLSCertPath() string {
	return fmt.Sprintf("%s/app/.tls-cert.pem", util.GoChatDir())
}

func TLSKeyPath() string {
	return fmt.Sprintf("%s/app/.tls-key.pem", util.GoChatDir())
}
