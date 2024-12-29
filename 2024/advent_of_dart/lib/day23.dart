/*
 * Note about day 23:
 * I modified the input file by making each entry so that the lhs < rhs alphabetically (swapping if necessary), then sorting the entire file alphabetically.
 * This allows me to make assumptions while traversing the file that make it much simpler to solve the problem.
 * I have since deleted the code that pre-processed the inputs since it only needed to be run once.
 */

mixin Day23 {
  int day23a(List<String> lines) {
    int result = 0;

    Map<int, Set<int>> network = {};
    for (final line in lines) {
      final a = nodeToInt(line.substring(0,2));
      final b = nodeToInt(line.substring(3));
      network.update(a, (e) => e..add(b), ifAbsent: () => {b});
    }

    List<List<int>> sets = [];
    for (final nodeA in network.keys) {
      for (final nodeB in network[nodeA]!) {
        if (!network.containsKey(nodeB)) continue;
        for (final nodeC in network[nodeA]!.intersection(network[nodeB]!)) {
          sets.add([nodeA, nodeB, nodeC]);
        }
      }
    }

    for (final set in sets) {
      if (set.any((e) => intToNode(e).startsWith('t'))) {
        result++;
      }
    }
    print('found ${sets.length} total sets');

    return result;
  }

  int day23b(List<String> lines) {
    Map<int, Set<int>> network = {};
    for (final line in lines) {
      final a = nodeToInt(line.substring(0,2));
      final b = nodeToInt(line.substring(3));
      network.update(a, (e) => e..add(b), ifAbsent: () => {b});
    }

    // goal is to find the largest network where every node
    // is connected to every other node
    
    Set<int> current = {};
    for (final nodeA in network.entries) {
      for (final nodeB in network[nodeA.key]!) {
        final largest = largestNetwork(nodeB, {nodeA.key}, network);
        if (largest.length > current.length) {
          print('found new largest network: ${largest.map((e) => intToNode(e)).join(',')}');
          current = largest;
        }
      }
    }

    print('largest network: ${current.map((e) => intToNode(e)).join(',')}');
    return -1;
  }

  Set<int> largestNetwork(int node, Set<int> containingAll, Map<int, Set<int>> network) {
    for (final priorNode in containingAll) { // current node must exist in all prior nodes
      if (!network.containsKey(priorNode) || !network[priorNode]!.contains(node)) {
        return containingAll;
      }
    }

    Set<int> result = containingAll..add(node);
    if (!network.containsKey(node)) return result; // base case - no more nodes to check

    for (final nextNode in network[node]!) {
      final next = largestNetwork(nextNode, containingAll..add(node), network);
      if (next.length > result.length) result = next;
    }
    return result;
  }

  int nodeToInt(String node) {
    if (node.length != 2) throw('node is not 2 characters: $node');
    return node.codeUnits.reduce((value, element) => value * 256 + element);
  }

  String intToNode(int value) {
    List<int> codeUnits = [];
    while (value > 0) {
      codeUnits.insert(0, value % 256);
      value = value ~/ 256;
    }
    String result = String.fromCharCodes(codeUnits);
    if (result.length != 2) throw('$value is not 2 characters: $result');
    return result;
  }
}
    