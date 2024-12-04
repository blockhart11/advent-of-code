Advent of Code 2024 in Dart

To run the challenge for a given day, use the Dart command line tool, e.g.:
`dart run advent_of_dart 01`

CLI Flags:

* -t: runs the challenge in test mode, which uses the daily sample input and compares to the expected output. See `input/01_sample.txt` for an example. The expected output is always the first line of that file (part 1 output on the left, then part 2 on the right).
* -b: runs part 2 of the daily challenge. This can be used in conjunction with -t to run against the sample input for part 2.

Example:

The following command solves day 1, part 2 for its sample input

`dart run advent_of_dart 01 -tb`

Output:

`Test run of day 01, part b: expected 31, got 31`