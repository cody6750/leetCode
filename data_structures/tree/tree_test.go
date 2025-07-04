package datastructures

import (
	"testing"
)

func Test_testTree(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTree()
		})
	}
}

func TestTree_BFS(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
		want bool
	}{
		{
			tr: initTree(),
			args: args{
				v: 10,
			},
			want: false,
		},
		{
			tr: initTree(),
			args: args{
				v: 8,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.BFS(tt.args.v); got != tt.want {
				t.Errorf("Tree.BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Delete(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		{
			tr: initTree(),
			args: args{
				v: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Delete(tt.args.v)
		})
	}
}

func TestNode_InOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringify(initTree().root, 0)
			initTree().root.PreOrderTraverse()
		})
	}
}

func TestNode_MinDepth(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want int
	}{
		{
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := initTree()
			stringify(tree.root, 0)
			if got := tree.root.MinDepth(); got != tt.want {
				t.Errorf("Node.MinDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_MaxDepth(t *testing.T) {
	tests := []struct {
		name string
		n    *Node
		want int
	}{
		{
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := initTree()
			stringify(tree.root, 0)
			if got := tree.root.MaxDepth(); got != tt.want {
				t.Errorf("Node.MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
