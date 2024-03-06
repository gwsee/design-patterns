<?php
namespace php\gpt\fly_weight;

interface FlyweightInterface {
    public function render($extrinsicState);
}

class ConcreteFlyweight implements FlyweightInterface {
    private $intrinsicState;

    public function __construct($intrinsicState) {
        $this->intrinsicState = $intrinsicState;
    }

    public function render($extrinsicState) {
        // 实际中，根据内部状态和外部状态完成某些功能，例如渲染
        return sprintf('Intrinsic State: %s, Extrinsic State: %s', $this->intrinsicState, $extrinsicState);
    }
}

class FlyweightFactory {
    private $pool = [];

    public function getFlyweight($intrinsicState) {
        if (!isset($this->pool[$intrinsicState])) {
            $this->pool[$intrinsicState] = new ConcreteFlyweight($intrinsicState);
        }

        return $this->pool[$intrinsicState];
    }
}

$flyweightFactory = new FlyweightFactory();

$flyweightA = $flyweightFactory->getFlyweight('stateA');
echo $flyweightA->render('extraState1');
echo "\n";
$flyweightB = $flyweightFactory->getFlyweight('stateA');
echo $flyweightB->render('extraState2');

// 由于stateA的内部状态相同，$flyweightA 和 $flyweightB 指向的是同一个实例

//Intrinsic State: stateA, Extrinsic State: extraState1
//Intrinsic State: stateA, Extrinsic State: extraState2

