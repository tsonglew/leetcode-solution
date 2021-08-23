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
        return json.dumps(self.preorder(root))    

    def preorder(self, root):
        if not root:
            return [None]
        return [root.val, *self.preorder(root.left), *self.preorder(root.right)]

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        self.i = 0
        self.data = json.loads(data)
        return self.generate()
    
    def generate(self):
        if self.i >= len(self.data):
            return
        if self.data[self.i] is None:
            self.i += 1
            return
        n = TreeNode(self.data[self.i])
        self.i += 1
        n.left = self.generate()
        n.right = self.generate()
        return n
        
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
