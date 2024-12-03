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
    // initFAC()
    repeat(readInt()) {
        readInt()
        println(run(readSlice<Long>().toLongArray()))
    }
}

// const val MOD = 998244353L

const val MOD = 1000000007L
// const val maxn = 1000000
const val maxn = 400000

var FAC = LongArray(maxn+1){0}
var IFAC = LongArray(maxn+1){0}

fun initFAC() {
    FAC[0] = 1
    FAC[1] = 1
    IFAC[0]= 1
    IFAC[1]=1
    for (i in 2 until maxn+1){
        FAC[i] = (i * FAC[i-1])%MOD
        IFAC[i] = mod_inverse(FAC[i])
    }
}

fun pow(x:Long, n:Long):Long {
	var x = x % MOD
	var res = 1L
    var n = n
	while (n > 0 ){
		if (n%2 == 1L ){
			res = (res * x) % MOD
		}
		x = (x * x) % MOD
		n = n / 2
	}
	return res
}

fun mod_inverse(x:Long):Long {
	return pow(x, MOD-2)
}

fun comb(n:Long, k:Long):Long {
	if (n < 0 || k > n ){
		return 0
	}
	val inv = mod(IFAC[k.toInt()] * IFAC[(n-k).toInt()])
	return (FAC[n.toInt()] * inv) % MOD
}

fun mod(v:Long ) :Long {
	var res = v % MOD
	if (res < 0 ){
		res += MOD
	}
	return res
}


fun run(a:LongArray):Long{
    var n = a.size
    fun cal(i:Int, j:Int):Long{
        var s = 0L
        for (k in i until j+1){
            s+=2*a[k]
        }
        if ((j-i+1)%2==0){
            return s
        }
        var maxv = 0L
        for (k in i until j+1 step 2){
            maxv = maxOf(maxv, a[k])
        }
        return s- maxv
    }
    var res = 0L
    var start = 0
    for (i in 0 until n){
        if (a[i]==0L){
            res+=cal(start, i-1)
            start=i+1
        }
    }
    res += cal(start, n-1)
    return res
}