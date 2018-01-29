/**
 * @author Zero
 * Created on 2017/7/5.
 */


fun main(args: Array<String>) {

    val t = Thread(Runnable {
        print("Hello: "+ Thread.currentThread().name)
    })

    t.start()

    val t2 = Thread({
        print("Hello: "+ Thread.currentThread().name)
    })

    t2.start()



}