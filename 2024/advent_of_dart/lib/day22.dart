import 'dart:collection';

const prune = 16777216;

mixin Day22 {
  int day22a(List<String> lines) {
    int result = 0;

    for (final line in lines) {
      int next = int.parse(line);
      for (int i = 0; i < 2000; i++) {
        next = secretNumber(next);
      }
      print('$line: $next');
      result += next;
    }

    return result;
  }

  int secretNumber(int input) {
    //     Calculate the result of multiplying the secret number by 64. Then, mix this result into the secret number. Finally, prune the secret number.
    int result = ((input * 64) ^ input) % prune;
    // Calculate the result of dividing the secret number by 32. Round the result down to the nearest integer. Then, mix this result into the secret number. Finally, prune the secret number.
    result = ((result ~/ 32) ^ result) % prune;
    // Calculate the result of multiplying the secret number by 2048. Then, mix this result into the secret number. Finally, prune the secret number.
    return ((result * 2048) ^ result) % prune;
  }

  int day22b(List<String> lines) {
    Map<(int, int, int, int), Map<int, int>> c = {};

    for (int id = 0; id < lines.length; id++) {
      int cur = int.parse(lines[id]); // first secret number
      Queue<int> sequence = Queue();
      int i = 0;
      while (i < 4) { // seed the change list with the next 4 numbers
        final next = secretNumber(cur);
        sequence.add(next%10 - cur%10);
        cur = next;
        i++;
      }
      updateCache(c, sequence, id, cur%10);

      while (i < 2000) {
        final next = secretNumber(cur);
        sequence.add(next%10-cur%10);
        sequence.removeFirst();
        updateCache(c, sequence, id, next%10);
        cur = next;
        i++;
      }
    }

    // find best result in map
    ((int, int, int, int), int) best = ((0,0,0,0), -1);
    for (final entry in c.entries) {
      final score = entry.value.values.fold(0, (prev, next) => prev + next);
      if (score > best.$2) {
        best = ((entry.key), score);
      }
    }

    print(best);

    return best.$2;
  }

  void updateCache(Map<(int, int, int, int), Map<int, int>> c, Queue q, int id, int val) {
    (int, int, int, int) key = (q.elementAt(0), q.elementAt(1), q.elementAt(2), q.elementAt(3));
    Map<int, int>? cVal = c[key];
    if (cVal == null) {
      c[key] = {id: val};
    } else {
      if (cVal.containsKey(id)) return; // we've already seen this sequence for this number.
      cVal[id] = val;
    }
  }
}