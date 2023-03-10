# defer 

在 Go 语言中，defer 语句必须跟在函数调用后面。这是因为 defer 语句实际上是把一个函数压入一个栈中，而这个栈是由编译器自动生成的。当函数返回时，这个栈会被依次弹出，执行栈中存放的函数。

如果 defer 语句后面没有紧跟着一个函数调用，那么编译器就无法生成合法的栈，因此会产生编译错误。而在你提供的代码中，defer 语句后面并没有紧跟着任何函数调用，所以会导致编译错误。

正确的做法是将 defer 和 recover 语句放在一个函数中，而不是单独使用。这样，defer 语句就可以跟在一个函数调用后面，生成一个合法的栈。