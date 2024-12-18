import 'dart:math';

import 'package:advent_of_dart/utils/point2d.dart';

const maxPath = 100 * 100;

class Node {
  bool corrupted = false;
  bool highlight = false;
  int shortestPath = maxPath;
  Node();
  @override
  String toString() {
    return highlight ? '@' : corrupted ? '#' : '.';
  }
}

mixin Day18 {
  int day18a(List<String> lines) {
    int result = 0;

    List<List<Node>> map = [];

    int size = int.parse(lines[0]);
    map.addAll(List.generate(size, (_) => List.generate(size, (_) => Node())));

    final fallenBytes =
        size == 7 ? 12 : 1024; // numbers provided by the part a prompt

    for (int i = 1; i <= fallenBytes; i++) {
      final byte = lines[i].split(',').map((e) => int.parse(e)).toList();
      map[byte[1]][byte[0]].corrupted = true;
    }

    for (final row in map) {
      print(row.join());
    }

    // find shortest path from starting point
    result = shortestPath(map, Point2D(0, 0), 0, '');

    return result;
  }

  int shortestPath(List<List<Node>> map, Point2D pos, int cost, String path) {
    // x and y for easy access
    final x = pos.x;
    final y = pos.y;

    if (x == map[y].length - 1 && y == map.length - 1) {
      // print('escape path, maybe (costs $cost): $path');
      return cost;
    }
    final n = map[y][x];
    if (cost >= n.shortestPath) return maxPath;
    n.shortestPath = cost;

    return minList([
      // try up, right, down, left
      (y > 0 && !map[y - 1][x].corrupted)
          ? shortestPath(map, Point2D(x, y - 1), cost + 1, '$path^')
          : maxPath,
      (x < map[y].length - 1 && !map[y][x + 1].corrupted)
          ? shortestPath(map, Point2D(x + 1, y), cost + 1, '$path>')
          : maxPath,
      (y < map.length - 1 && !map[y + 1][x].corrupted)
          ? shortestPath(map, Point2D(x, y + 1), cost + 1, '${path}v')
          : maxPath,
      (x > 0 && !map[y][x - 1].corrupted)
          ? shortestPath(map, Point2D(x - 1, y), cost + 1, '$path<')
          : maxPath,
    ]);
  }

  bool canEscape(List<List<Node>> map, Point2D pos, List<List<bool>> visited) {
    // x and y for easy access
    final x = pos.x;
    final y = pos.y;

    if (x == map[y].length - 1 && y == map.length - 1) { // made it!
      return true;
    }

    if (visited[y][x]) return false;
    visited[y][x] = true;

    return [
      // try up, right, down, left
      (y > 0 && !map[y - 1][x].corrupted)
          ? canEscape(map, Point2D(x, y - 1), visited)
          : false,
      (x < map[y].length - 1 && !map[y][x + 1].corrupted)
          ? canEscape(map, Point2D(x + 1, y), visited)
          : false,
      (y < map.length - 1 && !map[y + 1][x].corrupted)
          ? canEscape(map, Point2D(x, y + 1), visited)
          : false,
      (x > 0 && !map[y][x - 1].corrupted)
          ? canEscape(map, Point2D(x - 1, y), visited)
          : false,
    ].any((e) => e);
  }

  int minList(List<int> nums) {
    return nums.fold(maxPath, (prev, next) => min(prev, next));
  }

  int day18b(List<String> lines) {
    List<List<Node>> map = [];
    List<Point2D> bytes = [];

    int size = int.parse(lines[0]);
    map.addAll(List.generate(size, (_) => List.generate(size, (_) => Node())));

    final fallenBytes =
        size == 7 ? 12 : 1024; // starting number provided by the part a prompt

    lines.removeAt(0);
    for (int i = 0; i < lines.length; i++) {
      final byte = lines[i].split(',').map((e) => int.parse(e)).toList();
      bytes.add(Point2D(byte[0], byte[1]));
      if (i < fallenBytes) map[byte[1]][byte[0]].corrupted = true;
    }

    // let's try a sort of binary search
    int floor = fallenBytes;
    int ceil = lines.length - 1;
    int next = floor + ((ceil - floor) / 2).floor();
    corruptRange(map, bytes, floor, next);
    while (floor != ceil) {
      print('attempting byte $next (${bytes[next]}), floor = $floor, ceil = $ceil');
      List<List<bool>> visited = List.generate(size, (_) => List.filled(size, false));
      final pathExists = canEscape(map, Point2D(0, 0), visited);
      if (pathExists) {
        print('valid path found');
        if (ceil - next <= 1) break;
        // there is still a valid path. move the floor up
        floor = next;
        next = floor + ((ceil - floor) / 2).floor();
        corruptRange(map, bytes, floor, next);
      } else {
        print('no valid path');
        if (next - floor <= 1) break;
        // no valid path remains. move the ceil down
        ceil = next;
        next = floor + ((ceil - floor) / 2).floor();
        corruptRange(map, bytes, next+1, ceil, true);
      }
    }
    final byte = bytes[next];
    map[byte.y][byte.x].highlight = true;
    for (final row in map) {
      print(row.join());
    }
    print('finished search at $next (${bytes[next]})');

    return 0;
  }

  // corrupts bytes in range [from-to], inclusive
  void corruptRange(List<List<Node>> map, List<Point2D> bytes, int from, int to,
      [bool clear = false]) {
    if (to < from) {
      corruptRange(map, bytes, to, from, clear); // invert
      return;
    }

    for (int i = from; i <= to; i++) {
      final byte = bytes[i];
      map[byte.y][byte.x].corrupted = !clear;
    }
  }
}
