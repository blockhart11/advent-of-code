mixin Day03 {
  int day03a(List<String> lines) {
    int result = 0;
    List<String> mults = [];

    for (final line in lines) {
      RegExp regex = RegExp(r"mul\(\d+,\d+\)");
      Iterable<RegExpMatch> matches = regex.allMatches(line);
      for (final match in matches) {
        mults.add(match.group(0)!);
      }
    }

    for (final mult in mults) {
      result += mul(mult);
    }

    return result;
  }

  int day03b(List<String> lines) {
    int result = 0;
    List<String> ops = [];
    bool mulEnabled = true;

    for (final line in lines) {
      RegExp regex = RegExp(r"mul\(\d+,\d+\)|do\(\)|don\\'t\(\)");
      Iterable<RegExpMatch> matches = regex.allMatches(line);
      for (final match in matches) {
        ops.add(match.group(0)!);
      }
    }

    for (final op in ops) {
      if (op.startsWith('don\'t(')) {
        mulEnabled = false;
      } else if (op.startsWith('do(')) {
        mulEnabled = true;
      } else if (mulEnabled) { // op.startsWith('mul(') == true
        result += mul(op);
      }
    }

    return result;
  }

  int mul(String op) {
    final sides = op.split(',');
    final lhs = int.parse(sides[0].split('(')[1]);
    final rhs = int.parse(sides[1].substring(0, sides[1].length-1));
    return lhs * rhs;
  }
}
