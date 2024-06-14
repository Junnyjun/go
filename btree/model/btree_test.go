package model

import "testing"

func TestNewRootNode(t *testing.T) {
	root := NewBTree()

	if root != nil {
		t.Errorf("root.keys[0] does not match")
	}

}

//NewRootNode 테스트
//1. keys의 첫번째 값으로 value가는지.
//2. parent가 nil
//3. child가 nil
//4.isLeaf가 false
