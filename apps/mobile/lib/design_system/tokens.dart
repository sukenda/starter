import 'package:flutter/material.dart';

abstract final class DsColors {
  static const brand = Color(0xFF4F46E5);
  static const brandHover = Color(0xFF4338CA);
  static const canvas = Color(0xFFF8FAFC);
  static const surface = Color(0xFFFFFFFF);
  static const surfaceSubtle = Color(0xFFF1F5F9);
  static const text = Color(0xFF0F172A);
  static const textSecondary = Color(0xFF475569);
  static const border = Color(0xFFE2E8F0);
  static const danger = Color(0xFFB91C1C);
  static const success = Color(0xFF15803D);
}

abstract final class DsSpace {
  static const xs = 4.0, sm = 8.0, md = 16.0, lg = 24.0, xl = 32.0;
}

abstract final class DsRadius {
  static const sm = 6.0, md = 10.0, lg = 14.0, xl = 20.0;
}
