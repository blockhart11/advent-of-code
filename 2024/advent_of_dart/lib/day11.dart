mixin Day11 {
  int day11a(List<String> lines) {
    int result = 0;

    List<String> stones = lines[0].split(' ');
    Map<String, Map<int, int>> cache = {}; // key is the stone id, value is cached map of blinksLeft and solution

    for (final stone in stones) {
      result += day11aHelper(stone, 25, cache);
    }

    return result;
  }

  int day11b(List<String> lines) {
    int result = 0;

    List<String> stones = lines[0].split(' ');
    Map<String, Map<int, int>> cache = {}; // key is the stone id, value is cached map of blinksLeft and solution

    for (final stone in stones) {
      result += day11aHelper(stone, 75, cache);
    }

    return result;
  }

  int day11aHelper(String stone, int blinksLeft, Map<String, Map<int, int>> cache) {
    if (blinksLeft == 0) {
      return 1; // base case
    }

    // check cache
    if (!cache.containsKey(stone)) {
      cache[stone] = {}; // initialize new map entry
    } else if (cache[stone]!.containsKey(blinksLeft)) {
      return cache[stone]![blinksLeft]!; // cache hit!
    }

    int result = 0;

    if (int.parse(stone) == 0) { // use int to remove padding zeroes
      result = day11aHelper('1', blinksLeft - 1, cache);
      cache[stone]![blinksLeft] = result;
    } else if (stone.length.isEven) {
      String lhs = stone.substring(0, stone.length ~/ 2);
      String rhs = stone.substring(stone.length ~/ 2).replaceFirst(RegExp(r'^0+'), '');
      if (rhs == '') rhs = '0';

      result =
        day11aHelper(lhs, blinksLeft - 1, cache)
      + day11aHelper(rhs, blinksLeft - 1, cache);
      cache[stone]![blinksLeft] = result;
    } else {
      result = day11aHelper((int.parse(stone) * 2024).toString(), blinksLeft - 1, cache);
      cache[stone]![blinksLeft] = result;
    }

    return result;
  }

}
