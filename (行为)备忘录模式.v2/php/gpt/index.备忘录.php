<?php
namespace php\gpt\memento;

// Memento (备忘录)
class Memento
{
    private $state;

    public function __construct($state)
    {
        $this->state = $state;
    }

    public function getState()
    {
        return $this->state;
    }
}

// Originator (发起人)
class Originator
{
    private $state;

    public function setState($state)
    {
        $this->state = $state;
    }

    public function getState()
    {
        return $this->state;
    }

    public function saveStateToMemento()
    {
        return new Memento($this->state);
    }

    public function getStateFromMemento(Memento $memento)
    {
        $this->state = $memento->getState();
    }
}

// Caretaker (看护者)
class Caretaker
{
    private $mementoList = array();

    public function add(Memento $state)
    {
        $this->mementoList[] = $state;
    }

    public function get($index)
    {
        return $this->mementoList[$index];
    }
}

// 示例使用
$originator = new Originator();
$caretaker = new Caretaker();

// 改变状态
$originator->setState("State #1");
$originator->setState("State #2");
$caretaker->add($originator->saveStateToMemento());

// 改变再次状态
$originator->setState("State #3");

// 我们想恢复到State #2
$originator->getStateFromMemento($caretaker->get(0));
echo $originator->getState(); // 输出 State #2
