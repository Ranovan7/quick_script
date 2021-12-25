import random

width = 4
length = 25
tiling = ["X", " "]

class DungeonMap:

    def __init__(self, width, length):
        self.width = width
        self.length = length
        self.map = []
        self.first_col()
        self.generate()

    @property
    def presentation(self):
        return [[col[w] for col in self.map] for w in range(self.width)]

    def first_col(self):
        filling = []
        for i in range(width):
            filling.append(random.choice(tiling))

            if filling.count("X") >= 3:
                filling[-1] = " "

            if filling.count(" ") >= 3:
                filling[-1] = "X"

        random.shuffle(filling)
        self.map.append([" " for _ in range(self.width)])
        self.map.append(filling)

    def generate(self):
        for _ in range(self.length - 1):
            filling = self.randomize_col()
            self.map.append([" " for _ in range(self.width)])
            self.map.append(filling)

    def randomize_col(self):
        res = []

        res.append(random.choice(tiling))
        res.append(random.choice(tiling))

        if res[1] == "X" and res[0] == "X":
            res.append(" ")
        else:
            res.append(random.choice(tiling))

        if res[1] == "X" and res[2] == "X":
            res.append(" ")
            res[0] = " "
        else:
            res.append(random.choice(tiling))

        return res

dmap = DungeonMap(width, length)
for m in dmap.presentation:
    print(" ".join(m))
