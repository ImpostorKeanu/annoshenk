package main

import (
    "github.com/go-yaml/yaml"
    "io"
    "os"
)

var (
    defaultDelimiter = "="
)

// Item is a commandline item, like a flag.
//
// Reconstructed format: <Flag><Delimiter><Value>
type Item struct {
    // Flag is the leading value that is the leading part of the
    // Item.
    Flag *string
    // Delimiter is the character sequence delimiting Flag and
    // Value.
    Delimiter *string
    // Default is the default value supplied to Value.
    Default *string
    // Value is the value part of the Item.
    Value *string
    // Comments is a slice of comments about the flag.
    Comments []string
    // ExampleValues is a slice of additional examples that can
    // be supplied.
    ExampleValues []string
    // Exclude determines if the Item should be omitted from
    // the command as it's being executed.
    Exclude bool
}

// Items is a container to hold a series of item objects that
// is later built to form a command.
type Items []Item

// Build the command from each Item.
func (items *Items) Build() (cmd []string) {
    for _, item := range *items {
        if !item.Exclude {
            cmd = append(cmd, item.Build()...)
        }
    }
    return cmd
}

// Build a flag out of the item object.
func (item *Item) Build() (flag []string) {
    // Set flag base
    flag = append(flag, *item.Flag)
    switch item.Delimiter {
    case nil:
        // Delimited flags become a solitary string value
        if item.Value != nil {
            flag = append(flag, *item.Value)
        } else if item.Default != nil {
            flag = append(flag, *item.Default)
        }
    default:
        // Non-delimited flags become a tuple
        flag[0] += *item.Delimiter
        if item.Value != nil {
            flag[0] += *item.Value
        } else if item.Default != nil {
            flag[0] += *item.Default
        }
    }
    return flag
}

func parseItemsFile(fPath string) (items Items, err error) {

    //=================================
    // READ AND UNMARSHAL THE ITEM FILE
    //=================================

    var f *os.File
    if f, err = os.Open(fPath); err != nil {
        return items, err
    }
    var b []byte
    items = Items{}
    if b, err = io.ReadAll(f); err != nil {
        return items, err
    }
    if err = yaml.Unmarshal(b, &items); err != nil {
        return items, err
    }

    //==============================================
    // UPDATE ITEMS THAT REQUIRE A DEFAULT DELIMITER
    //==============================================

    for i, item := range items {
        if item.Flag != nil && (item.Value != nil || item.Default != nil) && item.Delimiter == nil {
            items[i].Delimiter = &defaultDelimiter
        }
    }

    return items, err
}
