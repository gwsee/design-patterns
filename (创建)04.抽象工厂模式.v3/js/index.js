// 抽象工厂类
class AbstractFactory {
  createProductA() {
    throw new Error('AbstractFactory: createProductA is not implemented.');
  }
  createProductB() {
    throw new Error('AbstractFactory: createProductB is not implemented.');
  }
}

// 具体产品A1
class ProductA1 {
  use() {
    return 'ProductA1 is being used';
  }
}

// 具体产品A2
class ProductA2 {
  use() {
    return 'ProductA2 is being used';
  }
}

// 具体产品B1
class ProductB1 {
  use() {
    return 'ProductB1 is being used';
  }
}

// 具体产品B2
class ProductB2 {
  use() {
    return 'ProductB2 is being used';
  }
}

// 具体工厂1
class ConcreteFactory1 extends AbstractFactory {
  createProductA() {
    return new ProductA1();
  }
  createProductB() {
    return new ProductB1();
  }
}

// 具体工厂2
class ConcreteFactory2 extends AbstractFactory {
  createProductA() {
    return new ProductA2();
  }
  createProductB() {
    return new ProductB2();
  }
}

// 客户端代码
function client(factory) {
  const productA = factory.createProductA();
  const productB = factory.createProductB();
  
  console.log(productA.use());
  console.log(productB.use());
}

// 使用具体工厂1
client(new ConcreteFactory1()); // "ProductA1 is being used", "ProductB1 is being used"
// 使用具体工厂2
client(new ConcreteFactory2()); // "ProductA2 is being used", "ProductB2 is being used"
