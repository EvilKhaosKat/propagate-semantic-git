package main

import (
	"fmt"
	"github.com/libgit2/git2go/v31"
	"os"
)

func main() {
	path := os.Args[1]

	repo, err := git.OpenRepository(path)
	if err != nil {
		panic(err)
	}

	iter, err := repo.NewBranchIterator(git.BranchLocal)
	if err != nil {
		panic(err)
	}

	err = iter.ForEach(func(branch *git.Branch, branchType git.BranchType) error {
		name, err := branch.Name()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", name)

		return nil
	})

	if err != nil {
		panic(err)
	}
}
