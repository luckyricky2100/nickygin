package classification

// import (
// 	"fmt"
// 	"strings"
// )

// type Category struct {
// 	ID       int
// 	Name     string
// 	ParentID int
// 	Children []*Category
// }

// var categories = []*Category{}

// func findCategory(id int) *Category {
// 	for _, c := range categories {
// 		if c.ID == id {
// 			return c
// 		}
// 	}
// 	return nil
// }

// func addCategory(id int, name string, parentID int) {
// 	parent := findCategory(parentID)
// 	category := &Category{ID: id, Name: name, ParentID: parentID}
// 	categories = append(categories, category)
// 	if parent != nil {
// 		parent.Children = append(parent.Children, category)
// 	}
// }

// func deleteCategory(id int) {
// 	for i, c := range categories {
// 		if c.ID == id {
// 			categories = append(categories[:i], categories[i+1:]...)
// 			break
// 		}
// 	}
// }

// func updateCategory(id int, name string) {
// 	for _, c := range categories {
// 		if c.ID == id {
// 			c.Name = name
// 			break
// 		}
// 	}
// }

// func printCategories(c *Category, level int) {
// 	if c == nil {
// 		for _, c := range categories {
// 			if c.ParentID == 0 {
// 				printCategories(c, 0)
// 			}
// 		}
// 	} else {
// 		fmt.Println(strings.Repeat(" ", level*2), c.Name)
// 		for _, child := range c.Children {
// 			printCategories(child, level+1)
// 		}
// 	}
// }

// func main() {
// 	addCategory(1, "Electronics", 0)
// 	addCategory(2, "Computers", 1)
// 	addCategory(3, "Mobiles", 1)
// 	addCategory(4, "Laptops", 2)
// 	addCategory(5, "Desktops", 2)
// 	printCategories(nil, 0)
// 	deleteCategory(5)
// 	updateCategory(4, "Notebooks")
// 	printCategories(nil, 0)
// }
