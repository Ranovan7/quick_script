# How To Run

### C
    gcc <app_name>.c -o <app_name>
    ./<app_name>

### C++
    g++ -o hello <app_name>.cpp
    ./<app_name>

### Cython
    code on file with .pyx extension
    python3 setup.py build_ext --build-lib ./cython_dir
    use cython module on a run.py
    python3 run.py

### GDScript
    godot3 -s <app_name>.gd

### Go
    go build <app_name>.go
    ./<app_name>

### Haskell
    ghc --make <app_name>.hs
    ./<app_name>

### Nim
    nim c -r --verbosity:0 <app name>.nim

### NodeJS
    node <filename>

### Python
    python3 <filename>.py

### Ruby
    rby <app_name>.rb

### Rust
    rustc <app_name>.rs
    ./<app_name>
