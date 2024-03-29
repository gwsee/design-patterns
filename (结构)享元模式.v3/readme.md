**享元**
```
它通过共享对象的方式来优化代码、减少应用程序中对象的数量。
    这个模式特别适合处理大量对象（例如文本编辑器中的字符、游戏中的树木等），
    它们的许多状态可以被外部化，从而使得多个对象可以共享相同的可变数据。
    
    
```

##实现方式
```
1：将需要改写为享元的类成员变量拆分为两个部分：
    内在状态： 包含不变的、 可在许多对象中重复使用的数据的成员变量。
    外在状态： 包含每个对象各自不同的情景数据的成员变量

2：保留类中表示内在状态的成员变量， 并将其属性设置为不可修改。 
    这些变量仅可在构造函数中获得初始数值。

3：找到所有使用外在状态成员变量的方法， 为在方法中所用的每个成员变量新建一个参数， 并使用该参数代替成员变量。

4：你可以有选择地创建工厂类来管理享元缓存池， 它负责在新建享元时检查已有的享元。 
    如果选择使用工厂， 客户端就只能通过工厂来请求享元， 它们需要将享元的内在状态作为参数传递给工厂。

5：客户端必须存储和计算外在状态 （情景） 的数值， 因为只有这样才能调用享元对象的方法。 
    为了使用方便， 外在状态和引用享元的成员变量可以移动到单独的情景类中。
```