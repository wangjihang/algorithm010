package n_tree_preorder

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPreorderIterative(t *testing.T) {
	root := preinit()
	Convey("TestPreorderIterative", t, func() {
		result := PreorderIterative(root)
		So(result, ShouldResemble, []int{1, 3, 5, 6, 2, 4})
	})
}

func BenchmarkPreorderRecursive(b *testing.B) {
	root := preinit()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PreorderRecursive(root)
	}
}

func BenchmarkPreorderRecursive2(b *testing.B) {
	root := preinit()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PreorderRecursive2(root)
	}
}

func preinit() *Node {
	return &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 2},
			{Val: 4},
		},
	}
}
