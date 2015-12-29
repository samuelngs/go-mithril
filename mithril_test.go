package mithril

import "testing"

func TestCreateElement(t *testing.T) {
	el := M("div#id-value.lol.class-name[data1='value1'][data2=\"value2\"][data3]")
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
	if el := M("div#id-value"); el.Attr("id") != "id-value" {
		t.Fatalf("Expected element ID to be id-value but it was %s", el.Attr("id"))
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
	el := M("ul", []Attribute{
		NewStringAttr("id", "listview"),
	})
	if el.Attr("id") != "listview" {
		t.Fatalf("Expected element ID to be a listview but it was %s", el.Attr("id"))
	}
}
