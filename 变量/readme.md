## Go变量
### 1.什么是变量
Go语言是静态类型语言，因此`变量有明确类型的`，编译器也会检查变量类型的正确性。从计算机系统实现角度来看，变量是一段或多段用来存储数据的内存。
### 2.Go变量类型
Go语言的基本类型有：
* bool
* string
* int、int8、int16、int32、int64
* uint、uint8、uint16、uint32、uint64、uintptr
* byte // uint8 的别名
* rune // int32 的别名 代表一个 Unicode 码
* float32、float64
* complex64、complex128

Go语言提供了25个关键字和此外，30 多个预定义的名字，比如 int 和 true 等不能作为变量名：

``` go
//关键字
break       default        func         interface        select
case        defer         go           map              struct
chan        else           goto           package          switch
const        fallthrough    if             range            type
continue    for             import        return            var
//预定义字符
true false iota nil

        int int8 int16 int32 int64
        uint uint8 uint16 uint32 uint64 uintptr
        float32 float64 complex128 complex64
        bool byte rune string error

        make len cap new append copy close delete
        complexrealimag
        panic recover
```

同时，当一个Go变量被声明之后，`系统会自动赋零值`：如int为0，bool为false，string为空字符串等。`所有的内存在 Go 中都是经过初始化的。`
变量命名一般使用驼峰式，首个单词小写，后续每个新单词首字母大写，如：helloWorld。

### 3.Go变量的声明
Go使用关键字var声明变量，格式为：
```go
var a int // var 开头，后置变量名 变量类型，结尾无须分号，变量会初始附零值，此处输出0
var b int = 123 // 定义同时并赋值
```
批量声明格式为：
```go
//声明多个变量
     	var a,b int

     	var (
     		c int
     		d string
     		e []float32
     		f func() bool
     		g struct {
     			x int
     		}
     	)
```
简化格式为：
```go
	//简化格式定义变量   名字 := 表达式
	//需要注意的是，简短模式（short variable declaration）有以下限制：
	//定义变量，同时显式初始化。
	//不能提供数据类型。
	//只能用在函数内部。
	a := 123
	b,c:=1, "abc"
	fmt.Println(a,b,c)
```
### 3.GO变量转换
```go
    a, b := 10, 20

    //可以通过多重赋值进行交换
    a, b = b, a

    fmt.Println(a)
    fmt.Println(b)
```
