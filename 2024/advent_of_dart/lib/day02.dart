mixin Day02 {
  int day02a(List<String> lines) {
    int result = 0;

    for (String line in lines) {
      final levels = line.split(" ").map((level) => int.parse(level)).toList();
      if (isSafe(levels) == -1) {
        result++;
      }
    }

    return result;
  }

  int day02b(List<String> lines) {
    int result = 0;

    for (String line in lines) {
      final levels = line.split(" ").map((level) => int.parse(level)).toList();
      final safeCheck = isSafe(levels);
      if (safeCheck == -1) {
        result++;
      } else {
        List<int> retryOne = List.from(levels);
        retryOne.removeAt(safeCheck - 1);
        if (isSafe(retryOne) == -1) {
          result++;
        } else {
          List<int> retryTwo = List.from(levels);
          retryTwo.removeAt(safeCheck);
          if (isSafe(retryTwo) == -1) {
            result++;
          }
        }
      }
    }

    return result;
  }

  // Returns the index of the level that first failed the safety check, or -1 if safe.
  int isSafe(List<int> line) {
    int prev = line[0];
    bool increasing = false;
    if (line[1] > prev) {
      increasing = true;
    }
    for (int i = 1; i < line.length; i++) {
      final diff = (line[i] - prev).abs();
      if (1 > diff || diff > 3) {
        return i;
      }
      if (increasing && line[i] < prev || !increasing && line[i] > prev) {
        return i;
      }
      prev = line[i];
    }
    return -1;
  }
}
