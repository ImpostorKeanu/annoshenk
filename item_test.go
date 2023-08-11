package main

import (
    "testing"
)

var (
    testDocPath = "test_doc.yml"
)

func TestItems_Build(t *testing.T) {
    if items, err := parseItemsFile(testDocPath); err != nil {
        t.Log("Failed to parse file of items")
        t.Fatal(err)
    } else {
        t.Log(items.Build())
    }
}
