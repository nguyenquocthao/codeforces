import java.util.*
import kotlin.math.*
import divceil

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
        var (k,m) =readSlice<Int>()
        m%=(3*k)
        println(max(0, 2*k -m ))

    }


}

class Tree(p: List<Int>, mask:Long){
    var mparent: Array<Int>
    var mchild: Array<MutableList<Int>>
    init{
        val n = p.size
        mparent = Array(n){0}; mchild = Array(n){mutableListOf<Int>()}; var mtrue = Array(n){it}
        fun gettrue(i:Int):Int{
            if (i==0){
                return 0
            }
            if ((mask shr mtrue[i]) and 1L == 0L){
                mtrue[i] = gettrue(p[i])
            }
            return mtrue[i]
        }
        for ((i,v) in p.withIndex()){
            if (i==0 || (mask shr i) and 1L == 0L){
                continue
            }
            val p = gettrue(v)
            mparent[i] = p
            mchild[p].add(i)
        }
    }

    fun getde(x:Int):Long{
        var res = 0L
        for (j in mchild[x]){
            res = res or (1L shl j) or getde(j)
        }
        return res
    }
    fun getan(x:Int):Long{
        var res = (1L shl mparent[x])
        if (res==1L){
            return 0L
        }
        return res or getan(mparent[x])
    }
}