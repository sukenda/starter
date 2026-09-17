import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:starter_mobile/design_system/tokens.dart';
import 'package:starter_mobile/features/auth/application/auth_controller.dart';

class LoginPage extends ConsumerStatefulWidget {
  const LoginPage({super.key});

  @override
  ConsumerState<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends ConsumerState<LoginPage> {
  final _formKey = GlobalKey<FormState>();
  final _email = TextEditingController();
  final _password = TextEditingController();

  @override
  void dispose() {
    _email.dispose();
    _password.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;
    FocusScope.of(context).unfocus();
    await ref
        .read(authControllerProvider.notifier)
        .login(_email.text.trim(), _password.text);
  }

  @override
  Widget build(BuildContext context) {
    final auth = ref.watch(authControllerProvider);
    final text = Theme.of(context).textTheme;

    return Scaffold(
      body: SafeArea(
        child: Center(
          child: SingleChildScrollView(
            padding: const EdgeInsets.all(DsSpace.lg),
            child: ConstrainedBox(
              constraints: const BoxConstraints(maxWidth: 420),
              child: Card(
                child: Padding(
                  padding: const EdgeInsets.all(DsSpace.lg),
                  child: Form(
                    key: _formKey,
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.stretch,
                      children: [
                        Text(
                          'FULLSTACK STARTER',
                          style: text.labelSmall?.copyWith(
                            color: DsColors.brand,
                            fontWeight: FontWeight.w700,
                            letterSpacing: 1.2,
                          ),
                        ),
                        const SizedBox(height: DsSpace.sm),
                        Text(
                          'Welcome back',
                          style: text.headlineMedium?.copyWith(
                            fontWeight: FontWeight.w700,
                          ),
                        ),
                        const SizedBox(height: DsSpace.sm),
                        Text(
                          'Sign in to continue to your workspace.',
                          style: text.bodyMedium?.copyWith(
                            color: DsColors.textSecondary,
                          ),
                        ),
                        const SizedBox(height: DsSpace.xl),
                        TextFormField(
                          controller: _email,
                          keyboardType: TextInputType.emailAddress,
                          autofillHints: const [AutofillHints.username],
                          decoration: const InputDecoration(labelText: 'Email'),
                          validator: (value) =>
                              value == null || !value.contains('@')
                              ? 'Enter a valid email.'
                              : null,
                        ),
                        const SizedBox(height: DsSpace.md),
                        TextFormField(
                          controller: _password,
                          obscureText: true,
                          autofillHints: const [AutofillHints.password],
                          decoration: const InputDecoration(
                            labelText: 'Password',
                          ),
                          validator: (value) => value == null || value.isEmpty
                              ? 'Password is required.'
                              : null,
                        ),
                        if (auth.hasError) ...[
                          const SizedBox(height: DsSpace.md),
                          Text(
                            'Unable to sign in. Check your credentials and try again.',
                            style: TextStyle(
                              color: Theme.of(context).colorScheme.error,
                            ),
                          ),
                        ],
                        const SizedBox(height: DsSpace.lg),
                        FilledButton(
                          onPressed: auth.isLoading ? null : _submit,
                          child: Text(
                            auth.isLoading ? 'Signing in…' : 'Sign in',
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
