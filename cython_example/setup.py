from setuptools import setup
from Cython.Build import cythonize

setup(
    name='Hello world',
    ext_modules=cythonize("hello.pyx"),
    zip_safe=False,
)
