// 抽象创建者类或函数
class Creator {
    // 工厂方法，需要被子类实现
    factoryMethod() {
      throw new Error('Factory method not implemented');
    }
  
    // 创造者类的操作
    create() {
      const product = this.factoryMethod();
      return `Creator: The same creator's code has just worked with ${product.operation()}`;
    }
  }
  
  // 具体产品类
  class ConcreteProduct1 {
    operation() {
      return '{Result of the ConcreteProduct1}';
    }
  }
  
  class ConcreteProduct2 {
    operation() {
      return '{Result of the ConcreteProduct2}';
    }
  }
  
  // 具体创建者类，覆盖工厂方法以返回不同类型的产品
  class ConcreteCreator1 extends Creator {
    factoryMethod() {
      return new ConcreteProduct1();
    }
  }
  
  class ConcreteCreator2 extends Creator {
    factoryMethod() {
      return new ConcreteProduct2();
    }
  }
  
  // 客户端代码
  function clientCode(creator) {
    console.log(creator.create());
  }
  
  const creator1 = new ConcreteCreator1();
  clientCode(creator1); // Output: Creator: The same creator's code has just worked with {Result of the ConcreteProduct1}
  
  const creator2 = new ConcreteCreator2();
  clientCode(creator2); // Output: Creator: The same creator's code has just worked with {Result of the ConcreteProduct2}
  