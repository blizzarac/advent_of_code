package main

import (
	"testing"
	"fmt"
)



func Test_cd_root_to_root(t *testing.T) {
	dir := directory{name: "/", parent: nil, children: nil}

	result := cd(&dir, "/")

	if result.name != dir.name {
		t.Errorf("Test_cd_root_to_root: Expected %s, got %s FAIL", dir.name, result.name)
	} else {
		fmt.Printf("Test_cd_root_to_root: Expected %s, got %s SUCCESS\n", dir.name, result.name)
	}
}

func Test_cd_root_to_parent(t *testing.T) {
	dir := directory{name: "/", parent: nil, children: nil}
	child := directory{name: "child", parent: &dir, children: nil}
	dir.children = append(dir.children, &child)

	result := cd(&child, "..")
	if result.name != dir.name {
		t.Errorf("Test_cd_root_to_parent: Expected %s, got %s FAIL", dir.name, result.name)
	} else {
		fmt.Printf("Test_cd_root_to_parent: Expected %s, got %s SUCCESS\n", dir.name, result.name)
	}
}

func Test_cd_root_to_child(t *testing.T) {
	dir := directory{name: "/", parent: nil, children: nil}
	child := directory{name: "child", parent: &dir, children: nil}
	dir.children = append(dir.children, &child)

	result := cd(&dir, "child")

	if result.name != child.name {
		t.Errorf("Test_cd_root_to_child: Expected %s, got %s FAIL", child.name, result.name)
	} else {
		fmt.Printf("Test_cd_root_to_child: Expected %s, got %s SUCCESS\n", child.name, result.name)
	}
}

func Test_ls(t *testing.T) {
	dir := directory{name: "/", parent: nil, children: nil}
	child := directory{name: "child", parent: &dir, children: nil}
	dir.children = append(dir.children, &child)

	result := ls(&dir)

	if len(result) != 1 {
		t.Errorf("Test_ls: Expected 1, got %d FAIL", len(result))
	} else {
		fmt.Printf("Test_ls: Expected 1, got %d SUCCESS\n", len(result))
	}

	if result[0] != "dir " +child.name {
		t.Errorf("Test_ls: Expected dir %s, got %s FAIL", child.name, result)
	} else {
		fmt.Printf("Test_ls: Expected dir %s, got %s SUCCESS\n", child.name, result)
	}
}
