import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:starter_mobile/features/auth/data/auth_repository.dart';
import 'package:starter_mobile/features/auth/domain/session.dart';

final authControllerProvider = AsyncNotifierProvider<AuthController, SessionUser?>(
  AuthController.new,
);

class AuthController extends AsyncNotifier<SessionUser?> {
  @override
  Future<SessionUser?> build() => ref.read(authRepositoryProvider).restore();

  Future<bool> login(String email, String password) async {
    state = const AsyncLoading();
    state = await AsyncValue.guard(
      () => ref.read(authRepositoryProvider).login(email, password),
    );
    return !state.hasError;
  }

  Future<void> logout() async {
    await ref.read(authRepositoryProvider).logout();
    state = const AsyncData(null);
  }
}
