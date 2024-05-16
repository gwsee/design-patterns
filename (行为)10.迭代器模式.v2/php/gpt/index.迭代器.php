<?php

//迭代器模式是一种行为设计模式，
//    它提供一种方法顺序访问一个聚合对象中的各个元素，而又无需暴露该对象的内部表示。
//    在 PHP 中，迭代器模式通常通过实现 Iterator 接口或使用内置的 Iterator 类来实现。
//Iterator 接口定义了以下方法，任何实现该接口的 PHP 类都必须提供这些方法的具体实现：
//current() 返回当前元素。
//next() 移动到下一个元素。
//key() 返回当前元素的键。
//valid() 检查当前位置是否有效。
//rewind() 将迭代器重置到第一个元素。

namespace php\gpt\iterator;

class Book
{
    private $author;
    private $title;

    public function __construct($title, $author)
    {
        $this->author = $author;
        $this->title = $title;
    }

    public function getAuthor()
    {
        return $this->author;
    }

    public function getTitle()
    {
        return $this->title;
    }

    public function getAuthorAndTitle()
    {
        return $this->getTitle() . ' by ' . $this->getAuthor();
    }
}

class BookList implements \Iterator
{
    private $books = [];
    private $currentIndex = 0;

    public function addBook(Book $book)
    {
        $this->books[] = $book;
    }

    public function removeBook(Book $book)
    {
        foreach ($this->books as $key => $value) {
            if ($value->getAuthorAndTitle() == $book->getAuthorAndTitle()) {
                unset($this->books[$key]);
            }
        }
        $this->books = array_values($this->books);
    }

    public function current()
    {
        return $this->books[$this->currentIndex];
    }

    public function key()
    {
        return $this->currentIndex;
    }

    public function next()
    {
        $this->currentIndex++;
    }

    public function valid()
    {
        return isset($this->books[$this->currentIndex]);
    }

    public function rewind()
    {
        $this->currentIndex = 0;
    }
}

// 使用迭代器
$bookList = new BookList();
$bookList->addBook(new Book("1984", "George Orwell"));
$bookList->addBook(new Book("The Hitchhiker's Guide to the Galaxy", "Douglas Adams"));

foreach ($bookList as $book) {
    echo $book->getAuthorAndTitle(), "\n";
}
