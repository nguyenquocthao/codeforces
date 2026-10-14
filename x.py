# Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

# The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

# The number of nodes in the tree is in the range [0, 1000].
# -109 <= Node.val <= 109
# -1000 <= targetSum <= 1000

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def pathSum(self, root: TreeNode | None, targetSum: int) -> int:
        res, m = 0, defaultdict(int)
        m[0]=1
        def dp(node, s):
            nonlocal res
            if not node: return
            s += node.val
            res += m[s-targetSum]
            m[s] += 1
            dp(node.left, s)
            dp(node.right, s)
            m[s] -= 1
        dp(root,0)
        return res
            