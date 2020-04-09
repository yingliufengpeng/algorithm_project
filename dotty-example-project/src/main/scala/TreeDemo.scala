object TreeDemo {

    enum Tree[+A] { // 定义一个Tree结构的泛型基类
       case Branch(left: Tree[A], right: Tree[A]) // 顶级具体子类Branch分支节点
       case Leaf(v: A)  // 定义具体子类叶子节点
       case Empty       // 哨兵对象,具体做递归函数终止条件


       // 通用方法,显示显示字符串
       override def toString(): String = this match {
           case Branch(left, right) => s"$left $right"
           case Leaf(v) => s"$v"
           case Empty => ""
       }

       // 把一棵树映射为另一颗树
       def map[B](f: A => B): Tree[B] = this match {
           case Branch(left, right) => Branch(left.map(f), right.map(f))
           case Leaf(a) => Leaf(f(a))
           case Empty => Empty
       }

       // 对一棵树加入扩展节点的操作
       def flatMap[B](f: A => Tree[B]): Tree[B] = this match {
           case Branch(left, right) => Branch(left.flatMap(f), right.flatMap(f))
           case Leaf(v) => Branch(f(v), Empty)
           case Empty => Empty
       }

       // 对数进行过滤操作,只需要自己想要子节点
       def filter(cond: A => Boolean): Tree[A] = this match {
           case Branch(left, right) => 
                (left.filter(cond), right.filter(cond)) match {
                    case (Empty, p) => Branch(Empty, p)
                    case (p, Empty) => Branch(p, Empty)
                    case (p1, p2) => Branch(p1, p2)
                }
           case Leaf(v) if (cond(v)) => Leaf(v)
           case _ => Empty
       }

    }


    def test: Unit = {
        import Tree.{Branch, Leaf}

        val tree = Branch(Branch(Leaf(3), Leaf(4)), Leaf(4))
        println(s"tree is $tree")

        val tree2 = tree.filter(_ > 3)
        println(s"tree2 is $tree2")

        val tree3 = tree.map(_ * 2)
        println(s"tree3 is $tree3")

        val tree4 = tree.flatMap(e => Branch(Leaf(e + 10), Leaf(e + 11)))
        println(s"tree4 is $tree4")

    }

}