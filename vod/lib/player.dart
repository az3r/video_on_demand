import 'dart:developer';

import 'package:chewie/chewie.dart';
import 'package:flutter/material.dart';
import 'package:video_player/video_player.dart';

class Player extends StatefulWidget {
  final String uri;
  const Player({super.key, required this.uri});

  @override
  PlayerState createState() => PlayerState();
}

class PlayerState extends State<Player> {
  late ChewieController _controller;

  @override
  void initState() {
    super.initState();
    _controller = ChewieController(
      autoInitialize: true,
      videoPlayerController: VideoPlayerController.networkUrl(
        Uri.parse(widget.uri),
      ),
    );
  }

  @override
  void dispose() {
    super.dispose();
    _controller.videoPlayerController.dispose();
    _controller.dispose();
  }

  @override
  void didUpdateWidget(covariant Player oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.uri != widget.uri) {
      _controller.videoPlayerController.dispose();
      _controller.dispose();
      _controller = ChewieController(
        autoInitialize: true,
        videoPlayerController: VideoPlayerController.networkUrl(
          Uri.parse(widget.uri),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return ValueListenableBuilder(
      valueListenable: _controller.videoPlayerController,
      builder: (context, value, child) {
        if (value.isInitialized) {
          return AspectRatio(
            aspectRatio: value.aspectRatio,
            child: Chewie(controller: _controller),
          );
        } else {
          return CircularProgressIndicator();
        }
      },
    );
  }
}
