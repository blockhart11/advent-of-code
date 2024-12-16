import 'package:advent_of_dart/utils/direction.dart';

class Point2D {
  int x, y;
  Point2D([this.x = -1, this.y = -1]);

  Point2D.from(Point2D other, [Dir2D? dir])
      : x = other.x,
        y = other.y {
    switch (dir) {
      case null:
        break;
      case Dir2D.up:
        y -= 1;
      case Dir2D.down:
        y += 1;
      case Dir2D.left:
        x -= 1;
      case Dir2D.right:
        x += 1;
    }
  }

  void set(int x, int y) {
    this.x = x;
    this.y = y;
  }

  @override
  String toString() {
    return '($x,$y)';
  }
}
