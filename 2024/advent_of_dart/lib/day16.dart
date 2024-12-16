import 'package:advent_of_dart/utils/direction.dart';
import 'package:advent_of_dart/utils/point2d.dart';

class Path {
  int cost = 0;
  Point2D pos = Point2D();
  Dir2D dir = Dir2D.right;
  String path = '';
  bool valid = true;
  bool complete = false;
  Set<Path> altPaths = {};

  Path(Point2D start)
      : pos = Point2D.from(start); // copy the point, don't steal it

  Path.from(Path other) {
    cost = other.cost;
    pos = Point2D.from(other.pos);
    dir = other.dir;
    path = other.path;
    valid = other.valid;
    altPaths = Set.from(other.altPaths);
  }

  Path.invalid() : valid = false;

  void invalidate() {
    valid = false;
    complete = false;
    altPaths = {};
  }

  void turn([bool ccw = false]) {
    dir = rotate(dir, ccw);
    cost += 1000;
  }

  void forward() {
    cost++;
    path += toSymbol(dir);
    pos = Point2D.from(pos, dir);
  }

  static Path min(List<Path> paths) {
    Path? minPath;
    for (Path p in paths) {
      if (p.valid && p.complete && (minPath == null || p.cost < minPath.cost)) {
        minPath = p;
      }
    }
    return minPath ?? Path.invalid();
  }

  static List<Path> mins(List<Path> paths) {
    List<Path> result = [];
    int minCost = 0;
    for (Path p in paths) {
      if (p.valid) {
        if (result.isEmpty || (result.isNotEmpty && p.cost < minCost)) {
          result = [p];
          minCost = p.cost;
        } else if (p.cost == minCost) {
          result.add(p);
        }
      }
    }
    return result;
  }

  void joinPaths(Path alt) {
    altPaths.addAll({alt, ...alt.altPaths});
  }
}

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

  void recordPaths(List<Path> paths) {
    for (final path in paths) {
      var pos = Point2D.from(start);
      nodeAt(pos)!.visit();
      for (int i = 0; i < path.path.length; i++) {
        switch (path.path[i]) {
          case '^':
            pos.y--;
          case 'v':
            pos.y++;
          case '<':
            pos.x--;
          case '>':
            pos.x++;
          default:
            throw ('Unknown direction ${path.path[i]}');
        }
        nodeAt(pos)!.visit();
      }
    }
  }

  int printMaze() {
    int result = 0;
    for (var row in nodes) {
      String line = '';
      for (var node in row) {
        if (node.type == NodeType.wall) {
          line += '#';
        } else if (node.type == NodeType.path) {
          if (node.visited) {
            line += 'O';
            result++;
          } else {
            line += '.';
          }
        } else if (node.type == NodeType.goal) {
          line += 'E';
          result++;
        } else {
          line += '?';
        }
      }
      print(line);
    }
    return result;
  }

  // returns node at pos if inDirection is null. Otherwise moves and returns next node if valid
  Node? nodeAt(Point2D pos) {
    if (pos.x >= 0 &&
        pos.y >= 0 &&
        pos.x < nodes[pos.y].length &&
        pos.y < nodes.length) {
      return nodes[pos.y][pos.x];
    }
    return null;
  }

  // moving forward costs 1 point, turning 90 degrees costs 1000 points
  List<Path> findCheapestPaths() {
    return cheapestPathsHelper(Path(Point2D.from(start)));
  }

  List<Path> cheapestPathsHelper(Path path) {
    if (path.cost > 99460) return []; // short circuit using answer from part a

    final node = nodeAt(path.pos)!;
    if (node.worthIt(path.dir, path)) {
      final nextNode = nodeAt(Point2D.from(path.pos, path.dir));
      if (nextNode != null) {
        switch (nextNode.type) {
          case NodeType.wall: // hit a wall. try rotating
            Path p2 = Path.from(path)..turn();
            Path p3 = Path.from(path)..turn(true);
            return Path.mins(
                [...cheapestPathsHelper(p2), ...cheapestPathsHelper(p3)]);
          case NodeType.path:
            Path p2 = Path.from(path)..turn();
            Path p3 = Path.from(path)..turn(true);
            return Path.mins([
              ...cheapestPathsHelper(path..forward()), // move forward
              ...cheapestPathsHelper(p2), // and rotate cw
              ...cheapestPathsHelper(p3) // and rotate ccw
            ]);
          case NodeType.goal:
            return [path..forward()];
          case NodeType.unknown:
            throw ('fell into the void of the maze?');
          default:
            throw ('unknown node type');
        }
      } else {
        throw ('fell off the map at ${path.pos} somehow');
      }
    } else {
      // node has a cheaper path already
      return [];
    }
  }
}

enum NodeType { wall, path, goal, unknown }

class Node {
  int? up, down, left, right; // best scores facing each direction
  NodeType type = NodeType.unknown;
  bool visited = false;
  Node(this.type);

  bool worthIt(Dir2D dir, Path p) {
    switch (dir) {
      case Dir2D.up:
        if (up == null || up! >= p.cost) {
          up = p.cost;
          return true;
        }
      case Dir2D.down:
        if (down == null || down! >= p.cost) {
          down = p.cost;
          return true;
        }
      case Dir2D.left:
        if (left == null || left! >= p.cost) {
          left = p.cost;
          return true;
        }
      case Dir2D.right:
        if (right == null || right! >= p.cost) {
          right = p.cost;
          return true;
        }
    }
    return false;
  }

  void visit() {
    visited = true;
  }
}

mixin Day16 {
  int day16a(List<String> lines) {
    Maze maze = Maze(lines);
    final paths = maze.findCheapestPaths();
    maze.recordPaths(paths);
    int nodesVisited = maze.printMaze();
    print('$nodesVisited nodes were visited');
    for (final path in paths) {
      print(path.path);
    }
    return paths[0].cost;
  }

  int day16b(List<String> lines) {
    Maze maze = Maze(lines);
    final paths = maze.findCheapestPaths();
    maze.recordPaths(paths);
    int nodesVisited = maze.printMaze();
    print('$nodesVisited nodes were visited');
    return nodesVisited;
  }
}
