class User(val name: String) {
    override fun toString(): String {
        return name
    }

}

object UserRegister {

    val register: MutableList<User> = mutableListOf()

    fun register(user: User) {
        register.add(user)
    }

    fun list(): List<User> {
        return register
    }

}


fun example1() {
    UserRegister.register(User("Zero"))
    UserRegister.register(User("May"))
    UserRegister.register(User("Frank"))
    UserRegister.register(User("JY"))
    UserRegister.register(User("Cindy"))

    println(UserRegister.list())
}


///静态方法

// 工厂模式
class Car {
    companion object Factory {
        fun create(): Car = Car()
    }
}


///单纯的静态方法

class Dog(val name: String) {

    fun eat(food: Any): Unit {
        println("$name eat food : $food")
    }

    companion object {
        fun create(name: String): Dog = Dog(name)

        fun kill(dog: Dog): Boolean {
            println("kill ${dog.name}")
            return true
        }
    }
}

fun example2() {
    val car1 = Car.Factory.create()
    val car2 = Car.create()


    val dog = Dog.create("小狗")
    dog.eat("feces")
    Dog.kill(dog)
}


///////////////////内嵌类///////////////////////
class Outer1 {
    private val bar: Int = 1

    class Nested { //不加inner声明, 就像java中的static class
        fun foo() {
            //不能访问bar
        }
    }
}


class Outer2 {
    private val bar: Int = 1

    inner class Nested { //使用inner声明, 等价java中默认的class
        fun foo() {
            bar //可以访问bar
        }
    }
}

fun example3() {
    Outer1.Nested().foo() //不需要创建Outer1对象
    Outer2().Nested().foo() //需要创建Outer2对象
}


