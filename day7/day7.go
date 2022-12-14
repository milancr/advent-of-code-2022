package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

const maxSize int = 100000
const disk int = 70000000
const requiredSpace int = 30000000

// Entry struct classifies each line
type Entry struct {
	name  string
	isDir bool
	size  int
}

func (e *Entry) String() string {
	return fmt.Sprintf("{name: %q, isDir: %t, size: %d}", e.name, e.isDir, e.size)
}

var entries map[string]*Entry

func parseCMD(cmd, pwd string) string {
	if cmd == "$ ls" {
		return pwd // ignore
	}
	if cmd == "$ cd .." {
		return path.Dir(pwd)
	}
	fields := strings.Fields(cmd)
	return path.Join(pwd, fields[len(fields)-1])
}

func parseLine(cmd, pwd string) string {
	if len(cmd) == 0 {
		return pwd
	}

	if cmd[0] == '$' {
		return parseCMD(cmd, pwd)
	}

	if strings.HasPrefix(cmd, "dir") {
		name := path.Join(pwd, cmd[4:])
		entries[name] = &Entry{name, true, 0}
	} else {
		fields := strings.Fields(cmd)
		size, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		name := path.Join(pwd, fields[1])

		entries[name] = &Entry{name, false, size}
	}
	return pwd
}

func calcSize(dir string) (size int) {
	for name, entry := range entries {
		if name == dir {
			continue
		}

		if path.Dir(name) != dir {
			continue
		}

		if entry.size == 0 {
			entry.size = calcSize(name)
		}
		size += entry.size
	}
	return size
}

func main() {
	entries = make(map[string]*Entry)

	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	root := "/"
	pwd := root
	entries[pwd] = &Entry{pwd, true, -1}

	// part 1
	for _, line := range strings.Split(string(file), "\n") {
		pwd = parseLine(line, pwd)
	}
	entries[root].size = calcSize(root)
	total := 0

	for _, entry := range entries {
		if entry.isDir && entry.size <= maxSize {
			total += entry.size
		}
	}

	fmt.Println(total)

	// part 2
	curBest := &Entry{"", false, disk}
	spaceUsed := entries[root].size
	for _, entry := range entries {
		if entry.isDir && disk-spaceUsed+entry.size > requiredSpace && entry.size < curBest.size {
			curBest = entry
		}
	}

	fmt.Println(curBest.size)
}
