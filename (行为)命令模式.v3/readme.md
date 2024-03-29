**命令**
```
是一种行为设计模式，它将请求封装成一个对象，从而使你可用不同的请求对客户进行参数化；
对请求排队或记录请求日志，以及支持可撤销的操作
```

```
是一种行为设计模式， 它可将请求转换为一个包含与请求相关的所有信息的独立对象。 
该转换让你能根据不同的请求将方法参数化、 延迟请求执行或将其放入队列中， 且能实现可撤销操作。
```
##实现方式
```
声明仅有一个执行方法的命令接口。

抽取请求并使之成为实现命令接口的具体命令类。 
    每个类都必须有一组成员变量来保存请求参数和对于实际接收者对象的引用。 
    所有这些变量的数值都必须通过命令构造函数进行初始化。
找到担任发送者职责的类。 在这些类中添加保存命令的成员变量。 

    发送者只能通过命令接口与其命令进行交互。 
    发送者自身通常并不创建命令对象， 而是通过客户端代码获取。
修改发送者使其执行命令， 而非直接将请求发送给接收者。

客户端必须按照以下顺序来初始化对象：
    创建接收者。
    创建命令， 如有需要可将其关联至接收者。
    创建发送者并将其与特定命令关联。
```
