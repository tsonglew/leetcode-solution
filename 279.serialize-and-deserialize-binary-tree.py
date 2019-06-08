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
        if root is None:
            return 'n'
        return ','.join([str(root.val), self.serialize(root.left), self.serialize(root.right)]) 
        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        self.value_list = data.split(',')
        return self.deserialize_from_list()
    
    def deserialize_from_list(self):
        if len(self.value_list) == 0:
            return
        v = self.value_list.pop(0)
        if v == 'n':
            return None
        new_node = TreeNode(int(v))
        new_node.left = self.deserialize_from_list()
        new_node.right = self.deserialize_from_list()
        return new_node
        
        
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))