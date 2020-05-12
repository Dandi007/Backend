from functools import partial


def callback(a, b, c, d):
    print(a, b, c, d)


def do_it(a, b, c, func):
    func(a, b, c)


cb = partial(callback, d=20)
do_it(1, 2, 3, cb)
