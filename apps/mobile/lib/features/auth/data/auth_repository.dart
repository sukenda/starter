import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:starter_mobile/core/network/api_client.dart';
import 'package:starter_mobile/features/auth/data/token_storage.dart';
import 'package:starter_mobile/features/auth/domain/session.dart';

final authRepositoryProvider = Provider<AuthRepository>((ref) => AuthRepository(
      dio: ref.watch(dioProvider),
      storage: ref.watch(tokenStorageProvider),
    ));

class AuthRepository {
  AuthRepository({required Dio dio, required TokenStorage storage})
      : _dio = dio,
        _storage = storage;
  final Dio _dio;
  final TokenStorage _storage;

  Future<SessionUser> login(String email, String password) async {
    final response = await _dio.post<Map<String, dynamic>>(
      '/api/v1/auth/login',
      data: {'email': email, 'password': password},
    );
    final data = response.data!['data'] as Map<String, dynamic>;
    final tokens = SessionTokens.fromJson(data['tokens'] as Map<String, dynamic>);
    await _storage.save(tokens);
    return SessionUser.fromJson(data['user'] as Map<String, dynamic>);
  }

  Future<SessionUser?> restore() async {
    final accessToken = await _storage.readAccessToken();
    if (accessToken == null) return null;
    try {
      return await me(accessToken);
    } on DioException {
      await _storage.clear();
      return null;
    }
  }

  Future<SessionUser> me(String accessToken) async {
    final response = await _dio.get<Map<String, dynamic>>(
      '/api/v1/auth/me',
      options: Options(headers: {'Authorization': 'Bearer $accessToken'}),
    );
    final data = response.data!['data'] as Map<String, dynamic>;
    return SessionUser.fromJson(data);
  }

  Future<void> logout() async {
    final accessToken = await _storage.readAccessToken();
    if (accessToken != null) {
      try {
        await _dio.post<void>('/api/v1/auth/logout',
            options: Options(headers: {'Authorization': 'Bearer $accessToken'}));
      } on DioException {
        // Local credentials must still be removed when the network is unavailable.
      }
    }
    await _storage.clear();
  }
}
