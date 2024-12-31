class Pinset {
  int a, b, c, d, e;

  Pinset(this.a, this.b, this.c, this.d, this.e);

  bool fits(Pinset other) {
    return (
      a + other.a < 6
      && b + other.b < 6
      && c + other.c < 6
      && d + other.d < 6
      && e + other.e < 6
    );
  }
}

mixin Day25 {
  int day25a(List<String> lines) {
    int result = 0;

    List<Pinset> locks = [];
    List<Pinset> keys = [];

    for (int i = 0; i < lines.length; i+=7) {
      List<int> pins = [];
      if (lines[i].startsWith('.')) {
        // this is a key
        i++;
        for (int j = 0; j < 5; j++) { // set each pin
          for (int k = i; ;k++) {
            if (lines[k][j] == '#') {
              pins.add(5-(k-i));
              break;
            }
          }
        }
        keys.add(Pinset(pins[0], pins[1], pins[2], pins[3], pins[4]));
      } else {
        // this is a lock
        i++;
        for (int j = 0; j < 5; j++) { // set each pin
          for (int k = i; ;k++) {
            if (lines[k][j] == '.') {
              pins.add(k-i);
              break;
            }
          }
        }
        locks.add(Pinset(pins[0], pins[1], pins[2], pins[3], pins[4]));
      }
    }

    // brute force
    print('testing ${locks.length} locks against ${keys.length} keys');
    for (final lock in locks) {
      for (final key in keys) {
        if (lock.fits(key)) {
          result++;
        }
      }
    }

    return result;
  }

  int day25b(List<String> lines) {
    int result = 0;

    // do the thing

    return result;
  }
}
    