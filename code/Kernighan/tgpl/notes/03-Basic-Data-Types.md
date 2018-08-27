# Chapter 3: Basic Data Types

<!-- TOC -->

- [3.1. Integers](#31-integers)
- [3.2. Floating-Point Numbers](#32-floating-point-numbers)
- [3.3. Complex Numbers](#33-complex-numbers)
- [3.4. Booleans](#34-booleans)
- [3.5. Strings](#35-strings)
  - [3.5.1 String Literals](#351-string-literals)
  - [3.5.2 Unicode](#352-unicode)
  - [3.5.3 UTF-8](#353-utf-8)
  - [3.5.4 Strings and Byte Slices](#354-strings-and-byte-slices)
  - [3.5.5 Conversions between Strings and Numbers](#355-conversions-between-strings-and-numbers)
- [3.6. Constants](#36-constants)
  - [3.6.1 The Constant Generator `iota`](#361-the-constant-generator-iota)
  - [3.6.2 Untyped Constants](#362-untyped-constants)

<!-- /TOC -->


It’s all bits at the bottom, of course, but computers operate fundamentally on fixed-size numbers called *words*, which are interpreted as integers, floating-point numbers, bit sets, or memory addresses, then combined into larger aggregates that represent packets, pixels, portfolios, poetry, and everything else. Go offers a variety of ways to organize data, with a spectrum of data types that at one end match the features of the hardware and at the other end provide what programmers need to conveniently represent complicated data structures.

Go’s types fall into four categories: *basic types*, *aggregate types*, *reference types*, and *interface types*. Basic types, the topic of this chapter, include numbers, strings, and booleans. Aggregate types—arrays (§4.1) and structs (§4.4)—form more complicated data types by combining val- ues of several simpler ones. Reference types are a diverse group that includes pointers (§2.3.2), slices (§4.2), maps (§4.3), functions (Chapter 5), and channels (Chapter 8), but what they have in common is that they all refer to program variables or state indirectly, so that the effect of an operation applied to one reference is observed by all copies of that reference. Finally, we’ll talk about interface types in Chapter 7.

## 3.1. Integers 

Go’s numeric data types include several sizes of integers, floating-point numbers, and complex numbers. Each numeric type determines the size and signedness of its values. Let’s begin with integers.

Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of signed integers—8, 16, 32, and 64 bits—represented by the types `int8`, `int16`, `int32`, and `int64`, and corresponding unsigned versions `uint8`, `uint16`, `uint32`, and `uint64`.

There are also two types called just `int` and `uint` that are the natural or most efficient size for signed and unsigned integers on a particular platform; int is by far the most widely used numeric type. Both these types have the same size, either 32 or 64 bits, but one must not make assumptions about which; different compilers may make different choices even on identical hardware.

The type `rune` is an synonym for `int32` and conventionally indicates that a value is a Unicode code point. The two names may be used interchangeably. Similarly, the type `byte` is an synonym for `uint8`, and emphasizes that the value is a piece of raw data rather than a small numeric quantity.

Finally, there is an unsigned integer type `uintptr`, whose width is not specified but is sufficient to hold all the bits of a pointer value. The `uintptr` type is used only for low-level programming, such as at the boundary of a Go program with a C library or an operating system. We’ll see examples of this when we deal with the unsafe package in Chapter 13.

Regardless of their size, `int`, `uint`, and `uintptr` are different types from their explicitly sized siblings. Thus `int` is not the same type as `int32`, even if the natural size of integers is 32 bits, and an explicit conversion is required to use an `int` value where an `int32` is needed, and vice versa.

Signed numbers are represented in *Two's-complement* [*0*](https://www.youtube.com/watch?v=lKTsv6iVxV4&t=188s) [*1*](https://www.reddit.com/r/compsci/comments/26jnqu/im_having_trouble_understanding_twos_compliment/) [*2*](https://stackoverflow.com/a/1049880/7687024)  form, in which the high-order bit is reserved for the sign of the number and the range of values of an *n*-bit number is from $−2^{n−1}$ to $2^{n−1}−1$. Unsigned integers use the full range of bits for non-negative values and thus have the range 0 to $2^n−1$. For instance, the range of `int8` is −128 to 127, whereas the range of `uint8` is 0 to 255.

Go’s binary [operators](https://golang.org/ref/spec#Operators) for arithmetic, logic, and comparison are listed here in order of decreasing precedence:
```
  *    /    %    <<    >>    &    &^    
  +    -    |    ^
  ==   !=   <    <=    >     >=
  &&
  ||
```
There are only five levels of precedence for binary operators. Operators at the same level associate to the left, so parentheses may be required for clarity, or to make the operators evaluate in the intended order in an expression like `mask & (1 << 28)`.

Each operator in the first two lines of the table above, for instance `+`, has a corresponding assignment operator like `+=` that may be used to abbreviate an assignment statement.

The integer arithmetic operators `+,` `-,` `*,` and `/` may be applied to integer, floating-point, and complex numbers, but the remainder operator `%` applies only to integers. The behavior of `%` for negative numbers varies across programming languages. In Go, the sign of the remainder is always the same as the sign of the dividend, so `-5 % 3` and `-5 % -3` are both `-2`. The behavior of `/` depends on whether its operands are integers, so `5.0/4.0` is `1.25`, but `5/4` is `1` because integer division truncates the result toward zero.

If the result of an arithmetic operation, whether signed or unsigned, has more bits than can be represented in the result type, it is said to *overflow*. The high-order bits that do not fit are silently discarded. If the original number is a signed type, the result could be negative if the leftmost bit is a 1, as in the `int8` example here:
```go
  var u uint8 = 255
  fmt.Println(u, u+1, u*u) // "255 0 1"
  var i int8 = 127
  fmt.Println(i, i+1, i*i) // "127 -128 1"
```
Two integers of the same type may be compared using the binary comparison operators below; the type of a comparison expression is a boolean.
```
  ==    equal to
  !=    not equal to
  <     less than
  <=    less than or equal to
  >     greater than
  >=    greater than or equal to
```

In fact, all values of basic type—booleans, numbers, and strings—are *comparable*, meaning that two values of the same type may be compared using the `==` and `!=` operators. Furthermore, integers, floating-point numbers, and strings are *ordered* by the comparison operators. The values of many other types are not comparable, and no other types are ordered. As we encounter each type, we’ll present the rules governing the *comparability* of its values.

There are also unary addition and subtraction operators:
```
  +    unary positive (no effect)
  -    unary negation
```
For integers, `+x` is a shorthand for `0+x` and `-x` is a shorthand for `0-x`; for floating-point and complex numbers, `+x` is just `x` and `-x` is the negation of `x`.

Go also provides the following bitwise binary operators, the first four of which treat their operands as bit patterns with no concept of arithmetic carry or sign:
```
  &    bitwise AND
  |    bitwise OR
  ^    bitwise XOR
  &^   bit clear (AND NOT) << left shift
  >>   right shift
```

The operator `^` is bitwise exclusive OR (XOR) when used as a binary operator, but when used as a unary prefix operator it is bitwise negation or complement; that is, it returns a value with each bit in its operand inverted. The `&^` operator is bit clear (AND NOT): in the expression  `z = x &^ y`, each bit of `z` is `0` if the corresponding bit of `y` is `1`; otherwise it equals the corresponding bit of `x`.

The code below shows how bitwise operations can be used to interpret a `uint8` value as a compact and efficient set of 8 independent bits. It uses `Printf`’s `%b` verb to print a number’s binary digits; `08` modifies `%b` (an adverb!) to pad the result with zeros to exactly 8 digits.
```go
  var x uint8 = 1<<1 | 1<<5
  var y uint8 = 1<<1 | 1<<2
  fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
  fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
  fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
  fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
  fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
  fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
  for i := uint(0); i < 8; i++ {
      if x&(1<<i) != 0 { // membership test
          fmt.Println(i) // "1", "5"
      } 
  }
  fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
  fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
```
(Section 6.5 shows an implementation of integer sets that can be much bigger than a byte.)

In the shift operations `x<<n` and `x>>n`, the `n` operand determines the number of bit positions to shift and must be unsigned; the x operand may be unsigned or signed. Arithmetically, a left shift `x<<n` is equivalent to multiplication by $2^n$ and a right shift `x>>n` is equivalent to the floor of division by $2^n$.

Left shifts fill the vacated bits with zeros, as do right shifts of unsigned numbers, but right shifts of signed numbers fill the vacated bits with copies of the sign bit. For this reason, it is important to use unsigned arithmetic when you’re treating an integer as a bit pattern.

Although Go provides unsigned numbers and arithmetic, we tend to use the signed `int` form even for quantities that can’t be negative, such as the length of an array, though `uint` might seem a more obvious choice. Indeed, the built-in len function returns a signed int, as in this loop which announces prize medals in reverse order:
```go
  medals := []string{"gold", "silver", "bronze"}
  for i := len(medals) - 1; i >= 0; i-- {
      fmt.Println(medals[i]) // "bronze", "silver", "gold"
  }
```

The alternative would be calamitous. If `len` returned an unsigned number, then `i` too would be a `uint`, and the condition `i >= 0` would always be true by definition. After the third iteration, in which `i == 0`, the `i--` statement would cause `i` to become not `−1`, but the maximum uint value (for example, $2^{64}−1$), and the evaluation of `medals[i]` would fail at run time, or *panic* (§5.9), by attempting to access an element outside the bounds of the slice.

For this reason, unsigned numbers tend to be used only when their bitwise operators or peculiar arithmetic operators are required, as when implementing bit sets, parsing binary file formats, or for hashing and cryptography. They are typically not used for merely non-negative quantities.

In general, an explicit conversion is required to convert a value from one type to another, and binary operators for arithmetic and logic (except shifts) must have operands of the same type. Although this occasionally results in longer expressions, it also eliminates a whole class of problems and makes programs easier to understand.

As an example familiar from other contexts, consider this sequence:
```go
  var apples int32 = 1
  var oranges int16 = 2
  var compote int = apples + oranges // compile error
```
Attempting to compile these three declarations produces an error message:
```
  invalid operation: apples + oranges (mismatched types int32 and int16)
```
This type mismatch can be fixed in several ways, most directly by converting everything to a common type:
```go
  var compote = int(apples) + int(oranges)
```

As described in Section 2.5, for every type `T`, the conversion operation `T(x)` converts the value `x` to type `T` if the conversion is allowed. Many integer-to-integer conversions do not entail any change in value; they just tell the compiler how to interpret a value. But a conversion that narrows a big integer into a smaller one, or a conversion from integer to floating-point or vice versa, may change the value or lose precision:
```go
  f := 3.141 // a float64
  i := int(f)
  fmt.Println(f, i)   // "3.141 3"
  f = 1.99
  fmt.Println(int(f)) // "1"
```
Float to integer conversion discards any fractional part, truncating toward zero. You should avoid conversions in which the operand is out of range for the target type, because the behavior depends on the implementation:
```go
  f := 1e100  // a float64
  i := int(f) // result is implementation-dependent
```
Integer literals of any size and type can be written as ordinary decimal numbers, or as octal numbers if they begin with `0`, as in `0666`, or as hexadecimal if they begin with `0x` or `0X`, as in `0xdeadbeef`. Hex digits may be upper or lower case. Nowadays octal numbers seem to be used for exactly one purpose—file permissions on POSIX systems—but hexadecimal numbers are widely used to emphasize the bit pattern of a number over its numeric value.

When printing numbers using the `fmt` package, we can control the radix and format with the `%d`, `%o`, and `%x` verbs, as shown in this example:
```go
  o := 0666
  fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
  x := int64(0xdeadbeef)
  fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
  // Output:
  // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
```

Note the use of two `fmt` tricks. Usually a `Printf` format string containing multiple `%` verbs would require the same number of extra operands, but the `[1]` "adverbs" after `%` tell `Printf` to use the first operand over and over again. Second, the `#` adverb for `%o` or `%x` or `%X` tells `Printf` to emit a `0` or `0x` or `0X` prefix respectively.

Rune literals are written as a character within single quotes. The simplest example is an ASCII character like `'a'`, but it’s possible to write any Unicode code point either directly or with numeric escapes, as we will see shortly.

Runes are printed with `%c`, or with `%q` if quoting is desired:
```go
  ascii := 'a'
  unicode := '⾙'
  newline := '\n'
  fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'" 
  fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 ⾙ '⾙'" 
  fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
```

## 3.2. Floating-Point Numbers 

Go provides two sizes of floating-point numbers, `float32` and `float64`. Their arithmetic properties are governed by the [IEEE 754](https://en.wikipedia.org/wiki/IEEE_754) standard implemented by all modern CPUs.

Values of these numeric types range from tiny to huge. The limits of floating-point values can be found in the [`math`](https://golang.org/pkg/math) package. The constant `math.MaxFloat32`, the largest `float32`, is about `3.4e38`, and `math.MaxFloat64` is about `1.8e308`. The smallest positive values are near `1.4e-45` and `4.9e-324`, respectively.

A `float32` provides approximately six decimal digits of precision, whereas a `float64` provides about 15 digits; `float64` should be preferred for most purposes because `float32` computations accumulate error rapidly unless one is quite careful, and the smallest positive integer that cannot be exactly represented as a `float32` is not large:
```go
  var f float32 = 16777216 // 1 << 24
  fmt.Println(f == f+1)    // "true"!
```
Floating-point numbers can be written literally using decimals, like this:
```go
  const e = 2.71828 // (approximately)
```
Digits may be omitted before the decimal point (`.707`) or after it (`1.`). Very small or very large numbers are better written in scientific notation, with the letter `e` or `E` preceding the decimal exponent:
```go
  const Avogadro = 6.02214129e23
  const Planck   = 6.62606957e-34
```
Floating-point values are conveniently printed with `Printf’s` `%g` verb, which chooses the most compact representation that has adequate precision, but for tables of data, the `%e` (exponent) or `%f` (no exponent) forms may be more appropriate. All three verbs allow field width and numeric precision to be controlled.
```go
  for x := 0; x < 8; x++ {
      fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
  }
```
The code above prints the powers of *e* with three decimal digits of precision, aligned in an eight-character field:
|   `x`   |      $e^x$      |
| ------- | --------------- |
| x = 0 | $e^x$ = 1.000   |
| x = 1 | $e^x$ = 2.718   |
| x = 2 | $e^x$ = 7.389   |
| x = 3 | $e^x$ = 20.086  |
| x = 4 | $e^x$ = 54.598  |
| x = 5 | $e^x$ = 148.413 |
| x = 6 | $e^x$ = 403.429 |
| x = 7 | $e^x$ = 1096.63 |

In addition to a large collection of the usual mathematical functions, the `math` package has functions for creating and detecting the special values defined by IEEE 754: the positive and negative infinities, which represent numbers of excessive magnitude and the result of division by zero; and NaN ("not a number"), the result of such mathematically dubious operations as `0/0` or `Sqrt(-1)`.
```go
  var z float64
  fmt.Println(z, -z, 1/z, -1/z, z/z) //  "0 -0 +Inf -Inf NaN"
```

The function `math.IsNaN` tests whether its argument is a not-a-number value, and `math.NaN` returns such a value. It’s tempting to use NaN as a sentinel value in a numeric computation, but testing whether a specific computational result is equal to NaN is fraught with peril because any comparison with NaN always yields false:
```go
  nan := math.NaN()
  fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
```

If a function that returns a floating-point result might fail, it’s better to report the failure separately, like this:
```go
  func compute() (value float64, ok bool) {
      // ...
      if failed {
          return 0, false
      }
      return result, true
  }
```

The next program illustrates floating-point graphics computation. It plots a function of two variables `z = f(x, y)` as a wire mesh 3-D surface, using Scalable Vector Graphics (SVG), a standard XML notation for line drawings. Figure 3.1 shows an example of its output for the function `sin(r)/r`, where `r` is `sqrt(x*x+y*y)`.
![Figure 3.1](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig3.1.png)
<img src="" alt="">
```go
// tgpl.io/ch3/surface
// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
```

Notice that the function `corner` returns two values, the coordinates of the corner of the cell.

The explanation of how the program works requires only basic geometry, but it’s fine to skip over it, since the point is to illustrate floating-point computation. The essence of the program is mapping between three different coordinate systems, shown in Figure 3.2. The first is a 2-D grid of 100x100 cells identified by integer coordinates (*i*, *j*), starting at (0, 0) in the far back corner. We plot from the back to the front so that background polygons may be obscured by foreground ones.

The second coordinate system is a mesh of 3-D floating-point coordinates (*x*, *y*, *z*), where *x* and *y* are linear functions of *i* and *j*, translated so that the origin is in the center, and scaled by the constant `xyrange`. The height *z* is the value of the surface function *f(x, y)*.

The third coordinate system is the 2-D image canvas, with (0, 0) in the top left corner. Points in this plane are denoted (*sx*, *sy*). We use an isometric projection to map each 3-D point (*x*, *y*, *z*) onto the 2-D canvas. A point appears farther to the right on the canvas the greater its *x* value or the *smaller* its y value. And a point appears farther down the canvas the greater its *x* value or *y* value, and the smaller its *z* value. The vertical and horizontal scale factors for *x* and *y* are derived from the sine and cosine of a 30° angle. The scale factor for *z*, 0.4, is an arbitrary parameter.

For each cell in the 2-D grid, the main function computes the coordinates on the image canvas of the four corners of the polygon *ABCD*, where B corresponds to (*i*, *j*) and *A*, *C*, and *D* are its neighbors, then prints an SVG instruction to draw it.

![Figure 3.2](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig3.2.png)

### Exercises
- **Exercise 3.1**: If the function `f` returns a non-finite `float64` value, the SVG file will contain invalid `<polygon>` elements (although many SVG renderers handle this gracefully). Modify the program to skip invalid polygons.
- **Exercise 3.2**: Experiment with visualizations of other functions from the `math` package. Can you produce an egg box, moguls, or a saddle?
- **Exercise 3.3**: Color each polygon based on its height, so that the peaks are colored red (`#ff0000`) and the valleys blue (`#0000ff`).
- **Exercise 3.4**: Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and writes SVG data to the client. The server must set the `Content-Type` header like this:
```go
  w.Header().Set("Content-Type", "image/svg+xml")
```
(This step was not required in the Lissajous example because the server uses standard heuristics to recognize common formats like PNG from the first 512 bytes of the response and generates the proper header.) Allow the client to specify values like height, width, and color as HTTP request parameters.


## 3.3. Complex Numbers 

Go provides two sizes of complex numbers, `complex64` and `complex128`, whose components are `float32` and `float64` respectively. The built-in function `complex` creates a complex number from its real and imaginary components, and the built-in `real` and `imag` functions extract those components:
```go
  var x complex128 = complex(1, 2) // 1+2i
  var y complex128 = complex(3, 4) // 3+4i
  fmt.Println(x*y)                 // "(-5+10i)"
  fmt.Println(real(x*y))           // "-5"
  fmt.Println(imag(x*y))           // "10"
```

If a floating-point literal or decimal integer literal is immediately followed by `i`, such as `3.141592i` or `2i`, it becomes an *imaginary literal*, denoting a complex number with a zero real component:
```go
  fmt.Println(1i * 1i) // "(-1+0i)", $i^2$ = -1
```

Under the rules for constant arithmetic, complex constants can be added to other constants (integer or floating point, real or imaginary), allowing us to write complex numbers naturally, like `1+2i`, or equivalently, `2i+1`. The declarations of `x` and `y` above can be simplified:
```go
  x := 1 + 2i
  y := 3 + 4i
```

Complex numbers may be compared for equality with `==` and `!=`. Two complex numbers are equal if their real parts are equal and their imaginary parts are equal.

The `math/cmplx` package provides library functions for working with complex numbers, such as the complex square root and exponentiation functions.
```go
  fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"
```

The following program uses `complex128` arithmetic to generate a [Mandelbrot set](https://en.wikipedia.org/wiki/Mandelbrot_set).
```go
// tgpl.io/ch3/mandelbrot
// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
```

The two nested loops iterate over each point in a 1024x1024 grayscale raster image represent- ing the −2 to +2 portion of the complex plane. The program tests whether repeatedly squar- ing and adding the number that point represents eventually "escapes" the circle of radius 2. If so, the point is shaded by the number of iterations it took to escape. If not, the value belongs to the Mandelbrot set, and the point remains black. Finally, the program writes to its standard output the PNG-encoded image of the iconic fractal, shown in Figure 3.3.

### Exercises

- **Exercise 3.5**: Implement a full-color Mandelbrot set using the function `image.NewRGBA` and the type `color.RGBA` or `color.YCbCr`.
- **Exercise 3.6**: Supersampling is a technique to reduce the effect of pixelation by computing the color value at several points within each pixel and taking the average. The simplest method is to divide each pixel into four "subpixels." Implement it.
- **Exercise 3.7**: Another simple fractal uses Newton’s method to find complex solutions to a function such as $z^4−1 = 0$. Shade each starting point by the number of iterations required to get close to one of the four roots. Color each point by the root it approaches.

![Figure 3.3](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/mandelbrot.png)

- **Exercise 3.8**: Rendering fractals at high zoom levels demands great arithmetic precision. Implement the same fractal using four different representations of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat`. (The latter two types are found in the `math/big` package. `Float` uses arbitrary but bounded-precision floating-point; `Rat` uses unbounded-precision rational numbers.) How do they compare in performance and memory usage? At what zoom levels do rendering artifacts become visible?

- **Exercise 3.9**: Write a web server that renders fractals and writes the image data to the client. Allow the client to specify the *x*, *y*, and zoom values as parameters to the HTTP request.

## 3.4. Booleans 

A value of type `bool`, or *boolean*, has only two possible values, `true` and `false`. The conditions in `if` and `for` statements are booleans, and comparison operators like `==` and `<` produce a boolean result. The unary operator `!` is logical negation, so `!true` is `false`, or, one might say, `(!true==false)==true`, although as a matter of style, we always simplify redundant boolean expressions like `x==true` to `x`.

Boolean values can be combined with the `&&` (AND) and `||` (OR) operators, which have *shortcircuit* behavior: if the answer is already determined by the value of the left operand, the right operand is not evaluated, making it safe to write expressions like this:
```go
  s != "" && s[0] == 'x'
```
where `s[0]` would panic if applied to an empty string.

Since `&&` has higher precedence than `||` (mnemonic: `&&` is boolean multiplication, `||` is boolean addition), no parentheses are required for conditions of this form:
```go
  if 'a' <= c && c <= 'z' ||
     'A' <= c && c <= 'Z' ||
     '0' <= c && c <= '9' {
     // ...ASCII letter or digit...
  }
```
There is no implicit conversion from a boolean value to a numeric value like 0 or 1, or vice versa. It’s necessary to use an explicit `if`, as in
```go
  i := 0 if b {
      i=1
  }
```
It might be worth writing a conversion function if this operation were needed often:
```go
  // btoi returns 1 if b is true and 0 if false.
  func btoi(b bool) int {
      if b { 
          return 1
      }
      return 0
  }
```
The inverse operation is so simple that it doesn’t warrant a function, but for symmetry here it is:
```go

  // itob reports whether i is non-zero.
  func itob(i int) bool { return i != 0 }
```


## 3.5. Strings 

A string is an immutable sequence of bytes. Strings may contain arbitrary data, including bytes with value 0, but usually they contain human-readable text. Text strings are conventionally interpreted as UTF-8-encoded sequences of Unicode code points (runes), which we’ll explore in detail very soon.

The built-in len function returns the number of bytes (not runes) in a string, and the *index* operation `s[i]` retrieves the *i*-th byte of string `s`, where `0 ≤ i < len(s)`.
```go
  s := "hello, world"
  fmt.Println(len(s))     // "12"
  fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')
```
Attempting to access a byte outside this range results in a panic:
```go
  c := s[len(s)] // panic: index out of range
```
The *i*-th byte of a string is not necessarily the *i*-th *character* of a string, because the UTF-8 encoding of a non-ASCII code point requires two or more bytes. Working with characters is discussed shortly.

The *substring* operation `s[i:j]` yields a new string consisting of the bytes of the original string starting at index `i` and continuing up to, but not including, the byte at index `j`. The result contains `j-i` bytes.
```go
  fmt.Println(s[0:5]) // "hello"
```
Again, a panic results if either index is out of bounds or if `j` is less than `i`.

Either or both of the `i` and `j` operands may be omitted, in which case the default values of `0` (the start of the string) and `len(s)` (its end) are assumed, respectively.
```go
  fmt.Println(s[:5]) // "hello"
  fmt.Println(s[7:]) // "world"
  fmt.Println(s[:])  // "hello, world"
```

The `+` operator makes a new string by concatenating two strings:
```go
  fmt.Println("goodbye" + s[5:]) // "goodbye, world"
```
Strings may be compared with comparison operators like `==` and `<`; the comparison is done byte by byte, so the result is the natural lexicographic ordering.

String values are immutable: the byte sequence contained in a string value can never be changed, though of course we can assign a new value to a string *variable*. To append one string to another, for instance, we can write
```go
  s := "left foot"
  t := s
  s += ", right foot"
```
This does not modify the string that `s` originally held but causes `s` to hold the new string formed by the `+=` statement; meanwhile, `t` still contains the old string.
```go
  fmt.Println(s) // "left foot, right foot"
  fmt.Println(t) // "left foot"
```
Since strings are immutable, constructions that try to modify a string’s data in place are not allowed:
```go
  s[0] = 'L' // compile error: cannot assign to s[0]
```
Immutability means that it is safe for two copies of a string to share the same underlying memory, making it cheap to copy strings of any length. Similarly, a string `s` and a substring like `s[7:]` may safely share the same data, so the substring operation is also cheap. No new memory is allocated in either case. Figure 3.4 illustrates the arrangement of a string and two of its substrings sharing the same underlying byte array.

![Figure 3.4](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig3.4.png)


### 3.5.1 String Literals 

A string value can be written as a string literal, a sequence of bytes enclosed in double quotes:
```go
  "Hello, 世界"
```

Because Go source files are always encoded in UTF-8 and Go text strings are conventionally interpreted as UTF-8, we can include Unicode code points in string literals.

Within a double-quoted string literal, *escape sequences* that begin with a backslash `\` can be used to insert arbitrary byte values into the string. One set of escapes handles ASCII control codes like newline, carriage return, and tab:

| Escape Sequence |                                                 Result                                       |
| --------------- | -------------------------------------------------------------------------------------------- |
| `\a`            | "alert" or bell                                                                              |
| `\b`            | backspace                                                                                    |
| `\f`            | form feed \n newline                                                                         |
| `\r`            | carriage return                                                                              |
| `\t`            | tab                                                                                          |
| `\v`            | vertical tab                                                                                 |
| `\'`            | single quote (only in the rune literal `'\''`) \" double quote (only within `"..."` literals) |
| `\\`            | backslash                                                                                    |

Arbitrary bytes can also be included in literal strings using hexadecimal or octal escapes. A *hexadecimal escape* is written *\xhh*, with exactly two hexadecimal digits *h* (in upper or lower case). An *octal escape* is written *\ooo* with exactly three octal digits *o* (0 through 7) not exceeding `\377`. Both denote a single byte with the specified value. Later, we’ll see how to encode Unicode code points numerically in string literals.

A *raw string literal* is written `...`, using backquotes instead of double quotes. Within a raw string literal, no escape sequences are processed; the contents are taken literally, including backslashes and newlines, so a raw string literal may spread over several lines in the program source. The only processing is that carriage returns are deleted so that the value of the string is the same on all platforms, including those that conventionally put carriage returns in text files.

Raw string literals are a convenient way to write regular expressions, which tend to have lots of backslashes. They are also useful for HTML templates, JSON literals, command usage messages, and the like, which often extend over multiple lines.
```go
  const GoUsage = `Go is a tool for managing Go source code.
      Usage:
         go command [arguments]
      ...`
```


### 3.5.2 Unicode

Long ago, life was simple and there was, at least in a parochial view, only one character set to deal with: ASCII, the American Standard Code for Information Interchange. ASCII, or more precisely US-ASCII, uses 7 bits to represent 128 ‘‘characters’’: the upper- and lower-case letters of English, digits, and a variety of punctuation and device-control characters. For much of the early days of computing, this was adequate, but it left a very large fraction of the world’s population unable to use their own writing systems in computers. With the growth of the Internet, data in myriad languages has become much more common. How can this rich vari- ety be dealt with at all and, if possible, efficiently?

The answer is [Unicode](https://unicode.org/), which collects all of the characters in all of the world’s writing systems, plus accents and other diacritical marks, control codes like tab and carriage return, and plenty of esoterica, and assigns each one a standard number called a *Unicode code point* or, in Go terminology, a *rune*.

Unicode version 8 defines code points for over 120,000 characters in well over 100 languages and scripts. How are these represented in computer programs and data? The natural data type to hold a single rune is `int32`, and that’s what Go uses; it has the synonym `rune` for precisely this purpose.

We could represent a sequence of runes as a sequence of `int32` values. In this representation, which is called UTF-32 or UCS-4, the encoding of each Unicode code point has the same size, 32 bits. This is simple and uniform, but it uses much more space than necessary since most computer-readable text is in ASCII, which requires only 8 bits or 1 byte per character. All the characters in widespread use still number fewer than 65,536, which would fit in 16 bits. Can we do better?


### 3.5.3 UTF-8

UTF-8 is a variable-length encoding of Unicode code points as bytes. UTF-8 was invented by Ken Thompson and Rob Pike, two of the creators of Go, and is now a Unicode standard. It uses between 1 and 4 bytes to represent each rune, but only 1 byte for ASCII characters, and only 2 or 3 bytes for most runes in common use. The high-order bits of the first byte of the encoding for a rune indicate how many bytes follow. A high-order 0 indicates 7-bit ASCII, where each rune takes only 1 byte, so it is identical to conventional ASCII. A high-order `110` indicates that the rune takes 2 bytes; the second byte begins with `10`. Larger runes have analogous encodings.

|               binary               |     range      |      description      |
| ---------------------------------- | -------------- | --------------------- |
| 0xxxxxx                            | runes 0−127    | (ASCII)               |
| 11xxxxx 10xxxxxx                   | 128−2047       | (values <128 unused)  |
| 110xxxx 10xxxxxx 10xxxxxx          | 2048−65535     | (values <2048 unused) |
| 1110xxx 10xxxxxx 10xxxxxx 10xxxxxx | 65536−0x10ffff | (other values unused) |

A variable-length encoding precludes direct indexing to access the *n*-th character of a string, but UTF-8 has many desirable properties to compensate. The encoding is compact, compatible with ASCII, and self-synchronizing: it’s possible to find the beginning of a character by backing up no more than three bytes. It’s also a prefix code, so it can be decoded from left to right without any ambiguity or lookahead. No rune’s encoding is a substring of any other, or even of a sequence of others, so you can search for a rune by just searching for its bytes, without worrying about the preceding context. The lexicographic byte order equals the Unicode code point order, so sorting UTF-8 works naturally. There are no embedded NUL (zero) bytes, which is convenient for programming languages that use NUL to terminate strings.

Go source files are always encoded in UTF-8, and UTF-8 is the preferred encoding for text strings manipulated by Go programs. The `unicode` package provides functions for working with individual runes (such as distinguishing letters from numbers, or converting an uppercase letter to a lower-case one), and the `unicode/utf8` package provides functions for encoding and decoding runes as bytes using UTF-8.

Many Unicode characters are hard to type on a keyboard or to distinguish visually from similar-looking ones; some are even invisible. Unicode escapes in Go string literals allow us to specify them by their numeric code point value. There are two forms, *\uhhhh* for a 16-bit value and *\Uhhhhhhhh* for a 32-bit value, where each *h* is a hexadecimal digit; the need for the 32-bit form arises very infrequently. Each denotes the UTF-8 encoding of the specified code point. Thus, for example, the following string literals all represent the same six-byte string:
```go
  "世界" 
  "\xe4\xb8\x96\xe7\x95\x8c"
  "\u4e16\u754c" 
  "\U00004e16\U0000754c"
```

The three escape sequences above provide alternative notations for the first string, but the values they denote are identical.

Unicode escapes may also be used in rune literals. These three literals are equivalent:
```go
  '世' 
  '\u4e16' 
  '\U00004e16'
```

A rune whose value is less than 256 may be written with a single hexadecimal escape, such as `'\x41'` for `'A'`, but for higher values, a `\u` or `\U` escape must be used. Consequently, `'\xe4\xb8\x96'` is not a legal rune literal, even though those three bytes are a valid UTF-8 encoding of a single code point.

Thanks to the nice properties of UTF-8, many string operations don’t require decoding. We can test whether one string contains another as a prefix:
```go
  func HasPrefix(s, prefix string) bool {
      return len(s) >= len(prefix) && s[:len(prefix)] == prefix
  }
```
or as a suffix:
```go
  func HasSuffix(s, suffix string) bool {
      return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
  }
```
or as a substring:
```go
  func Contains(s, substr string) bool {
      for i := 0; i < len(s); i++ {
          if HasPrefix(s[i:], substr) {
              return true
          }
      }
      return false
  }
```
using the same logic for UTF-8-encoded text as for raw bytes. This is not true for other encodings. (The functions above are drawn from the `strings` package, though its implementation of `Contains` uses a hashing technique to search more efficiently.)


On the other hand, if we really care about the individual Unicode characters, we have to use other mechanisms. Consider the string from our very first example, which includes two East Asian characters. Figure 3.5 illustrates its representation in memory. The string contains 13 bytes, but interpreted as UTF-8, it encodes only nine code points or runes:
```go
  import "unicode/utf8"
  s := "Hello, 世界"
  fmt.Println(len(s)) // "13" fmt.Println(utf8.RuneCountInString(s)) // "9"
```
To process those characters, we need a UTF-8 decoder. The `unicode/utf8` package provides one that we can use like this:
```go
  for i := 0; i < len(s); {
      r, size := utf8.DecodeRuneInString(s[i:])
      fmt.Printf("%d\t%c\n", i, r)
      i += size
  }
```
Each call to `DecodeRuneInString` returns `r`, the rune itself, and `size`, the number of bytes occupied by the UTF-8 encoding of `r`. The size is used to update the byte index `i` of the next rune in the string. But this is clumsy, and we need loops of this kind all the time. Fortunately, Go’s `range` loop, when applied to a string, performs UTF-8 decoding implicitly. The output of the loop below is also shown in Figure 3.5; notice how the index jumps by more than 1 for each non-ASCII rune.

![Figure 3.5](https://raw.githubusercontent.com/dunstontc/learn-go/master/code/Kernighan/tgpl/assets/fig3.5.png)

```go
  for i, r := range "Hello, BF" {
    fmt.Printf("%d\t%q\t%d\n", i, r, r)
  }
```
We could use a simple `range` loop to count the number of runes in a string, like this:
```go
  n := 0
  for _, _ = range s {
      n++ 
  }
```
As with the other forms of range loop, we can omit the variables we don’t need:
```go
  n := 0
  for range s {
      n++ 
  }
```
Or we can just call `utf8.RuneCountInString(s)`.

We mentioned earlier that it is mostly a matter of convention in Go that text strings are interpreted as UTF-8-encoded sequences of Unicode code points, but for correct use of `range` loops on strings, it’s more than a convention, it’s a necessity. What happens if we range over a string containing arbitrary binary data or, for that matter, UTF-8 data containing errors?

Each time a UTF-8 decoder, whether explicit in a call to `utf8.DecodeRuneInString` or implicit in a `range` loop, consumes an unexpected input byte, it generates a special Unicode *replacement character*, `'\uFFFD'`, which is usually printed as a white question mark inside a black hexagonal or diamond-like shape `$`. When a program encounters this rune value, it’s often a sign that some upstream part of the system that generated the string data has been careless in its treatment of text encodings.

UTF-8 is exceptionally convenient as an interchange format but within a program runes may be more convenient because they are of uniform size and are thus easily indexed in arrays and slices.

A `[]rune` conversion applied to a UTF-8-encoded string returns the sequence of Unicode code points that the string encodes:
```go
  // "program" in Japanese katakana
  s := "プログラム"
  fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0" 
  r := []rune(s)
  fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
```
(The verb `% x` in the first `Printf` inserts a space between each pair of hex digits.)

If a slice of runes is converted to a string, it produces the concatenation of the UTF-8 encodings of each rune:
```go
  fmt.Println(string(r)) // "プログラム"
```
Converting an integer value to a string interprets the integer as a rune value, and yields the UTF-8 representation of that rune:
```go
  fmt.Println(string(65)) // "A", not "65" 
  fmt.Println(string(0x4eac)) // "京" 
```
If the rune is invalid, the replacement character is substituted:
```go
  fmt.Println(string(1234567)) // "�"
```


### 3.5.4 Strings and Byte Slices
### 3.5.5 Conversions between Strings and Numbers
## 3.6. Constants
### 3.6.1 The Constant Generator `iota`
### 3.6.2 Untyped Constants

