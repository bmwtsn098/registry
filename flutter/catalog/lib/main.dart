import 'package:flutter/material.dart';
import 'theme.dart';
import 'routes.dart';

void main() {
  runApp(Application());
}

class Application extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'API Hub',
      theme: appTheme(),
      initialRoute: "/",
      onGenerateRoute: generateRoute,
    );
  }
}
