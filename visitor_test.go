package dag

import (
	"testing"

	"github.com/go-test/deep"
)

type testVisitor struct {
	Values []string
}

func (pv *testVisitor) Visit(v Vertexer) {
	_, value := v.Vertex()
	pv.Values = append(pv.Values, value.(string))
}

// schematic diagram:
//
//	v5
//	^
//	|
//	v4
//	^
//	|
//	v2 --> v3
//	^
//	|
//	v1
func getTestWalkDAG() *DAG {
	dag := NewDAG()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v2)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v2, v4)
	_ = dag.AddEdge(v4, v5)

	return dag
}

// schematic diagram:
//
//	v4 --> v5
//	       ^
//	       |
//	v1 --> v3
//	       ^
//	       |
//	      v2
func getTestWalkDAG2() *DAG {
	dag := NewDAG()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v3)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v3, v5)
	_ = dag.AddEdge(v4, v5)

	return dag
}

// schematic diagram:
//
//	v4 --> v5
//
//
//	v1 --> v3
//	       ^
//	       |
//	      v2
func getTestWalkDAG3() *DAG {
	dag := NewDAG()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v3)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v4, v5)

	return dag
}

// schematic diagram:
//
//	v4     v5
//	^      ^
//	|      |
//	v2 --> v3
//	^
//	|
//	v1
func getTestWalkDAG4() *DAG {
	dag := NewDAG()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v2)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v3, v5)
	_ = dag.AddEdge(v2, v4)

	return dag
}

// schematic diagram:
//
//	v5
//	^
//	|
//	v3 <-- v4
//	^      ^
//	|      |
//	v1     v2
func getTestWalkDAG5() *DAG {
	dag := NewDAG()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v3)
	_ = dag.AddEdge(v2, v4)
	_ = dag.AddEdge(v4, v3)
	_ = dag.AddEdge(v3, v5)

	return dag
}

// schematic diagram:
//
//	v5
//	^
//	|
//	v3 <-- v4 <-- v6 <--v7
//	^      ^
//	|      |
//	v1     v2
func getTestWalkDAG6() *DAG {
	dag := NewDAG()

	v1, v2, v3, v4, v5, v6, v7 := "1", "2", "3", "4", "5", "6", "7"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddVertexByID(v6, "v6")
	_ = dag.AddVertexByID(v7, "v7")
	_ = dag.AddEdge(v1, v3)
	_ = dag.AddEdge(v2, v4)
	_ = dag.AddEdge(v4, v3)
	_ = dag.AddEdge(v3, v5)
	_ = dag.AddEdge(v6, v4)
	_ = dag.AddEdge(v7, v6)

	return dag
}

func TestDFSWalk(t *testing.T) {
	cases := []struct {
		dag      *DAG
		expected []string
	}{
		{
			dag:      getTestWalkDAG(),
			expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			dag:      getTestWalkDAG2(),
			expected: []string{"v1", "v3", "v5", "v2", "v4"},
		},
		{
			dag:      getTestWalkDAG3(),
			expected: []string{"v1", "v3", "v2", "v4", "v5"},
		},
		{
			dag:      getTestWalkDAG4(),
			expected: []string{"v1", "v2", "v3", "v5", "v4"},
		},
		{
			dag:      getTestWalkDAG5(),
			expected: []string{"v1", "v3", "v5", "v2", "v4"},
		},
	}

	for _, c := range cases {
		pv := &testVisitor{}
		c.dag.DFSWalk(pv)

		expected := c.expected
		actual := pv.Values
		if deep.Equal(expected, actual) != nil {
			t.Errorf("DFSWalk() = %v, want %v", actual, expected)
		}
	}
}

func TestLongestChain(t *testing.T) {
	cases := []struct {
		dag      *DAG
		expected []string
	}{
		{
			dag:      getTestWalkDAG(),
			expected: []string{"1"},
		},
		{
			dag:      getTestWalkDAG4(),
			expected: []string{"1"},
		},
		{
			dag:      getTestWalkDAG5(),
			expected: []string{"2", "1"},
		},
		{
			dag:      getTestWalkDAG6(),
			expected: []string{"7", "2", "1"},
		},
	}

	for _, c := range cases {
		actual := c.dag.FindNodeWithLongestChain()

		expected := c.expected
		if deep.Equal(expected, actual) != nil {
			t.Errorf("LongestChain() = %v, want %v", actual, expected)
		}
	}
}

func TestBFSWalk(t *testing.T) {
	cases := []struct {
		dag      *DAG
		expected []string
	}{
		{
			dag:      getTestWalkDAG(),
			expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			dag:      getTestWalkDAG2(),
			expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
		{
			dag:      getTestWalkDAG3(),
			expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
		{
			dag:      getTestWalkDAG4(),
			expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			dag:      getTestWalkDAG5(),
			expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
	}

	for _, c := range cases {
		pv := &testVisitor{}
		c.dag.BFSWalk(pv)

		expected := c.expected
		actual := pv.Values
		if deep.Equal(expected, actual) != nil {
			t.Errorf("BFSWalk() = %v, want %v", actual, expected)
		}
	}
}

func TestOrderedWalk(t *testing.T) {
	cases := []struct {
		dag      *DAG
		expected []string
	}{
		{
			dag:      getTestWalkDAG(),
			expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			dag:      getTestWalkDAG2(),
			expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
		{
			dag:      getTestWalkDAG3(),
			expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
		{
			dag:      getTestWalkDAG4(),
			expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			dag:      getTestWalkDAG5(),
			expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
	}

	for _, c := range cases {
		pv := &testVisitor{}
		c.dag.OrderedWalk(pv)

		expected := c.expected
		actual := pv.Values
		if deep.Equal(expected, actual) != nil {
			t.Errorf("OrderedWalk() = %v, want %v", actual, expected)
		}
	}
}
