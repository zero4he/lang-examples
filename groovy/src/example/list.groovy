package example
/**
 * @author Zero
 * Created on 2017/7/4.
 */

ints = [0, 1, 2, 3]

newInts = ints.findAll { it > 0 }.toList()
println(newInts) //[1, 2, 3]

newInts += 9
println(newInts) //[1, 2, 3, 9]

println(ints) //[0, 1, 2, 3]

newInts.each {
    println("-> $it")
}

//ints = [0, 1, 2, 3]
println(ints[1..3]) //[1, 2, 3], 从0开始, 闭合区间, 即包括第1个也包括第3个(这里的第一个实际上就是第二个)
println(ints[1..-2]) //[1, 2]  -2表示倒数第二(包括)
