import 'package:flutter_test/flutter_test.dart';
import 'package:starter_mobile/core/config/app_config.dart';

void main() {
  test('API base URL has a development default', () {
    expect(AppConfig.apiBaseUrl, isNotEmpty);
  });
}
