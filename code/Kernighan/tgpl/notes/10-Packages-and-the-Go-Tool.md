# Chapter 10: Packages and the Go Tool

<!-- TOC -->

- [10.1. Introduction](#101-introduction)
- [10.2. Import Paths](#102-import-paths)
- [10.3. The Package Declaration](#103-the-package-declaration)
- [10.4. Import Declarations](#104-import-declarations)
- [10.5. Blank Imports](#105-blank-imports)
- [10.6. Packages and Naming](#106-packages-and-naming)
- [10.7. The Go Tool](#107-the-go-tool)

<!-- /TOC -->

A modest-size program today might contain 10,000 functions. Yet its author need think about only a few of them and design even fewer, because the vast majority were written by others and made available for reuse through `packages`.

Go comes with over 100 standard packages that provide the foundations for most applications. The Go community, a thriving ecosystem of package design, sharing, reuse, and improvement, has published many more, and you can find a searchable index of them at `http://godoc.org`. In this chapter, we’ll show how to use existing packages and create new ones.

Go also comes with the go tool, a sophisticated but simple-to-use command for managing workspaces of Go packages. Since the beginning of the book, we’ve been showing how to use the go tool to download, build, and run example programs. In this chapter, we’ll look at the tool’s underlying concepts and tour more of its capabilities, which include printing documentation and querying metadata about the packages in the workspace. In the next chapter we’ll explore its testing features.

## 10.1. Introduction
## 10.2. Import Paths
## 10.3. The Package Declaration 
## 10.4. Import Declarations 
## 10.5. Blank Imports
## 10.6. Packages and Naming 
## 10.7. The Go Tool
### 10.7.1. Workspace Organization
### 10.7.2. Downloading Packages
### 10.7.3. Building Packages
