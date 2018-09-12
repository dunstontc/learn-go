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



### 10.7.2. Downloading Packages
### 10.7.3. Building Packages
### 10.7.4. Documenting Packages
### 10.7.5. Internal Packages
### 10.7.6. Querying Packages
