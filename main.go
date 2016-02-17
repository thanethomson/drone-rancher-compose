/*
Simple plugin for Drone that allows execution of the Rancher Compose CLI.
*/
package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
  "github.com/drone/drone-plugin-go/plugin"
)

// We accept a list of Rancher Compose commands as input parameters.
type ComposerCommands struct {
  Commands []string `json:"commands"`
}

// Cleans out empty strings from the given array of strings.
func CleanSlice(arr []string) (result []string) {
  for _, v := range arr {
    if len(v) > 0 {
      result = append(result, v)
    }
  }
  return
}


func main() {
  var err error
  var path, wd string

  repo := plugin.Repo{}
  build := plugin.Build{}
  workspace := plugin.Workspace{}
  cmds := ComposerCommands{}

  plugin.Param("repo", &repo)
  plugin.Param("build", &build)
  plugin.Param("workspace", &workspace)
  plugin.Param("vargs", &cmds)
  err = plugin.Parse()

  if err != nil {
    fmt.Printf("Error while attempting to parse input: %s\n", err)
    os.Exit(1)
  }

  fmt.Printf("Got commands list:\n")
  for _, c := range cmds.Commands {
    fmt.Printf("- %s\n", c)
  }

  // let's try to find the Rancher Compose executable
  path, err = exec.LookPath("rancher-compose")
  if err != nil {
    fmt.Printf("Could not find rancher-compose executable on path: %s\n", err)
    os.Exit(1)
  }

  // chdir to the repo path
  if workspace.Path != "" {
    fmt.Printf("Using working directory: %s\n", workspace.Path)
    os.Chdir(workspace.Path)
  } else {
    wd, err = os.Getwd()
    if err != nil {
      fmt.Printf("Using working directory: %s\n", wd)
    } else {
      fmt.Printf("WARNING: Error while attempting to get working dir: %s\n", err)
    }
  }

  // execute each of our Rancher Compose commands in sequence
  for _, c := range cmds.Commands {
    args := CleanSlice(strings.Split(c, " "))
    fmt.Printf("Executing Rancher Compose: %s %s\n", path, strings.Join(args[:], " "))

    cmd := exec.Command(path, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()

    // exit if any command fails
    if err != nil {
      fmt.Printf("Error while executing command: %s\n", err)
      os.Exit(1)
    }
  }

  fmt.Printf("Successfully executed commands\n")

}
