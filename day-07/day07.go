package day07

import (
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2022/util"
)

func Part1(input string) *util.Solution {
	const dirSizeThreshold = 100000

	lines := strings.Split(input, "\n")

	dirMap := parseInput(lines)

	var totalSize int = 0

	for _, dir := range dirMap {
		size := dir.EffectiveSize()

		if size > dirSizeThreshold {
			continue
		}

		totalSize += size
	}

	return util.NewSolution(1, totalSize)
}

func Part2(input string) *util.Solution {
	const requiredSpace = 30000000
	const diskSpace = 70000000

	lines := strings.Split(input, "\n")

	dirMap := parseInput(lines)
	sizeList := make(DirectorySizeList, 0)

	totalSize := dirMap["/"].EffectiveSize()
	spaceToReclaim := requiredSpace - (diskSpace - totalSize)

	for path, dir := range dirMap {
		size := dir.EffectiveSize()

		if size < spaceToReclaim {
			continue
		}

		sizeList = append(sizeList, DirectorySize{path, size})
	}

	sort.Sort(sizeList)

	return util.NewSolution(2, sizeList[0].size)
}

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

// Implement sort interface for Directory array
type DirectorySize struct {
	path string
	size int
}

type DirectorySizeList []DirectorySize

func (list DirectorySizeList) Len() int {
	return len(list)
}

func (list DirectorySizeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list DirectorySizeList) Less(i, j int) bool {
	return list[i].size < list[j].size
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
