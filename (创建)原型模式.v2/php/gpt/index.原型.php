<?php
namespace php\gpt\prototype;

class Book {
    private $title;
    private $author;

    public function __construct($title, $author) {
        $this->title = $title;
        $this->author = $author;
    }

    public function getTitle() {
        return $this->title;
    }

    public function getAuthor() {
        return $this->author;
    }

    // 定义__clone魔法方法
    public function __clone() {
        // 在这里可以实现深拷贝逻辑
        // 如果Book对象有引用其他对象的属性，则在这里可以手动复制那些对象
        // 以确保每个副本都有独立的实例
    }
}

// 创建一个新书对象
$originalBook = new Book("The Great Gatsby", "F. Scott Fitzgerald");

// 克隆书对象，制作一个副本
$clonedBook = clone $originalBook;

// 下面两个输出相同的结果，因为它们拥有相同的数据
echo $clonedBook->getTitle(); // 输出: The Great Gatsby
echo "\n";
echo $clonedBook->getAuthor(); // 输出: F. Scott Fitzgerald

// 但是$originalBook和$clonedBook是两个不同的实例
//
//    The Great Gatsby
//    F. Scott Fitzgerald
