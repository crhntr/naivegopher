package naivegopher

import (
	"sort"
	"testing"
)

func TestNewClassifier(t *testing.T) {
	NewClassifier()
}

func TestFindOrInsert(t *testing.T) {
	c := NewClassifier()
	if len(c.Categories) != 0 {
		t.Fail()
	}
	c.findOrInsert("foo")
	if len(c.Categories) != 1 {
		t.Fail()
	}
	c.findOrInsert("bar")
	if len(c.Categories) != 2 {
		t.Fail()
	}
	c.findOrInsert("bar")
	if len(c.Categories) != 2 {
		t.Fail()
	}
}

func TestCategoryNames(t *testing.T) {
	c := NewClassifier()
	if len(c.CategoryNames()) != 0 {
		t.Fail()
	}
	c.findOrInsert("foo")
	if len(c.CategoryNames()) != 1 {
		t.Fail()
	}
	c.findOrInsert("bar")
	if len(c.CategoryNames()) != 2 {
		t.Fail()
	}
	c.findOrInsert("bar")
	if len(c.CategoryNames()) != 2 {
		t.Fail()
	}
}

func TestFindCategory(t *testing.T) {
	c := NewClassifier()
	if c.FindCategory("foo") >= 0 {
		t.Fail()
	}
	c.findOrInsert("foo")
	if c.FindCategory("foo") < 0 {
		t.Fail()
	}
	c.findOrInsert("bar")
	if c.FindCategory("bar") != c.FindCategory("bar") {
		t.Fail()
	}
}

func TestSort(t *testing.T) {
	c := NewClassifier()
	c.Categories = append(c.Categories,
		Category{Name: "foo"},
		Category{Name: "bar"},
		Category{Name: "que"},
		Category{Name: "baz"},
		Category{Name: "asd"},
		Category{Name: "jack"},
	)
	sort.Sort(c)
}
