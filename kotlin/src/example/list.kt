/**
 * @author Zero
 * Created on 2017/7/4.
 */
fun main(args: Array<String>) {
    val ints = listOf(0, 1, 2, 3)
    var newInts = ints.filter { it > 0 }.toList() //这里使用var声明变量, 因为等下需要用到字符重载

    println(newInts) //[1, 2, 3]

    println(newInts.plus(9)) //[1, 2, 3, 9]  , 返回了一个新的list

    newInts += 4 //等同于newInts.plus(8) , 因为plus是 operator fun

    println(newInts) //[1, 2, 3, 4]

    newInts.forEach {
        println("->$it")
    }

    //val ints = listOf(1, 2, 4, 6)
    println(ints.slice(1..3)) //[1, 2, 3] , 从0开始, 左闭右开区间, slice不是operator fun

    //


}
