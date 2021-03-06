import random
from typing import List, Tuple
import matplotlib.pyplot as plt

SIDES = ['up', 'down', 'right', 'left']
SELECTOR = ['random', 'first_mid_last']

def out_of_bounds(node, side, limit):
    if side == 'up':
        return (node[0] - 1) < 0
    elif side == 'down':
        return (node[0] + 1) > limit
    elif side == 'right':
        return (node[1] + 1) > limit
    elif side == 'left':
        return (node[1] - 1) < 0
    else:
        return True

def get_neighbor(node, side):
    if side == 'up':
        return (node[0] - 1, node[1])
    elif side == 'down':
        return (node[0] + 1, node[1])
    elif side == 'right':
        return (node[0], node[1] + 1)
    elif side == 'left':
        return (node[0], node[1] - 1)
    else:
        return None

class Maze:
    maze: List[List[str]] = []
    init_rooms: List[tuple[int, int]] = []
    wall_list: List[tuple[int, int]] = []
    path: List[tuple[int, int]] = []
    selector: str = "random"

    def __init__(self, dim, selector = "random"):
        self.dim = dim  # has to be odd numbers
        self.selector = selector

    def grid_202(self):
        for i in range(self.dim):
            self.maze.append([])
            for j in range(self.dim):
                if i % 2 == 0:
                    # walls and pillars only
                    if j % 2 == 0:
                        self.maze[i].append("X")
                    else:
                        self.maze[i].append("W")
                else:
                    # walls and room only
                    if j % 2 == 0:
                        self.maze[i].append("W")
                    else:
                        self.maze[i].append(" ")
                        self.init_rooms.append((i, j))
        self.center_rooms = self.get_center_rooms()
        return self

    def prim_generation(self):
        self.init_gen()
        wall = None
        while len(self.wall_list) > 0:
            wall = self.wall_selector(wall)
            adj_rooms = self.get_adjacent_rooms(wall)
            unmarked_rooms = self.unmarked_rooms(adj_rooms)
            if len(unmarked_rooms) == 1:
                self.path.append(wall)
                self.path.append(unmarked_rooms[0])
                self.init_rooms.remove(unmarked_rooms[0])
                self.update_wall_list(unmarked_rooms[0])
            self.wall_list.remove(wall)
        return self

    def init_gen(self):
        picked = self.pick_room()
        self.path.append(picked)
        self.update_wall_list(picked)

    def wall_selector(self, current_wall):
        if not current_wall:
            return random.choice(self.wall_list)

        if self.selector == "first_mid_last":
            mid = round(len(self.wall_list)/2)
            return self.wall_list[random.choice([0, mid, -1])]
        else:   # random
            random.shuffle(self.wall_list)
            return self.wall_list[0]

    def show(self):
        x = []
        y = []
        for i in range(self.dim):
            for j in range(self.dim):
                if (i, j) in self.path:
                    self.maze[i][j] = ' '
                else:
                    self.maze[i][j] = 'X'
                    x.append(j)
                    y.append(i)
        self.show_scatter(x, y)

    def show_scatter(self, x, y):
        plt.scatter(x, y, marker='s')
        plt.show()

    def get_center_rooms(self):
        result = []
        mid_high = round(self.dim/2) + 7
        mid_low = round(self.dim/2) - 7
        for room in self.init_rooms:
            if room[0] >= mid_low and room[0] <= mid_high:
                if room[1] >= mid_low and room[1] <= mid_high:
                    result.append(room)
        return result

    def pick_room(self):
        picked = random.choice(self.center_rooms)
        self.init_rooms.remove(picked)
        self.center_rooms.remove(picked)
        return picked

    def update_wall_list(self, picked):
        for side in SIDES:
            if not out_of_bounds(picked, side, self.dim):
                self.wall_list.append(get_neighbor(picked, side))

    def get_adjacent_rooms(self, wall):
        result = []
        for side in SIDES:
            if not out_of_bounds(wall, side, self.dim):
                neighbor = get_neighbor(wall, side)
                if neighbor in self.init_rooms:
                    result.append(neighbor)
        return result

    def unmarked_rooms(self, rooms):
        unmarked_rooms = []
        for room in rooms:
            if room not in self.path:
                unmarked_rooms.append(room)
        return unmarked_rooms

if __name__ == "__main__":
    n = 77  # has to be odd numbers
    # selector = 'random'
    selector = 'first_mid_last'

    print(f"Maze Generation for {n} x {n}")
    maze = Maze(dim=n, selector=selector)

    print("Generated Maze")
    maze.grid_202().prim_generation().show()
