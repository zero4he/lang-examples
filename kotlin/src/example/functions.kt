/**
 * @author Zero
 * Created on 2017/7/5.
 */
fun printMessage(message: String): Unit {                           // 1
    println(message)
}

fun printMessageWithPrefix(message: String, prefix: String = "") { // 2
    println("[$prefix] $message")
}

fun sum(x: Int, y: Int): Int {                                      // 3
    return x + y
}

fun multiply(x: Int, y: Int) = x * y                                // 4

fun main(args: Array<String>) {
    printMessage("Hello")                                           // 5
    printMessageWithPrefix("Hello", "Log")                          // 6
    printMessageWithPrefix("Hello")                                 // 7
    printMessageWithPrefix(prefix = "Log", message = "Hello")       // 8
    sum(1, 2)
}


