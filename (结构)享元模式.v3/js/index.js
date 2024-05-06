class Color {
    constructor(name) {
        this.name = name;
    }
}

class ColorFactory {
    constructor() {
        this.colors = {};
    }
    create(name) {
        let color = this.colors[name];
        if (color) {
            return color;
        } else {
            color = new Color(name);
            this.colors[name] = color;
        }
        return color;
    }
}

// 使用享元模式
const factory = new ColorFactory();

const red1 = factory.create('Red');
const red2 = factory.create('Red');
const blue1 = factory.create('Blue');

console.log(red1 === red2); // true 因为它们是共享的
console.log(red1 === blue1); // false 因为它们代表不同的颜色
