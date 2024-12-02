mixin Day01 {
  int day01a(List<String> lines) {
    int result = 0;
    List<int> lhs = [];
    List<int> rhs = [];

    // parse input lines into 2 arrays
    for (var line in lines) {
      final args = line.split(" ");
      lhs.add(int.parse(args[0]));
      rhs.add(int.parse(args[1]));
    }

    // sort the arrays
    lhs.sort();
    rhs.sort();

    // calculate distance for each line
    for (int i = 0; i < lhs.length; i++) {
      result += (lhs[i] - rhs[i]).abs();
    }

    return result;
  }

  int day01b(List<String> lines) {
    int result = 0;
    List<int> lhs = [];
    Map<int, int> rhsCount = {};

    // parse input lines into 2 arrays
    for (var line in lines) {
      final args = line.split(" ");
      lhs.add(int.parse(args[0]));
      final rhs = int.parse(args[1]);
      if (rhsCount.containsKey(rhs)) {
        rhsCount[rhs] = rhsCount[rhs]! + 1;
      } else {
        rhsCount[rhs] = 1;
      }
    }

    for (int left in lhs) {
      if (rhsCount.containsKey(left)) {
        result += left * rhsCount[left]!;
      }
    }

    return result;
  }
}
