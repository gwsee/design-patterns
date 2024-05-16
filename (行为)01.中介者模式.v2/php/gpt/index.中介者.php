<?php
namespace php\gpt\mediator;

interface MediatorInterface {
    public function notify($sender, $event);
}

class ConcreteMediator implements MediatorInterface {
    private $components = [];

    public function register($component) {
        $this->components[] = $component;
        $component->setMediator($this);
    }

    public function notify($sender, $event) {
        foreach ($this->components as $component) {
            if ($component !== $sender) {
                $component->handleEvent($event);
            }
        }
    }
}

class BaseComponent {
    protected $mediator;

    public function setMediator(MediatorInterface $mediator) {
        $this->mediator = $mediator;
    }

    public function handleEvent($event) {
        // Handle the event.
    }
}
class Component1 extends BaseComponent {
    public function handleEvent($event) {
        if ($event == "A") {
            $this->doA();
        } else if ($event == "B") {
            $this->doB();
        }
    }

    public function doA() {
        echo "Component 1 does A.\n";
        $this->mediator->notify($this, "A");
    }

    public function doB() {
        echo "Component 1 does B.\n";
        $this->mediator->notify($this, "B");
    }
}

class Component2 extends BaseComponent {
    public function handleEvent($event) {
        if ($event == "C") {
            $this->doC();
        } else if ($event == "D") {
            $this->doD();
        }
    }

    public function doC() {
        echo "Component 2 does C.\n";
        $this->mediator->notify($this, "C");
    }

    public function doD() {
        echo "Component 2 does D.\n";
        $this->mediator->notify($this, "D");
    }
}
$mediator = new ConcreteMediator();
$component1 = new Component1();
$component2 = new Component2();

$mediator->register($component1);
$mediator->register($component2);

$component1->doA();
$component2->doD();
