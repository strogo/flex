package flex

import "testing"

func TestNested_overflowing_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 200)
	NodeStyleSetHeight(rootChild0Child0, 200)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, -100, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0Child0))
}

func TestNested_overflowing_child_in_constraint_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 100)
	NodeStyleSetHeight(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 200)
	NodeStyleSetHeight(rootChild0Child0, 200)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, -100, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0Child0))
}

func TestParent_wrap_child_size_overflowing_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	rootChild0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0, 100)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild0Child0, 100)
	NodeStyleSetHeight(rootChild0Child0, 200)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0Child0))

	NodeCalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 100, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0Child0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0Child0))
	assertFloatEqual(t, 100, NodeLayoutGetWidth(rootChild0Child0))
	assertFloatEqual(t, 200, NodeLayoutGetHeight(rootChild0Child0))
}
