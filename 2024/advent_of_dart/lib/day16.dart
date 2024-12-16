import 'package:advent_of_dart/utils/direction.dart';
import 'package:advent_of_dart/utils/point2d.dart';

class Maze {
  List<List<Node>> nodes = [];
  Point2D start = Point2D();
  Maze(List<String> input) {
    for (int y = 0; y < input.length; y++) {
      nodes.add([]);
      for (int x = 0; x < input[y].length; x++) {
        switch (input[y][x]) {
          case '#':
            nodes[y].add(Node(NodeType.wall));
          case '.':
            nodes[y].add(Node(NodeType.path));
          case 'S':
            nodes[y].add(Node(NodeType.path));
            start.set(x, y);
          case 'E':
            nodes[y].add(Node(NodeType.goal));
          default:
            throw ('Unknown node type ${input[y][x]}');
        }
      }
    }
  }

  Node nodeAt(Point2D pos) {
    return nodes[pos.y][pos.x];
  }


  // moving forward costs 1 point, turning 90 degrees costs 1000 points
  int findCheapestPath() {
    return cheapestPathHelper(start, Dir2D.right, 0);
  }

  int cheapestPathHelper(Point2D pos, Dir2D dir, int cost) {
    // try going forward, turning cw, and turning ccw
    final node = nodeAt(pos);
    if (node.worthIt(dir, cost)) {

    }
  }
}

enum NodeType { wall, path, goal, unknown }

class Node {
  int? up, down, left, right; // best scores facing each direction
  NodeType type = NodeType.unknown;
  Node(this.type);

  bool worthIt(Dir2D dir, int cost) {
    switch (dir) {
      case Dir2D.up:
        if (up == null || up! > cost) {
          up = cost;
          return true;
        }
      case Dir2D.down:
        if (down == null || down! > cost) {
          down = cost;
          return true;
        }
      case Dir2D.left:
        if (left == null || left! > cost) {
          left = cost;
          return true;
        }
      case Dir2D.right:
        if (right == null || right! > cost) {
          right = cost;
          return true;
        }
    }
    return false;
  }
}

mixin Day16 {
  int day16a(List<String> lines) {
    Maze maze = Maze(lines);
    return maze.findCheapestPath();
  }

  int day16b(List<String> lines) {
    int result = 0;

    // do the thing

    return result;
  }
}
