<?php
namespace php\gpt\decorative;
interface Component {
    public function operation(): string;
}
class ConcreteComponent implements Component {
    public function operation(): string {
        return "ConcreteComponent";
    }
}
abstract class Decorator implements Component {
    protected $component;

    public function __construct(Component $component) {
        $this->component = $component;
    }

    public function operation(): string {
        return $this->component->operation();
    }
}
class ConcreteDecoratorA extends Decorator {
    public function operation(): string {
        return "ConcreteDecoratorA(" . parent::operation() . ")";
    }
}

class ConcreteDecoratorB extends Decorator {
    public function operation(): string {
        return "ConcreteDecoratorB(" . parent::operation() . ")";
    }
}
function clientCode(Component $component) {
    // ...
    echo "RESULT: " . $component->operation();
    // ...
}

// 简单的组件
$simple = new ConcreteComponent();
echo "Client: I've got a simple component:\n";
clientCode($simple);
echo "\n\n";

// 用 ConcreteDecoratorA 装饰
$decorator1 = new ConcreteDecoratorA($simple);
echo "Client: Now I've got a decorated component:\n";
clientCode($decorator1);
echo "\n\n";

// 用 ConcreteDecoratorB 装饰
$decorator2 = new ConcreteDecoratorB($decorator1);
echo "Client: Now I've got a more decorated component:\n";
clientCode($decorator2);
echo "\n\n";

// 输出结果：
// Client: I've got a simple component:
// RESULT: ConcreteComponent

// Client: Now I've got a decorated component:
// RESULT: ConcreteDecoratorA(ConcreteComponent)

// Client: Now I've got a more decorated component:
// RESULT: ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent))
