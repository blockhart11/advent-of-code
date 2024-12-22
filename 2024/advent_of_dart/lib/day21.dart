import 'dart:math';

enum Key { kA, k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, kU, kR, kD, kL }

Key toKey(String c) {
  switch (c) {
    case 'A':
      return Key.kA;
    case '0':
      return Key.k0;
    case '1':
      return Key.k1;
    case '2':
      return Key.k2;
    case '3':
      return Key.k3;
    case '4':
      return Key.k4;
    case '5':
      return Key.k5;
    case '6':
      return Key.k6;
    case '7':
      return Key.k7;
    case '8':
      return Key.k8;
    case '9':
      return Key.k9;
    case '^':
      return Key.kU;
    case '>':
      return Key.kR;
    case 'v':
      return Key.kD;
    case '<':
      return Key.kL;
    default:
      throw ArgumentError('Invalid key character: $c');
  }
}

String keyToString(Key key) {
  switch (key) {
    case Key.kA:
      return 'A';
    case Key.k0:
      return '0';
    case Key.k1:
      return '1';
    case Key.k2:
      return '2';
    case Key.k3:
      return '3';
    case Key.k4:
      return '4';
    case Key.k5:
      return '5';
    case Key.k6:
      return '6';
    case Key.k7:
      return '7';
    case Key.k8:
      return '8';
    case Key.k9:
      return '9';
    case Key.kU:
      return '^';
    case Key.kR:
      return '>';
    case Key.kD:
      return 'v';
    case Key.kL:
      return '<';
    default:
      throw ArgumentError('Invalid key: $key');
  }
}

class Controller {
  String name;
  Key pos = Key.kA;
  Controller? controller;
  Map<Key, Map<Key, List<String>>> pad;

  Controller(this.name, this.pad, [this.controller]);

  // TODO - Cache this shit!
  int moveAndPress(List<Key> keys) {
    pos = Key.kA;

    if (controller == null) {
      // print('$name enters sequence: ${keys.map((e) => keyToString(e)).join('')} in ${keys.length} moves');
      return keys.length;
    }

    // print('$name is requesting button sequence: ${keys.map((e) => keyToString(e)).join('')}');
    print('$name is requesting a long (${keys.length}) button sequence');
    List<List<Key>> moveListOpts = [[]];
    for (final key in keys) {
      final opts = pad[pos]![key]!;
      if (opts.length == 2) {
        print('$name: 2 options found from $pos to $key: $opts -- trying both...');
        List<List<Key>> clonedOpts = [];
        for (final moveListOpt in moveListOpts) {
          List<Key> newOpt = List.from(moveListOpt);
          newOpt.addAll(opts[1].split('').map((e) => toKey(e)).toList());
          newOpt.add(Key.kA);
          clonedOpts.add(newOpt);
          moveListOpt.addAll(opts[0].split('').map((e) => toKey(e)).toList());
          moveListOpt.add(Key.kA);
        }
        moveListOpts.addAll(clonedOpts);
      } else if (opts.length == 1) {
        for (final moveListOpt in moveListOpts) {
          moveListOpt.addAll(opts[0].split('').map((e) => toKey(e)).toList());
          moveListOpt.add(Key.kA);
        }
      }
      pos = key;
    }
    int best = 10000000000;
    for (final moveList in moveListOpts) {
      best = min(best, controller!.moveAndPress(moveList));
    }
    return best;
  }
}

mixin Day21 {
  int day21a(List<String> lines) {
    int result = 0;

    // set up robots and codes
    Controller me = Controller('Brady', arrowPad, null);
    Controller robot3 = Controller('coldBot', arrowPad, me);
    Controller robot2 = Controller('radiationBot', arrowPad, robot3);
    Controller robot1 = Controller('numpadBot', numPad, robot2);

    // process each code
    for (final line in lines) {
      print('finding shortest sequence for code $line...');
      // calculate shortest sequence
      final keys = line.split('').map((e) => toKey(e)).toList();
      final numMoves = robot1.moveAndPress(keys);

      // add to result
      result += numMoves * int.parse(line.substring(0, 3));
    }

    return result;
    // 161323 is too high
    // 157908 was correct
  }

  int day21b(List<String> lines) {
    int result = 0;

    // set up robots and codes
    Controller me = Controller('Brady', arrowPad, null);
    Controller r25 = Controller('r25', arrowPad, me);
    Controller r24 = Controller('r24', arrowPad, r25);
    Controller r23 = Controller('r23', arrowPad, r24);
    Controller r22 = Controller('r22', arrowPad, r23);
    Controller r21 = Controller('r21', arrowPad, r22);
    Controller r20 = Controller('r20', arrowPad, r21);
    Controller r19 = Controller('r19', arrowPad, r20);
    Controller r18 = Controller('r18', arrowPad, r19);
    Controller r17 = Controller('r17', arrowPad, r18);
    Controller r16 = Controller('r16', arrowPad, r17);
    Controller r15 = Controller('r15', arrowPad, r16);
    Controller r14 = Controller('r14', arrowPad, r15);
    Controller r13 = Controller('r13', arrowPad, r14);
    Controller r12 = Controller('r12', arrowPad, r13);
    Controller r11 = Controller('r11', arrowPad, r12);
    Controller r10 = Controller('r10', arrowPad, r11);
    Controller r9 = Controller('r09', arrowPad, r10);
    Controller r8 = Controller('r08', arrowPad, r9);
    Controller r7 = Controller('r07', arrowPad, r8);
    Controller r6 = Controller('r06', arrowPad, r7);
    Controller r5 = Controller('r05', arrowPad, r6);
    Controller r4 = Controller('r04', arrowPad, r5);
    Controller r3 = Controller('r03', arrowPad, r4);
    Controller r2 = Controller('r02', arrowPad, r3);
    Controller r1 = Controller('r01', arrowPad, r2);
    Controller r0 = Controller('numpadBot', numPad, r1);

    // process each code
    for (final line in lines) {
      print('finding shortest sequence for code $line...');
      // calculate shortest sequence
      final keys = line.split('').map((e) => toKey(e)).toList();
      final numMoves = r0.moveAndPress(keys);

      // add to result
      result += numMoves * int.parse(line.substring(0, 3));
    }
    return result;
  }
}

const Map<Key, Map<Key, List<String>>> numPad = {
  Key.kA: {
    Key.k0: ['<'],
    Key.k1: ['^<<'],
    Key.k2: ['<^'],
    // Key.k2: ['<^', '^<'],
    Key.k3: ['^'],
    Key.k4: ['^^<<'],
    Key.k5: ['<^^'],
    // Key.k5: ['<^^', '^^<'],
    Key.k6: ['^^'],
    Key.k7: ['^^^<<'],
    Key.k8: ['<^^^'],
    // Key.k8: ['<^^^', '^^^<'],
    Key.k9: ['^^^'],
    Key.kA: [''],
  },
  Key.k0: {
    Key.k0: [''],
    Key.k1: ['^<'],
    Key.k2: ['^'],
    Key.k3: ['^>'],
    // Key.k3: ['>^', '^>'],
    Key.k4: ['^^<'],
    Key.k5: ['^^'],
    Key.k6: ['^^>'],
    // Key.k6: ['>^^', '^^>'],
    Key.k7: ['^^^<'],
    Key.k8: ['^^^'],
    Key.k9: ['^^^>'],
    // Key.k9: ['>^^^', '^^^>'],
    Key.kA: ['>'],
  },
  Key.k1: {
    Key.k0: ['>v'],
    Key.k1: [''],
    Key.k2: ['>'],
    Key.k3: ['>>'],
    Key.k4: ['^'],
    Key.k5: ['^>'],
    // Key.k5: ['>^', '^>'],
    Key.k6: ['^>>'],
    // Key.k6: ['>>^', '^>>'],
    Key.k7: ['^^'],
    Key.k8: ['^^>'],
    // Key.k8: ['>^^', '^^>'],
    Key.k9: ['^^>>'],
    // Key.k9: ['>>^^', '^^>>'],
    Key.kA: ['>>v'],
  },
  Key.k2: {
    Key.k0: ['v'],
    Key.k1: ['<'],
    Key.k2: [''],
    Key.k3: ['>'],
    Key.k4: ['<^'],
    // Key.k4: ['<^', '^<'],
    Key.k5: ['^'],
    Key.k6: ['^>'],
    // Key.k6: ['>^', '^>'],
    Key.k7: ['<^^'],
    // Key.k7: ['<^^', '^^<'],
    Key.k8: ['^^'],
    Key.k9: ['^^>'],
    // Key.k9: ['>^^', '^^>'],
    Key.kA: ['v>'],
    // Key.kA: ['>v', 'v>'],
  },
  Key.k3: {
    Key.k0: ['<v'],
    // Key.k0: ['<v', 'v<'],
    Key.k1: ['<<'],
    Key.k2: ['<'],
    Key.k3: [''],
    Key.k4: ['<<^'],
    // Key.k4: ['<<^', '^<<'],
    Key.k5: ['<^'],
    // Key.k5: ['<^', '^<'],
    Key.k6: ['^'],
    Key.k7: ['<<^^'],
    // Key.k7: ['<<^^', '^^<<'],
    Key.k8: ['<^^'],
    // Key.k8: ['<^^', '^^<'],
    Key.k9: ['^^'],
    Key.kA: ['v'],
  },
  Key.k4: {
    Key.k0: ['>vv'],
    Key.k1: ['v'],
    Key.k2: ['v>'],
    // Key.k2: ['>v', 'v>'],
    Key.k3: ['v>>'],
    // Key.k3: ['>>v', 'v>>'],
    Key.k4: [''],
    Key.k5: ['>'],
    Key.k6: ['>>'],
    Key.k7: ['^'],
    Key.k8: ['^>'],
    // Key.k8: ['>^', '^>'],
    Key.k9: ['^>>'],
    // Key.k9: ['>>^', '^>>'],
    Key.kA: ['>>vv'],
  },
  Key.k5: {
    Key.k0: ['vv'],
    Key.k1: ['<v'],
    // Key.k1: ['<v', 'v<'],
    Key.k2: ['v'],
    Key.k3: ['v>'],
    // Key.k3: ['>v', 'v>'],
    Key.k4: ['<'],
    Key.k5: [''],
    Key.k6: ['>'],
    Key.k7: ['<^'],
    // Key.k7: ['<^', '^<'],
    Key.k8: ['^'],
    Key.k9: ['^>'],
    // Key.k9: ['>^', '^>'],
    Key.kA: ['vv>'],
    // Key.kA: ['>vv', 'vv>'],
  },
  Key.k6: {
    Key.k0: ['<vv'],
    // Key.k0: ['<vv', 'vv<'],
    Key.k1: ['<<v'],
    // Key.k1: ['<<v', 'v<<'],
    Key.k2: ['<v'],
    // Key.k2: ['<v', 'v<'],
    Key.k3: ['v'],
    Key.k4: ['<<'],
    Key.k5: ['<'],
    Key.k6: [''],
    Key.k7: ['<<^'],
    // Key.k7: ['<<^', '^<<'],
    Key.k8: ['<^'],
    // Key.k8: ['<^', '^<'],
    Key.k9: ['^'],
    Key.kA: ['vv'],
  },
  Key.k7: {
    Key.k0: ['>vvv'],
    Key.k1: ['vv'],
    Key.k2: ['vv>'],
    // Key.k2: ['>vv', 'vv>'],
    Key.k3: ['vv>>'],
    // Key.k3: ['>>vv', 'vv>>'],
    Key.k4: ['v'],
    Key.k5: ['v>'],
    // Key.k5: ['>v', 'v>'],
    Key.k6: ['v>>'],
    // Key.k6: ['>>v', 'v>>'],
    Key.k7: [''],
    Key.k8: ['>'],
    Key.k9: ['>>'],
    Key.kA: ['>>vvv'],
  },
  Key.k8: {
    Key.k0: ['vvv'],
    Key.k1: ['<vv'],
    // Key.k1: ['<vv', 'vv<'],
    Key.k2: ['vv'],
    Key.k3: ['vv>'],
    // Key.k3: ['>vv', 'vv>'],
    Key.k4: ['<v'],
    // Key.k4: ['<v', 'v<'],
    Key.k5: ['v'],
    Key.k6: ['v>'],
    // Key.k6: ['>v', 'v>'],
    Key.k7: ['<'],
    Key.k8: [''],
    Key.k9: ['>'],
    Key.kA: ['vvv>'],
    // Key.kA: ['>vvv', 'vvv>'],
  },
  Key.k9: {
    Key.k0: ['<vvv'],
    // Key.k0: ['<vvv', 'vvv<'],
    Key.k1: ['<<vv'],
    // Key.k1: ['<<vv', 'vv<<'],
    Key.k2: ['<vv'],
    // Key.k2: ['<vv', 'vv<'],
    Key.k3: ['vv'],
    Key.k4: ['<<v'],
    // Key.k4: ['<<v', 'v<<'],
    Key.k5: ['<v'],
    // Key.k5: ['<v', 'v<'],
    Key.k6: ['v'],
    Key.k7: ['<<'],
    Key.k8: ['<'],
    Key.k9: [''],
    Key.kA: ['vvv'],
  }
};

const Map<Key, Map<Key, List<String>>> arrowPad = {
  Key.kA: {
    Key.kA: [''],
    Key.kU: ['<'],
    Key.kR: ['v'],
    // Key.kD: ['<v', 'v<'],
    Key.kD: ['<v'],
    Key.kL: ['v<<'],
  },
  Key.kU: {
    Key.kA: ['>'],
    Key.kU: [''],
    // Key.kR: ['v>', '>v'],
    Key.kR: ['v>'],
    Key.kD: ['v'],
    Key.kL: ['v<'],
  },
  Key.kR: {
    Key.kA: ['^'],
    // Key.kU: ['<^', '^<'],
    Key.kU: ['<^'],
    Key.kR: [''],
    Key.kD: ['<'],
    Key.kL: ['<<'],
  },
  Key.kD: {
    // Key.kA: ['^>', '>^'],
    Key.kA: ['^>'],
    Key.kU: ['^'],
    Key.kR: ['>'],
    Key.kD: [''],
    Key.kL: ['<'],
  },
  Key.kL: {
    Key.kA: ['>>^'],
    Key.kU: ['>^'],
    Key.kR: ['>>'],
    Key.kD: ['>'],
    Key.kL: [''],
  }
};

// expect <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
// actual v<<A>>^AvA^Av<<A>>^AAv<A<A>>^AAvAA^<A>Av<A^>AA<A>Av<A<A>>^AAA<Av>A^A