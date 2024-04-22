import 'dart:io';

import 'package:ar_flutter_plugin/managers/ar_location_manager.dart';
import 'package:ar_flutter_plugin/managers/ar_session_manager.dart';
import 'package:ar_flutter_plugin/managers/ar_object_manager.dart';
import 'package:ar_flutter_plugin/managers/ar_anchor_manager.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:ar_flutter_plugin/ar_flutter_plugin.dart';
import 'package:ar_flutter_plugin/datatypes/config_planedetection.dart';
import 'package:ar_flutter_plugin/datatypes/node_types.dart';
import 'package:ar_flutter_plugin/models/ar_node.dart';
import 'package:vector_math/vector_math_64.dart';
import 'package:path_provider/path_provider.dart';

class ArAppDemo extends StatefulWidget {
  const ArAppDemo({super.key});

  @override
  State<ArAppDemo> createState() => _ArAppDemoState();
}

class _ArAppDemoState extends State<ArAppDemo> {
  ARSessionManager? arSessionManager;
  ARObjectManager? arObjectManager;
  //String localObjectReference;
  ARNode? localObjectNode;
  //String webObjectReference;
  ARNode? webObjectNode;
  ARNode? fileSystemNode;
  HttpClient? httpClient;
  //use to rotate model
  var rotateAxis = [0.0, 0.0, 0.0];
  final scaleNode = 0.2;
  final originPosition = Vector3(0.0, -0.1, -0.2);
  var isLoading = true;

  @override
  void dispose() {
    super.dispose();
    arSessionManager!.dispose();
  }

  @override
  void initState() {
    super.initState();

    //Download model to file system
    httpClient = HttpClient();
    _downloadFile(
        "https://cdn.discordapp.com/attachments/1069987642378305656/1230814278471122965/AnubisRaw.glb?ex=6634b01f&is=66223b1f&hm=871fffe42e5e22f0a49d2e056b16b87dc1a7415ee46be5f28f0003b5d4501bd3&",
        "LocalGLB.glb");
  }

  @override
  Widget build(BuildContext context) {
    return (!isLoading)
        ? Scaffold(
            appBar: AppBar(
              title: const Text('Local & Web Objects'),
            ),
            // ignore: avoid_unnecessary_containers
            body: Container(
                child: Stack(children: [
                  ARView(
                    onARViewCreated: onARViewCreated,
                    planeDetectionConfig:
                        PlaneDetectionConfig.horizontalAndVertical,
                  ),
                  Align(
                      alignment: FractionalOffset.bottomCenter,
                      child: Column(
                          mainAxisAlignment: MainAxisAlignment.end,
                          children: [
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                              children: [
                                ElevatedButton(
                                  onPressed: () => {scaleNodeFileSys(0.02)},
                                  child: const Text("+"),
                                ),
                                ElevatedButton(
                                  onPressed: () => {scaleNodeFileSys(-0.02)},
                                  child: const Text("-"),
                                ),
                                ElevatedButton(
                                    onPressed: () => {moveNodeFileSys(0, 0, -0.1)},
                                    child: const Icon(Icons.arrow_circle_up)),
                                ElevatedButton(
                                    onPressed: () => {moveNodeFileSys(0, 0, 0.1)},
                                    child: const Icon(Icons.arrow_circle_down)),
                              ],
                            ),
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                              children: [
                                ElevatedButton(
                                  onPressed: () => {moveNodeFileSys(-0.1, 0, 0)},
                                  child: const Icon(Icons.arrow_left),
                                ),
                                ElevatedButton(
                                    onPressed: () => {moveNodeFileSys(0.1, 0, 0)},
                                    child: const Icon(Icons.arrow_right)),
                                ElevatedButton(
                                    onPressed: () => {moveNodeFileSys(0, 0.1, 0)},
                                    child: const Icon(Icons.arrow_drop_up)),
                                ElevatedButton(
                                    onPressed: () => {moveNodeFileSys(0, -0.1, 0)},
                                    child: const Icon(Icons.arrow_drop_down)),
                              ],
                            ),
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                              children: [
                                ElevatedButton(
                                    onPressed: () => {rotateNodeFileSys(1, -0.5)},
                                    child: const Icon(Icons.change_circle_rounded)),
                                ElevatedButton(
                                    onPressed: () => {rotateNodeFileSys(0, -0.5)},
                                    child:
                                        const Icon(Icons.change_circle_outlined)),
                                ElevatedButton(
                                    onPressed: () => {rotateNodeFileSys(0, 0.5)},
                                    child:
                                        const Icon(Icons.change_circle_outlined)),
                                ElevatedButton(
                                    onPressed: () => {rotateNodeFileSys(1, 0.5)},
                                    child: const Icon(Icons.change_circle_rounded)),
                              ],
                            ),
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                              children: [
                                ElevatedButton(
                                    onPressed:
                                        onFileSystemObjectAtOriginButtonPressed,
                                    child: const Text(
                                        "Add/Remove Filesystem\nObject at Origin")),
                                ElevatedButton(
                                    onPressed: resetNodeFileSys,
                                    child: const Text("Reset")),
                              ],
                            ),
                          ]))
                ])))
        : const Scaffold(
          body: Text("check"),
        );
  }

  void onARViewCreated(
      ARSessionManager arSessionManager,
      ARObjectManager arObjectManager,
      ARAnchorManager arAnchorManager,
      ARLocationManager arLocationManager) {
    this.arSessionManager = arSessionManager;
    this.arObjectManager = arObjectManager;

    this.arSessionManager!.onInitialize(
          showFeaturePoints: false,
          showPlanes: true,
          customPlaneTexturePath: "Images/triangle.png",
          showWorldOrigin: false,
          handleTaps: false,
        );
    this.arObjectManager!.onInitialize();
  }

  Future<File> _downloadFile(String url, String filename) async {
    var request = await httpClient!.getUrl(Uri.parse(url));
    var response = await request.close();
    var bytes = await consolidateHttpClientResponseBytes(response);
    String dir = (await getApplicationDocumentsDirectory()).path;
    File file = File('$dir/$filename');
    // write file to local
    try {
      await file.writeAsBytes(bytes);
    } on Exception catch (_) {
      throw Exception("can't write file");
    }
    debugPrint("Downloading finished, path: " '$dir/$filename');
    isLoading = false;
    return file;
  }

  Future<void> onFileSystemObjectAtOriginButtonPressed() async {
    if (fileSystemNode != null) {
      arObjectManager!.removeNode(fileSystemNode!);
      fileSystemNode = null;
    } else {
      var newNode = ARNode(
          type: NodeType.fileSystemAppFolderGLB,
          uri: "LocalGLB.glb",
          scale: Vector3(scaleNode, scaleNode, scaleNode),
          position: originPosition);
      bool? didAddFileSystemNode = await arObjectManager!.addNode(newNode);
      fileSystemNode = (didAddFileSystemNode!) ? newNode : null;
    }
  }

  Future<void> moveNodeFileSys(double x, double y, double z) async {
    if (fileSystemNode != null) {
      var newTranslation = fileSystemNode!.position;
      newTranslation = Vector3(
          newTranslation.x + x, newTranslation.y + y, newTranslation.z + z);
      final newTransform = Matrix4.identity();
      newTransform.setTranslation(newTranslation);
      newTransform.scale(fileSystemNode!.scale);
      saveRotate(newTransform);

      fileSystemNode!.transform = newTransform;
    }
  }

  void saveRotate(Matrix4 tranForm) {
    for (int i = 0; i < 3; i++) {
      if (rotateAxis[i] != 0) {
        var newRotationAxis = Vector3(0, 0, 0);
        newRotationAxis[i] = 1.0;
        tranForm.rotate(newRotationAxis, rotateAxis[i]);
      }
    }
  }

  Future<void> rotateNodeFileSys(int xyz, double angle) async {
    if (fileSystemNode != null) {
      final newTransform = Matrix4.identity();
      newTransform.scale(fileSystemNode!.scale);
      newTransform.setTranslation(fileSystemNode!.position);

      var newRotationAxis = Vector3(0, 0, 0);
      newRotationAxis[xyz] = 1.0;
      rotateAxis[xyz] += angle;
      saveRotate(newTransform);

      fileSystemNode!.transform = newTransform;
    }
  }

  Future<void> resetNodeFileSys() async {
    final newTransform = Matrix4.identity();
    // newTransform.scale(fileSystemNode!.scale);
    newTransform.setTranslation(originPosition);
    var newRotationAxis = Vector3(1, 1, 1);
    newTransform.rotate(newRotationAxis, 0);
    newTransform.scale(scaleNode);
    rotateAxis = [0, 0, 0]; // reset rotate
    saveRotate(newTransform);

    fileSystemNode!.transform = newTransform;
  }

  Future<void> scaleNodeFileSys(double upDown) async {
    if (fileSystemNode != null) {
      var newTranslation = fileSystemNode!.position;
      final newTransform = Matrix4.identity();
      newTransform.setTranslation(newTranslation);
      saveRotate(newTransform);
      var transformScale = Vector3(upDown, upDown, upDown);
      newTransform.scale(fileSystemNode!.scale + transformScale);

      fileSystemNode!.transform = newTransform;
    }
  }
}
