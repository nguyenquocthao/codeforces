package myutil

import java.util.*
import myutil.MOD

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

fun bitLength(v:Int):Int{
    if (v==0){
        return 0
    }
    return 32 - v.countLeadingZeroBits()
}

fun bitLength64(v:Long):Int{
    if (v==0L){
        return 0
    }
    return 64 - v.countLeadingZeroBits()
}


fun modpos(a:Int,b:Int):Int{
    var x = a%b
    if (x<0){
        return x+b
    }
    return x
}

fun divceil(a:Int, b:Int):Int{
    if (a%b==0){
        return a/b
    } else{
        return a/b + 1
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

fun bsearch(lo:Int, hi:Int, f:(Int)->Boolean):Int{
    var l = lo
    var r = hi-1
    var res = hi
    while (l<=r){
        var mid = (l+r)/2
        if (f(mid)){
            res=mid
            r=mid-1
        } else {
            l=mid+1
        }
    }
    return res
}