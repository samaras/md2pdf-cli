# Flutter Tutorial: Creating an App with Bottom Navigation and ViewPager

In this tutorial, we will learn how to create a Flutter app with a bottom navigation bar and a ViewPager to navigate between different screens. This will allow users to easily switch between various sections or features within the app.

## Prerequisites

Before we begin, ensure you have Flutter installed on your machine. You can follow the official installation guide [here](https://flutter.dev/docs/get-started/install).

## Creating a New Flutter Project

Let's start by creating a new Flutter project. Open your terminal or command prompt and execute the following commands:

```bash
flutter create bottom_navigation_viewpager
cd bottom_navigation_viewpager
```

## Adding Dependencies

We'll need to add dependencies for the ViewPager and bottom navigation bar. Open your `pubspec.yaml` file and add the following dependencies:

```yaml
dependencies:
  flutter:
    sdk: flutter
  flutter_swiper: ^1.1.6  # For ViewPager
  cupertino_icons: ^1.0.2
```

Save the file and run `flutter pub get` in your terminal to install the dependencies.

## Creating Screens

Create separate Dart files for each screen you want to include in your app. For demonstration purposes, let's create three screens: HomeScreen, ProfileScreen, and SettingsScreen.

### Home Screen

Create a file named `home_screen.dart` and add the following code:

```dart
import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Home'),
      ),
      body: Center(
        child: Text('Home Screen'),
      ),
    );
  }
}
```

### Profile Screen

Create a file named `profile_screen.dart` and add the following code:

```dart
import 'package:flutter/material.dart';

class ProfileScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Profile'),
      ),
      body: Center(
        child: Text('Profile Screen'),
      ),
    );
  }
}
```

### Settings Screen

Create a file named `settings_screen.dart` and add the following code:

```dart
import 'package:flutter/material.dart';

class SettingsScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Settings'),
      ),
      body: Center(
        child: Text('Settings Screen'),
      ),
    );
  }
}
```

## Implementing Bottom Navigation Bar

Now, let's implement the bottom navigation bar in the `main.dart` file.

```dart
import 'package:flutter/material.dart';
import 'package:bottom_navigation_viewpager/home_screen.dart';
import 'package:bottom_navigation_viewpager/profile_screen.dart';
import 'package:bottom_navigation_viewpager/settings_screen.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Bottom Navigation ViewPager',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: BottomNavigation(),
    );
  }
}

class BottomNavigation extends StatefulWidget {
  @override
  _BottomNavigationState createState() => _BottomNavigationState();
}

class _BottomNavigationState extends State<BottomNavigation> {
  int _selectedIndex = 0;

  final List<Widget> _screens = [
    HomeScreen(),
    ProfileScreen(),
    SettingsScreen(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _screens[_selectedIndex],
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
        items: [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: 'Settings',
          ),
        ],
      ),
    );
  }
}
```

## Adding ViewPager

We'll use the `flutter_swiper` package to implement ViewPager functionality.

```dart
import 'package:flutter/material.dart';
import 'package:flutter_swiper/flutter_swiper.dart';
import 'package:bottom_navigation_viewpager/home_screen.dart';
import 'package:bottom_navigation_viewpager/profile_screen.dart';
import 'package:bottom_navigation_viewpager/settings_screen.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Bottom Navigation ViewPager',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: BottomNavigation(),
    );
  }
}

class BottomNavigation extends StatefulWidget {
  @override
  _BottomNavigationState createState() => _BottomNavigationState();
}

class _BottomNavigationState extends State<BottomNavigation> {
  int _selectedIndex = 0;

  final List<Widget> _screens = [
    HomeScreen(),
    ProfileScreen(),
    SettingsScreen(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Swiper(
        itemBuilder: (BuildContext context, int index) {
          return _screens[index];
        },
        itemCount: _screens.length,
        loop: false,
        index: _selectedIndex,
        onIndexChanged: _onItemTapped,
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
        items: [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: 'Settings',
          ),
        ],
      ),
    );
  }
}
```

## Conclusion

Congratulations! You have successfully created a Flutter app with a bottom navigation bar and a ViewPager to navigate between different screens. Feel free to customize and extend this app according to your requirements. Happy coding!
