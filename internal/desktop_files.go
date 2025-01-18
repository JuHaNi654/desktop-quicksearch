package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	left_bracket  = []byte{91}
	right_bracket = []byte{93}
	separator     = "="
)

type Group struct {
	Name   string
	Values map[string]string
}

func newGroup() *Group {
	return &Group{
		Values: make(map[string]string),
	}
}

func loadDesktopFile(r io.Reader) []*Group {
	var groups []*Group
	s := bufio.NewScanner(r)
	s.Split(splitFileChunks)
	for s.Scan() {
		g := newGroup()
		for _, r := range bytes.Split(s.Bytes(), []byte("\n")) {
			if bytes.HasPrefix(r, left_bracket) && bytes.HasSuffix(r, right_bracket) {
				g.Name = string(r)
			} else {
				key, value := splitKeyValue(string(r), separator)
				g.Values[key] = value
			}
		}
		groups = append(groups, g)
	}

	return groups
}

func splitKeyValue(s, delimiter string) (string, string) {
	index := strings.Index(s, delimiter)
	if index == -1 {
		return s, ""
	}
	return s[:index], s[index+len(delimiter):]
}

func splitFileChunks(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), bytes.TrimSuffix(data, []byte("\n")), nil
	}

	return 0, nil, nil
}

type DesktopEntry struct {
	Desktop string `json:"desktop"`
	Name    string `json:"name"`
	Exec    string `json:"exec"`
	Type    string `json:"type"`
	Icon    string `json:"icon"`
}

var path = []string{
	"/usr/share/applications/",
	"/usr/local/share/applications/",
	"/var/lib/snapd/desktop/applications/",
}

func GetDesktopEntries() []*DesktopEntry {
	var entries []*DesktopEntry

	for _, p := range path {
		files, err := os.ReadDir(p)
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range files {
			filePath := fmt.Sprintf("%s%s", p, file.Name())
			f, err := os.Open(filePath)
			if err != nil {
				fmt.Println(err)
			}

			settings := loadDesktopFile(f)
			for _, s := range settings {
				name, nameOk := s.Values["Name"]
				exec, execOk := s.Values["Exec"]
        noDisplay, noDisplayOk := s.Values["NoDisplay"]
        if noDisplayOk && noDisplay == "true" {
          continue
        }

				if s.Name == "[Desktop Entry]" && nameOk && execOk {
					entries = append(entries, &DesktopEntry{
						Desktop: file.Name(),
						Name:    name,
						Exec:    exec,
						Type:    s.Values["Type"],
						Icon:    s.Values["Icon"],
					})
				}
			}
		}
	}

	return entries
}
