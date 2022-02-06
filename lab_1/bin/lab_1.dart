import 'dart:io';

void main(List<String> arguments) {
  final a = [
    input('Enter x for the set A: '),
    input('Enter y for the set A: '),
  ];
  final b = [1, 2, 3];
  final c = [
    input('Enter a for the set C: '),
    input('Enter b for the set C: '),
  ];
  print('Set A: $a');
  print('Set B: $b');
  print('Set C: $c');

  final axb = cartesian([a, b]); // 1
  final bxa = cartesian([b, a]); // 2
  final axa = cartesian([a, a]); // 3
  final axbxc = cartesian([cartesian([a, b]), c]); // 5
  final arb = axb.where((pair) => (pair[0] - pair[1]) % 2 == 0); // 6

  print('A x B = $axb (${axb.length} elements)');
  print('B x A = $axb (${bxa.length} elements)');
  print('A x A = $axb (${axa.length} elements)');

  print('(A x B) x C = $axbxc (${axbxc.length} elements)');
  print('A R B = $arb (${arb.length} elements)');
}

// Using List (has duplicates) instead of Set (no duplicates)
// to keep sync with example that uses std::vector
List<List<T>> cartesian<T>(List<List<T>> args) {
  final head = args[0];
  final tail = args.skip(1).toList();
  final remainder = tail.isNotEmpty ? cartesian(tail) : [[]];
  List<List<T>> product = [];
  for (final h in head) {
    for (final r in remainder) {
      product.add([h, ...r]);
    }
  }
  return product;
}

// Parses integer from input/waits for correct output
int input(String message) {
  while(true) {
    stdout.write(message);
    try {
      final input = stdin.readLineSync();
      return int.parse(input ?? '');
    } on FormatException {
      print('Unable to parse number. Enter an integer.');
    }
  }
}
