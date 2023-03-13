package localgit

import (
	"context"
	"fmt"
	"strings"

	"github.com/ldez/go-git-cmd-wrapper/v2/checkout"
	"github.com/ldez/go-git-cmd-wrapper/v2/git"
)

// Checkout new branch
// TODO catch error and do something
func Checkout(branchName string) {
	out, _ := git.Checkout(checkout.NewBranch(branchName), git.CmdExecutor(cmdExecutorMock))
	fmt.Println(out)
}

func cmdExecutorMock(_ context.Context, name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
