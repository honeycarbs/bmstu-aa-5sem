import re
from itertools import chain
import sys

def separateConf(confs_array, file_line):
    conf_array = []
    confline = re.split('\s+', file_line)[:-1]
    for i in confline:
        conf_array.append(i)
    
    confs_array.append(conf_array)
    return confs_array



def main():
    filename = "classes12.txt"
    file = open(filename, "r")

    class1 = []
    class2 = []

    class_filled = False

    for line in file:
        if (len(line) == 1):
            class_filled = True
            continue
        if not class_filled:
            separateConf(class2, line)
        else:
            separateConf(class1, line)
        
    both = []

    for f in class1:
        for s in class2:
            if f[:3] == s[:3]:
                cur = []
                for ff, ss in zip(f[3:], s[3:]):
                    cur.append(ff)
                    cur.append(ss)
                both.append(list(x for x in chain(f[:3], cur)))
    
    for items in both:
        for i, item in enumerate(items):
            if i != len(items) - 1:
                sys.stdout.write('{} & '.format(item))
            else:
                sys.stdout.write('{} \\\\ \n'.format(item))

if __name__ == "__main__":
    main()