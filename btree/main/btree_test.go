package main

import "testing"

func TestNewRootNode(t *testing.T) {
	root := NewRootNode(0)

	if root == nil {
		t.Errorf("Root node is nil")
	}
}

//NewRootNode 테스트
//1. keys의 첫번째 값으로 value가는지.
//2. parent가 nil
//3. child가 ni
