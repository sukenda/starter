import 'dart:async';

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:starter_mobile/core/network/api_client.dart';
import 'package:starter_mobile/features/auth/data/auth_repository.dart';
import 'package:starter_mobile/features/auth/data/token_storage.dart';

final authenticatedDioProvider = Provider<Dio>((ref) {
  final dio = ref.watch(dioProvider);
  final storage = ref.watch(tokenStorageProvider);
  final authRepository = ref.watch(authRepositoryProvider);

  dio.interceptors.insert(
    0,
    QueuedInterceptorsWrapper(
      onRequest: (options, handler) async {
        if (_isAuthEndpoint(options.path)) {
          return handler.next(options);
        }
        final accessToken = await storage.readAccessToken();
        if (accessToken != null && accessToken.isNotEmpty) {
          options.headers['Authorization'] = 'Bearer $accessToken';
        }
        handler.next(options);
      },
      onError: (error, handler) async {
        final request = error.requestOptions;
        if (error.response?.statusCode != 401 ||
            _isAuthEndpoint(request.path) ||
            request.extra[_retryKey] == true) {
          return handler.next(error);
        }

        final tokens = await authRepository.refresh();
        if (tokens == null) return handler.next(error);

        request.extra[_retryKey] = true;
        request.headers['Authorization'] = 'Bearer ${tokens.accessToken}';
        try {
          final response = await dio.fetch<dynamic>(request);
          return handler.resolve(response);
        } on DioException catch (retryError) {
          return handler.next(retryError);
        }
      },
    ),
  );

  return dio;
});

const _retryKey = 'auth_refresh_retried';

bool _isAuthEndpoint(String path) =>
    path.endsWith('/api/v1/auth/login') ||
    path.endsWith('/api/v1/auth/refresh');
