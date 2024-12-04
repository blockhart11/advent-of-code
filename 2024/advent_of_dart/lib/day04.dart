mixin Day04 {
  int day04a(List<String> lines) {
    int result = 0;

    // convert input to 2d array for ease of access
    List<List<String>> input = [];
    for (String line in lines) {
      input.add(line.split(''));
    }

    for (int row = 0; row < input.length; row++) {
      // for each row
      for (int col = 0; col < input[row].length; col++) {
        // for each column
        if (input[row][col] != 'X') continue; // short circuit if not X
        // check for border collisions first
        final left = col >= 3;
        final right = col < input[row].length - 3;
        final up = row >= 3;
        final down = row < input.length - 3;

        if (left && up) {
          // test up left
          if (isMAS(input[row - 1][col - 1], input[row - 2][col - 2],
              input[row - 3][col - 3])) {
                print('XMAS left-up at ($row, $col)');
            result += 1;
          }
        }
        if (up) {
          // test up
          if (isMAS(
              input[row - 1][col], input[row - 2][col], input[row - 3][col])) {
                print('XMAS up at ($row, $col)');
            result += 1;
          }
        }
        if (right && up) {
          // test up right
          if (isMAS(input[row - 1][col + 1], input[row - 2][col + 2],
              input[row - 3][col + 3])) {
                print('XMAS right-up at ($row, $col)');
            result += 1;
          }
        }
        if (right) {
          // test right
          if (isMAS(
              input[row][col + 1], input[row][col + 2], input[row][col + 3])) {
                print('XMAS right at ($row, $col)');
            result += 1;
          }
        }
        if (right && down) {
          // test down right
          if (isMAS(input[row + 1][col + 1], input[row + 2][col + 2],
              input[row + 3][col + 3])) {
                print('XMAS right-down at ($row, $col)');
            result += 1;
          }
        }
        if (down) {
          // test down
          if (isMAS(
              input[row + 1][col], input[row + 2][col], input[row + 3][col])) {
                print('XMAS down at ($row, $col)');
            result += 1;
          }
        }
        if (left && down) {
          // test down left
          if (isMAS(input[row + 1][col - 1], input[row + 2][col - 2],
              input[row + 3][col - 3])) {
                print('XMAS left-down at ($row, $col)');
            result += 1;
          }
        }
        if (left) {
          // test left
          if (isMAS(
              input[row][col - 1], input[row][col - 2], input[row][col - 3])) {
                print('XMAS left at ($row, $col)');
            result += 1;
          }
        }
      }
    }

    return result;
  }

  int day04b(List<String> lines) {
    int result = 0;

    // convert input to 2d array for easier use
    List<List<String>> input = [];
    for (String line in lines) {
      input.add(line.split(''));
    }

    for (int row = 0; row < input.length; row++) {
      // for each row
      for (int col = 0; col < input[row].length; col++) {
        // for each column
        if (input[row][col] != 'A') continue; // short circuit if not A
        if (row == 0 || col == 0 || row == input[row].length -1 || col == input.length - 1) continue; // short circuit if on a border

        if (isXMAS(input[row-1][col-1], input[row-1][col+1], input[row+1][col-1], input[row+1][col+1])) result++;
      }
    }

    return result;
  }

  // pattern match for part a
  bool isMAS(String m, String a, String s) {
    return m == 'M' && a == 'A' && s == 'S';
  }

  // pattern match for part b
  bool isXMAS(String ul, ur, dl, dr) {
    return ((ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M'))
    && ((ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M'));
  }
}
