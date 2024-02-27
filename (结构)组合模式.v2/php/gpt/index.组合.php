<?php
namespace php\gpt\combine;
interface Component {
    public function operation();
}
class Leaf implements Component {
    private $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function operation() {
        echo "Leaf " . $this->name . " is operated.\n";
    }
}
class Composite implements Component {
    private $children = [];
    private $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function add(Component $component) {
        $this->children[] = $component;
    }

    public function remove(Component $component) {
        $this->children = array_filter($this->children, function($child) use ($component) {
            return $child !== $component;
        });
    }

    public function operation() {
        echo "Composite " . $this->name . " operates on its children:\n";
        foreach ($this->children as $child) {
            $child->operation();
        }
    }
}
// 创建叶子节点
$leaf1 = new Leaf('1');
$leaf2 = new Leaf('2');
$leaf3 = new Leaf('3');

// 创建复合节点，并添加叶子节点
$composite = new Composite('Comp1');
$composite->add($leaf1);
$composite->add($leaf2);
$composite->add($leaf3);

// 创建更高一级的复合节点
$rootComposite = new Composite('RootComp');
$rootComposite->add($composite);

// 执行操作
$rootComposite->operation();

/**
    Composite RootComp operates on its children:
    Composite Comp1 operates on its children:
    Leaf 1 is operated.
    Leaf 2 is operated.
    Leaf 3 is operated.
 */