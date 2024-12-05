mixin Day05 {
  int day05a(List<String> lines) {
    int result = 0;
    Map<int, List<int>> rules = {};

    // parse the rules
    int i = 0;
    while (lines[i] != '') {
      final args = lines[i].split('|');
      final lhs = int.parse(args[0]);
      final rhs = int.parse(args[1]);
      rules.update(lhs, (list) => list..add(rhs), ifAbsent: () => [rhs]);
      i++;
    }

    // iterate over each manual
    for (i = i + 1; i < lines.length; i++) {
      final pages = lines[i].split(',').map((e) => int.parse(e)).toList();
      var valid = true;
      for (int k = 0; k < pages.length; k++) {
        // check that each page precedes the ones it's supposed to
        if (!pages.sublist(k + 1).fold(true, (prev, page) {
          if (!prev) return prev;
          if (!rules.containsKey(page)) return prev;

          return !rules[page]!.contains(pages[k]);
        })) {
          valid = false;
          break;
        }
      }
      if (valid) result += pages[pages.length ~/ 2];
    }

    return result;
  }

  int day05b(List<String> lines) {
    int result = 0;
    Map<int, List<int>> rules = {};

    // parse the rules
    int i = 0;
    while (lines[i] != '') {
      final args = lines[i].split('|');
      final lhs = int.parse(args[0]);
      final rhs = int.parse(args[1]);
      rules.update(lhs, (list) => list..add(rhs), ifAbsent: () => [rhs]);
      i++;
    }

    // iterate over each manual
    for (i = i + 1; i < lines.length; i++) {
      final pages = lines[i].split(',').map((e) => int.parse(e)).toList();
      var valid = true;
      for (int k = 0; k < pages.length; k++) {
        for (int j = k + 1; j < pages.length; j++) {
          if (!rules.containsKey(pages[j])) continue;
          if (rules[pages[j]]!.contains(pages[k])) {
            valid = false;

            pages.insert(k, pages[j]);
            pages.removeAt(j+1);
            k--; // reset k so that this iteration runs again on the newly placed page
            break;
          }
        }
      }

      // if invalid, add the newly sorted manual to the result
      if (!valid) {
        result += pages[pages.length ~/ 2];
      }
    }

    return result;
  }
}

class Day05Rule {
  List<int> before = [];
}
