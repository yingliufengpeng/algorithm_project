import scala.reflect.ClassTag
import scala.collection.mutable.ArrayBuffer
object ArrayDemo {


    class Array2[A: ClassTag](val n: Int, val array: Array[A]) {

        assert(n > 0)

        def this(n: Int) = {
            this(n, Array.ofDim(n))
        }

        def this(n: Int, ini: Int => A) = {
            this(n, Array.tabulate(n)(ini))
        }

        def map[B: ClassTag](f: A => B): Array2[B] = new Array2[B](this.n, this.array.map(f))

        def flatMap[B: ClassTag](f: A => Array2[B]) = {
            val buf = ArrayBuffer.empty[B]
            this.array.foreach(buf ++= f(_).array)
            new Array2[B](buf.length, buf.toArray)
        }

        def filter[B >: A](cond: A => Boolean): Array2[A] = {

            val buf = ArrayBuffer.empty[A] 

            for {
                e <- this.array if cond(e) 
            } {
                buf += e 
            }

            new Array2[A](buf.length, buf.toArray)

        }

        def foreach(f: A => Unit): Unit = {

            this.array.foreach(f)
        }



        override def toString(): String = this.array.map(_.toString()).mkString(",")

        
    }
    object Array2  {


        def apply[A: ClassTag](arrs: A*) = {

            new Array2(arrs.length, arrs.toArray)

        }
    }

    def test: Unit = {
        
        val arr = new Array2[Int](3, _ + 1)

        val arr2 = arr.map(_ * 2)
        println(s"arr2 is $arr2")

        val arr3 = arr.flatMap[String](e => new Array2(1, index => (index + 3).toString()))
        println(s"arr3 is $arr3")

        val arr4 = arr.foreach(println)

        val arr5 = Array2(1, 2, 3, 4)
        println(s"arr5 is $arr5")


    

        

    }

}