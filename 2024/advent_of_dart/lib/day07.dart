import 'dart:io';
import 'dart:math';

mixin Day07 {
  int day07a(List<String> lines) {
    int result = 0;

    for (int lineIdx = 0; lineIdx < lines.length; lineIdx++) {
      List<String> sides = lines[lineIdx].split(':');
      int lhs = int.parse(sides[0]); // equation result
      List<int> rhs =
          sides[1].substring(1).split(" ").map(int.parse).toList(); // operands
      // print('testing operands: $rhs, answer: $lhs');

      // short circuit test 1 - addition
      int rhsSum = rhs.fold(0, (prev, op) => prev + op);
      if (lhs < rhsSum) {
        // print('short circuiting because sum of rhs > lhs\n$rhsSum >\n$lhs');
        continue;
      }
      // short circuit test 2 - multiplication
      int rhsProduct = rhs.fold(1, (prev, op) {
        if (op == 1) return prev + 1;
        return (prev > lhs) ? prev : prev * op;
      });
      if (lhs > rhsProduct) {
        // print('short circuiting because product of rhs < lhs\n$rhsProduct <\n$lhs');
        continue;
      }

      int opPerms = pow(2, (rhs.length - 1)).toInt();
      for (int i = 0; i < opPerms; i++) {
        // try each potential operator permutation
        int rhsResult = rhs[0];
        // print('testing with mask $i (${i.toRadixString(2)})');
        for (int j = 1; j < rhs.length; j++) {
          // perform a bitwise mask on i using j, add if zero, multiply if one
          if ((i & pow(2, (j - 1)).toInt()) == 0) {
            rhsResult += rhs[j];
            // print('j: $j, adding ${rhs[j]}, current result is $rhsResult');
          } else {
            rhsResult *= rhs[j];
            // print('j: $j, multiplying ${rhs[j]}, current result is $rhsResult');
          }
          if (rhsResult > lhs) break; // short circuit if already exceeds lhs
        }
        if (rhsResult == lhs) {
          print(
              'VALID solution found for $lhs using mask ${i.toRadixString(2)}');
          result += lhs;
          lines[lineIdx] += ':VALID';
          break;
        }
      }
    }

    // Uncomment this to write valid ones to file
    final f = File('input/07b.txt');
    if (!f.existsSync()) f.createSync();
    f.writeAsStringSync(lines.join('\n'), mode: FileMode.write);

    return result;
  }

  int day07b(List<String> lines) {
    int result = 0;

    for (final line in lines) {
      List<String> sides = line.split(':');
      int lhs = int.parse(sides[0]); // equation result
      if (sides.length == 3) {
        // short circuit: already found to be valid in part a
        print('reusing valid answer from part a: $lhs');
        result += lhs;
        continue;
      }

      List<int> rhs =
          sides[1].substring(1).split(" ").map(int.parse).toList(); // operands
      // print('testing operands: $rhs, answer: $lhs');

      // short circuit test 1 - addition
      // int rhsSum = rhs.fold(0, (prev, op) => prev + op);
      // if (lhs < rhsSum) {
      //   print('short circuiting because sum of rhs > lhs\n$rhsSum >\n$lhs');
      //   continue;
      // }

      bool foundSoln = false;
      int opPerms = pow(3, (rhs.length - 1)).toInt();
      for (int i = 0; i < opPerms; i++) {
        if (foundSoln) break;

        // try each potential operator permutation
        int rhsResult = rhs[0];
        String base3 = i.toRadixString(3).padLeft(rhs.length - 1, '0');
        // print('testing with mask $i ($base3)');
        for (int j = 0; j < base3.length; j++) {
          switch (int.parse(base3[j])) {
            case 0:
              rhsResult += rhs[j + 1];
            // print('j: $j, adding ${rhs[j+1]}, current result is $rhsResult');
            case 1:
              rhsResult *= rhs[j + 1];
            // print(
            // 'j: $j, multiplying ${rhs[j+1]}, current result is $rhsResult');
            case 2:
              rhsResult = int.parse('$rhsResult${rhs[j + 1]}');
            // print(
            // 'j: $j, concatting ${rhs[j+1]}, current result is $rhsResult');
          }
          if (rhsResult > lhs) {
            // short circuit if already exceeds lhs
            // print('INVALID: result ($rhsResult) exceeds answer');
            break;
          }
        }
        if (rhsResult == lhs) {
          print('VALID solution found for $lhs:$rhs using mask $base3');
          result += lhs;
          foundSoln = true;
          break;
        }
      }
    }

    return result;
  }
}
