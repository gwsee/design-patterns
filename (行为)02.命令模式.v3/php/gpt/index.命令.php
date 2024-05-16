<?php
namespace php\gpt\command;

interface Command {
    public function execute();
}
class Light {
    public function turnOn() {
        echo "The light is on.\n";
    }

    public function turnOff() {
        echo "The light is off.\n";
    }
}
class TurnOnLightCommand implements Command {
    protected $light;

    public function __construct(Light $light) {
        $this->light = $light;
    }

    public function execute() {
        $this->light->turnOn();
    }
}

class TurnOffLightCommand implements Command {
    protected $light;

    public function __construct(Light $light) {
        $this->light = $light;
    }

    public function execute() {
        $this->light->turnOff();
    }
}
class RemoteControl {
    public function submit(Command $command) {
        $command->execute();
    }
}
// Receiver
$light = new Light();

// ConcreteCommand
$turnOn = new TurnOnLightCommand($light);
$turnOff = new TurnOffLightCommand($light);

// Invoker
$remote = new RemoteControl();
$remote->submit($turnOn); // 输出：The light is on.
$remote->submit($turnOff); // 输出：The light is off.
