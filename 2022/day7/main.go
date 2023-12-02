package main

import (
	"log"
	"strconv"
	"errors"
	"fmt"
	"strings"
	"regexp"
	"os"
	"bufio"
)

type directory struct {
	name string
	parent *directory
	children []*directory
	files []*dir_file
}

type dir_file struct {
	name string
	size int
}

func cd(dir *directory, path string) (result *directory, err error) {
	if path == "/" {
		if dir.name == "/" {
			result = dir
			return
		} else {
			result, _ = cd(dir.parent, path)
			return
		}
	}

	if path == ".." {
		if dir.parent == nil {
			result = dir
			return
		} else {
			result = dir.parent
			return
		}
	}

	// path is a directory name
	var	foundChild *directory
	for _, child := range dir.children {
		if child.name == path {
			foundChild = child
			break
		}
	}

	if foundChild != nil {
		result = foundChild
	} else {
		err = errors.New("No such file or directory")
	}

	return
}

func ls(dir *directory) []string {
	var result []string

	for _, child := range dir.children {
		result = append(result, "dir " +child.name)
	}

	for _, file := range dir.files {
		result = append(result, fmt.Sprint(file.size) + " "  + file.name)
	}

	return result
}

func printDirectory(dir *directory, indent int) {
	fmt.Println(strings.Repeat(" ", indent), dir.name + "/")

	for _, file := range dir.files {
		printFile(file, indent+2)
	}

	for _, child := range dir.children {
		printDirectory(child, indent+2)
	}
}

func printFile(f *dir_file, indent int) {
	fmt.Println(strings.Repeat(" ", indent), f.size, f.name)
}

func calculateDirectorySize(dir *directory) (result int) {
	result = 0
	for _, file := range dir.files {
		result += file.size
	}

	for _, child := range dir.children {
		result += calculateDirectorySize(child)
	}

	return
}

func findDirectoriesBySize(dir *directory, min int, max int) (result []*directory) {
	if dir.children == nil {
		result = nil
	}
	for _, child := range dir.children {
		dirSize := calculateDirectorySize(child)
		if dirSize <= max && dirSize >= min {
			result = append(result, child)
		}

		result = append(result, findDirectoriesBySize(child, min, max)...)
	}

	return
}

func findDirectoriesBelowSize(dir *directory, size int) (result []*directory) {
	if dir.children == nil {
		result = nil
	}
	for _, child := range dir.children {
		if calculateDirectorySize(child) < size {
			result = append(result, child)
		}

		result = append(result, findDirectoriesBelowSize(child, size)...)
	}

	return
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var root *directory = &directory{name: "/", parent: nil, children: nil, files: nil}
	re := regexp.MustCompile("\\$\\s(?P<command>.+)\\s(?P<param>.+)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
//		log.Println("CurrentRoot: ", root.name)
		line := scanner.Text()

		if strings.HasPrefix(line, "$") {
			matches := re.FindStringSubmatch(line)
			commandIdx := re.SubexpIndex("command")
			paramIdx := re.SubexpIndex("param")

			if matches == nil {
				continue;
			}

			command := matches[commandIdx]
			if command == "cd" {
				param := matches[paramIdx]
				changed_root, _ := cd(root, param)
			//	log.Println("ChangedRoot: ", changed_root.name)
				root = changed_root
			}
		} else if strings.HasPrefix(line, "dir") {
			stringSlice := strings.Split(line, " ")
			child := directory{name: stringSlice[1], parent: root, children: nil}
			root.children = append(root.children, &child)
		} else {
			stringSlice := strings.Split(line, " ")
			file_size,_ := strconv.Atoi(stringSlice[0])
			new_file := dir_file{name: stringSlice[1], size: file_size}
			root.files = append(root.files, &new_file)
		}
	}


	root, _ = cd(root, "/")

	unused_space := 70000000 - calculateDirectorySize(root)
	needed_space := 30000000 - unused_space
	log.Println("Needed space", 30000000-unused_space)

	dirs_to_delete := findDirectoriesBySize(root, needed_space, 924933642 )

	result_size := 0
	for _, dir := range dirs_to_delete {
		log.Println("Deleting", dir.name, "size", calculateDirectorySize(dir))
		result_size += calculateDirectorySize(dir)
	}

	log.Println("Below SUM 100.000",result_size)
	log.Println("Root size", calculateDirectorySize(root))
}










