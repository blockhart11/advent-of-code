import 'dart:io';
import 'dart:mirrors';

Future<void> calculate(String day, [bool test = false, bool alt = false]) async {
  var mirror = reflect(Solutions());
  final sample = test ? "_sample" : "";
  final part = alt ? 'b' : 'a';
  final f = File('input/$day$sample.txt');
  List<String> lines = await f.readAsLines();

  if (test) {
    String expect = lines[0];
    if (!alt) {
      expect = expect.split(" ")[0];
    } else {
      expect = expect.split(" ")[1];
    }
    
    final result = mirror.invoke(Symbol('day$day$part'), [lines.sublist(1)]).reflectee;
    print('Test run of day $day, part $part: expected $expect, got $result');
  } else {
    final result = mirror.invoke(Symbol('day$day$part'), [lines]).reflectee;
    print('Day $day, part $part result: $result');
  }
}

class Solutions {
  int day01a(List<String> lines) {
    int result = 0;
    List<int> lhs = [];
    List<int> rhs = [];

    // parse input lines into 2 arrays
    for (var line in lines) {
      final args = line.split(" ");
      lhs.add(int.parse(args[0]));
      rhs.add(int.parse(args[1]));
    }

    // sort the arrays
    lhs.sort(); rhs.sort();

    // calculate distance for each line
    for (int i = 0; i < lhs.length; i++) {
      result += (lhs[i] - rhs[i]).abs();
    }

    return result;
  }
  int day01b(List<String> lines) {
    int result = 0;
    List<int> lhs = [];
    Map<int, int> rhsCount = {};

    // parse input lines into 2 arrays
    for (var line in lines) {
      final args = line.split(" ");
      lhs.add(int.parse(args[0]));
      final rhs = int.parse(args[1]);
      if (rhsCount.containsKey(rhs)) {
        rhsCount[rhs] = rhsCount[rhs]! + 1;
      } else {
        rhsCount[rhs] = 1;
      }
    }

    for (int left in lhs) {
      if (rhsCount.containsKey(left)) {
        result += left * rhsCount[left]!;
      }
    }

    return result;
  }
  
}