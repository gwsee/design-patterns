<?php
namespace php\gpt\state;

interface State {
    public function doAction(Context $context);
}
class StartState implements State {
    public function doAction(Context $context) {
        echo "Player is in start state\n";
        $context->setState($this);
    }

    public function __toString(){
        return "Start State";
    }
}

class StopState implements State {
    public function doAction(Context $context) {
        echo "Player is in stop state\n";
        $context->setState($this);
    }

    public function __toString(){
        return "Stop State";
    }
}
class Context {
    private $state;

    public function __construct() {
        $this->state = null;
    }

    public function getState() {
        return $this->state;
    }

    public function setState(State $state) {
        $this->state = $state;
    }
}

// 创建上下文对象
$context = new Context();

// 创建状态对象
$startState = new StartState();
$stopState = new StopState();

// 使用状态对象改变上下文状态
$startState->doAction($context);

// 我们可以通过上下文对象查看其状态变化
echo $context->getState();

$stopState->doAction($context);
echo $context->getState();

