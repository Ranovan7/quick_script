extends SceneTree

func _init():
    var noise = OpenSimplexNoise.new()

    noise.seed = randi()

    var map = [[],[],[],[]]

    for x in range(15):
        for y in range(4):
            map[y].append(round(noise.get_noise_2d(x, y) * 1000))

    for m in map:
        print(m)

    quit()
