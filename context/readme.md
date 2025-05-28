# Golang中的context库作用详解

`context`是Go语言标准库中一个非常重要的包，主要用于在多个goroutine之间传递请求相关的值、取消信号和超时信息。以下是它的主要作用：

## 核心功能

1. **传播取消信号**：
    - 允许父goroutine取消其所有子goroutine
    - 提供树状结构的取消传播机制

2. **传递请求范围的值**：
    - 安全地在调用链中传递请求特定的数据
    - 键值对存储，键需要满足可比性

3. **处理超时和截止时间**：
    - 可以设置绝对截止时间(`WithDeadline`)
    - 可以设置相对超时时间(`WithTimeout`)

## 主要使用场景

- **HTTP请求处理**：在Web服务器中传递请求上下文，处理客户端断开连接的情况
- **数据库查询**：为长时间运行的查询设置超时
- **微服务调用**：在服务间调用时传递跟踪ID、认证信息等
- **并发控制**：协调多个goroutine的执行和取消

## 常用方法

```go
// 创建上下文
ctx := context.Background()  // 空上下文，通常作为根
ctx := context.TODO()        // 不确定使用哪种上下文时使用

// 派生上下文
ctx, cancel := context.WithCancel(parentCtx)
ctx, cancel := context.WithTimeout(parentCtx, timeout)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
ctx := context.WithValue(parentCtx, key, value)

// 使用上下文
select {
case <-ctx.Done():
    // 上下文被取消或超时
    return ctx.Err()
default:
    // 正常执行
}
```

## 最佳实践

1. 将`context.Context`作为函数的第一个参数传递
2. 只传递请求范围的数据，而非函数参数
3. 在阻塞操作前检查`ctx.Done()`
4. 取消函数(`cancel`)应该被调用以释放资源
5. 上下文是不可变的，每次派生都会创建新的上下文

`context`包的设计体现了Go语言的并发哲学，提供了一种优雅的方式来管理goroutine的生命周期和跨API边界的请求范围数据。




## Context接口
`context.Context`是一个接口，该接口定义了四个需要实现的方法。具体签名如下：

type Context interface {   
   Deadline() (deadline time.Time, ok bool)  
   Done() <-chan struct{}  
   Err() error 
   Value(key interface{}) interface{}  
}    

其中：   
* `Deadline`方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；

* `Done`方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；

* `Err`方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
如果当前Context被取消就会返回Canceled错误；
如果当前Context超时就会返回DeadlineExceeded错误；

* `Value`方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；
Background()和TODO()

Go内置两个函数：`Background()`和`TODO()`，这两个函数分别返回一个实现了Context接口的background和todo。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。

Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。