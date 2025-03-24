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
        var (n,k) = readSlice<Int>()
        myPrint(*run(n,k).toTypedArray())

    }


}

fun run(n:Int, k:Int): IntArray{
    fun dp(start:Int, k2:Int):IntArray{
        if (k2==0 || start==n){
            return intArrayOf(-1)
        }
        if (k2==1){
            return (start..n).toList().toIntArray()
        }
        var nele = n-start+1
        var start2 = start + (nele/2)
        var x = dp(start2, k2-1)
        if (x[0]==-1){
            return x
        }
        var y = (start..(start2-1)).toList().toIntArray()
        var l = mutableListOf<Int>()
        for ((i,v) in x.withIndex()){
            l.add(v)
            if (i<y.size){
                l.add(y[i])
            }
        }
        return l.toIntArray()
    }
    return dp(1, k)
}