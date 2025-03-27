import 'package:dart/data_structure/array.dart';
import 'package:test/test.dart';

void main() {
  test('Array getLast', () {
    final arr = Array(data: [1, 2, 4, 5, 6]);

    expect(arr.getLast(), 6);
  });

  test('Array getFirst', () {
    final arr = Array<int>(data: []);
    expect(arr.getFirst(), null);
  });

  test('Array getSize', () {
    final arr = Array<int>(data: [2, 4, 5, 6]);
    expect(arr.getSize(), 4);
  });
}
