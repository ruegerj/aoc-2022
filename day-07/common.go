package day07

import (
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

type Directory struct {
	Name        string
	Size        int
	Directories []*Directory
}

func NewDirectory(name string) *Directory {
	dir := Directory{
		Name:        name,
		Size:        0,
		Directories: make([]*Directory, 0),
	}

	return &dir
}

func (dir *Directory) AppendFile(file int) {
	dir.Size += file
}

func (dir *Directory) AppendDirectory(directory *Directory) {
	dir.Directories = append(dir.Directories, directory)
}

func (dir *Directory) EffectiveSize() int {
	size := dir.Size

	for _, subDir := range dir.Directories {
		size += subDir.EffectiveSize()
	}

	return size
}

func parseInput(lines []string) map[string]*Directory {
	dirMap := make(map[string]*Directory, 0)
	dirMap["/"] = NewDirectory("/")

	path := append(make([]string, 0), "/")
	var currentDir *Directory

	for _, line := range lines {
		parts := strings.Split(line, " ")

		if parts[0] != "$" && parts[0] != "dir" {
			fileSize := util.MustParseInt(parts[0])
			currentDir.AppendFile(fileSize)
			continue
		}

		if parts[0] != "$" {
			newDir := NewDirectory(parts[1])
			path = append(path, newDir.Name)
			dirPath := strings.Join(path, "/")

			currentDir.AppendDirectory(newDir)
			dirMap[dirPath] = newDir
			continue
		}

		if parts[1] == "ls" {
			continue
		}

		if parts[2] == ".." {
			if len(path) > 1 {
				path = path[:len(path)-1]
			}
		}

		currentDir = dirMap[strings.Join(path, "/")]
	}

	return dirMap
}
