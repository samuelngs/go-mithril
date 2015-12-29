package mithril

import "testing"

func TestCreateElement(t *testing.T) {
	el := M("div#id-value.class-name[data1='value1'][data2=\"value2\"][data3]")
	if el == nil {
		t.Fatalf("Expected element to be a VirtualElement object but it was nil")
	}
}

func TestElementTag(t *testing.T) {
	if el := M("li"); el.Tag != "li" {
		t.Fatalf("Expected element tag to be \"li\" but it was %s", el.Tag)
	}
}

func TestElementID(t *testing.T) {
	if el := M("div#id-value"); el.Attrs.ID != "id-value" {
		t.Fatalf("Expected element ID to be id-value but it was %s", el.Attrs.ID)
	}
}

func TestCreateElementChildren(t *testing.T) {
	el := M(
		"div#obj1",
		M("div#obj2"),
	)
	if _, ok := el.Children.(*VirtualElement); !ok {
		t.Fatalf("Expected element.Children to be a VirtualElement")
	}
}

func TestCreateElementWithAttr(t *testing.T) {
	attr := NewAttributes()
	attr.ID = "listview"
	el := M("ul", attr)
	if el.Attrs.ID != "listview" {
		t.Fatalf("Expected element ID to be a listview but it was %s", el.Attrs.ID)
	}
}
