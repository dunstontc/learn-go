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
- **Exercise 1.11**: Try `fetchall` with longer argument lists, such as samples from the top million web sites available at `alexa.com`. How does the program behave if a web site just doesn’t respond? (Section 8.9 describes mechanisms for coping in such cases.)
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
- **Exercise 3.7**: Another simple fractal uses Newton’s method to find complex solutions to a function such as $z^4−1 = 0$. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.
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
