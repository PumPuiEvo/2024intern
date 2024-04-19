import 'package:arapp/model/shop_cread_model.dart';
import 'package:arapp/views/testOnlyAR.dart';
import 'package:flutter/material.dart';
import 'package:arapp/app_colors.dart';
import 'package:arapp/model/ar_objects.dart';
import 'package:arapp/views/ar_objects_screen.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  List<ShopCardModel> mycard = [
    ShopCardModel(Icons.shopping_bag, 'Shoe', false, ARObjects.shoe, true),
    ShopCardModel(Icons.apartment, 'Chair', false, ARObjects.chair, false),
    ShopCardModel(Icons.home, 'Chicken', false, ARObjects.chicken, true),
    ShopCardModel(Icons.grade, 'Figure', false, ARObjects.figure, true),
    ShopCardModel(Icons.animation, 'Fox', false, ARObjects.fox, true),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.background,
      appBar: AppBar(),
      body: Column(
        children: <Widget>[
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Text(
              "Let's shopping!",
              style: TextStyle(fontSize: 24, color: AppColors.black54),
            ),
          ),
          Expanded(
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: GridView.count(
                crossAxisCount: 2,
                crossAxisSpacing: 8,
                mainAxisSpacing: 8,
                children: mycard
                    .map(
                      (e) => InkWell(
                        onTap: () => onTap(e),
                        child: Card(
                          color: e.isActive ? AppColors.mainColor : null,
                          child: Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: <Widget>[
                              Icon(
                                e.icon,
                                size: 50,
                                color: e.isActive
                                    ? AppColors.white
                                    : AppColors.mainColor,
                              ),
                              const SizedBox(height: 10),
                              Text(
                                e.title,
                                style: TextStyle(
                                    color: e.isActive
                                        ? AppColors.white
                                        : AppColors.grey),
                              ),
                            ],
                          ),
                        ),
                      ),
                    )
                    .toList(),
              ),
            ),
          ),
          ElevatedButton(
            onPressed: () { Navigator.push(context, LocalAndWebObjectsWidget() as Route<Object?>); },
            child: const Text("testPage"))
        ],
      ),
    );
  }

  void onTap(ShopCardModel e) {
    setState(() {
      e.isActive = !e.isActive;
    });
    Navigator.push(
        context,
        MaterialPageRoute(
            builder: (context) => ARObjectsScreen(
                  object: e.object,
                  isLocal: e.isLocal,
                ))).then((value) {
      setState(() {
        e.isActive = !e.isActive;
      });
    });
  }
}