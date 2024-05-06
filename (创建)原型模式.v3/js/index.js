// // 旧版本
// // 构造函数
// function Person(name) {
//     this.name = name;
// }

// // 向 Person 的原型上添加方法
// Person.prototype.sayHello = function() {
//     console.log('Hello, my name is ' + this.name + '.');
// };

// // 创建两个 Person 实例
// var person1 = new Person('Alice');
// var person2 = new Person('Bob');

// // 由于 sayHello 方法是定义在原型上的，因此所有的 Person 实例都可以访问它。
// person1.sayHello(); // 输出：Hello, my name is Alice.
// person2.sayHello(); // 输出：Hello, my name is Bob.


class Person {
    constructor(name) {
        this.name = name;
    }
    
    sayHello() {
        console.log('Hello, my name is ' + this.name + '.');
    }
}

let person1 = new Person('Alice');
let person2 = new Person('Bob');

person1.sayHello(); // 输出：Hello, my name is Alice.
person2.sayHello(); // 输出：Hello, my name is Bob.
