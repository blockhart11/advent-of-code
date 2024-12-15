import "dart:io";

class Point2D {
  int x, y;
  Point2D([this.x = -1, this.y = -1]);
  void set(int x, int y) {
    this.x = x;
    this.y = y;
  }
}

mixin Day15 {
  int day15a(List<String> lines) {
    int result = 0;

// parse input
    List<List<String>> map = [];
    List<String> ops = [];
    Point2D robot = Point2D();
    int i = 0;
    while (lines[i].isNotEmpty) {
      map.add(lines[i].split(''));
      final robotPos = lines[i].indexOf('@');
      if (robotPos >= 0) {
        map[i][robotPos] = '.'; // no reason to keep the robot on the map
        robot.set(i, robotPos);
      }

      i++;
    }
    for (i = i + 1; i < lines.length; i++) {
      ops.addAll(lines[i].split(''));
    }

    for (String op in ops) {
      // for (final y in map) {
      //   print(y.join());
      // }
      // print('robot is at ${robot.x}, ${robot.y}, now moving $op');
      bool pushing = false;
      bool done = false;
      switch (op) {
        case '^': // up
          for (int y = robot.y - 1; !done && y >= 0; y--) {
            int x = robot.x;
            switch (map[y][x]) {
              case '.':
                if (pushing) {
                  map[y][x] = 'O'; // empty space now has a box
                }
                robot.y--; // robot moves up one
                map[robot.y][robot.x] = '.'; // robot's space is now empty
                done = true;
              case 'O':
                pushing = true;
              case '#':
                done = true; // hit a wall. break out of the outer loop
              default:
                throw ('invalid map tile at $x, $y: ${map[y][x]}');
            }
          }
        case '>': // right
          for (int x = robot.x + 1; !done && x < map[robot.y].length; x++) {
            int y = robot.y;
            switch (map[y][x]) {
              case '.':
                if (pushing) {
                  map[y][x] = 'O'; // empty space now has a box
                }
                robot.x++; // robot moves right one
                map[robot.y][robot.x] = '.'; // robot's space is now empty
                done = true;
              case 'O':
                pushing = true;
              case '#':
                done = true;
              default:
                throw ('invalid map tile at $x, $y: ${map[y][x]}');
            }
          }
        case 'v': // down
          for (int y = robot.y + 1; !done && y < map.length; y++) {
            int x = robot.x;
            switch (map[y][x]) {
              case '.':
                if (pushing) {
                  map[y][x] = 'O'; // empty space now has a box
                }
                robot.y++; // robot moves down one
                map[robot.y][robot.x] = '.'; // robot's space is now empty
                done = true;
              case 'O':
                pushing = true;
              case '#':
                done = true;
              default:
                throw ('invalid map tile at $x, $y: ${map[y][x]}');
            }
          }
        case '<': // left
          for (int x = robot.x - 1; !done && x >= 0; x--) {
            int y = robot.y;
            switch (map[y][x]) {
              case '.':
                if (pushing) {
                  map[y][x] = 'O'; // empty space now has a box
                }
                robot.x--; // robot moves left one
                map[robot.y][robot.x] = '.'; // robot's space is now empty
                done = true;
              case 'O':
                pushing = true;
              case '#':
                done = true;
              default:
                throw ('invalid map tile at $x, $y: ${map[y][x]}');
            }
          }
        default:
          throw ('invalid instruction $op');
      }
    }

    for (int y = 0; y < map.length; y++) {
      for (int x = 0; x < map[y].length; x++) {
        if (map[y][x] == 'O') result += (100 * y) + x;
      }
    }

    return result;
  }

  int day15b(List<String> lines) {
    int result = 0;

    // parse input
    List<List<String>> map = [];
    List<String> ops = [];
    Point2D robot = Point2D();
    int i = 0;
    while (lines[i].isNotEmpty) {
      map.add([]);
      for (int x = 0; x < lines[i].length; x++) {
        switch (lines[i][x]) {
          case '#':
            map[i].addAll(['#', '#']);
          case '.':
            map[i].addAll(['.', '.']);
          case 'O':
            map[i].addAll(['[', ']']);
          case '@':
            map[i].addAll(['.', '.']);
            robot.set(x * 2, i);
        }
      }
      i++;
    }
    if (robot.x == -1) throw ('robot not found');

    for (i = i + 1; i < lines.length; i++) {
      ops.addAll(lines[i].split(''));
    }

    for (i = 0; i < ops.length; i++) {
      final op = ops[i];
      // for (int y = 0; y < map.length; y++) {
      //   if (y == robot.y) {
      //     print(
      //         '${map[y].sublist(0, robot.x).join()}@${map[y].sublist(robot.x + 1).join()}');
      //   } else {
      //     print(map[y].join());
      //   }
      // }
      // print('step $i: move $op');
      // stdin.readLineSync();

      bool pushing = false;
      bool done = false;
      switch (op) {
        case '^': // up
          switch (map[robot.y - 1][robot.x]) {
            case '.':
              robot.y--;
            case '[':
              final from = Point2D(robot.x, robot.y - 1);
              if (canPushUp(map, from)) {
                pushUp(map, from);
                robot.y--;
              }
            case ']':
              final from = Point2D(robot.x - 1, robot.y - 1);
              if (canPushUp(map, from)) {
                pushUp(map, from);
                robot.y--;
              }
            case '#':
              break; // hit a wall
            default:
              throw ('invalid map tile at ${robot.x}, ${robot.y}: ${map[robot.y][robot.x]}');
          }
        case '>': // right
          for (int x = robot.x + 1; !done && x < map[robot.y].length; x++) {
            int y = robot.y;
            switch (map[y][x]) {
              case '.':
                robot.x++; // robot moves right one
                if (pushing) {
                  map[robot.y][robot.x] = '.'; // robot's space is now empty
                  while (x > robot.x) {
                    map[y][x--] = ']';
                    map[y][x--] = '[';
                  }
                }
                done = true;
              case '[':
                pushing = true;
                x++; // skip next character because we know it must be ']'
              case '#':
                done = true;
              default:
                throw ('invalid map tile at $x, $y: ${map[y][x]}');
            }
          }
        case 'v': // down
          switch (map[robot.y + 1][robot.x]) {
            case '.':
              robot.y++;
            case '[':
              final from = Point2D(robot.x, robot.y + 1);
              if (canPushDown(map, from)) {
                pushDown(map, from);
                robot.y++;
              }
            case ']':
              final from = Point2D(robot.x - 1, robot.y + 1);
              if (canPushDown(map, from)) {
                pushDown(map, from);
                robot.y++;
              }
            case '#':
              break; // hit a wall
            default:
              throw ('invalid map tile at ${robot.x}, ${robot.y}: ${map[robot.y][robot.x]}');
          }
        case '<': // left
          for (int x = robot.x - 1; !done && x >= 0; x--) {
            int y = robot.y;
            switch (map[y][x]) {
              case '.':
                robot.x--; // robot moves left one
                if (pushing) {
                  map[robot.y][robot.x] = '.'; // robot's space is now empty
                  while (x < robot.x) {
                    map[y][x++] = '[';
                    map[y][x++] = ']';
                  }
                }
                done = true;
              case ']':
                pushing = true;
                x--; // skip next character because we know it must be '['
              case '#':
                done = true;
              default:
                throw ('invalid map tile at $x, $y: ${map[y][x]}');
            }
          }
        default:
          throw ('invalid instruction $op');
      }
    }

    for (int y = 0; y < map.length; y++) {
      if (y == robot.y) {
        print(
            '${map[y].sublist(0, robot.x).join()}@${map[y].sublist(robot.x + 1).join()}');
      } else {
        print(map[y].join());
      }
    }
    print('robot is at ${robot.x}, ${robot.y}');

    for (int y = 0; y < map.length; y++) {
      for (int x = 0; x < map[y].length; x++) {
        if (map[y][x] == '[') result += (100 * y) + x;
      }
    }
    return result;
  }

  void pushUp(List<List<String>> map, Point2D from) {
    if (map[from.y - 1][from.x] == '#' || map[from.y - 1][from.x + 1] == '#') {
      throw ('tried to push up but cant');
    }
    switch (map[from.y - 1][from.x] + map[from.y - 1][from.x + 1]) {
      case '[]':
        pushUp(map, Point2D(from.x, from.y - 1));
      case '][':
        pushUp(map, Point2D(from.x - 1, from.y - 1));
        pushUp(map, Point2D(from.x + 1, from.y - 1));
      case '].':
        pushUp(map, Point2D(from.x - 1, from.y - 1));
      case '.[':
        pushUp(map, Point2D(from.x + 1, from.y - 1));
      case '..':
        break;
      default:
        throw ('error while pushing up');
    }
    map[from.y][from.x] = '.';
    map[from.y][from.x + 1] = '.';
    map[from.y - 1][from.x] = '[';
    map[from.y - 1][from.x + 1] = ']';
  }

  bool canPushUp(List<List<String>> map, Point2D from) {
    if (map[from.y - 1][from.x] == '#' || map[from.y - 1][from.x + 1] == '#') {
      return false;
    }
    switch (map[from.y - 1][from.x] + map[from.y - 1][from.x + 1]) {
      case '[]':
        return canPushUp(map, Point2D(from.x, from.y - 1));
      case '][':
        return canPushUp(map, Point2D(from.x - 1, from.y - 1)) &&
            canPushUp(map, Point2D(from.x + 1, from.y - 1));
      case '].':
        return canPushUp(map, Point2D(from.x - 1, from.y - 1));
      case '.[':
        return canPushUp(map, Point2D(from.x + 1, from.y - 1));
      case '..':
        return true;
      default:
        throw ('error while checking if can push up');
    }
  }

  void pushDown(List<List<String>> map, Point2D from) {
    if (map[from.y + 1][from.x] == '#' || map[from.y + 1][from.x + 1] == '#') {
      throw ('tried to push down but cant');
    }
    switch (map[from.y + 1][from.x] + map[from.y + 1][from.x + 1]) {
      case '[]':
        pushDown(map, Point2D(from.x, from.y + 1));
      case '][':
        pushDown(map, Point2D(from.x - 1, from.y + 1));
        pushDown(map, Point2D(from.x + 1, from.y + 1));
      case '].':
        pushDown(map, Point2D(from.x - 1, from.y + 1));
      case '.[':
        pushDown(map, Point2D(from.x + 1, from.y + 1));
      case '..':
        break;
      default:
        throw ('error while pushing down');
    }
    map[from.y][from.x] = '.';
    map[from.y][from.x + 1] = '.';
    map[from.y + 1][from.x] = '[';
    map[from.y + 1][from.x + 1] = ']';
  }

  bool canPushDown(List<List<String>> map, Point2D from) {
    if (map[from.y + 1][from.x] == '#' || map[from.y + 1][from.x + 1] == '#') {
      return false;
    }
    switch (map[from.y + 1][from.x] + map[from.y + 1][from.x + 1]) {
      case '[]':
        return canPushDown(map, Point2D(from.x, from.y + 1));
      case '][':
        return canPushDown(map, Point2D(from.x - 1, from.y + 1)) &&
            canPushDown(map, Point2D(from.x + 1, from.y + 1));
      case '].':
        return canPushDown(map, Point2D(from.x - 1, from.y + 1));
      case '.[':
        return canPushDown(map, Point2D(from.x + 1, from.y + 1));
      case '..':
        return true;
      default:
        throw ('error while checking if can push down');
    }
  }
}
