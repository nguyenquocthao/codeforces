import java.util.*
import kotlin.math.*
import kotlin.text.trim
import kotlin.comparisons.maxOf
// import divceil
// import myPrint
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
        var (n,k) = readSlice<Long>()
        println(run(readSlice<Long>(), k))
        // var (n,k) = readSlice<Int>()
        // myPrint(*run(n,k).toTypedArray())

    }


}


fun run(a: List<Long>, power: Long): Int{
    // println(a)
    // println(power)
    var a = a.sorted()
    fun dp(i:Int, p:Long, ngreen: Int, nblue: Int):Int{
        var p = p
        var i = i
        while (i<a.size && p>a[i]){
            p+=a[i]/2
            i+=1
        }
        if (i>=a.size){
            return i
        }
        var res = i
        if (ngreen>0){
            res = maxOf(res, dp(i, p*2, ngreen-1, nblue))
        }
        if (nblue>0){
            res=maxOf(res, dp(i, p*3, ngreen, nblue-1))
        }
        return res
    }
    return dp(0,power,2,1)
}
