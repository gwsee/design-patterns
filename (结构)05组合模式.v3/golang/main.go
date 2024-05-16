package main

/**
让我们试着用一个操作系统文件系统的例子来理解组合模式。
文件系统中有两种类型的对象： 文件和文件夹。 在某些情形中， 文件和文件夹应被视为相同的对象。 这就是组合模式发挥作用的时候了。

想象一下， 你需要在文件系统中搜索特定的关键词。
这一搜索操作需要同时作用于文件和文件夹上。 对于文件而言， 其只会查看文件的内容； 对于文件夹则会在其内部的所有文件中查找关键词。
*/

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
	/**
	Serching recursively for keyword rose in folder Folder2
	Searching for keyword rose in file File2
	Searching for keyword rose in file File3
	Serching recursively for keyword rose in folder Folder1
	Searching for keyword rose in file File1
	*/
}
