


## Zero Value

Any variable declared but not assigned a value defaults to a `zero value` unique to each type.

## Literals

    Pg 18: Literals in Go are untyped.
        * Covered more later. 

- Integer Literals

Base10, Base16, Octal

- Floating Point Literals
- Rune Literals - Represent characters and are surrounded by single quotes. 
  - ** Single and Double quoted text are not interchangeable. **
- String Literals
  - Double Quotes: "Greetings and Sultations"

## Booleans

    bool type
    Zero Value is false

    var flag bool   // no value assigned, set to false
    var isAwesome = true

## Numeric Types

    Zero Value for all integer types is 0.
    Do: Just use the int type, unless other reason
    NOTE: Integer div in Go follows truncation towards zero. 
    WARNING: Don't use floating points to represent exact decimal values, like money

    Type name   Value range
    int8            –128 to 127
    int16           –32768 to 32767
    int32           –2147483648 to 2147483647
    int64           –9223372036854775808 to 9223372036854775807
    byte (uint8)    0 to 255
    uint16          0 to 65536
    uint32          0 to 4294967295
    uint64          0 to 18446744073709551615

    Special Names
    byte            more common name for uint8
    int             On x86-64, is uint64
    uint            Unsigned version of int (values are always 0 or positive)

    Choosing which Integer to Use

    pg. 21

    For binary file format or network protocol using a specific int size, use the corresponding int type.
    For library functions meant to work w/ any integer type, write a pair of functions, one w/ int64 and
    the other with uint. At time of writing generics did not exist. Refer to page 21.

## Variable Declarations

Go has a lot of ways to declare variables. The declaration style communicates something about how
the variable is used.

    Verbose
    var x int = 10
    var x     = 10  // If 10 is meant to be an int, you can skip the type declaration

    var x int       // Assigns zero value for int
    var x, y int = 10, 20
    
    var x, y = 10, "Hello"  // x is int, y is string literal