# Exercises

## Chapter 1

- **Exercise 1.1**: Modify the `echo` program to also print `os.Args[0]`, the name of the command that invoked it.
- **Exercise 1.2**: Modify the `echo` program to print the index and value of each of its arguments, one per line.
- **Exercise 1.3**: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses `strings.Join`. (Section 1.6 illustrates part of the `time` package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
- **Exercise 1.4**: Modify dup2 to print the names of all files in which each duplicated line occurs.
- **Exercise 1.5**: Change the Lissajous program's color palette to green on black, for added authenticity. To create the web color `#RRGGBB`, use `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`, where each pair of hexadecimal digits represents the intensity of the red, green, or blue component of the pixel.
- **Exercise 1.6**: Modify the Lissajous program to produce images in multiple colors by adding more values to `palette` and then displaying them by changing the third argument of `SetColorIndex` in some interesting way.
- **Exercise 1.7**: The function call `io.Copy(dst,src)` reads from `src` and writes to `dst`. Use it instead of `ioutil.ReadAll` to copy the response body to `os.Stdout` without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of `io.Copy`.
- **Exercise 1.8**: Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix.
- **Exercise 1.9**: Modify fetch to also print the HTTP status code, found in resp.Status.
- **Exercise 1.10**: Find a web site that produces a large amount of data. Investigate caching by running `fetchall` twice in succession to see whether the reported time changes much. Do you get the same content each time? Modify `fetchall` to print its output to a file so it can be examined.
- **Exercise 1.11**: Try `fetchall` with longer argument lists, such as samples from the top million web sites available at `alexa.com`. How does the program behave if a web site just doesn't respond? (Section 8.9 describes mechanisms for coping in such cases.)
- **Exercise 1.12**: Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like `http://localhost:8000/?cycles=20` sets the number of cycles to 20 instead of the default 5. Use the `strconv.Atoi` function to convert the string parameter into an integer. You can see its documentation with go doc `strconv.Atoi`.

## Chapter 2

- **Exercise 2.1**: Add types, constants, and functions to `tempconv` for processing temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.
- **Exercise 2.2**: Write a general-purpose unit-conversion program analogous to cf that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.
- **Exercise 2.3**: Rewrite `PopCount` to use a loop instead of a single expression. Compare the performance of the two versions. (Section 11.4 shows how to compare the performance of different implementations systematically.)
- **Exercise 2.4**: Write a version of `PopCount` that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time. Compare its performance to the tablelookup version.
- **Exercise 2.5**: The expression `x&(x-1)` clears the rightmost non-zero bit of `x`. Write a version of `PopCount` that counts bits by using this fact, and assess its performance.

## Chapter 3

- **Exercise 3.1**: If the function `f` returns a non-finite `float64` value, the SVG file will contain invalid `<polygon>` elements (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.
- **Exercise 3.2**: Experiment with visualizations of other functions from the `math` package. Can you produce an egg box, moguls, or a saddle?
- **Exercise 3.3**: Color each polygon based on its height, so that the peaks are colored red (`#ff0000`) and the valleys blue (`#0000ff`).
- **Exercise 3.4**: Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the `Content-Type` header like this:
```go
  w.Header().Set("Content-Type", "image/svg+xml")
  // (This step was not required in the Lissajous example because the server uses standard heuristics to recognize common formats like PNG from the first 512 bytes of the response and generates the proper header.) Allow the client to specify values like height, width, and color as HTTP request parameters.
```
- **Exercise 3.5**: Implement a full-color Mandelbrot set using the function `image.NewRGBA` and the type `color.RGBA` or `color.YCbCr`.
- **Exercise 3.6**: Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels." Implement it.
- **Exercise 3.7**: Another simple fractal uses Newton's method to find complex solutions to a function such as $z^4−1 = 0$. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.
- **Exercise 3.8**: Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat`. (The latter two types are found in the `math/big` package. `Float` uses arbitrary but bounded-precision floating-point; `Rat` uses unbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?
- **Exercise 3.9**: Write a web server that renders fractals and writes the image data to the client. Allow the client to specify the *x*, *y*, and zoom values as parameters to the HTTP request.
- **Exercise 3.10**: Write a non-recursive version of `comma`, using `bytes.Buffer` instead of string concatenation.
- **Exercise 3.11**: Enhance `comma` so that it deals correctly with floating-point numbers and an optional sign.
- **Exercise 3.12**: Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.
- **Exercise 3.13**: Write `const` declarations for KB, MB, up through YB as compactly as you can.

## Chapter 4

- **Exercise 4.1**: Write a function that counts the number of bits that are different in two SHA256 hashes. (See `PopCount` from Section2.6.2.)
- **Exercise 4.2**: Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.
- **Exercise 4.3**: Rewrite `reverse` to use an array pointer instead of a slice.
- **Exercise 4.4**: Write a version of rotate that operates in a single pass.
- **Exercise 4.5**: Write an in-place function to eliminate adjacent duplicates in a `[]string` slice.
- **Exercise 4.6**: Write an in-place function that squashes each run of adjacent Unicode spaces (see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.
- **Exercise 4.7**: Modify reverse to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?
- **Exercise 4.8**: Modify `charcount` to count letters, digits, and so on in their Unicode categories, using functions like `unicode.IsLetter`.
- **Exercise 4.9**: Write a program `wordfreq` to report the frequency of each word in an input text file. `Callinput.Split(bufio.ScanWords)` before the first call to `Scan` to break the input into words instead of lines.
- **Exercise 4.10**: Modify `issues` to report the results in age categories, say less than a month old, less than a year old, and more than a year old.
- **Exercise 4.11**: Build a tool that lets users create, read, update, and delete GitHub issues from the command line, invoking their preferred text editor when substantial text input is required.
- **Exercise 4.12**: The popular web comic *xkcd* has a JSON interface. For example, a request to `https://xkcd.com/571/info.0.json` produces a detailed description of comic 571, one of many favorites. Download each URL (once!) and build an offline index. Write a tool `xkcd` that, using this index, prints the URL and transcript of each comic that matches a search term provided on the command line.
- **Exercise 4.13**: The JSON-based web service of the Open Movie Database lets you search `https://omdbapi.com/` for a movie by name and download its poster image. Write a tool `poster` that downloads the poster image for the movie named on the command line.
- **Exercise 4.14**: Create a web server that queries GitHub once and then allows navigation of the list of bug reports, milestones, and users.

## Chapter 5
- **Exercise 5.1**: Change the `findlinks` program to traverse the `n.FirstChild` linked list using recursive calls to `visit` instead of a loop.
- **Exercise 5.2**: Write a function to populate a mapping from element names (`p`, `div`, `span`, and so on) to the number of elements with that name in an HTML document tree.
- **Exercise 5.3**: Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into `<script>` or `<style>` elements, since their contents are not visible in a web browser.
- **Exercise 5.4**: Extend the `visit` function so that it extracts other kinds of links from the document, such as images, scripts, and style sheets.
Use short forms like `<img/>` instead of `<img></img>` when an element has no children. Write a test to ensure that the output can be parsed successfully. (See Chapter 11.)
- **Exercise 5.8**: Modify `forEachNode` so that the `pre` and `post` functions return a boolean result indicating whether to continue the traversal. Use it to write a function `ElementByID` with the following signature that finds the first HTML element with the specified `id` attribute. The function should stop the traversal as soon as a match is found.
```go
  func ElementByID(doc *html.Node, id string) *html.Node
```
- **Exercise 5.9**: Write a function `expand(s string, f func(string) string) string` that
replaces each substring `"$foo"` within `s` by the text returned by `f("foo")`.
- **Exercise 5.10**: Rewrite `topoSort` to use maps instead of slices and eliminate the initial sort. Verify that the results, though nondeterministic, are valid topological orderings.
- **Exercise 5.11**: The instructor of the linear algebra course decides that calculus is now a prerequisite. Extend the `topoSort` function to report cycles.
- **Exercise 5.12**: The `startElement` and `endElement` functions in `gopl.io/ch5/outline2` (§5.5) share a global variable, `depth`. Turn them into anonymous functions that share a variable local to the `outline` function.
- **Exercise 5.13**: Modify `crawl` to make local copies of the pages it finds, creating directories as necessary. Don't make copies of pages that come from a different domain. For example, if the original page comes from `golang.org`, save all files from there, but exclude ones from `vimeo.com`.
- **Exercise 5.14**: Use the `breadthFirst` function to explore a different structure. For example, you could use the course dependencies from the `topoSort` example (a directed graph), the file system hierarchy on your computer (a tree), or a list of bus or subway routes downloaded from your city government's web site (an undirected graph).
- **Exercise 5.15**: Write variadic functions `max` and `min`, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument.
- **Exercise 5.16**: Write a variadic version of `strings.Join`.
- **Exercise 5.17**: Write a variadic function `ElementsByTagName` that, given an HTML node tree and zero or more names, returns all the elements that match one of those names. Here are two example calls:
```go
  func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

  images := ElementsByTagName(doc, "img")
  headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
```
- **Exercise 5.18**: Without changing its behavior, rewrite the `fetch` function to use `defer` to close the writable file.
- **Exercise 5.19**: Use `panic` and `recover` to write a function that contains no `return` statement yet returns a non-zero value.

## Chapter 6

- **Exercise 6.1**: Implement these additional methods:
```go
  func (*IntSet) Len() int      // return the number of elements
  func (*IntSet) Remove(x int)  // remove x from the set
  func (*IntSet) Clear()        // remove all elements from the set
  func (*IntSet) Copy() *IntSet // return a copy of the set
```
- **Exercise 6.2**: Define a variadic (*IntSet).AddAll(...int) method that allows a list of values to be added, such as `s.AddAll(1, 2, 3)`.
- **Exercise 6.3**: `(*IntSet).UnionWith` computes the union of two sets using |, the word-parallel bitwise OR operator. Implement methods for `IntersectWith`, `DifferenceWith`, and `SymmetricDifference` for the corresponding set operations. (The symmetric difference of two sets contains the elements present in one set or the other but not both.)
- **Exercise 6.4**: Add a method `Elems` that returns a slice containing the elements of the set, sui
able for iterating over with a `range` loop.
- **Exercise 6.5**: The type of each word used by `IntSet` is `uint64`, but 64-bit arithmetic may be inefficient on a 32-bit platform. Modify the program to use the uint type, which is the most efficient unsigned integer type for the platform. Instead of dividing by 64, define a constant holding the effective size of `uint` in bits, 32 or 64. You can use the perhaps too-clever expression `32 << (^uint(0) >> 63)` for this purpose.

## Chapter 7

- **Exercise 7.1**: Using the ideas from `ByteCounter`, implement counters for words and for lines. You will find `bufio.ScanWords` useful.
- **Exercise 7.2**: Write a function `CountingWriter` with the signature below that, given an `io.Writer`, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new `Writer`.
```go
  func CountingWriter(w io.Writer) (io.Writer, *int64)
```
- **Exercise 7.3**: Write a `String` method for the `*tree` type in `gopl.io/ch4/treesort` (§4.4) that reveals the sequence of values in the tree.
- **Exercise 7.4**: The `strings.NewReader` function returns a value that satisfies the `io.Reader` interface (and others) by reading from its argument, a string. Implement a simple version of `NewReader` yourself, and use it to make the HTML parser (§5.2) take input from a string.
- **Exercise 7.5**: The LimitReader function in the io package accepts an `io.Reader` `r` and `a` number of bytes `n`, and returns another `Reader` that reads from `r` but reports an end-of-file condition after `n` bytes. Implement it.
```go
  func LimitReader(r io.Reader, n int64) io.Reader
```
- **Exercise 7.6**: Add support for Kelvin temperatures to `tempflag`.
- **Exercise 7.7**: Explain why the help message contains `°C` when the default value of `20.0` does not.
- **Exercise 7.8**: Many GUIs provide a table widget with a stateful multi-tier sort: the primary sort key is the most recently clicked column head, the secondary sort key is the second-most recently clicked column head, and so on. Define an implementation of `sort.Interface` for use by such a table. Compare that approach with repeated sorting using `sort.Stable`.
- **Exercise 7.9**: Use the `html/template` package (§4.6) to replace `printTracks` with a function that displays the `tracks` as an HTML table. Use the solution to the previous exercise to arrange that each click on a column head makes an HTTP request to sort the table.
- **Exercise 7.10**: The `sort.Interface` type can be adapted to other uses. Write a function `IsPalindrome(s sort.Interface) bool` that reports whether the sequence `s` is a palindrome, in other words, reversing the sequence would not change it. Assume that the elements at indices `i` and `j` are equal if `!s.Less(i, j) && !s.Less(j, i)`.
- **Exercise 7.11**: Add additional handlers so that clients can create, read, update, and delete database entries. For example, a request of the form `/update?item=socks&price=6` will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid. (Warning: this change introduces concurrent variable updates.)
- **Exercise 7.12**: Change the handler for `/list` to print its output as an HTML table, not text. You may find the `html/template` package (§4.6) useful.
- **Exercise 7.13**: Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the results, when parsed again, yield an equivalent tree.
- **Exercise 7.14**: Define a new concrete type that satisfies the `Expr` interface and provides a new operation such as computing the minimum value of its operands. Since the `Parse` function does not create instances of this new type, to use it you will need to construct a syntax tree directly (or extend the parser).
- **Exercise 7.15**: Write a program that reads a single expression from the standard input, prompts the user to provide values for any variables, then evaluates the expression in the resulting environment. Handle all errors gracefully.
- **Exercise 7.16**: Write a web-based calculator program.
- **Exercise 7.17**: Extend `xmlselect` so that elements may be selected not just by name, but by their attributes too, in the manner of CSS, so that, for instance, an element like `<div id="page" class="wide">` could be selected by a matching `id` or `class` as well as its name.
- **Exercise 7.18**: Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it. Nodes are of two kinds: `CharData` nodes represent text strings, and `Element` nodes represent named elements and their attributes. Each element node has a slice of child nodes.

You may find the following declarations helpful.
```go
    import "encoding/xml"

    type Node interface{} // CharData or *Element
     
    type CharData string

    type Element struct {
        Type     xml.Name
        Attr     []xml.Attr
        Children []Node
    }
```

## Chapter 8

- **Exercise 8.8**: Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any client that shouts nothing within 10 seconds.
- **Exercise 8.9**: Write a version of `du` that computes and periodically displays separate totals for each of the root directories.
- **Exercise 8.10**: HTTP requests may be cancelled by closing the optional Cancel channel in the http.Request struct. Modify the web crawler of Section 8.6 to support cancellation.
  - Hint: the `http.Get` convenience function does not give you an opportunity to customize a `Request`. Instead, create the request using `http.NewRequest`, set its `Cancel` field, then perform the request by calling `http.DefaultClient.Do(req)`.
- **Exercise 8.11**: Following the approach of `mirroredQuery` in Section 8.4.4, implement a variant of `fetch` that requests several URLs concurrently. As soon as the first response arrives, cancel the other requests.
- **Exercise 8.12**: Make the broadcaster announce the current set of clients to each new arrival. This requires that the `clients` set and the `entering` and `leaving` channels record the client name too.
- **Exercise 8.13**: Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes. Hint: calling `conn.Close()` in another goroutine unblocks active Read calls such as the one done by `input.Scan()`.
- **Exercise 8.14**: Change the chat server's network protocol so that each client provides its name on entering. Use that name instead of the network address when prefixing each message with its sender's identity.
- **Exercise 8.15**: Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it. Alternatively, add buffering to each client's outgoing message channel so that most messages are not dropped; the broadcaster should use a non-blocking send to this channel.

## Chapter 9


## Chapter 10
- **Exercise 10.1**: Extend the `jpeg` program so that it converts any supported input format to any output format, using `image.Decode` to detect the input format and a flag to select the output format.
- **Exercise 10.2**: Define a generic archive file-reading function capable of reading ZIP files (`archive/zip`) and POSIX tar files (`archive/tar`). Use a registration mechanism similar to the one described above so that support for each file format can be plugged in using blank imports.


## Chapter 11


## Chapter 12
- **Exercise 12.1**: Extend `Display` so that it can display maps whose keys are structs or arrays.
- **Exercise 12.2**: Make `display` safe to use on cyclic data structures by bounding the number of steps it takes before abandoning the recursion. (In Section 13.3, we'll see another way to detect cycles.)
- **Exercise 12.3**: Implement the missing cases of the encode function. Encode booleans as t and nil, floating-point numbers using Go's notation, and complex numbers like `1+2i` as `#C(1.02.0)`. Interfaces can be encoded as a pair of a type name and a value, for instance `("[]int"(123))`, but beware that this notation is ambiguous: the `reflect.Type.String` method may return the same string for different types.
- **Exercise 12.4**: Modify encode to pretty-print the S-expression in the style shown above. 
- **Exercise 12.5**: Adapt encode to emit JSON instead of S-expressions. Test your encoder using the standard decoder, `json.Unmarshal`.
- **Exercise 12.6**: Adapt encode so that, as an optimization, it does not encode a field whose value is the zero value of its type.
- **Exercise 12.7**: Create a streaming API for the S-expression decoder, following the style of `json.Decoder` (§4.5).
- **Exercise 12.8**: The `sexpr.Unmarshal` function, like `json.Marshal`, requires the complete input in a byte slice before it can begin decoding. Define a `sexpr.Decoder` type that, like `json.Decoder`, allows a sequence of values to be decoded from an `io.Reader`. Change `sexpr.Unmarshal` to use this new type.
- **Exercise 12.9**: Write a token-based API for decoding S-expressions, following the style of `xml.Decoder` (§7.14). You will need five types of tokens: `Symbol`, `String`, `Int`, `StartList`, and `EndList`.
- **Exercise 12.10**: Extend `sexpr.Unmarshal` to handle the booleans, floating-point numbers, and interfaces encoded by your solution to Exercise 12.3. (Hint: to decode interfaces, you will need a mapping from the name of each supported type to its `reflect.Type`.)
- **Exercise 12.11**: Write the corresponding `Pack` function. Given a struct value, `Pack` should return a URL incorporating the parameter values from the struct.
- **Exercise 12.12**: Extend the field tag notation to express parameter validity requirements. For example, a string might need to be a valid email address or credit-card number, and an integer might need to be a valid US ZIP code. Modify `Unpack` to check these requirements.
- **Exercise 12.13**: Modify the S-expression encoder (§12.4) and decoder (§12.6) so that they honor the `sexpr:"..."` field tag in a similar manner to `encoding/json` (§4.5).
