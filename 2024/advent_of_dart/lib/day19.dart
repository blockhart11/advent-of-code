import 'dart:math';

mixin Day19 {
  int day19a(List<String> lines) {
    int result = 0;

    Map<String,List<String>> towels = {};
    for (final towel in lines[0].split(', ')) {
      towels.update(towel[0], (d) => d..add(towel), ifAbsent: () => [towel]);
    }

    for (final key in towels.keys) {
      List<String> patterns = towels[key]!;
      patterns.sort((a, b) => a.length.compareTo(b.length));
      print('$key: $patterns');
      for (int i = 0; i < patterns.length; i++) {
        if (canMakePattern(patterns[i], towels, patterns[i].length-1)) {
          // we don't need this towel
          patterns.removeAt(i--);
        }
      }
    }

    print('reduced pattern list:');
    for (final key in towels.keys) {
      print('$key: ${towels[key]}');
    }

    for (final pattern in lines.sublist(2)) {
      if (canMakePattern(pattern, towels)) {
        result++;
      }
    }

    return result;
  }

  bool canMakePattern(String p, Map<String, List<String>> towels, [int maxTowelLen = 10]) {
    if (p.isEmpty) return true; // base case
    return towels[p[0]]?.any((towel) {
      if (towel.length > maxTowelLen) return false;
      if (!p.startsWith(towel)) return false;
      return canMakePattern(p.substring(towel.length), towels);
    }) ?? false;
  }

  int day19b(List<String> lines) {
    int result = 0;

    Map<String, bool> towels = {};
    int maxLen = 0;
    for (final towel in lines[0].split(', ')) {
      towels[towel] = true;
      maxLen = max(maxLen, towel.length);
    }

    for (final pattern in lines.sublist(2)) {
      final subResult = count(towels, pattern, maxLen);
      print('$subResult patterns for: $pattern');
      result += subResult;
    }

    return result;
  }

  int count(Map<String, bool> towels, String pattern, [int maxLen = 8]) {
    Map<String, int> cache = {};
    return countHelper(cache, towels, pattern, maxLen);
  }

  int countHelper(Map<String, int> cache, Map<String, bool> towels, String pattern, [int maxLen = 8]) {
    if (pattern.isEmpty) return 1;
    if (cache.containsKey(pattern)) return cache[pattern]!;

    int result = 0;
    for (int i = 1; i <= pattern.length && i <= maxLen; i++) {
      final lhs = pattern.substring(0, i);
      if (towels.containsKey(lhs)) {
        final rhs = pattern.substring(i);
        final subResult = countHelper(cache, towels, pattern.substring(i), maxLen);
        cache[rhs] = subResult;
        result += subResult;
      }
    }
    return result;
  }
}
    