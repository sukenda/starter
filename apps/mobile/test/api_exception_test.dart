import 'package:flutter_test/flutter_test.dart';
import 'package:starter_mobile/core/network/api_exception.dart';

void main() {
  test('ApiException preserves normalized metadata', () {
    const exception = ApiException(
      code: 'validation_failed',
      message: 'Request is invalid.',
      statusCode: 422,
      requestId: 'request-1',
    );

    expect(exception.code, 'validation_failed');
    expect(exception.statusCode, 422);
    expect(exception.requestId, 'request-1');
  });
}
