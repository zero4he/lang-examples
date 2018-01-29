/**
 * @author Zero
 * Created on 2017/7/4.
 */


fun main(args: Array<String>) {

    val sum = { x: Int, y: Int -> x + y }
    println(sum(1, 4))

    val sum2: (Int, Int) -> Int = { x, y -> x + y }
    println(sum2(1, 4))


}


