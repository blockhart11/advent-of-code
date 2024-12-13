import 'dart:math';

class ClawMachine {
  int ax, ay, bx, by, px, py;

  ClawMachine()
      : ax = 0,
        ay = 0,
        bx = 0,
        by = 0,
        px = 0,
        py = 0;

  @override
  String toString() {
    return 'A: $ax, $ay, B: $bx, $by, Prize: $px, $py';
  }
}

mixin Day13 {
  int day13a(List<String> lines) {
    int result = 0;
    List<ClawMachine> machines = parseInput(lines);

    for (final m in machines) {
      int minCost = 10000000000000;
      final aMax = min(100, min(m.px ~/ m.ax, m.py ~/ m.ay));
      final bMax = min(100, min(m.px ~/ m.bx, m.py ~/ m.by));
      if (aMax < bMax) {
        // a has fewer possibilities. try those.
        for (int a = 0; a <= aMax; a++) {
          final xLeft = m.px - (a * m.ax);
          if (xLeft % m.bx == 0) {
            final b = xLeft ~/ m.bx;
            if ((a * m.ay) + (b * m.by) == m.py) {
              // solution found
              final cost = 3*a + b;
              print('solution for machine $m found\nA: $a, B: $b, cost: $cost');
              minCost = min(minCost, 3*a + b);
            }
          }
        }
      } else {
        // b has fewer possibilities. try those.
        for (int b = 0; b <= bMax; b++) {
          final xLeft = m.px - (b * m.bx);
          if (xLeft % m.ax == 0) {
            final a = xLeft ~/ m.ax;
            if ((a * m.ay) + (b * m.by) == m.py) {
              // solution found
              final cost = 3*a + b;
              print('solution for machine $m found\nA: $a, B: $b, cost: $cost');
              minCost = min(minCost, 3*a + b);
            }
          }
        }
      }
      if (minCost != 10000000000000) {
        result += minCost;
      }
    }

    return result;
  }

  int day13b(List<String> lines) {
    int result = 0;
    List<ClawMachine> machines = parseInput(lines, true);

    for (int i = 0; i < machines.length; i++) {
      final m = machines[i];
      final b = ((m.py - ((m.px*m.ay)/m.ax)) / (m.by - ((m.bx*m.ay)/m.ax))).round();
      final a = ((m.py - ((m.px*m.by)/m.bx)) / (m.ay - ((m.ax*m.by)/m.bx))).round();
      if (a*m.ax + b*m.bx == m.px
      && a*m.ay + b*m.by == m.py) {
          final cost = (3*a) + b;
          print('solution for machine $i found. A: $a, B: $b, cost: $cost');
          print('Prize: X=${a*m.ax + b*m.bx}, Y=${a*m.ay + b*m.by}');
          result += cost;
      }
    }

    return result;
  }

  List<ClawMachine> parseInput(List<String> lines, [bool addAGajillion = false]) {
    List<ClawMachine> machines = [];

    for (int i = 0; i < lines.length; i += 4) {
      final nextClaw = ClawMachine();
      final a = lines[i].split(' ');
      final b = lines[i + 1].split(' ');
      final p = lines[i + 2].split(' ');

      nextClaw.ax = int.parse(a[2].substring(2, a[2].length - 1));
      nextClaw.ay = int.parse(a[3].substring(2));
      nextClaw.bx = int.parse(b[2].substring(2, b[2].length - 1));
      nextClaw.by = int.parse(b[3].substring(2));
      nextClaw.px = int.parse(p[1].substring(2, p[1].length - 1));
      nextClaw.py = int.parse(p[2].substring(2));

      if (addAGajillion) {
        nextClaw.px += 10000000000000;
        nextClaw.py += 10000000000000;
      }

      machines.add(nextClaw);
    }
    return machines;
  }
}
