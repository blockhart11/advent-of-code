mixin Day10 {
  int day10a(List<String> lines) {
    int result = 0;

    List<List<(int, int, List<int>)>> map = lines
        .map((line) => line.split('').map((e) => (int.parse(e), 0, <int>[])).toList())
        .toList();

    int nineIdx = 0;

    for (int y = 0; y < map.length; y++) {
      for (int x = 0; x < map[y].length; x++) {
        if (map[y][x].$1 == 9) {
          // work backwards from 9
          map[y][x] = (9, 1, [nineIdx]);
          visitNeighborsDistinct(map, x, y, nineIdx);
          // print('visited 9 #$nineIdx at ($x, $y):');
          // for (final line in map) {
          //   print(line);
          // }
          nineIdx++;
        }
      }
    }

    for (final row in map) {
      for (final col in row) {
        if (col.$1 == 0) {
          result += col.$2;
        }
      }
    }

    return result;
  }

  int day10b(List<String> lines) {
    int result = 0;

    List<List<(int, int)>> map = lines
        .map((line) => line.split('').map((e) => (int.parse(e), 0)).toList())
        .toList();

    for (int y = 0; y < map.length; y++) {
      for (int x = 0; x < map[y].length; x++) {
        if (map[y][x].$1 == 9) {
          // work backwards from 9
          map[y][x] = (9, 1);
          visitNeighbors(map, x, y);
          // print('visited 9 #$nineIdx at ($x, $y):');
          // for (final line in map) {
          //   print(line);
          // }
        }
      }
    }

    for (final row in map) {
      for (final col in row) {
        if (col.$1 == 0) {
          result += col.$2;
        }
      }
    }

    return result;
  }

  void visitNeighborsDistinct(List<List<(int, int, List<int>)>> map, int x, int y, int idx) {
    final current = map[y][x].$1;
    // up
    if (y > 0) {
      final neighbor = map[y - 1][x];
      if (neighbor.$1 == current - 1 && !neighbor.$3.contains(idx)) {
        map[y - 1][x] = (neighbor.$1, neighbor.$2 + 1, neighbor.$3..add(idx));
        visitNeighborsDistinct(map, x, y - 1, idx);
      }
    }
    // right
    if (x < map[y].length - 1) {
      final neighbor = map[y][x + 1];
      if (neighbor.$1 == current - 1 && !neighbor.$3.contains(idx)) {
        map[y][x + 1] = (neighbor.$1, neighbor.$2 + 1, neighbor.$3..add(idx));
        visitNeighborsDistinct(map, x + 1, y, idx);
      }
    }
    // down
    if (y < map.length - 1) {
      final neighbor = map[y + 1][x];
      if (neighbor.$1 == current - 1 && !neighbor.$3.contains(idx)) {
        map[y + 1][x] = (neighbor.$1, neighbor.$2 + 1, neighbor.$3..add(idx));
        visitNeighborsDistinct(map, x, y + 1, idx);
      }
    }
    // left
    if (x > 0) {
      final neighbor = map[y][x - 1];
      if (neighbor.$1 == current - 1 && !neighbor.$3.contains(idx)) {
        map[y][x - 1] = (neighbor.$1, neighbor.$2 + 1, neighbor.$3..add(idx));
        visitNeighborsDistinct(map, x - 1, y, idx);
      }
    }
  }

  void visitNeighbors(List<List<(int, int)>> map, int x, int y) {
    final current = map[y][x].$1;
    // up
    if (y > 0) {
      final neighbor = map[y - 1][x];
      if (neighbor.$1 == current - 1) {
        map[y - 1][x] = (neighbor.$1, neighbor.$2 + 1);
        visitNeighbors(map, x, y - 1);
      }
    }
    // right
    if (x < map[y].length - 1) {
      final neighbor = map[y][x + 1];
      if (neighbor.$1 == current - 1) {
        map[y][x + 1] = (neighbor.$1, neighbor.$2 + 1);
        visitNeighbors(map, x + 1, y);
      }
    }
    // down
    if (y < map.length - 1) {
      final neighbor = map[y + 1][x];
      if (neighbor.$1 == current - 1) {
        map[y + 1][x] = (neighbor.$1, neighbor.$2 + 1);
        visitNeighbors(map, x, y + 1);
      }
    }
    // left
    if (x > 0) {
      final neighbor = map[y][x - 1];
      if (neighbor.$1 == current - 1) {
        map[y][x - 1] = (neighbor.$1, neighbor.$2 + 1);
        visitNeighbors(map, x - 1, y);
      }
    }
  }
}
