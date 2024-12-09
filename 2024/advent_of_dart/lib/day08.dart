mixin Day08 {
  int day08a(List<String> lines) {
    Map<String, List<(int, int)>> antennae =
        {}; // for each frequency, track each location it occurs
    for (int row = 0; row < lines.length; row++) {
      for (int col = 0; col < lines[row].length; col++) {
        final freq = lines[row][col];
        if (freq == '.') continue;
        antennae.update(freq, (current) => current..add((col, row)),
            ifAbsent: () => [(col, row)]);
      }
    }

    print('Map:\n $antennae');
    for (final key in antennae.keys) {
      print('$key: ${antennae[key]!.length}');
    }

    Map<(int, int), bool> antinodes = {};
    final height = lines.length;
    final width = lines[0].length;
    for (final antenna in antennae.entries) {
      final coords = antenna.value;
      for (int i = 0; i < coords.length; i++) {
        for (int j = i + 1; j < coords.length; j++) {
          // calculate antinodes of coords[i] and coords[j]. I'm too dumb to solve it elegantly
          final isLeft = coords[i].$1 < coords[j].$1;
          final dx = (coords[i].$1 - coords[j].$1).abs();
          final dy = (coords[i].$2 - coords[j].$2).abs();
          // note: i is never below j because the input is read from top to bottom.
          if (isLeft) {
            // i is left of j
            final nodeA = (coords[i].$1 - dx, coords[i].$2 - dy);
            if (nodeA.$1 >= 0 && nodeA.$2 >= 0) {
              print('antinode at (${nodeA.$1}, ${nodeA.$2})');
              antinodes[nodeA] = true;
            }
            final nodeB = (coords[j].$1 + dx, coords[j].$2 + dy);
            if (nodeB.$1 < width && nodeB.$2 < height) {
              print('antinode at (${nodeB.$1}, ${nodeB.$2})');
              antinodes[nodeB] = true;
            }
          } else {
            // i is right of j
            final nodeA = (coords[i].$1 + dx, coords[i].$2 - dy);
            if (nodeA.$1 < width && nodeA.$2 >= 0) {
              print('antinode at (${nodeA.$1}, ${nodeA.$2})');
              antinodes[nodeA] = true;
            }
            final nodeB = (coords[j].$1 - dx, coords[j].$2 + dy);
            if (nodeB.$1 >= 0 && nodeB.$2 < height) {
              print('antinode at (${nodeB.$1}, ${nodeB.$2})');
              antinodes[nodeB] = true;
            }
          }
        }
      }
    }

    return antinodes.length;
  }

  int day08b(List<String> lines) {
    Map<String, List<(int, int)>> antennae =
        {}; // for each frequency, track each location it occurs
    for (int row = 0; row < lines.length; row++) {
      for (int col = 0; col < lines[row].length; col++) {
        final freq = lines[row][col];
        if (freq == '.') continue;
        antennae.update(freq, (current) => current..add((col, row)),
            ifAbsent: () => [(col, row)]);
      }
    }

    print('Map:\n $antennae');
    for (final key in antennae.keys) {
      print('$key: ${antennae[key]!.length}');
    }

    // iterate over each antenna, find and record all unique antinode locations (don't double count!)
    Map<(int, int), bool> antinodes = {};
    final height = lines.length;
    final width = lines[0].length;
    for (final antenna in antennae.entries) {
      final coords = antenna.value;
      for (int i = 0; i < coords.length; i++) {
        antinodes[coords[i]] = true;
        for (int j = i + 1; j < coords.length; j++) {
          // calculate antinodes of coords[i] and coords[j]. I'm too dumb to solve it elegantly
          final isLeft = coords[i].$1 < coords[j].$1;
          final dx = (coords[i].$1 - coords[j].$1).abs();
          final dy = (coords[i].$2 - coords[j].$2).abs();
          // note: i is never below j because the input is read from top to bottom.
          if (isLeft) {
            // i is left of j
            var nodeA = (coords[i].$1 - dx, coords[i].$2 - dy);
            while (nodeA.$1 >= 0 && nodeA.$2 >= 0) {
              print('antinode at (${nodeA.$1}, ${nodeA.$2})');
              antinodes[nodeA] = true;
              nodeA = (nodeA.$1 - dx, nodeA.$2 - dy);
            }
            var nodeB = (coords[j].$1 + dx, coords[j].$2 + dy);
            while (nodeB.$1 < width && nodeB.$2 < height) {
              print('antinode at (${nodeB.$1}, ${nodeB.$2})');
              antinodes[nodeB] = true;
              nodeB = (nodeB.$1 + dx, nodeB.$2 + dy);
            }
          } else {
            // i is right of j
            var nodeA = (coords[i].$1 + dx, coords[i].$2 - dy);
            while (nodeA.$1 < width && nodeA.$2 >= 0) {
              print('antinode at (${nodeA.$1}, ${nodeA.$2})');
              antinodes[nodeA] = true;
              nodeA = (nodeA.$1 + dx, nodeA.$2 - dy);
            }
            var nodeB = (coords[j].$1 - dx, coords[j].$2 + dy);
            while (nodeB.$1 >= 0 && nodeB.$2 < height) {
              print('antinode at (${nodeB.$1}, ${nodeB.$2})');
              antinodes[nodeB] = true;
              nodeB = (nodeB.$1 - dx, nodeB.$2 + dy);
            }
          }
        }
      }
    }
    
    return antinodes.length;
  }
}
