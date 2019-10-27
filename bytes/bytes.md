# bytes

# 1：bytes.go 提供byte数组的操作接口

```
type byte = uint8  byte的原始类型
```

# 2: API

```go
func Equal(a, b []byte) bool {  
    return string(a) == string(b) //转换为string对比 因为string的底层是[]byte数组
}
```





# 3:总结 

大量调用    src\internal\bytealg底层包    调用汇编