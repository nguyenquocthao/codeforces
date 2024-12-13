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
    // initFAC()
    // repeat(1) {
    var (n, q) = readSlice<Int>()
    var health = readSlice<Int>().toIntArray()
    var strength = readSlice<Int>().toIntArray()
    var queries = Array(q) { readSlice<Int>() }
    // run(a,b,queries)

    var acc = IntArray(n+1)
    for ((i,v) in strength.withIndex()){
        acc[i+1]=acc[i]+v
    }
    var toadd = Array(n){mutableListOf<Int>()}
    for ((i,h) in health.withIndex()){
        for (k in 1 until K+1){
            var j = bsearch(0, i){idx -> acc[idx]>acc[i] - divceil(h, k)}
            if (j<i){
                toadd[j].add(i)
            }
        }
    }
    var toquery = Array(n){mutableListOf<Int>()}
    for ((i,qu) in queries.withIndex()){
        toquery[qu[0]-1].add(i)
    }
    var tree = FenwickTree(IntArray(n){1})
    var res = IntArray(q)
    var ml = IntArray(M)
    for (i in 0 until n){
        for(j in toadd[i]){
            tree.add(j, 1)
        }
        if (toquery[i].size==0){
            continue
        }
        var s = strength[i]
        for (j in 1 until M){
            if (i+j>=n){
                break
            }
            ml[j] = ml[j-1] + divceil(health[i+j], s)
            s+=strength[i+j]
        }
        for (qi in toquery[i]){
            var r = queries[qi][1]-1
            if (r-i < M){
                res[qi] = ml[r-i]
            } else {
                res[qi] = ml[M-1] + tree.rangeSum(i+M, r)
            }
        }
    }
    for(v in res){
        println(v)
    }


}


fun bsearch(lo: Int, hi: Int, f: (Int) -> Boolean): Int {
    var l = lo
    var r = hi - 1
    var res = hi
    while (l <= r) {
        var mid = (l + r) / 2
        if (f(mid)) {
            res = mid
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return res
}

class FenwickTree(a: IntArray) {
    var tree = intArrayOf(0) + a
    init {
        for (i in 1 until tree.size) {
            var j = i + (i and -i)
            if (j < tree.size) {
                tree[j] += tree[i]
            }
        }
    }
    fun add(i: Int, added: Int) {
        // a[i]+=added
        var i = i + 1
        while (i < tree.size) {
            tree[i] += added
            i += i and -i
        }
    }
    fun sum(i: Int): Int {
        var res = 0
        var i = i
        while (i > 0) {
            res += tree[i]
            i -= i and -i
        }
        return res
    }
    fun rangeSum(l: Int, r: Int): Int {
        // index (from 0) from l to r inclusive
        return sum(r + 1) - sum(l)
    }
}
