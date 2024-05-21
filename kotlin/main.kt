import java.util.*
import kotlin.collections.mutableListOf
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

// 1 0 0 0
// 0 0 0 1


fun main(){
    initFAC()
    var (n,k) = readSlice<Long>()
    val res = LongArray(n.toInt()+1){0L}
    res[0] = comb(2*n-2, k)
    for (i in 1 until n.toInt()){
        res[i] = mod(2*comb(n-i-1, k-1) + (n-i-1)*comb(n-i-2, k-2))
        res[0] = mod(res[0]-res[i])
        if (res[i]==0L){
            break
        }
    }
    println(res.joinToString(" "))




}