<?php
//php index.单例.php
class Singleton
{
    /**
     * @var Singleton $instance 单例模式的实例
     */
    private static $instance = null;

    /**
     * 构造函数为private，防止直接创建对象
     */
    private function __construct()
    {
        // 初始化代码
    }

    /**
     * 克隆函数为private，防止克隆对象
     */
    private function __clone()
    {
        // 禁止对象克隆
    }

    /**
     * 通过一个静态方法提供全局访问点
     *
     * @return Singleton 返回自身实例
     */
    public static function getInstance()
    {
        if (self::$instance === null) {
            self::$instance = new Singleton();
        }

        return self::$instance;
    }

    // 例子方法，你可以添加自己的业务方法
    public function doSomething()
    {
        echo "Doing something...\n";
    }
}

// 使用单例模式
$instance1 = Singleton::getInstance();
$instance1->doSomething();

$instance2 = Singleton::getInstance();
$instance2->doSomething();

// 检查两个实例是否相同
var_dump($instance1 === $instance2); // 输出 bool(true)
?>
