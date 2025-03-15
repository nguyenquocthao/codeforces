import java.util.*
import kotlin.math.*
import kotlin.text.trim
import divceil
import myPrint
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
        var c = readLine()!!.last()
        println(run(readLine()!!, c))

    }


}

fun run(s:String, c: Char):Int{
    if (c=='g'){
        return 0
    }
    var res=0
    var firstg = -1
    var firstc = -1
    for ((i,ch) in s.withIndex()){
        if (ch=='g'){
            if (firstg<0){
                firstg = i
            }
            if (firstc>=0){
                res=max(res, i-firstc)
                firstc = -1
            }
        } else if (ch==c && firstc<0){
            firstc = i
        }
    }
    if (firstc>=0){
        res = max(res, firstg + s.length - firstc)
    }
    return res

}