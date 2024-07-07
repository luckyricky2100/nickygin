package classification

import (
	"fmt"
	"strings"
)

// 预排序遍历树算法：modified preorder tree traversal algorithm
type NestedCategory struct {
	ID    int
	Name  string
	Left  int
	Right int
}

type NoestedTree struct {
	Categories []*NestedCategory
}

func (tree *NoestedTree) findNestedCategory(id int) *NestedCategory {
	for _, c := range tree.Categories {
		if c.ID == id {
			return c
		}
	}
	return nil
}

func (tree *NoestedTree) addNestedCategory(id int, name string, parentID int) {
	parent := tree.findNestedCategory(parentID)
	nestedCategory := &NestedCategory{ID: id, Name: name}
	if parent != nil {
		nestedCategory.Left = parent.Right
		nestedCategory.Right = nestedCategory.Left + 1

		for _, v := range tree.Categories {
			if v.Left > parent.Right {
				v.Left += 2
			}
			if v.Right > parent.Right {
				v.Right += 2
			}
		}

		parent.Right += 2
	} else {
		nestedCategory.Left = 1
		nestedCategory.Right = 2
	}

	tree.Categories = append(tree.Categories, nestedCategory)
}

func (tree *NoestedTree) deleteNestedCategory(id int) {
	for i, c := range tree.Categories {
		if c.ID == id {
			tree.Categories = append(tree.Categories[:i], tree.Categories[i+1:]...)
			break
		}
	}
}

func (tree *NoestedTree) updateNestedCategory(id int, name string) {
	for _, c := range tree.Categories {
		if c.ID == id {
			c.Name = name
			break
		}
	}
}

func printCategories(categories []*NestedCategory) {
	for _, c := range categories {
		fmt.Println(strings.Repeat(" ", (c.Right-c.Left)/2), c.Name, " ", c.Left, " ", c.Right)
	}
}

func NestedCategoryTest() {
	tree := NoestedTree{Categories: []*NestedCategory{}}
	tree.addNestedCategory(1, "Electronics", 0)
	tree.addNestedCategory(2, "Computers", 1)
	tree.addNestedCategory(3, "Mobiles", 1)
	tree.addNestedCategory(4, "Laptops", 2)
	tree.addNestedCategory(5, "Desktops", 2)
	printCategories(tree.Categories)
	fmt.Println("----------------------------------")
	tree.deleteNestedCategory(5)
	tree.updateNestedCategory(4, "Notebooks")
	printCategories(tree.Categories)
}

// func convertToPreorderNodes() {
// 	nestedCategory := []*NestedCategory{}
// 	convertSubtreeToPreorderNodes(findTreeNode(0), 1, nestedCategory)
// }

// func convertSubtreeToPreorderNodes(node *TreeNode, left int, nestedCategory []*NestedCategory) int {
// 	if node == nil {
// 		return left
// 	}
// 	right := left + 1
// 	for _, child := range node.Children {
// 		right = convertSubtreeToPreorderNodes(child, right)
// 	}
// 	preorderNodes = append(preorderNodes, &PreorderNode{ID: node.ID, Name: node.Name, Left: left, Right: right})
// 	return right + 1
// }
