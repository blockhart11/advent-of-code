import 'dart:io';
import 'dart:mirrors';

import 'day01.dart';
import 'day02.dart';
import 'day03.dart';
import 'day04.dart';
import 'day05.dart';
import 'day06.dart';
import 'day07.dart';
import 'day08.dart';
import 'day09.dart';
import 'day10.dart';
import 'day11.dart';
import 'day12.dart';
import 'day13.dart';
import 'day14.dart';
import 'day15.dart';
import 'day16.dart';
import 'day17.dart';
import 'day18.dart';
import 'day19.dart';
import 'day20.dart';
import 'day21.dart';
import 'day22.dart';
import 'day23.dart';
import 'day24.dart';
import 'day25.dart';

Future<void> calculate(String day,
    [bool test = false, bool alt = false]) async {
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

    final result =
        mirror.invoke(Symbol('day$day$part'), [lines.sublist(1)]).reflectee;
    print('Test run of day $day, part $part: expected $expect, got ${result != -1 ? result : 'whatever was printed above ^'}');
  } else {
    final result = mirror.invoke(Symbol('day$day$part'), [lines]).reflectee;
    print('Day $day, part $part result: ${result != -1 ? result : 'whatever was printed above ^'}');
  }
}

Future<void> create(String day) async {
  for (int i = int.parse(day); i <= 25; i++) {
    await File('input/$i.txt').create(); // input file
    await File('input/${i}_sample.txt').create(); // sample input file
    final fSoln = await File('lib/day$i.dart').create();
    fSoln.writeAsStringSync('''
mixin Day$i {
  int day${i}a(List<String> lines) {
    int result = 0;

    // do the thing

    return result;
  }

  int day${i}b(List<String> lines) {
    int result = 0;

    // do the thing

    return result;
  }
}
    ''');
  }
}

class Solutions
    with
        Day25,
        Day24,
        Day23,
        Day22,
        Day21,
        Day20,
        Day19,
        Day18,
        Day17,
        Day16,
        Day15,
        Day14,
        Day13,
        Day12,
        Day11,
        Day10,
        Day09,
        Day08,
        Day07,
        Day06,
        Day05,
        Day04,
        Day03,
        Day02,
        Day01 {}