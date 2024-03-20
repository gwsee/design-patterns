// 组件抽象类
class UIComponent {
    constructor(name) {
      this.name = name;
    }
    
    // 添加组件
    add() {}
    
    // 移除组件
    remove() {}
    
    // 获取子组件
    getChild() {}
    
    // 具体的操作方法
    operation() {}
  }
  
  // 叶节点组件
  class Leaf extends UIComponent {
    constructor(name) {
      super(name);
    }
    
    operation() {
      console.log(`Leaf ${this.name} is operated.`);
    }
  }
  
  // 组合节点组件
  class Composite extends UIComponent {
    constructor(name) {
      super(name);
      this.children = [];
    }
    
    add(component) {
      this.children.push(component);
    }
    
    remove(component) {
      const index = this.children.indexOf(component);
      if (index > -1) {
        this.children.splice(index, 1);
      }
    }
    
    getChild(index) {
      return this.children[index];
    }
    
    operation() {
      console.log(`Composite ${this.name}`);
      for (const child of this.children) {
        child.operation();
      }
    }
  }
  
  // 使用示例
  const root = new Composite('Root');
  const leaf1 = new Leaf('Leaf 1');
  const leaf2 = new Leaf('Leaf 2');
  
  const subComp = new Composite('Sub Composite');
  const leaf3 = new Leaf('Leaf 3');
  subComp.add(leaf3);
  
  root.add(leaf1);
  root.add(leaf2);
  root.add(subComp);
  
  root.operation();
  // 输出结果：
  // Composite Root
  // Leaf Leaf 1 is operated.
  // Leaf Leaf 2 is operated.
  // Composite Sub Composite
  // Leaf Leaf 3 is operated.
  