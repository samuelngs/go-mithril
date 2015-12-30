package mithril

import "testing"

func TestIDAttribute(t *testing.T) {
	attr := ID("awesome")
	if attr.Val() != "awesome" {
		t.Fatalf("Expected id to be \"awesome\" but it was %s", attr)
	}
}

func TestOverrideID(t *testing.T) {
	attr := ID("awesome")
	attr.Val("cool")
	if attr.Val() != "cool" {
		t.Fatalf("Expected id to be \"cool\" but it was %s", attr)
	}
}

func TestClassAttribute(t *testing.T) {
	attr := Class("class-1")
	if attr.Val() != "class-1" {
		t.Fatalf("Expected id to be \"cool\" but it was %s", attr)
	}
}

func TestMultipleClassnames(t *testing.T) {
	attr := Class("class-1", "class-2", "class-3")
	expected := "class-1 class-2 class-3"
	if attr.Val() != expected {
		t.Fatalf("Expected id to be \"%s\" but it was %s", expected, attr)
	}
}
