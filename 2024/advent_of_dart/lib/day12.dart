class EdgeVisitor {
  bool up = false;
  bool down = false;
  bool right = false;
  bool left = false;

  void set(Direction dir) {
    switch (dir) {
      case Direction.up:
        up = true;
      case Direction.right:
        right = true;
      case Direction.down:
        down = true;
      case Direction.left:
        left = true;
    }
  }

  bool visited(Direction dir) {switch (dir) {
      case Direction.up:
        return up;
      case Direction.right:
        return right;
      case Direction.down:
        return down;
      case Direction.left:
        return left;
    }
  }
}

enum Direction {
  up,
  down,
  left,
  right;
}

class Plot {
  String letter;
  Map<(int, int), EdgeVisitor> nodes = {};
  int area = 0;
  int perimiter = 0;

  Plot(this.letter, (int, int) node) {
    nodes[node] = EdgeVisitor();
    area = 1;
    perimiter = 4;
  }

  void addNode((int, int) node) {
    nodes[node] = EdgeVisitor();
    area++;
    // perimiter changes depending on how many neighbors
    // - 1 neighbor: +2
    // - 2 neighbors: no change
    // - 3 neighbors: -2
    // - 4 neighbors: -4
    int numNeighbors = 0;
    if (nodes.containsKey((node.$1 - 1, node.$2))) {
      // up
      numNeighbors++;
    }
    if (nodes.containsKey((node.$1, node.$2 + 1))) {
      // right
      numNeighbors++;
    }
    if (nodes.containsKey((node.$1 + 1, node.$2))) {
      // down
      numNeighbors++;
    }
    if (nodes.containsKey((node.$1, node.$2 - 1))) {
      // left
      numNeighbors++;
    }

    switch (numNeighbors) {
      case 1:
        perimiter += 2;
      case 2:
        break;
      case 3:
        perimiter -= 2;
      case 4:
        perimiter -= 4;
      default:
        throw ('invalid number of neighbors for plot $this at $node: $numNeighbors');
    }
  }

  // you can only call this once because I track the state within the object here. I know, it's bad.
  int numSides() {
    int result = 0;

    for (final node in nodes.keys) {
      // detect and walk each edge, clockwise I guess
      // up
      if (!nodes[node]!.up) {
        if (!nodes.containsKey((node.$1 - 1, node.$2))) {
          // top edge detected
          result += walkEdge(node, Direction.up);
        }
      }
      // right
      if (!nodes[node]!.right) {
        if (!nodes.containsKey((node.$1, node.$2 + 1))) {
          // bottom edge detected
          result += walkEdge(node, Direction.right);
        }
      }
      // down
      if (!nodes[node]!.down) {
        if (!nodes.containsKey((node.$1 + 1, node.$2))) {
          // bottom edge detected
          result += walkEdge(node, Direction.down);
        }
      }
      // left
      if (!nodes[node]!.left) {
        if (!nodes.containsKey((node.$1, node.$2 - 1))) {
          // left edge detected
          result += walkEdge(node, Direction.left);
        }
      }
    }

    print('$letter has $result sides');

    return result;
  }

  int walkEdge((int, int) node, Direction edge) {
    if (nodes[node]!.visited(edge)) return 0;
    nodes[node]!.set(edge);

    switch (edge) {
      case Direction.up:
        if (nodes.containsKey((node.$1, node.$2+1))) { // has right neighbor
          if (nodes.containsKey((node.$1-1, node.$2+1))) { // upper right neighbor
            // inner corner. follow left edge up.
            return 1 + walkEdge((node.$1-1, node.$2+1), Direction.left);
          } else {
            // continue right along upper edge
            return walkEdge((node.$1, node.$2+1), edge);
          }
        } else {
          // outer corner. stay on same node, follow right edge
          return 1 + walkEdge(node, Direction.right);
        }
      case Direction.right:
        if (nodes.containsKey((node.$1+1, node.$2))) { // has lower neighbor
          if (nodes.containsKey((node.$1+1, node.$2+1))) { // lower right neighbor
            // inner corner. follow upper edge to the right.
            return 1 + walkEdge((node.$1+1, node.$2+1), Direction.up);
          } else {
            // continue down along right edge
            return walkEdge((node.$1+1, node.$2), edge);
          }
        } else {
          // outer corner. stay on same node, follow lower edge
          return 1 + walkEdge(node, Direction.down);
        }
      case Direction.down:
        if (nodes.containsKey((node.$1, node.$2-1))) { // has left neighbor
          if (nodes.containsKey((node.$1+1, node.$2-1))) { // lower left neighbor
            // inner corner. follow right edge down.
            return 1 + walkEdge((node.$1+1, node.$2-1), Direction.right);
          } else {
            // continue along lower edge
            return walkEdge((node.$1, node.$2-1), edge);
          }
        } else {
          // outer corner. stay on same node, follow left edge
          return 1 + walkEdge(node, Direction.left);
        }
      case Direction.left:
        if (nodes.containsKey((node.$1-1, node.$2))) { // has upper neighbor
          if (nodes.containsKey((node.$1-1, node.$2-1))) { // upper left neighbor
            // inner corner. follow lower edge left.
            return 1 + walkEdge((node.$1-1, node.$2-1), Direction.down);
          } else {
            // continue along left edge
            return walkEdge((node.$1-1, node.$2), edge);
          }
        } else {
          // outer corner. stay on same node, follow upper edge
          return 1 + walkEdge(node, Direction.up);
        }
    }
  }

  @override
  String toString() {
    return '$letter: area: $area, perimiter: $perimiter';
  }
}

mixin Day12 {
  int day12a(List<String> lines) {
    int result = 0;

    // Pseudocode
    // - iterate over each character top to bottom, left to right
    // - for each letter, FOLLOW THE WHOLE PLOT, charting the entire thing for that letter.
    //   - record all locations, area, and give the plot a unique ID (letter + ID)
    // - if the letter is already part of an existing plot, skip it. It should be charted already.

    List<List<bool>> visitedNodes =
        List.generate(lines.length, (_) => List.filled(lines[0].length, false));
    Map<String, List<Plot>> plots = {};

    for (int y = 0; y < lines.length; y++) {
      for (int x = 0; x < lines[y].length; x++) {
        if (visitedNodes[y][x]) continue; // already processed this node
        visitedNodes[y][x] = true;
        String letter = lines[y][x];

        final plot = Plot(letter, (y, x));

        // add plot to map
        plots.update(letter, (plots) => plots..add(plot),
            ifAbsent: () => [plot]);

        // process entire plot recursively
        visitAllNeighbors(plot, y, x, lines, visitedNodes);
      }
    }

    for (final letter in plots.values) {
      result +=
          letter.fold(0, (prev, next) => prev + next.area * next.perimiter);
      print(letter);
    }

    return result;
  }

  int day12b(List<String> lines) {
    int result = 0;

    List<List<bool>> visitedNodes =
        List.generate(lines.length, (_) => List.filled(lines[0].length, false));
    Map<String, List<Plot>> plots = {};

    for (int y = 0; y < lines.length; y++) {
      for (int x = 0; x < lines[y].length; x++) {
        if (visitedNodes[y][x]) continue; // already processed this node
        visitedNodes[y][x] = true;
        String letter = lines[y][x];

        final plot = Plot(letter, (y, x));

        // add plot to map
        plots.update(letter, (plots) => plots..add(plot),
            ifAbsent: () => [plot]);

        // process entire plot recursively
        visitAllNeighbors(plot, y, x, lines, visitedNodes);
      }
    }

    for (final letter in plots.values) {
      result += letter.fold(0, (prev, next) => prev + next.area * next.numSides());
    }

    return result;
  }

  void visitAllNeighbors(
      Plot p, int y, int x, List<String> map, List<List<bool>> visitedNodes) {
    // visit up
    if (y > 0 && !visitedNodes[y - 1][x] && map[y - 1][x] == p.letter) {
      visitedNodes[y - 1][x] = true;
      p.addNode((y - 1, x));
      visitAllNeighbors(p, y - 1, x, map, visitedNodes);
    }
    // visit right
    if (x < map[y].length - 1 &&
        !visitedNodes[y][x + 1] &&
        map[y][x + 1] == p.letter) {
      visitedNodes[y][x + 1] = true;
      p.addNode((y, x + 1));
      visitAllNeighbors(p, y, x + 1, map, visitedNodes);
    }
    // visit down
    if (y < map.length - 1 &&
        !visitedNodes[y + 1][x] &&
        map[y + 1][x] == p.letter) {
      visitedNodes[y + 1][x] = true;
      p.addNode((y + 1, x));
      visitAllNeighbors(p, y + 1, x, map, visitedNodes);
    }
    // visit left
    if (x > 0 && !visitedNodes[y][x - 1] && map[y][x - 1] == p.letter) {
      visitedNodes[y][x - 1] = true;
      p.addNode((y, x - 1));
      visitAllNeighbors(p, y, x - 1, map, visitedNodes);
    }
  }
}
