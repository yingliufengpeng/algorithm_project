
object StackDemo {

    class Stack[A](val elems: List[A]) {

        def push(a: A): Stack[A] = Stack((this.elems :+ a):_*)

        def pop(): (Stack[A], Option[A]) = this.elems.lastOption match {
            case p: Some[A] => (Stack(this.elems.init:_*), p)
            case p => ((Stack(this.elems:_*), p))
        }

        def peek(): Option[A] = this.elems.lastOption

        override def toString(): String = this.elems.map(_.toString).mkString(",")
    }

    object Stack {

        def apply[A](arr: A*): Stack[A] = {

            new Stack[A](arr.toList)
        }
    }

    def test: Unit = {
        
        val stack = Stack(2, 3)
        println(s"stack is $stack")

        val s2 = stack.push(3)
        println(s"s2 is $s2")

        val (s3, somevalue) = stack.pop()
        println(s"s3 is $s3  value is ${somevalue.get}")

        val s4 = stack.peek()
        println(s"s4 is ${s4.get}")


    }


}