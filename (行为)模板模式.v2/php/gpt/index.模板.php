<?php
namespace php\gpt\template;
abstract class Game
{
    // 模板方法，定义了算法的骨架
    public final function play()
    {
        $this->setup();
        $this->startPlay();
        $this->endPlay();
    }

    // 准备游戏，由子类实现的具体细节
    abstract protected function setup();

    // 开始游戏，由子类实现的具体细节
    abstract protected function startPlay();

    // 结束游戏，由子类实现的具体细节
    abstract protected function endPlay();
}

class Football extends Game
{
    protected function setup()
    {
        echo "Setting up a football game.\n";
    }

    protected function startPlay()
    {
        echo "Starting the football game.\n";
    }

    protected function endPlay()
    {
        echo "Ending the football game.\n";
    }
}

class Basketball extends Game
{
    protected function setup()
    {
        echo "Setting up a basketball game.\n";
    }

    protected function startPlay()
    {
        echo "Starting the basketball game.\n";
    }

    protected function endPlay()
    {
        echo "Ending the basketball game.\n";
    }
}

// 使用
$football = new Football();
$football->play();

echo "------------------\n";

$basketball = new Basketball();
$basketball->play();
