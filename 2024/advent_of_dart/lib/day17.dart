import 'dart:math';

class Computer {
  int rA = 0;
  int rB = 0;
  int rC = 0;
  List<int> program = []; // (opcode, operand)

  Computer(this.rA, this.rB, this.rC);

  void setRegisters(int a, int b, int c) {
    rA = a;
    rB = b;
    rC = c;
  }

  void compile(String program) {
    this.program = program.split(',').map((e) => int.parse(e)).toList();
  }

  List<int> run({bool haltOnOut = false, bool mustMatch = false}) {
    // run compiled program
    List<int> out = [];
    int outTracker = 0;
    int rAStart = rA;
    // print(
    //     'running program $program with starting registers\nA: $rA\nB: $rB\nC: $rC');

    int i = 0; // instruction pointer
    while (i < program.length) {
      final opcode = program[i];
      final operand = program[i + 1];
      // print('op: $opcode, operand: $operand, A: $rA, B: $rB, C: $rC');
      switch (opcode) {
        case 0: // adv -- divide A / 2^combo-op, store in A
          rA = (rA / pow(2, combo(operand))).floor();
          // print('rA = rA / 2^combo --> $rA');
        case 1: // bxl -- bitwise XOR of B
          rB = rB ^ operand;
          // print('rB = rB ^ operand --> $rB');
        case 2: // bst -- operand mod 8
          rB = combo(operand) % 8;
          // print('rB = combo % 8 --> $rB');
        case 3: // jnz -- jump
          if (rA != 0) {
            i = operand;
            // print('jump to $i');
            continue;
          } else {
            // print('rA is zero. do not jump');
          }
        case 4: // bxc -- bitwise XOR of B and C
          rB = rB ^ rC;
          // print('rB = rB ^ rC --> $rB');
        case 5: // out -- mod 8, then output
          int nextOut = (combo(operand) % 8);
          if (haltOnOut) {
            return [nextOut];
          }
          if (mustMatch && nextOut != program[outTracker]) {
            return [];
          }
          out.add(nextOut);
          outTracker++;
        case 6: // bdv -- divide A / 2^op, store in B
          rB = (rA / pow(2, combo(operand))).floor();
          // print('rB = rA / 2^combo --> $rB');
        case 7: // cdv -- divide A / 2^op, store in C
          rC = (rA / pow(2, combo(operand))).floor();
          // print('rC = rA / 2^combo --> $rC');
        default:
          throw ('Unknown opcode $opcode');
      }
      i += 2;
    }

    return out;
  }

  int combo(int op) {
    switch (op) {
      case 0:
      case 1:
      case 2:
      case 3:
        return op;
      case 4:
        return rA;
      case 5:
        return rB;
      case 6:
        return rC;
      case 7:
        throw ('reserved operand $op found');
    }
    throw ('invalid operand $op');
  }
}

mixin Day17 {
  int day17a(List<String> lines) {
    int result = 0;
    final a = int.parse(lines[0].substring(12));
    final b = int.parse(lines[1].substring(12));
    final c = int.parse(lines[2].substring(12));
    final program = lines[4].substring(9);

    Computer computer = Computer(247839653009594, b, c)..compile(program);
    // Computer computer = Computer(a, b, c)..compile(program);
    print('Output: ${computer.run()}');

    return result;
  }

  int day17b(List<String> lines) {
    int result = 0;
    final a = int.parse(lines[0].substring(12));
    final b = int.parse(lines[1].substring(12));
    final c = int.parse(lines[2].substring(12));
    final program = lines[4].substring(9);
    // final program = lines[4].substring(9).split(',').map((e) => int.parse(e)).toList();

    Computer computer = Computer(a, b, c)..compile(program);

    // // final start = 35184372088832;
    final start = 247839571658938;
    final end = 281474976710656;

    for (int i = start; i < end; i++) {
      final b = (i % 8) ^ 1;
      if ((b ^ (i/pow(2, b)).floor()) % 8 != 4) continue;
      computer.setRegisters(i, 0, 0);
      // String out = computer.run();
      // if (i % 10000 == 0) print('running program with A = $i... ${out.replaceAll(',', '').length} digits output: $out');
      final out = computer.run(mustMatch: true);
      if (out.length == 16 && out.join(',') == program) {
        print('program outputs itself ($out) at A = $i');
        break;
      } else if (i % 10000000 == 0) {
        print('A = $i, ${((i - start)/(end - start)).toStringAsFixed(3)}% complete');
      }
    }

    // let's try this a more clever way...
    // int a = 0;
    // final computer = Computer(a, 0, 0)..compile(program);
    // for (int i = computer.program.length - 1; i >= 0; i--) {
    //   List<int> firstOut = [];
    //   while (true) {
    //     computer.setRegisters(a, 0, 0);
    //     firstOut = computer.run(haltOnOut: true);
    //     if (firstOut[0] == computer.program[i]) break;
    //     a++;
    //   }

    //   computer.setRegisters(a, 0, 0);
    //   print('output when A = $a: ${computer.run()}');
    //   a *= 8;
    // }

    return result;
  }
}
