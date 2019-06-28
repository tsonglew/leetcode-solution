# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
import copy

class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> int:
        self.cnt = 0
        def sub_sum(root: TreeNode, sum: int):
            if root is None:
                return []
            could_sum_to = []
            left_sub_sum = sub_sum(root.left, sum)
            right_sub_sum = sub_sum(root.right, sum)
            if root.val == sum:
                self.cnt += 1
            for i in left_sub_sum+right_sub_sum:
                if i == sum - root.val:
                    self.cnt += 1
                could_sum_to.append(root.val+i)
            could_sum_to.append(root.val)
            return could_sum_to
        sub_sum(root, sum)
        return self.cnt