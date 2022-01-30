# How To Run

### C
    gcc <app_name>.c -o <app_name>
    ./<app_name>

### C++
    g++ -o hello <app_name>.cpp
    ./<app_name>

### Cython
    python3 setup.py build_ext --build-lib ./cython_dir
    python3 run.py

### GDScript
    godot3 -s <app_name>.gd

### Go
    go run <app_name>.go

### Haskell
    ghc --make <app_name>.hs
    ./<app_name>

### Nim
    nim c -r --verbosity:0 <app name>.nim

### NodeJS/Javascript
    node <filename>

### Python
    python3 <filename>.py

### Ruby
    rby <app_name>.rb

### Rust
    rustc <app_name>.rs
    ./<app_name>

### Typescript
    npx ts-node <app_name>.ts
