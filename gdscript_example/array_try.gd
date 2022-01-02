extends SceneTree

func _init():
    var arr: Array = []
    var rand_count_info: Dictionary = {
        1: 100,
        2: 200,
        3: 300,
    }
    for k in rand_count_info:
        for _i in range(rand_count_info[k]):
            arr.append(k)
    print(arr.count(1))
    print(arr.count(2))
    print(arr.count(3))
    print()

    var arr2 = [1,2,3,4,5]
    arr2.shuffle()
    print(arr2)

    quit()
