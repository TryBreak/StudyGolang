# 标准库

- bufio
- log
- encoding/json &&
- regexp
- time
- strings/math/rand
- godoc -http

## jsoniter 比标准库快 6 倍

go get github.com/json-iterator/go

- 转为 jsonStr

```go
jsoniter.Marshal(group)

```

- 转为结构体

```go
var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
var animals []Animal
jsoniter.Unmarshal(jsonBlob, &animals)

m := make(map[string]interface{})  // json 解析常规格式

```

## 日志

zap
