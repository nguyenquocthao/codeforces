import java.util.*
import kotlin.math.*
import divceil
// import bitLength64


inline fun <reified T> readSlice(): List<T> {
    return readLine()!!.split(" ").map { parseToken<T>(it) }
}

inline fun <reified T> parseToken(token: String): T {
    return when (T::class) {
        Int::class -> token.toInt() as T
        Long::class -> token.toLong() as T
        Double::class -> token.toDouble() as T
        String::class -> token as T
        else -> throw IllegalArgumentException("Unsupported type")
    }
}

fun readInt(): Int {
    return readLine()!!.toInt()
}

fun myPrint(vararg args: Any?) {
    println(args.joinToString(" "))
}

fun divceil(a: Int, b: Int): Int {
    if (a % b == 0) {
        return a / b
    } else {
        return a / b + 1
    }
}

val M = 1000
// val M = 1
val K = 10

fun main() {
    repeat(readInt()){
        readInt()
        println(run(readSlice<Int>().toIntArray()))
        // println(KotlinVersion.CURRENT)

    }


}

fun run(a:IntArray):Int{
    if (a[0]!=a[1] && a[0]!=a[2]){
        return 1
    }
    for (i in 0 until a.size){
        if (a[i]!=a[0]){
            return i+1
        }
    }
    throw Exception("not found")
}