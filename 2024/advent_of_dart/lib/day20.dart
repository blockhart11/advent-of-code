import 'dart:math';

import 'package:advent_of_dart/utils/direction.dart';
import 'package:advent_of_dart/utils/point2d.dart';

class Node {
  int time = -1;
  Point2D pos;
  bool wall = false;
  Node(this.pos);

  int distance(Node other) {
    return pos.distance(other.pos);
  }
}

mixin Day20 {
  int day20a(List<String> lines) {
    int result = 0;

    List<List<Node>> map = [];
    Point2D start = Point2D();
    Point2D end = Point2D();
    print('loading map...');
    for (int y = 0; y < lines.length; y++) {
      map.add([]);
      for (int x = 0; x < lines[y].length; x++) {
        Node n = Node(Point2D(x, y));
        switch (lines[y][x]) {
          case '#':
            n.wall = true;
          case 'S':
            start.set(x, y);
          case 'E':
            end.set(x, y);
          case '.': // do nothing
            break;
          default:
            throw ('unexpected character in map: ${lines[y][x]}');
        }
        map[y].add(n);
      }
    }

    // find initial timing for each node in the track
    print('running race track...');
    Node cur = map[start.y][start.x];
    Node next = Node(Point2D());
    int time = 0;
    map[cur.pos.y][cur.pos.x].time = time++;
    while (!cur.pos.equals(end)) {
      final y = cur.pos.y;
      final x = cur.pos.x;
      // try to move up, right, down, and left to find the next piece of track
      if (y > 0 && !map[y - 1][x].wall && map[y - 1][x].time == -1) {
        next = map[y - 1][x];
      } else if (x < map[y].length - 1 &&
          !map[y][x + 1].wall &&
          map[y][x + 1].time == -1) {
        next = map[y][x + 1];
      } else if (y < map.length - 1 &&
          !map[y + 1][x].wall &&
          map[y + 1][x].time == -1) {
        next = map[y + 1][x];
      } else if (x > 0 && !map[y][x - 1].wall && map[y][x - 1].time == -1) {
        next = map[y][x - 1];
      } else {
        throw ('stuck at ${map[y][x].pos}!');
      }
      next.time = time++;
      cur = next;
    }

    // test all shortcuts
    print('track takes $time to complete normally. testing shortcuts...');
    Map<int, int> shortcuts = {};
    for (int y = 1; y < map.length - 1; y++) {
      for (int x = 1; x < map[y].length - 1; x++) {
        cur = map[y][x];
        if (cur.wall) continue;
        // check for shortcuts up, right, down, and left
        if (y > 2 && map[y - 1][x].wall) {
          Node next = map[y - 2][x];
          time = next.time - cur.time - 2;
          if (!next.wall && time > 0) {
            // shortcut found
            print('saved $time at ${cur.pos} going up');
            shortcuts.update(time, (t) => t + 1, ifAbsent: () => 1);
          }
        }
        if (x < map[y].length - 3 && map[y][x + 1].wall) {
          Node next = map[y][x + 2];
          time = next.time - cur.time - 2;
          if (!next.wall && time > 0) {
            // shortcut found
            print('saved $time at ${cur.pos} going right');
            shortcuts.update(time, (t) => t + 1, ifAbsent: () => 1);
          }
        }
        if (y < map.length - 3 && map[y + 1][x].wall) {
          Node next = map[y + 2][x];
          time = next.time - cur.time - 2;
          if (!next.wall && time > 0) {
            // shortcut found
            print('saved $time at ${cur.pos} going down');
            shortcuts.update(time, (t) => t + 1, ifAbsent: () => 1);
          }
        }
        if (x > 2 && map[y][x - 1].wall) {
          Node next = map[y][x - 2];
          time = next.time - cur.time - 2;
          if (!next.wall && time > 0) {
            // shortcut found
            print('saved $time at ${cur.pos} going left');
            shortcuts.update(time, (t) => t + 1, ifAbsent: () => 1);
          }
        }
      }
    }

    print(shortcuts);
    // count the number of cheats saving 100 or more
    for (final key in shortcuts.keys) {
      if (key >= 100) result += shortcuts[key]!;
    }

    return result;
  }

  int day20b(List<String> lines) {
    int result = 0;

    List<List<Node>> map = [];
    Point2D start = Point2D();
    Point2D end = Point2D();
    print('loading map...');
    for (int y = 0; y < lines.length; y++) {
      map.add([]);
      for (int x = 0; x < lines[y].length; x++) {
        Node n = Node(Point2D(x, y));
        switch (lines[y][x]) {
          case '#':
            n.wall = true;
          case 'S':
            start.set(x, y);
          case 'E':
            end.set(x, y);
          case '.': // do nothing
            break;
          default:
            throw ('unexpected character in map: ${lines[y][x]}');
        }
        map[y].add(n);
      }
    }

    // find initial timing for each node in the track
    print('running race track...');
    List<Dir2D> path = [];
    Node cur = map[start.y][start.x];
    Node next = Node(Point2D());
    int time = 0;
    map[cur.pos.y][cur.pos.x].time = time++;
    while (!cur.pos.equals(end)) {
      final y = cur.pos.y;
      final x = cur.pos.x;
      // try to move up, right, down, and left to find the next piece of track
      if (y > 0 && !map[y - 1][x].wall && map[y - 1][x].time == -1) {
        next = map[y - 1][x];
        path.add(Dir2D.up);
      } else if (x < map[y].length - 1 &&
          !map[y][x + 1].wall &&
          map[y][x + 1].time == -1) {
        next = map[y][x + 1];
        path.add(Dir2D.right);
      } else if (y < map.length - 1 &&
          !map[y + 1][x].wall &&
          map[y + 1][x].time == -1) {
        next = map[y + 1][x];
        path.add(Dir2D.down);
      } else if (x > 0 && !map[y][x - 1].wall && map[y][x - 1].time == -1) {
        next = map[y][x - 1];
        path.add(Dir2D.left);
      } else {
        throw ('stuck at ${map[y][x].pos}!');
      }
      next.time = time++;
      cur = next;
    }

    // run the track, test the shortcuts
    print('track takes $time to complete normally. testing shortcuts...');
    cur = map[start.y][start.x];
    Map<int, int> shortcutMap = {};
    int progress = 1;
    for (final dir in path) {
      print('testing node ${progress++} of ${path.length} at ${cur.pos}');
      final y = cur.pos.y;
      final x = cur.pos.x;
      // calculate all shortcuts up to distance 20
      final yFloor = max(0, y-20);
      final yCeil = min(map.length-1, y+20);
      // this is gonna be slow...
      List<Point2D> shortcuts = [];
      for (int testY = yFloor; testY <= yCeil; testY++) {
        final xFloor = max(0, (x-20)+(y-testY).abs());
        final xCeil = min(map[y].length-1, (x+20)-(y-testY).abs());
        for (int testX = xFloor; testX <= xCeil; testX++) {
          next = map[testY][testX];
          final dist = cur.distance(next);
          time = next.time - cur.time - dist;
          if (
            dist <= 20
            && !next.wall
            && time >= 100
            // && time >= 50 // test input
          ) {
            // print('found shortcut from ${cur.pos} to ${next.pos}, saves $time');
            shortcuts.add(next.pos);
            shortcutMap.update(time, (t) => t + 1, ifAbsent: () => 1);
          }
        }
      }
      if (shortcuts.isNotEmpty) {
        // print('found ${shortcuts.length} from ${cur.pos}: $shortcuts');
        // print('found ${shortcuts.length} from ${cur.pos}');
        result += shortcuts.length;
      }

      // move to next node
      switch (dir) {
        case Dir2D.up:
          cur = map[y - 1][x];
        case Dir2D.right:
          cur = map[y][x + 1];
        case Dir2D.down:
          cur = map[y + 1][x];
        case Dir2D.left:
          cur = map[y][x - 1];
      }
    }

    print(shortcutMap);

    return result;
    // higher than  0978833
    // 993178?
    // lower than   1106272
  }
}
