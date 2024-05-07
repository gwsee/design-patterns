class Product {
    constructor() {
        this.parts = [];
    }

    addPart(part) {
        this.parts.push(part);
    }

    listParts() {
        console.log(`Product parts: ${this.parts.join(', ')}`);
    }
}

class Builder {
    constructor() {
        this.product = new Product();
    }

    buildPartA() {} // 抽象方法
    buildPartB() {} // 抽象方法
    buildPartC() {} // 抽象方法

    getProduct() {
        return this.product;
    }
}

class ConcreteBuilder extends Builder {
    buildPartA() {
        this.product.addPart('PartA');
    }

    buildPartB() {
        this.product.addPart('PartB');
    }

    buildPartC() {
        this.product.addPart('PartC');
    }
}

class Director {
    construct(builder) {
        builder.buildPartA();
        builder.buildPartB();
        builder.buildPartC();
    }
}

// 使用
const director = new Director();
const builder = new ConcreteBuilder();
director.construct(builder);

const product = builder.getProduct();
product.listParts(); // Output: Product parts: PartA, PartB, PartC
