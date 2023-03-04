# slice 原理


go 的 slice 在函数参数传递中是值传递还是引用传递
在函数参数传递中，Go 的 slice 是引用传递。

当将 slice 传递给函数时，传递的是一个指向底层数组的指针，而不是整个 slice 的副本。因此，对 slice 的修改将影响原始 slice。

下面是一个简单的示例代码：

go
Copy code
package main

import "fmt"

func modifySlice(s []int) {
s[0] = 100
s = append(s, 200)
fmt.Println("Inside function:", s)
}

func main() {
numbers := []int{1, 2, 3, 4, 5}
fmt.Println("Before function call:", numbers)

    modifySlice(numbers)

    fmt.Println("After function call:", numbers)
}
输出为：

r
Copy code
Before function call: [1 2 3 4 5]
Inside function: [100 2 3 4 5 200]
After function call: [100 2 3 4 5]
在这个例子中，modifySlice 函数将第一个元素修改为 100，并在 slice 中添加了一个新元素 200。在函数调用后，原始的 numbers slice 也被修改为 [100 2 3 4 5]，因为 numbers 和 s 实际上引用了同一个底层数组。




hao wu
那为什么 我append 他引用传递就失效了
当你向 slice 中添加元素时，如果 slice 的底层数组容量不足，Go 会创建一个新的底层数组，并将原始数组的元素复制到新的数组中，然后添加新的元素。在这种情况下，新的底层数组地址与原始 slice 的底层数组地址不同，因此对新 slice 的修改不会影响原始 slice。

以下是一个示例代码：

go
Copy code
package main

import "fmt"

func modifySlice(s []int) {
s[0] = 100
s = append(s, 200)
fmt.Println("Inside function:", s)
}

func main() {
numbers := make([]int, 5, 5)
for i := 0; i < 5; i++ {
numbers[i] = i + 1
}
fmt.Println("Before function call:", numbers)

    modifySlice(numbers)

    fmt.Println("After function call:", numbers)
}
输出为：

r
Copy code
Before function call: [1 2 3 4 5]
Inside function: [100 2 3 4 5 200]
After function call: [100 2 3 4 5]
在这个例子中，我们使用 make 函数创建了一个具有容量为 5 的 slice。在函数调用中，我们首先将第一个元素修改为 100，然后将一个新元素 200 添加到 slice 中。在添加新元素时，底层数组已满，因此 Go 创建了一个新的底层数组，并将原始数组的元素复制到新数组中，然后添加新元素。新的 slice [100 2 3 4 5 200] 的底层数组与原始 slice 的底层数组不同，因此对新 slice 的修改不会影响原始 slice，输出结果为 [100 2 3 4 5]。




hao wu
那slice本质是什么传递
slice 本质上是一个包含底层数组地址、长度和容量的结构体。因此，在函数参数传递中，Go 实际上是将 slice 的指针传递给函数，即传递底层数组的地址。

因为 slice 包含底层数组地址，所以当 slice 作为函数参数传递时，实际上传递的是指向底层数组的指针，而不是 slice 的副本。因此，对 slice 的任何修改都会反映在原始 slice 上。

需要注意的是，当向 slice 中添加元素时，如果底层数组的容量不足，Go 会为 slice 创建一个新的底层数组，并将原始数组的元素复制到新的数组中，然后添加新元素。在这种情况下，新的底层数组地址与原始 slice 的底层数组地址不同，因此对新的 slice 的修改不会影响原始 slice。