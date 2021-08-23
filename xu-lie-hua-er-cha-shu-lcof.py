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
        r = {}
        self.nums = defaultdict(lambda: 0)
        r['pre'] = self.preorder(root)
        self.nums = defaultdict(lambda: 0)
        r['in'] = self.inorder(root)
        r = json.dumps(r)
        return r
        

    def preorder(self, root):
        if not root:
            return []
        v = root.val
        self.nums[root.val] += 1
        if self.nums[root.val] > 1:
            v = f"{root.val}|{self.nums[root.val]}"
        return [v, *self.preorder(root.left), *self.preorder(root.right)]
    
    def inorder(self, root):
        if not root:
            return []
        v = root.val
        self.nums[root.val] += 1
        if self.nums[root.val] > 1:
            v = f"{root.val}|{self.nums[root.val]}"
        return [*self.inorder(root.left), v, *self.inorder(root.right)]


    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        inorder = json.loads(data)['in']
        preorder = json.loads(data)['pre']
        return self.generate(inorder, preorder)

    def generate(self, inorder, preorder):
        if len(inorder) == 0:
            return None
        v = preorder[0]
        inorder_idx = inorder.index(v)
        n = TreeNode(v) if isinstance(v, int) else TreeNode(int(v.split('|')[0]))
        n.left = self.generate(inorder[:inorder_idx], preorder[1:inorder_idx+1])
        n.right = self.generate(inorder[inorder_idx+1:], preorder[inorder_idx+1:])
        return n
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
