import 'package:flutter_test/flutter_test.dart';
import 'package:starter_mobile/core/network/api_client.dart';

void main() {
  test('Dio provider is available for dependency injection', () {
    expect(dioProvider, isNotNull);
  });
}
