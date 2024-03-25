// Receiver
class Bulb {
    turnOn() {
      console.log('Bulb has been lit');
    }
    turnOff() {
      console.log('Darkness!');
    }
  }
  
  // Command interface
  class Command {
    execute() {}
    undo() {}
    redo() {}
  }
  
  // ConcreteCommand
  class TurnOnCommand extends Command {
    constructor(bulb) {
      super();
      this.bulb = bulb;
    }
  
    execute() {
      this.bulb.turnOn();
    }
  
    undo() {
      this.bulb.turnOff();
    }
  
    redo() {
      this.execute();
    }
  }
  
  class TurnOffCommand extends Command {
    constructor(bulb) {
      super();
      this.bulb = bulb;
    }
  
    execute() {
      this.bulb.turnOff();
    }
  
    undo() {
      this.bulb.turnOn();
    }
  
    redo() {
      this.execute();
    }
  }
  
  // Invoker
  class RemoteControl {
    submit(command) {
      command.execute();
    }
  }
  
  // Client
  const bulb = new Bulb();
  const turnOnCommand = new TurnOnCommand(bulb);
  const turnOffCommand = new TurnOffCommand(bulb);
  
  const remote = new RemoteControl();
  remote.submit(turnOnCommand);
  remote.submit(turnOffCommand);
  