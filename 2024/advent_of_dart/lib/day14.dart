class Robot {
  int px, py, vx, vy;
  Robot(this.px, this.py, this.vx, this.vy);
}

mixin Day14 {
  int day14a(List<String> lines) {
    int result = 0;
    final int duration = 7138;
    int width = 101;
    int height = 103;
    if (lines.length <= 20) {
      // detect if we're running the test input
      width = 11;
      height = 7;
    }

    List<Robot> robots = [];

    // parse input
    for (final line in lines) {
      final splitLine = line.split(' ');
      final p = splitLine[0].split(',');
      final v = splitLine[1].split(',');

      robots.add(Robot(int.parse(p[0].substring(2)), int.parse(p[1]),
          int.parse(v[0].substring(2)), int.parse(v[1])));
    }

    // safety zones
    int nw = 0;
    int ne = 0;
    int sw = 0;
    int se = 0;
    int zoneX = width ~/ 2;
    int zoneY = height ~/ 2;

    // run robots for duration seconds
    for (final robot in robots) {
      robot.px = (robot.px + duration * robot.vx) % width;
      robot.py = (robot.py + duration * robot.vy) % height;
      if (robot.px < zoneX) {
        // west
        if (robot.py < zoneY) {
          // north
          nw++;
        } else if (robot.py > zoneY) {
          // south
          sw++;
        }
      } else if (robot.px > zoneX) {
        // east
        if (robot.py < zoneY) {
          // north
          ne++;
        } else if (robot.py > zoneY) {
          // south
          se++;
        }
      }
    }

    draw(robots, width, height);
    result = nw * ne * sw * se;
    return result;
  }

  int day14b(List<String> lines) {
    int width = 101;
    int height = 103;
    if (lines.length <= 20) {
      // detect if we're running the test input
      width = 11;
      height = 7;
    }

    List<Robot> robots = [];

    // parse input
    for (final line in lines) {
      final splitLine = line.split(' ');
      final p = splitLine[0].split(',');
      final v = splitLine[1].split(',');

      robots.add(Robot(int.parse(p[0].substring(2)), int.parse(p[1]),
          int.parse(v[0].substring(2)), int.parse(v[1])));
    }

    // safety zones
    int zoneX = width ~/ 2;
    int zoneY = height ~/ 2;
    int minRating = 229839456; // answer to part a :shrug:
    int minRatingTime = 0;

    // run robots for 101 * 103 seconds
    for (int i = 0; i < 101 * 103; i++) {
      int nw = 0;
      int ne = 0;
      int sw = 0;
      int se = 0;
      for (final robot in robots) {
        final x = (robot.px + i * robot.vx) % width;
        final y = (robot.py + i * robot.vy) % height;
        if (x < zoneX) {
          // west
          if (y < zoneY) {
            // north
            nw++;
          } else if (y > zoneY) {
            // south
            sw++;
          }
        } else if (x > zoneX) {
          // east
          if (y < zoneY) {
            // north
            ne++;
          } else if (y > zoneY) {
            // south
            se++;
          }
        }
      }
      final rating = nw * ne * sw * se;
      if (rating < minRating) {
        minRating = rating;
        minRatingTime = i;
        print('maybe after $minRatingTime seconds?');
      }
    }

    return minRatingTime;
  }

  void draw(List<Robot> robots, int width, int height) {
    List<List<String>> map = List.generate(height, (int index) => List.filled(width, ' '));
    for (final robot in robots) {
      map[robot.py][robot.px] = '#';
    }
    for (final row in map) {
      print(row.join());
    }

  }
}
