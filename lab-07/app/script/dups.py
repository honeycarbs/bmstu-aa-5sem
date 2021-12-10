from os import write
import re

file = open("seg-compars.txt", "r")
ready = open("alph-seg.dat", "w")

entrs = []

for i, line in enumerate(file):
    l = re.split('\s+', line)[:-1]
    l.append(i)

    entrs.append(l)

entrs.sort(key = lambda x: x[1])

for i in entrs:
    print(i)

for n, i in enumerate(entrs):
    ready.write(str(n) + " " + str(i[1]) + "\n")

file.close()
ready.close()