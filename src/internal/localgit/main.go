package localgit

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/v2/branch"
	"github.com/ldez/go-git-cmd-wrapper/v2/checkout"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
	"github.com/ldez/go-git-cmd-wrapper/v2/types"
)

func isDryRun() bool {
	return os.Getenv("DRY_RUN") == "true"
}

// Checkout new branch
// TODO catch error and do something
func NewBranch(targetB string) {
	fmt.Printf("Currently on branch: '%s'\n", GetCurrentBranch())
	fmt.Printf("Creating new branch '%s'\n", targetB)
	if isDryRun() {
		out, _ := git.Checkout(checkout.NewBranchForce(targetB), git.CmdExecutor(cmdExecutorMock))
		fmt.Println(out)
	} else {
		out, err := git.Checkout(checkout.NewBranchForce(targetB))
		fmt.Println(out)
		if err != nil {
			fmt.Printf("Unable to create branch %s error: %s", targetB, err)
		}
	}
}

func ChangeBranch(targetB string) {
	fmt.Printf("Changing branch from '%s' to '%s'\n", GetCurrentBranch(), targetB)
	if isDryRun() {
		out, _ := git.Checkout(checkout.Branch(targetB), git.CmdExecutor(cmdExecutorMock))
		fmt.Println(out)
	} else {
		out, _ := git.Checkout(checkout.Branch(targetB))
		fmt.Println(out)
	}
}

func GetCurrentBranch() string {
	branch, _ := git.Branch(branch.ShowCurrent)
	branch = strings.Replace(branch, "\n", "", -1)
	return branch
}

func CherryPickCommit(hash string) {
	fmt.Printf("Cherry-picking commit: '%s'\n", hash)
	out, _ := CherryPick(CherryPickByHash(hash), git.CmdExecutor(cmdExecutorMock))
	fmt.Println(out)
}

func cmdExecutorMock(_ context.Context, name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}

func command(ctx context.Context, name string, options ...types.Option) (string, error) {
	g := types.NewCmd(name)
	g.ApplyOptions(options...)

	return g.Exec(ctx, g.Base, g.Debug, g.Options...)
}

// Cherry-pick
func CherryPick(options ...types.Option) (string, error) {
	return command(context.Background(), "cherry-pick", options...)
}

func CherryPickByHash(commitHash string) func(*types.Cmd) {
	return func(g *types.Cmd) {
		g.AddOptions(commitHash)
	}
}
