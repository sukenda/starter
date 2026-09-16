class ApiException implements Exception {
  const ApiException({
    required this.code,
    required this.message,
    this.statusCode,
    this.requestId,
    this.details,
  });

  final String code;
  final String message;
  final int? statusCode;
  final String? requestId;
  final Object? details;

  @override
  String toString() => 'ApiException($code): $message';
}
