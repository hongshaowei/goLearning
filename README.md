# goLearning

Documentation
* Language Spec
* Effective Go
* Go Standard Library
* Go forum - for asking questions to get help

安裝 Go 時，需要設置 GOROOT 和 GOPATH 這兩個環境變數，環境變數來自 Unix 文化圈，對某些初學者，尤其是 Windows 使用者，會產生混淆。這兩個環境變數都是設置目錄路徑；其中 GOROOT 是 Go 開發工具本身所在的位置，Go 開發工具和標準函式庫等都放在這個路徑之下；GOPATH 則是專案和第三方套件所在的位置，不可以和 GOROOT 重覆，通常建議把專案放在 GOPATH 所在的目錄之下，編譯專案時較不會出問題。
* bin/：存放專案所產生的執行檔
* pkg/：存放專案所產生的函式庫檔
* src/：存放專案原始碼

Basic Go Commands
As my understanding, build is just to create an executable, install creates executable and put into your env(workspace/bin/)
* go build
    * run inside a main folder creates an executable inside main, other folders not main won't be able to go build/run cause no entry points(main)
* go install
    * run in main folder and creates the executable in workspace/bin/ and a packagename.a file in workspace/pkg/darwin_amd64/path/as/structure/
* go clean
    * clean other kinds of files not .go

Terminology
* Lexical Elements
    * of or relating to the words or vocabulary of a language(see go spec)
    * definition to words and vocabulary in go language
* Literal Values
    * just values… @@
* Runes
    * character
        * "", `` for string, and string is a collection of runes (`` represents inside is raw string)
        * '' for rune or so called char
        * if assign a variable 'a' and print it, you’ll get the decimal number it represents
        * convert it to string: string(variable) or use "" to surround the char make it a string
    * an integer value identifying a unicode code point
    * with int 32 (4 bytes) we have numbers for each of the code point (the number that represents a char in code scheme)
        * alias for int 32
        * UTF-8 is a 4 bytes coding scheme
        * with int 32 (4 bytes) we have numbers for each of the code points
* Printing UTF-8
* Variadic
    * The final incoming parameter in a function signature may have a type prefixed with ... . A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.
    * The argument is a slice or array and passed into a function that takes …T, you can pass the arg with postfix … as arg… , it’s kinda like open it up and pull each elements out rather than use the single slice or array object


Keywords
類別	關鍵字
程式宣告	import, package
程式實體宣告和定義	chan, const, func, interface, map, struct, type, var
程式流程控制	go, select, break, case, continue, default, defer, else, fallthrough, for, goto, if, range, return, switch
rang
* loop over a object like array or slice and returns index and value
* for i, v range slice {} or nothing to do with index you can : for _, v range slice {}

Data Type
Basic Type	string, bool, byte, rune, int/uint, int8/uint8, int16/uint16, int32/uint32, int64/uint64, float32, float64, complex64, complex128
Composite Type	Array, Struct, Function, Interface, Slice, Map, Channel, Pointer

Variables
keyword 
* var
shorthand 
* :=
* can only be used inside function
zero value
* nil
type format verb
* %T

Reminder
* {} - braces, curly braces, curlies, mustaches
* [] - brackets
* () - parentheses, parens

Parameters and Arguments
when defining a function
- func xxx(parameters)
when calling a function
- xxx(arguments)

Callback 
Simply speaking, callback is passing a func as an argument.
Then while func A was called and func B is the argument, then B will be called by A

Time to use Anonymous function
A function cannot be declared inside a function,
so when it is needed to, you can use 
- var := func() type {},
to make the original function identifier as a variable been declared inside another function
