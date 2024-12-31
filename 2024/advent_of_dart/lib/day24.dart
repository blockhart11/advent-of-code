import 'dart:math';

import 'package:advent_of_dart/utils/tree.dart';

class Wire {
  String id;
  bool? val;

  Wire(this.id, [this.val]);

  @override
  String toString() {
    return '$id${val != null ? (val! ? '(1)' : '(0)') : ''}';
  }
}

enum Op { and, or, xor }

Op opFromString(String op) {
  switch (op) {
    case 'AND':
      return Op.and;
    case 'OR':
      return Op.or;
    case 'XOR':
      return Op.xor;
  }
  throw ('no op for $op');
}

String opToSymbol(Op op) {
  switch (op) {
    case Op.and:
      return '&';
    case Op.or:
      return '|';
    case Op.xor:
      return '^';
  }
}

class Gate {
  Op op;
  Wire lhs;
  Wire rhs;
  Wire out;
  int? level;

  Gate(this.op, this.lhs, this.rhs, this.out);

  // returns false if the input wires are not ready
  bool setOut() {
    if (
      lhs.val == null
      || rhs.val == null
      || out.val != null
    ) return false;
    switch (op) {
      case Op.and:
        out.val = lhs.val! && rhs.val!;
      case Op.or:
        out.val = lhs.val! || rhs.val!;
      case Op.xor:
        out.val = lhs.val! ^ rhs.val!;
    }
    return true;
  }

  @override
  String toString() {
    return '$lhs ${opToSymbol(op)} $rhs ${out.id.startsWith('z') ? '-------------> $out' : '--> $out'}';
  }
}

mixin Day24 {
  int day24a(List<String> lines) {
    int result = 0;
    List<Wire> wires = [];
    List<Gate> gates = [];

    int i = 0;
    for (; lines[i].isNotEmpty; i++) {
      final data = lines[i].split(': ');
      wires.add(Wire(data[0], data[1] == '1'));
    }
    i++;
    for (; i < lines.length; i++) {
      // load gates
      final data = lines[i].split(' ');
      final lhs = wires.firstWhere((e) => e.id == data[0], orElse: () {
        Wire w = Wire(data[0]);
        wires.add(w);
        return w;
      });
      final rhs = wires.firstWhere((e) => e.id == data[2], orElse: () {
        Wire w = Wire(data[2]);
        wires.add(w);
        return w;
      });
      final out = wires.firstWhere((e) => e.id == data[4], orElse: () {
        Wire w = Wire(data[4]);
        wires.add(w);
        return w;
      });
      final op = opFromString(data[1]);
      gates.add(Gate(op, lhs, rhs, out));
    }

    List<List<Gate>> gTree = [];
    for (int i = 0; gates.isNotEmpty; i++) {
      print('processing gates, ${gates.length} left...');
      for (int j = 0; j < gates.length; j++) {
        if (gates[j].setOut()) {
          (gTree.length == i) ? gTree.add([gates[j]]) : gTree[i].add(gates[j]);
          gates.removeAt(j);
          j--;
        }
      }
    }

    // print the gate tree
    print('gate tree:');
    for (int i = 0; i < gTree.length; i++) {
      print('\n***LEVEL $i***\n');
      print(gTree[i].map((e) => e.toString()).join('\n'));
    }

    List<Wire> zWires = [];
    for (final wire in wires) {
      if (wire.id.startsWith('z')) {
        zWires.add(wire);
      }
    }
    zWires.sort((lhs, rhs) => lhs.id.compareTo(rhs.id));

    for (int i = 0; i < zWires.length; i++) {
      // print('${zWires[i].id}: ${zWires[i].val! ? '1' : '0'}');
      if (zWires[i].val!) {
        result += pow(2, i).toInt();
      }
    }

    return result;
  }

  // Note: This code DOES NOT SOLVE the puzzle. I worked it out manually once I figured out the structure of the input machine. It was not easy.
  int day24b(List<String> lines) {
    int result = 0;
    List<Wire> wires = [];
    List<Gate> gates = [];

    int i = 0;
    for (; lines[i].isNotEmpty; i++) {
      final data = lines[i].split(': ');
      wires.add(Wire(data[0], data[1] == '1'));
    }
    i++;
    for (; i < lines.length; i++) {
      // load gates
      final data = lines[i].split(' ');
      final lhs = wires.firstWhere((e) => e.id == data[0], orElse: () {
        Wire w = Wire(data[0]);
        wires.add(w);
        return w;
      });
      final rhs = wires.firstWhere((e) => e.id == data[2], orElse: () {
        Wire w = Wire(data[2]);
        wires.add(w);
        return w;
      });
      final out = wires.firstWhere((e) => e.id == data[4], orElse: () {
        Wire w = Wire(data[4]);
        wires.add(w);
        return w;
      });
      final op = opFromString(data[1]);
      gates.add(Gate(op, lhs, rhs, out));
    }

    // print the zWires and how they're computed
    List<Wire> zWires = [];
    for (final wire in wires) {
      if (wire.id.startsWith('z')) {
        zWires.add(wire);
      }
    }
    zWires.sort((lhs, rhs) => lhs.id.compareTo(rhs.id));

    for (int i = 0; i < zWires.length; i++) {
      TreeNode<Wire> t = TreeNode(zWires[i]);
      addChildrenRecursive(t, gates);
      printTree(t);
    }

    return result;
  }

  void addChildrenRecursive(TreeNode<Wire> t, List<Gate> gates) {
    final gate = gates.firstWhere((e) => e.out.id == t.value.id);
    final lhsChild = TreeNode(gate.lhs);
    final rhsChild = TreeNode(gate.rhs);
    t.addChild(lhsChild);
    t.addChild(rhsChild);
    if (!lhsChild.value.id.startsWith(RegExp('x|y'))) addChildrenRecursive(lhsChild, gates);
    if (!rhsChild.value.id.startsWith(RegExp('x|y'))) addChildrenRecursive(rhsChild, gates);
  }
}
