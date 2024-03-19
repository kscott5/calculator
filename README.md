# [Understand how to use packages, variables, and functions in Go](https://learn.microsoft.com/en-us/training/modules/go-variables-functions-packages/)
Learn about the basic data types in Go and about how to declare variables, write functions, and use packages.
36 min    Module    7 Units

## [Learning about packages and modules](https://learn.microsoft.com/en-us/training/modules/go-variables-functions-packages/4-packages/?ns-enrollment-type=learningpath&ns-enrollment-id=learn.languages.go-first-steps)
Packages in Go are like libraries or modules in other programming languages. You can package your code and reuse it somewhere else. The source code of a package can be distributed in more than one .go file. So far, we've been writing the main package and have made a few references to other native packages.
10 min

In this section, you'll learn what a package is. You'll also learn how to create one and how to consume external packages.

## Reference a local package (a module) without github.com account

> [!NOTE]
> Go mod provides access to operations on modules.
>
> usage: go mod edit [editing flags] [-fmt|-print|-json] [go.mod]
>
> Edit provides a command-line interface for editing go.mod,
> for use primarily by tools or scripts. It reads only go.mod;
> it does not look up information about the modules involved.
> By default, edit reads and writes the go.mod file of the main module,
> but a different target file can be specified after the editing flags.
>
> See [https://golang.org/ref/mod#go-mod-edit](https://golang.org/ref/mod#go-mod-edit) for more about 'go mod edit'.


```shell
mkdir $GOPATH/src/helloworld
cd $GOPATH/src/helloworld
go mod edit -require=gitub.com/$USER/calculator@v0.0.0
go mod edit -replace github.com/$USER/calculator=../calculator
```



