package model

import "testing"

func TestNewRootNode(t *testing.T) {
	root := NewRootNode(0)

	if root.keys[0] != 0 {
		t.Errorf("root.keys[0] should be 0")
	}
	if root.parent != nil {
		t.Errorf("root.parent should be nil")
	}
	if len(root.child) != 0 {
		t.Errorf("root.child should be nil")
	}
	if root.isLeaf != false {
		t.Errorf("root.isLeaf should be false")
	}
}

//NewRootNode 테스트
//1. keys의 첫번째 값으로 value가는지.
//2. parent가 nil
//3. child가 nil
//4.isLeaf가 false
