class SessionUser {
  const SessionUser({
    required this.id,
    required this.email,
    required this.name,
  });

  final String id;
  final String email;
  final String name;

  factory SessionUser.fromJson(Map<String, dynamic> json) => SessionUser(
    id: json['id'] as String,
    email: json['email'] as String,
    name: json['name'] as String,
  );
}

class SessionTokens {
  const SessionTokens({
    required this.accessToken,
    required this.refreshToken,
  });

  final String accessToken;
  final String refreshToken;

  factory SessionTokens.fromJson(Map<String, dynamic> json) => SessionTokens(
    accessToken: json['access_token'] as String,
    refreshToken: json['refresh_token'] as String,
  );
}
