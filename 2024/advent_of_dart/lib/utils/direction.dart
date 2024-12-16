enum Dir2D { up, down, left, right }

// rotate 90 degrees clockwise, or counter-clockwise if ccw is true
Dir2D rotate(Dir2D dir, [bool ccw = false]) {
  switch (dir) {
    case Dir2D.up:
      return ccw ? Dir2D.left : Dir2D.right;
    case Dir2D.down:
      return ccw ? Dir2D.right : Dir2D.left;
    case Dir2D.left:
      return ccw ? Dir2D.down : Dir2D.up;
    case Dir2D.right:
      return ccw ? Dir2D.up : Dir2D.down;
  }
}

String toSymbol(Dir2D dir) {
  switch (dir) {
    case Dir2D.up:
      return '^';
    case Dir2D.down:
      return 'v';
    case Dir2D.left:
      return '<';
    case Dir2D.right:
      return '>';
  }
}