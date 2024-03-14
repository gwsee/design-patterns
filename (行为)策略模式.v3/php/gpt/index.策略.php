<?php
// 命名空间
namespace php\gpt\strategy;

// 策略接口，定义了一个操作方法
interface StrategyInterface {
    public function doOperation($num1, $num2);
}

// 具体策略实现：加法
class OperationAdd implements StrategyInterface {
    public function doOperation($num1, $num2) {
        return $num1 + $num2;
    }
}

// 具体策略实现：减法
class OperationSubtract implements StrategyInterface {
    public function doOperation($num1, $num2) {
        return $num1 - $num2;
    }
}

// 具体策略实现：乘法
class OperationMultiply implements StrategyInterface {
    public function doOperation($num1, $num2) {
        return $num1 * $num2;
    }
}

// 上下文类，它接受一个策略对象，并使用它执行操作
class Context {
    private $strategy;

    public function __construct(StrategyInterface $strategy) {
        $this->strategy = $strategy;
    }

    public function executeStrategy($num1, $num2) {
        return $this->strategy->doOperation($num1, $num2);
    }
}

// 使用策略模式的例子
$contextAdd = new Context(new OperationAdd());
echo "10 + 5 = " . $contextAdd->executeStrategy(10, 5) . "\n";

$contextSubtract = new Context(new OperationSubtract());
echo "10 - 5 = " . $contextSubtract->executeStrategy(10, 5) . "\n";

$contextMultiply = new Context(new OperationMultiply());
echo "10 * 5 = " . $contextMultiply->executeStrategy(10, 5) . "\n";
