package main

import (
	"os/exec"
	"strings"
)

const pkgfile = "./package.json"

type packageJSON struct {
	Version string `json:"version"`
}

func parseNodeVersion() (string, error) {
	nodePath, err := exec.LookPath("node")
	out, err := exec.Command(nodePath, "--version").Output()
	if err != nil {
		return "", err
	}
	return strings.Replace(string(out), "\n", "", -1), nil
}

func segmentNode(p *powerline) {
	// stat, err := os.Stat(pkgfile)
	nodeVersion, err := parseNodeVersion()
	if err == nil && nodeVersion != "" {
		p.appendSegment("node-version", segment{
			content:    nodeVersion + " \u2B22",
			foreground: p.theme.NodeFg,
			background: p.theme.NodeBg,
		})
	}
}
