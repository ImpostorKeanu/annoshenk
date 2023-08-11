package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
    "strings"
)

var (
    cmdDesc    = "Execute arbitrary commands via configuration file."
    head, tail []string
    rootCmd    = cobra.Command{
        Use:     "shenk <CONFIG_FILE.yml> [--head|-H]... [--tail|-t]...",
        Args:    cobra.MinimumNArgs(1),
        Short:   cmdDesc,
        Long:    cmdDesc,
        Example: "shenk nmap.yml -t targets/*",
        Run:     run,
    }
)

func init() {
    rootCmd.Flags().StringArrayVarP(&head, "head", "H", nil,
        "Flags prefixed to the command prior to execution.")
    rootCmd.Flags().StringArrayVarP(&tail, "tail", "t", nil,
        "Flags suffixed to the command prior to execution.")
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        eLog.Fatalf("Failed to execute command: %v", err)
    }
}

func run(c *cobra.Command, args []string) {

    iLog.Print("Someone's feeling stabby! 🔪")

    //======================
    // PARSE THE CONFIG FILE
    //======================

    iLog.Printf("Parsing config file: %s", args[0])
    var items Items
    var err error
    if items, err = parseItemsFile(args[0]); err != nil {
        eLog.Fatalf("Failed to parse config file: %v", err)
    }

    //========================
    // PARSE CONFIG FILE ITEMS
    //========================

    iLog.Print("Translating configuration items")
    parts := items.Build()

    //============================
    // PROCESS HEAD AND TAIL PARTS
    //============================

    var name string
    if head != nil {
        name = head[0]
        if len(head) > 1 {
            parts = append(head[1:], parts...)
        }
    } else {
        name = parts[0]
        if len(parts) > 1 {
            parts = parts[1:]
        }
    }

    // Append tail to remaining parts
    if tail != nil {
        parts = append(parts, tail...)
    }

    //====================
    // EXECUTE THE COMMAND
    //====================

    wLog.Printf("Command name: %v", name)
    wLog.Printf("Command flags & arguments: %v", strings.Join(parts, " "))
    fmt.Print("\n\n")
    cmd := exec.Command(name, parts...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if cmd.Run(); err != nil {
        eLog.Fatalf("Failed to execute command: %v", err)
    }

    fmt.Print("\n\n")
    iLog.Printf("Execution complete")
}
