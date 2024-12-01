import 'package:advent_of_dart/advent_of_dart.dart' as aod;
import 'package:args/args.dart';

void main(List<String> arguments) async {
  final flagParser = ArgParser()
    ..addFlag('test', abbr: 't', help: 'Triggers a test run.')
    ..addFlag('b', abbr: 'b', help: 'Runs part 2.');

  final parsedArgs = flagParser.parse(arguments);

  aod.calculate(arguments[0], parsedArgs.wasParsed('test'), parsedArgs.wasParsed('b'));
}
