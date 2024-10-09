import java.util.*
import kotlin.math.*

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

fun main() {
    repeat(readInt()) {
        var (n,p) = readSlice<Int>()
        println(run(p, readSlice<Int>().toIntArray()))
    }
}

fun run(p:Int, a:IntArray): Int{
    val n = a.size
    a.sort()
    fun dp(i: Int, p: Int, ngreen: Int, nblue: Int): Int{
        if (i==n){
            return i
        }
        var p = p
        for (j in i until n){
            if (p>a[j]){
                p += a[j]/2
            } else {
                var res = j
                if (ngreen>0){
                    res = maxOf(res, dp(j, p*2, ngreen-1, nblue))
                }
                if (nblue>0){
                    res = maxOf(res, dp(j, p*3, ngreen, nblue-1))
                }
                return res
            }
        }
        return n
    }
    return dp(0,p,2,1)
}