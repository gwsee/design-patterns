<?php

namespace php\gpt\builder;

class Car {
    public $brand;
    public $model;
    public $color;

    public function show() {
        echo "Car: {$this->brand}, {$this->model}, {$this->color}\n";
    }
}
interface CarBuilderInterface {
    public function setBrand($brand);
    public function setModel($model);
    public function setColor($color);
    public function getResult();
}
class CarBuilder implements CarBuilderInterface {
    private $car;

    public function __construct() {
        $this->car = new Car();
    }

    public function setBrand($brand) {
        $this->car->brand = $brand;
        return $this;
    }

    public function setModel($model) {
        $this->car->model = $model;
        return $this;
    }

    public function setColor($color) {
        $this->car->color = $color;
        return $this;
    }

    public function getResult() {
        return $this->car;
    }
}
class CarDirector {
    public function build(CarBuilder $builder) {
        return $builder
            ->setBrand('Tesla')
            ->setModel('Model S')
            ->setColor('Red')
            ->getResult();
    }
}
// 直接使用生成器
$builder = new CarBuilder();
$car = $builder
    ->setBrand('BMW')
    ->setModel('X6')
    ->setColor('White')
    ->getResult();
$car->show(); // 输出: Car: BMW, X6, White

// 使用导演类
$director = new CarDirector();
$car = $director->build(new CarBuilder());
$car->show(); // 输出: Car: Tesla, Model S, Red
