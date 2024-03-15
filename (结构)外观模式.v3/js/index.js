// 假设有一个复杂的子系统，包含多个方法
class ComplexSubsystemA {
    operation1() {
      console.log('ComplexSubsystemA: Operation 1');
    }
    operation2() {
      console.log('ComplexSubsystemA: Operation 2');
    }
  }
  
  class ComplexSubsystemB {
    operation1() {
      console.log('ComplexSubsystemB: Operation 1');
    }
    operation2() {
      console.log('ComplexSubsystemB: Operation 2');
    }
  }
  
  // 创建一个外观类，为复杂子系统提供简化的方法
  class Facade {
    constructor(subsystemA, subsystemB) {
      this.subsystemA = subsystemA || new ComplexSubsystemA();
      this.subsystemB = subsystemB || new ComplexSubsystemB();
    }
  
    simpleOperation() {
      console.log('Facade: Simplified operation.');
      this.subsystemA.operation1();
      this.subsystemB.operation1();
    }
  }
  
  // 客户端代码通过外观类接口使用子系统
  const facade = new Facade();
  facade.simpleOperation();
/**
Facade: Simplified operation.
ComplexSubsystemA: Operation 1
ComplexSubsystemB: Operation 1
*/