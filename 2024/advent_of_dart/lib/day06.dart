import 'package:advent_of_dart/utils/direction.dart';

class Position {
  int x;
  int y;
  Dir2D orientation = Dir2D.up;

  Position(this.x, this.y);
  void setPos(int x, int y) {
    this.x = x;
    this.y = y;
  }

  void rotate() {
    switch (orientation) {
      case Dir2D.up:
        orientation = Dir2D.right;
      case Dir2D.down:
        orientation = Dir2D.left;
      case Dir2D.right:
        orientation = Dir2D.down;
      case Dir2D.left:
        orientation = Dir2D.up;
    }
  }
}

class Tile {
  bool visited = false;
  bool visitedUp = false;
  bool visitedDown = false;
  bool visitedRight = false;
  bool visitedLeft = false;
  bool obstructed = false;

  Tile([this.obstructed = false]);
}

mixin Day06 {
  int day06a(List<String> lines) {
    int result = 0;
    var pos = Position(-1, -1);
    List<List<Tile>> map = parseMap(lines, pos);

    while (inBounds(pos, map.length, map[0].length)) {
      if (!map[pos.y][pos.x].visited) {
        print('visiting (${pos.x},${pos.y})');
        result++;
        map[pos.y][pos.x].visited = true;
      } else {
        print('already visited (${pos.x},${pos.y})');
      }

      move(map, pos);
    }

    return result;
  }

  int day06b(List<String> lines) {
    int result = 0;
    var startPos = Position(-1, -1);
    List<List<Tile>> map = parseMap(lines, startPos);

    for (int row = 0; row < lines.length; row++) {
      for (int col = 0; col < lines[0].length; col++) {
        if (map[col][row].obstructed) continue; // already blocked. no loop here.
        if (startPos.x == col && startPos.y == row) continue; // starting position not allowed.

        Map<(int, int), Tile> visited = {};
        map[col][row].obstructed = true; // obstruct this tile temporarily
        bool foundLoop = false;
        var pos = Position(startPos.x, startPos.y); // reset to starting position
        while (inBounds(pos, map.length, map[0].length) && !foundLoop) {
          // check for a loop
          switch (pos.orientation) {
            case Dir2D.up:
              if (visited.containsKey((pos.y, pos.x)) && visited[(pos.y,pos.x)]!.visitedUp) {
                foundLoop = true;
              } else {
                visited.update((pos.y,pos.x),(tile) => tile..visitedUp = true, ifAbsent: () => Tile()..visitedUp = true);
              }
            case Dir2D.down:
              if (visited.containsKey((pos.y, pos.x)) && visited[(pos.y,pos.x)]!.visitedDown) {
                foundLoop = true;
              } else {
                visited.update((pos.y,pos.x),(tile) => tile..visitedDown = true, ifAbsent: () => Tile()..visitedDown = true);
              }
            case Dir2D.right:
              if (visited.containsKey((pos.y, pos.x)) && visited[(pos.y,pos.x)]!.visitedRight) {
                foundLoop = true;
              } else {
                visited.update((pos.y,pos.x),(tile) => tile..visitedRight = true, ifAbsent: () => Tile()..visitedRight = true);
              }
            case Dir2D.left:
              if (visited.containsKey((pos.y, pos.x)) && visited[(pos.y,pos.x)]!.visitedLeft) {
                foundLoop = true;
              } else {
                visited.update((pos.y,pos.x),(tile) => tile..visitedLeft = true, ifAbsent: () => Tile()..visitedLeft = true);
              }
            default: // error
              throw ('what Dir2D is this even');
          }

          if (foundLoop) {
            print('Found loop by obstructing (${pos.x},${pos.y})');
            result++;
            break;
          }

          move(map, pos);
        }
        map[col][row].obstructed = false; // clear obstruction
      }
    }
    return result;
  }

  List<List<Tile>> parseMap(List<String> input, Position posPtr) {
    List<List<Tile>> map = [];
    for (int y = 0; y < input.length; y++) {
      List<String> chars = input[y].split('');
      int x = 0;
      map.add(chars.map((e) {
        if (e == '^') posPtr.setPos(x, y);
        x++;
        return Tile(e == '#');
      }).toList());
    }
    return map;
  }

  bool inBounds(Position pos, int height, int width) {
    return pos.x >= 0 && pos.y >= 0 && pos.x < width && pos.y < height;
  }

  void move(List<List<Tile>> map, Position pos) {
    switch (pos.orientation) {
      case Dir2D.up:
        if (pos.y == 0 || !map[pos.y - 1][pos.x].obstructed) {
          pos.setPos(pos.x, pos.y - 1); // move
        } else {
          pos.rotate();
        }
      case Dir2D.down:
        if (pos.y == map.length - 1 || !map[pos.y + 1][pos.x].obstructed) {
          pos.setPos(pos.x, pos.y + 1); // move
        } else {
          pos.rotate();
        }
      case Dir2D.left:
        if (pos.x == 0 || !map[pos.y][pos.x - 1].obstructed) {
          pos.setPos(pos.x - 1, pos.y); // move
        } else {
          pos.rotate();
        }
      case Dir2D.right:
        if (pos.x == map[0].length - 1 || !map[pos.y][pos.x + 1].obstructed) {
          pos.setPos(pos.x + 1, pos.y); // move
        } else {
          pos.rotate();
        }
    }
  }
}
