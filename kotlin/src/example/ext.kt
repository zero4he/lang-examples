/**
 * @author Zero
 * Created on 2017/8/9.
 */
object Utils {

    //java invoke: Utils.INSTANCE.sayHi("8888");
    fun sayHi(msg: String) {
        print(msg)
    }


}

//可以写在别的文件中
//java 中无法通过 Utils.INSTANCE.sayHello("hello") 调用
fun Utils.sayHello(msg: String): String {
    println(msg)
    return msg
}


fun main(args: Array<String>) {
    Utils.sayHi("---")
    Utils.sayHello("hello")

}