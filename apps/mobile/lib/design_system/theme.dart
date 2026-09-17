import 'package:flutter/material.dart';
import 'package:starter_mobile/design_system/tokens.dart';

abstract final class StarterTheme {
  static ThemeData get light {
    final scheme =
        ColorScheme.fromSeed(
          seedColor: DsColors.brand,
          brightness: Brightness.light,
        ).copyWith(
          primary: DsColors.brand,
          surface: DsColors.surface,
          error: DsColors.danger,
        );
    return ThemeData(
      useMaterial3: true,
      colorScheme: scheme,
      scaffoldBackgroundColor: DsColors.canvas,
      inputDecorationTheme: InputDecorationTheme(
        filled: true,
        fillColor: DsColors.surface,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(DsRadius.md),
          borderSide: const BorderSide(color: DsColors.border),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(DsRadius.md),
          borderSide: const BorderSide(color: DsColors.border),
        ),
      ),
      cardTheme: CardThemeData(
        color: DsColors.surface,
        elevation: 0,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(DsRadius.lg),
          side: const BorderSide(color: DsColors.border),
        ),
      ),
      filledButtonTheme: FilledButtonThemeData(
        style: FilledButton.styleFrom(
          minimumSize: const Size(0, 48),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(DsRadius.md),
          ),
          textStyle: const TextStyle(fontWeight: FontWeight.w600),
        ),
      ),
    );
  }
}
