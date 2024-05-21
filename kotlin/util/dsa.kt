package myutil

import java.util.*


class Node(val lo: Int, val hi: Int, var left: Node?, var right: Node?, var maxv: Int, var minv: Int, var n: Int) {

    fun fix() {

    }



    fun decrease(i: Int, v: Int) {
        if (lo == i && hi == i) {
            maxv -= v
            minv -= v
        } else {
            if (i <= left!!.hi) {
                left!!.decrease(i, v)
            } else {
                right!!.decrease(i, v)
            }
            fix()
        }
    }
}

fun createNode(a: List<Int>): Node {
    // myPrint(81, *a.toTypedArray())
    val n = a.size
    fun create(lo: Int, hi: Int): Node {
        return if (lo == hi) {
            Node(lo, hi, null, null, a[lo], a[lo], 1)
        } else {
            val mid = (lo + hi) / 2
            val res = Node(lo, hi, create(lo, mid), create(mid + 1, hi), 0, 0, 0)
            res.fix()
            res
        }
    }
    return create(0, n - 1)
}
// for (ch in lo.toInt()..hi.toInt() step added){ // Use step to increment/decrement ch
//     if (!indexes.containsKey(ch.toChar())){
//         continue
//     }
//     res.addAll(indexes[ch.toChar()]!!)
// }


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



class FenwickTree(a:IntArray){
    var tree = intArrayOf(0) + a
    init{
        for (i in 1 until tree.size){
            var j = i + (i and -i)
            if (j<tree.size){
                tree[j] += tree[i]
            }
        }
    }
    fun add(i:Int, added:Int){
        // a[i]+=added
        var i= i+1
        while (i<tree.size){
            tree[i] += added
            i+=i and -i
        }
    }
    fun sum(i:Int):Int{
        var res = 0
        var i = i
        while (i>0){
            res += tree[i]
            i-=i and -i
        }
        return res
    }
    fun rangeSum(l:Int, r:Int):Int{
        // index (from 0) from l to r inclusive
        return sum(r+1) - sum(l)
    }

}