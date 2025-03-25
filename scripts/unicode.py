#!/usr/bin/env python3
"""Testing how to print Unicode characters correctly on the terminal."""

import sys


class Colour:
    __black = 0
    __red = 1
    __green = 2
    __yellow = 3
    __blue = 4
    __magenta = 5
    __cyan = 6
    __white = 7

    def csi(self, n: int) -> str:
        return f"[38:5:{n}m"

    def def_fg(self) -> str:
        return "[0m"

    def black(self) -> str:
        return self.csi(Colour.__black)

    def red(self) -> str:
        return self.csi(Colour.__red)

    def green(self) -> str:
        return self.csi(Colour.__green)

    def yellow(self) -> str:
        return self.csi(Colour.__yellow)

    def blue(self) -> str:
        return self.csi(Colour.__blue)

    def magenta(self) -> str:
        return self.csi(Colour.__magenta)

    def cyan(self) -> str:
        return self.csi(Colour.__cyan)

    def white(self) -> str:
        return self.csi(Colour.__white)


class Printer:
    def __init__(self):
        self.__colour = Colour()

    def def_fg(self) -> None:
        sys.stdout.write(self.__colour.def_fg())

    def black(self) -> None:
        sys.stdout.write(self.__colour.black())

    def red(self) -> None:
        sys.stdout.write(self.__colour.red())

    def green(self) -> None:
        sys.stdout.write(self.__colour.green())

    def yellow(self) -> None:
        sys.stdout.write(self.__colour.yellow())

    def blue(self) -> None:
        sys.stdout.write(self.__colour.blue())

    def magenta(self) -> None:
        sys.stdout.write(self.__colour.magenta())

    def cyan(self) -> None:
        sys.stdout.write(self.__colour.cyan())

    def white(self) -> None:
        sys.stdout.write(self.__colour.white())


digits = (
    "    5    10        20        30        40        50        60        70        80"
)

ruler = (
    "....'....|....'....|....'....|....'....|....'....|....'....|....'....|....'....|"
)


printer = Printer()

printer.yellow()
print(digits)
print(ruler)
printer.def_fg()
print("    This is a Unicode string.")
print("    This has a 󱨎in it.")

print("    This has a ⚉ in it.")
