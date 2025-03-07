# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        v = self.preorder(root)
        print(v)
        return ",".join(v)

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        if not data:
            return None
        l = data.split(",")
        return self.reveal(l)

    def reveal(self, l):
        v = l.pop(0)
        if v == "n":
            return None
        root = TreeNode(v)
        root.left = self.reveal(l)
        root.right = self.reveal(l)
        return root

    def preorder(self, root):
        if not root:
            return ["n"]
        return [str(root.val)] + self.preorder(root.left) + self.preorder(root.right)


# Your Codec object will be instantiated and called as such:
# ser = Codec()
# deser = Codec()
# ans = deser.deserialize(ser.serialize(root))
