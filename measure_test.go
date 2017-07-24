package flex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func _measure3(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	measureCount, ok := node.Context.(*int)
	if ok {
		(*measureCount)++
	}

	return Size{Width: 10, Height: 10}
}

func _simulate_wrapping_text(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if widthMode == MeasureModeUndefined || width >= 68 {
		return Size{Width: 68, Height: 16}
	}

	return Size{Width: 50, Height: 32}
}

func _measure_assert_negative(node *Node, width float32, widthMode MeasureMode, height float32, heightMode MeasureMode) Size {
	if width < 0 {
		panic(fmt.Sprintf("width is %.2f and should be >= 0", width))
	}
	if height < 0 {
		panic(fmt.Sprintf("height is %.2f should be >= 0, height", height))
	}
	// EXPECT_GE(width, 0);
	//EXPECT_GE(height, 0);

	return Size{
		Width: 0, Height: 0,
	}
}

func TestDont_measure_single_grow_shrink_child(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure)
	NodeStyleSetFlexGrow(rootChild0, 1)
	NodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
}

func TestMeasure_absolute_child_with_no_constraints(t *testing.T) {
	root := NewNode()

	rootChild0 := NewNode()
	YGNodeInsertChild(root, rootChild0, 0)

	measureCount := 0

	rootChild0Child0 := NewNode()
	NodeStyleSetPositionType(rootChild0Child0, PositionTypeAbsolute)
	rootChild0Child0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0Child0, _measure3)
	YGNodeInsertChild(rootChild0, rootChild0Child0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 1, measureCount)
}

func TestDont_measure_when_min_equals_max(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	NodeStyleSetMinWidth(rootChild0, 10)
	NodeStyleSetMaxWidth(rootChild0, 10)
	NodeStyleSetMinHeight(rootChild0, 10)
	NodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_percentages(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	NodeStyleSetMinWidthPercent(rootChild0, 10)
	NodeStyleSetMaxWidthPercent(rootChild0, 10)
	NodeStyleSetMinHeightPercent(rootChild0, 10)
	NodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_mixed_width_percent(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	NodeStyleSetMinWidthPercent(rootChild0, 10)
	NodeStyleSetMaxWidthPercent(rootChild0, 10)
	NodeStyleSetMinHeight(rootChild0, 10)
	NodeStyleSetMaxHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestDont_measure_when_min_equals_max_mixed_height_percent(t *testing.T) {
	root := NewNode()
	NodeStyleSetAlignItems(root, AlignFlexStart)
	NodeStyleSetWidth(root, 100)
	NodeStyleSetHeight(root, 100)

	measureCount := 0

	rootChild0 := NewNode()
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)
	NodeStyleSetMinWidth(rootChild0, 10)
	NodeStyleSetMaxWidth(rootChild0, 10)
	NodeStyleSetMinHeightPercent(rootChild0, 10)
	NodeStyleSetMaxHeightPercent(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assert.Equal(t, 0, measureCount)
	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))
}

func TestMeasure_enough_size_should_be_in_single_line(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 100)

	rootChild0 := NewNode()
	NodeStyleSetAlignSelf(rootChild0, AlignFlexStart)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 68, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 16, NodeLayoutGetHeight(rootChild0))
}

func TestMeasure_not_enough_size_should_wrap(t *testing.T) {
	root := NewNode()
	NodeStyleSetWidth(root, 55)

	rootChild0 := NewNode()
	NodeStyleSetAlignSelf(rootChild0, AlignFlexStart)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)

	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, NodeLayoutGetHeight(rootChild0))
}

func TestMeasure_zero_space_should_grow(t *testing.T) {
	root := NewNode()
	NodeStyleSetHeight(root, 200)
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetFlexGrow(root, 0)

	measureCount := 0

	rootChild0 := NewNode()
	NodeStyleSetFlexDirection(rootChild0, FlexDirectionColumn)
	NodeStyleSetPadding(rootChild0, EdgeAll, 100)
	rootChild0.Context = &measureCount
	NodeSetMeasureFunc(rootChild0, _measure3)

	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, 282, Undefined, DirectionLTR)

	assertFloatEqual(t, 282, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
}

func TestMeasure_flex_direction_row_and_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetPadding(root, EdgeLeft, 25)
	NodeStyleSetPadding(root, EdgeTop, 25)
	NodeStyleSetPadding(root, EdgeRight, 25)
	NodeStyleSetPadding(root, EdgeBottom, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 0, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 75, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_column_and_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetPadding(root, EdgeAll, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 57, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_row_no_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

func TestMeasure_flex_direction_row_no_padding_align_items_flexstart(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)
	NodeStyleSetAlignItems(root, AlignFlexStart)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 50, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

func TestMeasure_with_fixed_size(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetPadding(root, EdgeAll, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeStyleSetWidth(rootChild0, 10)
	NodeStyleSetHeight(rootChild0, 10)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 10, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 35, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

func TestMeasure_with_flex_shrink(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetPadding(root, EdgeAll, 25)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 25, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 25, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 25, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

func TestMeasure_no_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetMargin(root, EdgeTop, 20)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 50)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _simulate_wrapping_text)
	NodeStyleSetFlexShrink(rootChild0, 1)
	YGNodeInsertChild(root, rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	NodeStyleSetWidth(rootChild1, 5)
	NodeStyleSetHeight(rootChild1, 5)
	YGNodeInsertChild(root, rootChild1, 1)
	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, NodeLayoutGetLeft(root))
	assertFloatEqual(t, 20, NodeLayoutGetTop(root))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(root))
	assertFloatEqual(t, 50, NodeLayoutGetHeight(root))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild0))
	assertFloatEqual(t, 0, NodeLayoutGetTop(rootChild0))
	assertFloatEqual(t, 50, NodeLayoutGetWidth(rootChild0))
	assertFloatEqual(t, 32, NodeLayoutGetHeight(rootChild0))

	assertFloatEqual(t, 0, NodeLayoutGetLeft(rootChild1))
	assertFloatEqual(t, 32, NodeLayoutGetTop(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetWidth(rootChild1))
	assertFloatEqual(t, 5, NodeLayoutGetHeight(rootChild1))
}

/*
#if GTEST_HAS_DEATH_TEST
TEST(YogaDeathTest, cannot_add_child_to_node_with_measure_func) {
  root := YGNodeNew();
  YGNodeSetMeasureFunc(root, _measure3);

  rootChild0 := YGNodeNew();
  ASSERT_DEATH(YGNodeInsertChild(root, rootChild0, 0), "Cannot add child.*");
  YGNodeFree(rootChild0);
  ;
}

TEST(YogaDeathTest, cannot_add_nonnull_measure_func_to_non_leaf_node) {
  root := YGNodeNew();
  rootChild0 := YGNodeNew();
  YGNodeInsertChild(root, rootChild0, 0);

  ASSERT_DEATH(YGNodeSetMeasureFunc(root, _measure3), "Cannot set measure function.*");
  ;
}
#endif
*/

func TestCan_nullify_measure_func_on_any_node(t *testing.T) {
	root := NewNode()
	YGNodeInsertChild(root, NewNode(), 0)

	NodeSetMeasureFunc(root, nil)
	assert.True(t, root.Measure == nil)
}

func TestCant_call_negative_measure(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionColumn)
	NodeStyleSetWidth(root, 50)
	NodeStyleSetHeight(root, 10)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	NodeStyleSetMargin(rootChild0, EdgeTop, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
}

func TestCant_call_negative_measure_horizontal(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	NodeStyleSetFlexDirection(root, FlexDirectionRow)
	NodeStyleSetWidth(root, 10)
	NodeStyleSetHeight(root, 20)

	rootChild0 := NewNodeWithConfig(config)
	NodeSetMeasureFunc(rootChild0, _measure_assert_negative)
	NodeStyleSetMargin(rootChild0, EdgeStart, 20)
	YGNodeInsertChild(root, rootChild0, 0)

	NodeCalculateLayout(root, Undefined, Undefined, DirectionLTR)
}
