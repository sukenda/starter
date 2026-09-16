import 'package:flutter_test/flutter_test.dart';
import 'package:starter_mobile/core/storage/secure_store.dart';

class _MemorySecureStore implements SecureStore {
  final Map<String, String> _values = {};

  @override
  Future<void> clear() async => _values.clear();

  @override
  Future<void> delete(String key) async => _values.remove(key);

  @override
  Future<String?> read(String key) async => _values[key];

  @override
  Future<void> write(String key, String value) async => _values[key] = value;
}

void main() {
  test('SecureStore contract supports credential lifecycle', () async {
    final store = _MemorySecureStore();
    await store.write('token', 'secret');
    expect(await store.read('token'), 'secret');
    await store.delete('token');
    expect(await store.read('token'), isNull);
  });
}
