import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:starter_mobile/core/config/app_config.dart';
import 'package:starter_mobile/core/network/api_exception.dart';

final dioProvider = Provider<Dio>((ref) {
  final dio = Dio(
    BaseOptions(
      baseUrl: AppConfig.apiBaseUrl,
      connectTimeout: const Duration(seconds: 10),
      receiveTimeout: const Duration(seconds: 15),
      sendTimeout: const Duration(seconds: 15),
      headers: const {'Accept': 'application/json'},
    ),
  );

  dio.interceptors.add(
    InterceptorsWrapper(
      onError: (error, handler) {
        final response = error.response;
        final body = response?.data;
        final apiError = body is Map<String, dynamic> ? body['error'] : null;
        if (apiError is Map<String, dynamic>) {
          return handler.reject(
            DioException(
              requestOptions: error.requestOptions,
              response: response,
              error: ApiException(
                code: apiError['code'] as String? ?? 'request_failed',
                message: apiError['message'] as String? ?? 'Request failed.',
                statusCode: response?.statusCode,
                requestId: apiError['request_id'] as String?,
                details: apiError['details'],
              ),
              type: error.type,
            ),
          );
        }
        handler.next(error);
      },
    ),
  );

  return dio;
});
