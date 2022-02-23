extends SceneTree

func _init():
    print("Hello World!")

    var example = 100 if true else 10
    print(example)

    var parseJSON = JSON.parse('{"example":{"gold":false,"silver":false,"bronze":false}}')
    print(parseJSON.result)

    quit()
