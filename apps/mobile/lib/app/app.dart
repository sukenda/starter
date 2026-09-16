import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:starter_mobile/app/router.dart';
import 'package:starter_mobile/design_system/theme.dart';

class StarterApp extends ConsumerWidget {
  const StarterApp({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return MaterialApp.router(
      title: 'Starter',
      debugShowCheckedModeBanner: false,
      routerConfig: ref.watch(routerProvider),
      theme: StarterTheme.light,
    );
  }
}
