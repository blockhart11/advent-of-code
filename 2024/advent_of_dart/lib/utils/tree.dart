class TreeNode<T> {
  T value;
  List<TreeNode<T>> children;

  TreeNode(this.value, [List<TreeNode<T>>? children])
      : children = children ?? [];

  void addChild(TreeNode<T> child) {
    children.add(child);
  }
}

void printTree(TreeNode node, [String prefix = '', bool isLast = true]) {
  print('$prefix${isLast ? '\'-' : '--'} ${node.value}');
  for (int i = 0; i < node.children.length; i++) {
    printTree(node.children[i], '$prefix${isLast ? '  ' : '| '}', i == node.children.length - 1);
  }
}