# Chapter 10: Packages and the Go Tool

<!-- TOC -->

- [10.1. Introduction](#101-introduction)
- [10.2. Import Paths](#102-import-paths)
- [10.3. The Package Declaration](#103-the-package-declaration)
- [10.4. Import Declarations](#104-import-declarations)
- [10.5. Blank Imports](#105-blank-imports)
- [10.6. Packages and Naming](#106-packages-and-naming)
- [10.7. The Go Tool](#107-the-go-tool)
  - [10.7.1. Workspace Organization](#1071-workspace-organization)
  - [10.7.2. Downloading Packages](#1072-downloading-packages)
  - [10.7.3. Building Packages](#1073-building-packages)
  - [10.7.4. Documenting Packages](#1074-documenting-packages)
  - [10.7.5. Internal Packages](#1075-internal-packages)
  - [10.7.6. Querying Packages](#1076-querying-packages)

<!-- /TOC -->

A modest-size program today might contain 10,000 functions. Yet its author need think about only a few of them and design even fewer, because the vast majority were written by others and made available for reuse through *packages*.

Go comes with over 100 standard packages that provide the foundations for most applications. The Go community, a thriving ecosystem of package design, sharing, reuse, and improvement, has published many more, and you can find a searchable index of them at `http://godoc.org`. In this chapter, we'll show how to use existing packages and create new ones.

Go also comes with the go tool, a sophisticated but simple-to-use command for managing workspaces of Go packages. Since the beginning of the book, we've been showing how to use the go tool to download, build, and run example programs. In this chapter, we'll look at the tool's underlying concepts and tour more of its capabilities, which include printing documentation and querying metadata about the packages in the workspace. In the next chapter we'll explore its testing features.


## 10.1. Introduction

The purpose of any package system is to make the design and maintenance of large programs practical by grouping related features together into units that can be easily understood and changed, independent of the other packages of the program. This *modularity* allows packages to be shared and reused by different projects, distributed within an organization, or made available to the wider world.

Each package defines a distinct name space that encloses its identifiers. Each name is associated with a particular package, letting us choose short, clear names for the types, functions, and so on that we use most often, without creating conflicts with other parts of the program.

Packages also provide *encapsulation* by controlling which names are visible or exported outside the package. Restricting the visibility of package members hides the helper functions and types behind the package's API, allowing the package maintainer to change the implementation with confidence that no code outside the package will be affected. Restricting visibility also hides variables so that clients can access and update them only through exported functions that preserve internal invariants or enforce mutual exclusion in a concurrent program.

When we change a file, we must recompile the file's package and potentially all the packages that depend on it. Go compilation is notably faster than most other compiled languages, even when building from scratch. There are three main reasons for the compiler's speed. First, all imports must be explicitly listed at the beginning of each source file, so the compiler does not have to read and process an entire file to determine its dependencies. Second, the dependencies of a package form a directed acyclic graph, and because there are no cycles, packages can be compiled separately and perhaps in parallel. Finally, the object file for a compiled Go package records export information not just for the package itself, but for its dependencies too. When compiling a package, the compiler must read one object file for each import but need not look beyond these files.


## 10.2. Import Paths

Each package is identified by a unique string called its *import path*. Import paths are the strings that appear in import declarations.
```go
    import (
        "fmt"
        "math/rand"
        "encoding/json"
        "golang.org/x/net/html"
        "github.com/go-sql-driver/mysql"
    )
```
As we mentioned in Section 2.6.1, the Go language specification doesn't define the meaning of these strings or how to determine a package's import path, but leaves these issues to the tools. In this chapter, we'll take a detailed look at how the go tool interprets them, since that's what the majority of Go programmers use for building, testing, and so on. Other tools do exist, though. For example, Go programmers using Google's internal multi-language build system follow different rules for naming and locating packages, specifying tests, and so on, that more closely match the conventions of that system.

For packages you intend to share or publish, import paths should be globally unique. To avoid conflicts, the import paths of all packages other than those from the standard library should start with the Internet domain name of the organization that owns or hosts the package; this also makes it possible to find packages. For example, the declaration above imports an HTML parser maintained by the Go team and a popular third-party MySQL database driver.


## 10.3. The Package Declaration 

A `package` declaration is required at the start of every Go source file. Its main purpose is to determine the default identifier for that package (called the *package name*) when it is imported by another package.

For example, every file of the `math/rand` package starts with `package` rand, so when you import this package, you can access its members as `rand.Int`, `rand.Float64`, and so on.
```go
    package main

    import (
        "fmt"
        "math/rand"
    )

    func main() {
        fmt.Println(rand.Int())
    }
```
Conventionally, the package name is the last segment of the import path, and as a result, two packages may have the same name even though their import paths necessarily differ. For example, the packages whose import paths are `math/rand` and `crypto/rand` both have the name rand. We'll see how to use both in the same program in a moment.

There are three major exceptions to the "last segment" convention. The first is that a package defining a command (an executable Go program) always has the name `main`, regardless of the package's import path. This is a signal to go build (§10.7.3) that it must invoke the linker to make an executable file.

The second exception is that some files in the directory may have the suffix `_test` on their package name if the file name ends with `_test.go`. Such a directory may define *two* packages: the usual one, plus another one called an *external test package*. The `_test` suffix signals to `go test` that it must build both packages, and it indicates which files belong to each package. External test packages are used to avoid cycles in the import graph arising from dependencies of the test; they are covered in more detail in Section 11.2.4.

The third exception is that some tools for dependency management append version number suffixes to package import paths, such as `"gopkg.in/yaml.v2"`. The package name excludes the suffix, so in this case it would be just `yaml`.


## 10.4. Import Declarations 


A Go source file may contain zero or more import declarations immediately after the `package` declaration and before the first non-import declaration. Each import declaration may specify the import path of a single package, or multiple packages in a parenthesized list. The two forms below are equivalent but the second form is more common.
```go
    import "fmt"
    import "os"
       
    import (
        "fmt"
        "os"
    )
```
Imported packages may be grouped by introducing blank lines; such groupings usually indicate different domains. The order is not significant, but by convention the lines of each group are sorted alphabetically. (Both `gofmt` and `goimports` will group and sort for you.)
```go
    import (
        "fmt"
        "html/template"
        "os"
        "golang.org/x/net/html"
        "golang.org/x/net/ipv4"
    )
```
If we need to import two packages whose names are the same, like `math/rand` and `crypto/rand`, into a third package, the import declaration must specify an alternative name for at least one of them to avoid a conflict. This is called a *renaming import*.
```go
    import (
        "crypto/rand"
        mrand "math/rand" // alternative name mrand avoids conflict
    )
```
The alternative name affects only the importing file. Other files, even ones in the same package, may import the package using its default name, or a different name.

A renaming import may be useful even when there is no conflict. If the name of the imported package is unwieldy, as is sometimes the case for automatically generated code, an abbreviated name may be more convenient. The same short name should be used consistently to avoid confusion. Choosing an alternative name can help avoid conflicts with common local variable names. For example, in a file with many local variables named path, we might import the standard `"path"` package as `pathpkg`.

Each import declaration establishes a dependency from the current package to the imported package. The `go build` tool reports an error if these dependencies form a cycle.


## 10.5. Blank Imports

It is an error to import a package into a file but not refer to the name it defines within that file. However, on occasion we must import a package merely for the side effects of doing so: evaluation of the initializer expressions of its package-level variables and execution of its `init` functions (§2.6.2). To suppress the "unused import" error we would otherwise encounter, we must use a renaming import in which the alternative name is `_`, the blank identifier. As usual, the blank identifier can never be referenced.
```go
    import _ "image/png" // register PNG decoder
```
This is known as a *blank import*. It is most often used to implement a compile-time mechanism whereby the main program can enable optional features by blank-importing additional packages. First we'll see how to use it, then we'll see how it works.

The standard library's `image` package exports a `Decode` function that reads bytes from an `io.Reader`, figures out which image format was used to encode the data, invokes the appropriate decoder, then returns the resulting `image.Image.` Using `image.Decode`, it's easy to build a simple image converter that reads an image in one format and writes it out in another:
```go
// gopl.io/ch10/jpeg
// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
```
If we feed the output of `gopl.io/ch3/mandelbrot` (§3.3) to the converter program, it detects the PNG input format and writes a JPEG version of Figure 3.3.
```
    $ go build gopl.io/ch3/mandelbrot
    $ go build gopl.io/ch10/jpeg
    $ ./mandelbrot | ./jpeg >mandelbrot.jpg
    Input format = png
```
Notice the blank import of `image/png`. Without that line, the program compiles and links as usual but can no longer recognize or decode input in PNG format:
```
    $ go build gopl.io/ch10/jpeg
    $ ./mandelbrot | ./jpeg >mandelbrot.jpg
    jpeg: image: unknown format
```
Here's how it works. The standard library provides decoders for GIF, PNG, and JPEG, and users may provide others, but to keep executables small, decoders are not included in an application unless explicitly requested. The `image.Decode` function consults a table of supported formats. Each entry in the table specifies four things: the name of the format; a string that is a prefix of all images encoded this way, used to detect the encoding; a function `Decode` that decodes an encoded image; and another function `DecodeConfig` that decodes only the image metadata, such as its size and color space. An entry is added to the table by calling `image.RegisterFormat`, typically from within the package initializer of the supporting package for each format, like this one in `image/png`:
```go
    package png // image/png

    func Decode(r io.Reader) (image.Image, error)
    func DecodeConfig(r io.Reader) (image.Config, error)

    func init() {
        const pngHeader = "\x89PNG\r\n\x1a\n"
        image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
    }
```
The effect is that an application need only blank-import the package for the format it needs to make the `image.Decode` function able to decode it.

The `database/sql` package uses a similar mechanism to let users install just the database drivers they need. For example:
```go
    import (
        "database/mysql"
        _ "github.com/lib/pq"              // enable support for Postgres
        _ "github.com/go-sql-driver/mysql" // enable support for MySQL
    )
  
    db, err = sql.Open("postgres", dbname) // OK
    db, err = sql.Open("mysql", dbname)    // OK
    db, err = sql.Open("sqlite3", dbname)  // returns error: unknown driver "sqlite3"
```

### Exercises
- **Exercise 10.1**: Extend the `jpeg` program so that it converts any supported input format to any output format, using `image.Decode` to detect the input format and a flag to select the output format.
- **Exercise 10.2**: Define a generic archive file-reading function capable of reading ZIP files (`archive/zip`) and POSIX tar files (`archive/tar`). Use a registration mechanism similar to the one described above so that support for each file format can be plugged in using blank imports.


## 10.6. Packages and Naming 

In this section, we'll offer some advice on how to follow Go's distinctive conventions for naming packages and their members.

When creating a package, keep its name short, but not so short as to be cryptic. The most frequently used packages in the standard library are named `bufio`, `bytes`, `flag`, `fmt`, `http`, `io`, `json`, `os`, `sort`, `sync`, and `time`.

Be descriptive and unambiguous where possible. For example, don't name a utility package `util` when a name such as `imageutil` or `ioutil` is specific yet still concise. Avoid choosing package names that are commonly used for related local variables, or you may compel the package's clients to use renaming imports, as with the `path` package.

Package names usually take the singular form. The standard packages `bytes`, `errors`, and `strings` use the plural to avoid hiding the corresponding predeclared types and, in the case of `go/types`, to avoid conflict with a keyword.

Avoid package names that already have other connotations. For example, we originally used the name temp for the temperature conversion package in Section 2.5, but that didn't last long. It was a terrible idea because "temp" is an almost universal synonym for "temporary." We went through a brief period with the name `temperature`, but that was too long and didn't say what the package did. In the end, it became `tempconv`, which is shorter and parallel with `strconv`.

Now let's turn to the naming of package members. Since each reference to a member of another package uses a qualified identifier such as `fmt.Println`, the burden of describing the package member is borne equally by the package name and the member name. We need not mention the concept of formatting in `Println` because the package name `fmt` does that already. When designing a package, consider how the two parts of a qualified identifier work together, not the member name alone. Here are some characteristic examples:
```
    bytes.Equal        flag.Int        http.Get        json.Marshal
```
We can identify some common naming patterns. The `strings` package provides a number of independent functions for manipulating strings:
```go
    package strings

    func Index(needle, haystack string) int

    type Replacer struct{ /* ... */ }
    func NewReplacer(oldnew ...string) *Replacer

    type Reader struct{ /* ... */ }
    func NewReader(s string) *Reader
```
The word `string` does not appear in any of their names. Clients refer to them as `strings.Index`, `strings.Replacer`, and so on.

Other packages that we might describe as *single-type packages*, such as `html/template` and `math/rand`, expose one principal data type plus its methods, and often a New function to create instances.
```go
    package rand // "math/rand"

    type Rand struct{ /* ... */ }
    func New(source Source) *Rand
```
This can lead to repetition, as in `template.Template` or `rand.Rand`, which is why the names of these kinds of packages are often especially short.

At the other extreme, there are packages like `net/http` that have a lot of names without a lot of structure, because they perform a complicated task. Despite having over twenty types and many more functions, the package's most important members have the simplest names: `Get`, `Post`, `Handle`, `Error`, `Client`, `Server`.


## 10.7. The Go Tool

The rest of this chapter concerns the `go` tool, which is used for downloading, querying, formatting, building, testing, and installing packages of Go code.

The go tool combines the features of a diverse set of tools into one command set. It is a package manager (analogous to `apt` or `rpm`) that answers queries about its inventory of packages, computes their dependencies, and downloads them from remote version-control systems. It is a build system that computes file dependencies and invokes compilers, assemblers, and linkers, although it is intentionally less complete than the standard Unix `make`. And it is a test driver, as we will see in Chapter 11.

Its command-line interface uses the "Swiss army knife" style, with over a dozen subcommands, some of which we have already seen, like `get`, `run`, `build`, and `fmt`. You can run `go help` to see the index of its built-in documentation, but for reference, we've listed the most commonly used commands below:
```
    $ go
    Go is a tool for managing Go source code.
    ...
        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

    Use "go help <command>" for more information about a command.
    ...
```
To keep the need for configuration to a minimum, the go tool relies heavily on conventions. For example, given the name of a Go source file, the tool can find its enclosing package, because each directory contains a single package and the import path of a package corresponds to the directory hierarchy in the workspace. Given the import path of a package, the tool can find the corresponding directory in which it stores object files. It can also find the URL of the server that hosts the source code repository.


### 10.7.1. Workspace Organization

The only configuration most users ever need is the `GOPATH` environment variable, which specifies the root of the workspace. When switching to a different workspace, users update the value of `GOPATH`. For instance, we set GOPATH to `$HOME/gobook` while working on this book:
```
    $ export GOPATH=$HOME/gobook
    $ go get gopl.io/...
```
After you download all the programs for this book using the command above, your workspace will contain a hierarchy like this one:
```
    GOPATH/ 
        src/
            gopl.io/
                .git/
                ch1/
                    helloworld/
                        main.go 
                    dup/
                        main.go
                    ...
            golang.org/x/net/
                .git/
                html/
                    parse.go
                    node.go 
                    ...
        bin/
            helloworld
            dup 
        pkg/
            darwin_amd64/
                ...
```
GOPATH has three subdirectories. The src subdirectory holds source code. Each package resides in a directory whose name relative to `$GOPATH/src` is the package's import path, such asgopl.io/ch1/helloworld. Observe that a single `GOPATH` workspace contains multiple version-control repositories beneath `src`, such as `gopl.io` or `golang.org`. The `pkg` subdirectory is where the build tools store compiled packages, and the `bin` subdirectory holds executable programs like `helloworld`.

A second environment variable, `GOROOT`, specifies the root directory of the Go distribution, which provides all the packages of the standard library. The directory structure beneath `GOROOT` resembles that of `GOPATH`, so, for example, the source files of the `fmt` package reside in the `$GOROOT/src/fmt` directory. Users never need to set `GOROOT` since, by default, the go tool will use the location where it was installed.

The `go env` command prints the effective values of the environment variables relevant to the toolchain, including the default values for the missing ones. `GOOS` specifies the target operating system (for example, `android`, `linux`, `darwin`, or `windows`) and `GOARCH` specifies the target processor architecture, such as `amd64`, `386`, or `arm`. Although `GOPATH` is the only variable you must set, the others occasionally appear in our explanations.
```
    $ go env
    GOPATH="/home/gopher/gobook"
    GOROOT="/usr/local/go"
    GOARCH="amd64"
    GOOS="darwin"
    ...
```


### 10.7.2. Downloading Packages

When using the `go` tool, a package's import path indicates not only where to find it in the local workspace, but where to find it on the Internet so that go get can retrieve and update it.

The go get command can download a single package or an entire subtree or repository using the `...` notation, as in the previous section. The tool also computes and downloads all the dependencies of the initial packages, which is why the `golang.org/x/net/html` package appeared in the workspace in the previous example.

Once `go get` has downloaded the packages, it builds them and then *installs* the libraries and commands. We'll look at the details in the next section, but an example will show how straightforward the process is. The first command below gets the `golint` tool, which checks for common style problems in Go source code. The second command runs `golint` on `gopl.io/ch2/popcount` from Section 2.6.2. It helpfully reports that we have forgotten to write a doc comment for the package:
```
    $ go get github.com/golang/lint/golint
    $ $GOPATH/bin/golint gopl.io/ch2/popcount
    src/gopl.io/ch2/popcount/main.go:1:1:
        package comment should be of the form "Package popcount ..."
```
The go get command has support for popular code-hosting sites like GitHub, Bitbucket, and Launchpad and can make the appropriate requests to their version-control systems. For less well-known sites, you may have to indicate which version-control protocol to use in the import path, such as Git or Mercurial. Run `go help` import path for the details.

The directories that `go get` creates are true clients of the remote repository, not just copies of the files, so you can use version-control commands to see a diff of local edits you've made or to update to a different revision. For example, the `golang.org/x/net` directory is a Git client:
```
    $ cd $GOPATH/src/golang.org/x/net
    $ git remote -v
    origin  https://go.googlesource.com/net (fetch)
    origin  https://go.googlesource.com/net (push)
```
Notice that the apparent domain name in the package's import path, `golang.org`, differs from the actual domain name of the Git server, `go.googlesource.com`. This is a feature of the go tool that lets packages use a custom domain name in their import path while being hosted by a generic service such as `googlesource.com` or `github.com`. HTML pages beneath `https://golang.org/x/net/html` include the metadata shown below, which redirects the go tool to the Git repository at the actual hosting site:
```
    $ go build gopl.io/ch1/fetch
    $ ./fetch https://golang.org/x/net/html | grep go-import
    <meta name="go-import" 
          content="golang.org/x/net git https://go.googlesource.com/net">
```
If you specify the `-u` flag, go get will ensure that all packages it visits, including dependencies, are updated to their latest version before being built and installed. Without that flag, packages that already exist locally will not be updated.

The `go get -u` command generally retrieves the latest version of each package, which is convenient when you're getting started but may be inappropriate for deployed projects, where precise control of dependencies is critical for release hygiene. The usual solution to this problem is to *vendor* the code, that is, to make a persistent local copy of all the necessary dependencies, and to update this copy carefully and deliberately. Prior to Go 1.5, this required changing those packages' import paths, so our copy of `golang.org/x/net/html` would become `gopl.io/vendor/golang.org/x/net/html`. More recent versions of the go tool support vendoring directly, though we don't have space to show the details here. See *Vendor Directories* in the output of the `go help gopath` command.

#### Exercises
- **Exercise 10.3**: Using `fetch http://gopl.io/ch1/helloworld?go-get=1`, find out which service hosts the code samples for this book. (HTTP requests from `go get` include the `go-get` parameter so that servers can distinguish them from ordinary browser requests.)


### 10.7.3. Building Packages

The `go build` command compiles each argument package. If the package is a library, the result is discarded; this merely checks that the package is free of compile errors. If the package is named `main`, `go build` invokes the linker to create an executable in the current directory; the name of the executable is taken from the last segment of the package's import path.

Since each directory contains one package, each executable program, or *command* in Unix terminology, requires its own directory. These directories are sometimes children of a directory named `cmd`, such as the `golang.org/x/tools/cmd/godoc` command which serves Go package documentation through a web interface (§10.7.4).

Packages may be specified by their import paths, as we saw above, or by a relative directory name, which must start with a `.` or `..` segment even if this would not ordinarily be required. If no argument is provided, the current directory is assumed. Thus the following commands build the same package, though each writes the executable to the directory in which `go build` is run:
```
    $ cd $GOPATH/src/gopl.io/ch1/helloworld
    $ go build
```
and:
```
    $ cd anywhere
    $ go build gopl.io/ch1/helloworld
```
and:
```
    $ cd $GOPATH
    $ go build ./src/gopl.io/ch1/helloworld
```
but not:
```
    $ cd $GOPATH
    $ go build src/gopl.io/ch1/helloworld
    Error: cannot find package "src/gopl.io/ch1/helloworld".
```
Packages may also be specified as a list of file names, though this tends to be used only for small programs and one-off experiments. If the package name is `main`, the executable name comes from the basename of the first `.go` file.
```
    $ cat quoteargs.go
    package main

    import (
        "fmt"
        "os"
    )

    func main() {
        fmt.Printf("%q\n", os.Args[1:])
    }
    $ go build quoteargs.go
    $ ./quoteargs one "two three" four\ five
    ["one" "two three" "four five"]
```
Particularly for throwaway programs like this one, we want to run the executable as soon as we've built it. The `go run` command combines these twos teps:
```
    $ go run quoteargs.go one "two three" four\ five
    ["one" "two three" "four five"]
```
The first argument that doesn't end in `.go` is assumed to be the beginning of the list of arguments to the Go executable.

By default, the `go build` command builds the requested package and all its dependencies, then throws away all the compiled code except the final executable, if any. Both the dependency analysis and the compilation are surprisingly fast, but as projects grow to dozens of packages and hundreds of thousands of lines of code, the time to recompile dependencies can become noticeable, potentially several seconds, even when those dependencies haven't changed at all.

The `go install` command is very similar to go build, except that it saves the compiled code for each package and command instead of throwing it away. Compiled packages are saved beneath the `$GOPATH/pkg` directory corresponding to the src directory in which the source resides, and command executables are saved in the `$GOPATH/bin` directory. (Many users put `$GOPATH/bin` on their executable search path.) Thereafter, `go build` and `go install` do not run the compiler for those packages and commands if they have not changed, making subsequent builds much faster. For convenience, `go build -i` installs the packages that are dependencies of the build target.

Since compiled packages vary by platform and architecture, `go install` saves them beneath a subdirectory whose name incorporates the values of the `GOOS` and `GOARCH` environment variables. For example, on a Mac the `golang.org/x/net/html` package is compiled and installed in the file `golang.org/x/net/html.a` under `$GOPATH/pkg/darwin_amd64`.

It is straightforward to *cross-compile* a Go program, that is, to build an executable intended for a different operating system or CPU. Just set the `GOOS` or `GOARCH` variables during the build. The `cross` program prints the operating system and architecture for which it was built:
```go
// gopl.io/ch10/cross
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
```
The following commands produce 64-bit and 32-bit executables respectively:
```
    $ go build gopl.io/ch10/cross
    $ ./cross
    darwin amd64
    $ GOARCH=386 go build gopl.io/ch10/cross
    $ ./cross
    darwin 386
```
Some packages may need to compile different versions of the code for certain platforms or processors, to deal with low-level portability issues or to provide optimized versions of important routines, for instance. If a file name includes an operating system or processor architecture name like `net_linux.go` or `asm_amd64.s`, then the `go` tool will compile the file only when building for that target. Special comments called *build tags* give more fine-grained control. For example, if a file contains this comment:
```go
    // +build linux darwin
```
before the package declaration (and its doc comment), `go build` will compile it only when building for Linux or Mac OS X, and this comment says never to compile the file:
```go
    // +build ignore
```
For more details, see the *Build Constraints* section of the `go/build` package's documentation:
```
    $ go doc go/build
```


### 10.7.4. Documenting Packages

Go style strongly encourages good documentation of package APIs. Each declaration of an exported package member and the package declaration itself should be immediately preceded by a comment explaining its purpose and usage.

Go *doc comments* are always complete sentences, and the first sentence is usually a summary that starts with the name being declared. Function parameters and other identifiers are mentioned without quotation or markup. For example, here's the doc comment for `fmt.Fprintf`:
```go
    // Fprintf formats according to a format specifier and writes to w.
    // It returns the number of bytes written and any write error encountered.
    func Fprintf(w io.Writer, format string, a ...interface{}) (int, error)
```
The details of `Fprintf`'s formatting are explained in a doc comment associated with the `fmt` package itself. A comment immediately preceding a `package` declaration is considered the doc comment for the package as a whole. There must be only one, though it may appear in any file. Longer package comments may warrant a file of their own; `fmt`'s is over 300 lines. This file is usually called `doc.go`.

Good documentation need not be extensive, and documentation is no substitute for simplicity. Indeed, Go's conventions favor brevity and simplicity in documentation as in all things, since documentation, like code, requires maintenance too. Many declarations can be explained in one well-worded sentence, and if the behavior is truly obvious, no comment is needed.

Throughout the book, as space permits, we've preceded many declarations by doc comments, but you will find better examples as you browse the standard library. Two tools can help you do that.

The `go doc` tool prints the declaration and doc comment of the entity specified on the command line, which may be a package:
```
    $ go doc time
    package time // import "time"
    Package time provides functionality for measuring and displaying time.
    const Nanosecond Duration = 1 ...
    func After(d Duration) <-chan Time
    func Sleep(d Duration)
    func Since(t Time) Duration
    func Now() Time
    type Duration int64
    type Time struct { ... } ...many more...
```
or a package member:
```
    $ go doc time.Since
    func Since(t Time) Duration
        Since returns the time elapsed since t.
        It is shorthand for time.Now().Sub(t).
```
or a method:
```
    $ go doc time.Duration.Seconds
    func (d Duration) Seconds() float64
        Seconds returns the duration as a floating-point number of seconds.
```
The tool does not need complete import paths or correct identifier case. This command prints the documentation of `(*json.Decoder).Decode` from the `encoding/json` package:
```
  $ go doc json.decode
  func (dec *Decoder) Decode(v interface{}) error
      Decode reads the next JSON-encoded value from its input and stores
      it in the value pointed to by v.
```
The second tool, confusingly named `godoc`, serves cross-linked HTML pages that provide the same information as `go doc` and muc hmore. The `godoc` server at `https://golang.org/pkg` covers the standard library. Figure 10.1 shows the documentation for the `time` package, and in Section 11.6 we'll see `godoc`'s interactive display of example programs. The `godoc` server at `https://godoc.org` has a searchable index of thousands of open-source packages.

You can also run an instance of `godoc` in your workspace if you want to browse your own packages. Visit `http://localhost:8000/pkg` in your browser while running this command:
```
    $ godoc -http :8000
```
Its `-analysis=type` and `-analysis=pointer` flags augment the documentation and the source code with the results of advanced static analysis.


### 10.7.5. Internal Packages

The package is the most important mechanism for encapsulation in Go programs. Unexported identifiers are visible only within the same package, and exported identifiers are visible to the world.

Sometimes, though, a middle ground would be helpful, a way to define identifiers that are visible to a small set of trusted packages, but not to everyone. For example, when we're breaking up a large package into more manageable parts, we may not want to reveal the interfaces between those parts to other packages. Or we may want to share utility functions across several packages of a project without exposing them more widely. Or perhaps we just want to experiment with a new package without prematurely committing to its API, by putting it "on probation" with a limited set of clients.

![Figure 10.1](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig10.1.png)

To address these needs, the go build tool treats a package specially if its import path contains a path segment named `internal`. Such packages are called *internal packages*. An internal package may be imported only by another package that is inside the tree rooted at the parent of the internal directory. For example, given the packages below, `net/http/internal/chunked` can be imported from `net/http/httputil` or `net/http`, but not from `net/url`. However, `net/url` may import `net/http/httputil`.
```
    net/http
    net/http/internal/chunked
    net/http/httputil
    net/url
```


### 10.7.6. Querying Packages

The `go list` tool reports in formation about available packages. In its simplest form, `go list` tests whether a package is present in the workspace and prints its import path if so:
```
    $ go list github.com/go-sql-driver/mysql
    github.com/go-sql-driver/mysql
```
An argument to `go list` may contain the `"..."` wildcard, which matches any substring of a package's import path. We can use it to enumerate all the packages within a Go workspace:
```
    $ go list ...
    archive/tar
    archive/zip
    bufio
    bytes 
    cmd/addr2line 
    cmd/api 
    ...many more...
```
or within a specific subtree:
```
    $ go list gopl.io/ch3/...
    gopl.io/ch3/basename1
    gopl.io/ch3/basename2
    gopl.io/ch3/comma
    gopl.io/ch3/mandelbrot
    gopl.io/ch3/netflag
    gopl.io/ch3/printints
    gopl.io/ch3/surface
```
or related to a particular topic:
```
    $ go list ...xml...
    encoding/xml
    gopl.io/ch7/xmlselect
```
The `go list` command obtains the complete metadata for each package, not just the import path, and makes this information available to users or other tools in a variety of formats. The `-json` flag causes `go list` to print the entire record of each package in JSON format:
```
    $ go list -json hash
    {
        "Dir": "/home/gopher/go/src/hash",
        "ImportPath": "hash",
        "Name": "hash",
        "Doc": "Package hash provides interfaces for hash functions.",
        "Target": "/home/gopher/go/pkg/darwin_amd64/hash.a",
        "Goroot": true,
        "Standard": true,
        "Root": "/home/gopher/go",
        "GoFiles": [
            "hash.go"
        ],
        "Imports": [
          "io" 
        ],
        "Deps": [
            "errors",
            "io",
            "runtime",
            "sync",
            "sync/atomic",
            "unsafe"
        ] 
    }
```
The `-f` flag lets users customize the output format using the template language of package `text/template` (§4.6). This command prints the transitive dependencies of the `strconv` package, separated by spaces:
```
    $ go list -f '{{join .Deps " "}}' strconv
    errors math runtime unicode/utf8 unsafe
```
and this command prints the direct imports of each package in the `compress` subtree of the standard library:
```
    $ go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
    compress/bzip2 -> bufio io sort
    compress/flate -> bufio fmt io math sort strconv
    compress/gzip -> bufio compress/flate errors fmt hash hash/crc32 io time
    compress/lzw -> bufio errors fmt io
    compress/zlib -> bufio compress/flate errors fmt hash hash/adler32 io
```
The `go list` command is useful for both one-off interactive queries and for build and test automation scripts. We'll use it again in Section 11.2.4. For more information, including the set of available fields and their meaning, see the output of `go help list`.

In this chapter, we've explained all the important subcommands of the go tool — except one. In the next chapter, we'll see how the `go test` command is used for testing Go programs.

#### Exercises
- **Exercise 10.4**: Construct a tool that reports the set of all packages in the workspace that transitively depend on the packages specified by the arguments. Hint: you will need to run `go list twice`, once for the initial packages and once for all packages. You may want to parse its JSON output using the `encoding/json` package (§4.5).
