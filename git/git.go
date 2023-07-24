package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Checkout ...
func Checkout(gitRepo, commitID string) {
	cmd := exec.Command("git", "checkout", commitID)
	cmd.Dir = gitRepo
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to git checkout due to cmd.Start: %v\n", err)
		os.Exit(1)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Failed to git checkout due to cmd.Wait: %v\n", err)
		os.Exit(1)
	}
}

// // GetCommitID ...
// func GetCommitID(gitRepo string) string {
// 	cmd := exec.Command("git", "rev-parse", "HEAD")
// 	cmd.Dir = gitRepo
// 	out, err := cmd.Output()
// 	if err != nil {
// 		fmt.Printf("Failed to get commit ID due to: %v\n", err)
// 		os.Exit(1)
// 	}
// 	return string(out)
// }

// GetCommitIDs ...
func GetCommitIDs(gitRepo string) []string {
	cmd := exec.Command("git", "log", "--format=format:%H")
	cmd.Dir = gitRepo
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to get commit ID due to: %v\n", err)
		os.Exit(1)
	}
	return strings.Split(string(out), "\n")
}
