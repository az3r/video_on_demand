import 'dart:developer';

import 'package:chewie/chewie.dart';
import 'package:flutter/material.dart';
import 'package:video_player/video_player.dart';
import 'package:vod/player.dart';

void main() {
  runApp(const MainApp());
}

class MainApp extends StatefulWidget {
  const MainApp({super.key});

  @override
  State<MainApp> createState() => _MainAppState();
}

class _MainAppState extends State<MainApp> {
  late String _uri;
  late ChewieController controller;

  @override
  void initState() {
    super.initState();
    _uri = "http://10.0.2.2:3000/storages/2EZguViS/index.m3u8";
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        floatingActionButton: IconButton.filled(
          onPressed: () {
            setState(() {
              _uri = "http://10.0.2.2:3000/storages/UWfBrGwM/index.m3u8";
            });
          },
          icon: const Icon(Icons.play_arrow),
        ),
        body: Center(child: Player(uri: _uri)),
      ),
    );
  }
}
