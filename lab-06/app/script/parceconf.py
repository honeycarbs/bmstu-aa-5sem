import itertools
import re
import os
import sys

def get_config(filename):
    configs = []
    save_path = '../meta/'
    complete_name = os.path.join(save_path, filename)

    file = open(complete_name, "r")
    for line in file:
        if line != '':
            config = re.split('\s+', line)[:-1]
            configs.append(config[0])

    return configs[:3]

def coefsort(coeflist):
    for item in coeflist:
        item[3] = int(item[3])
    
    coeflist = sorted(coeflist, cmp=lambda a,b: cmp(a[3], b[3]), reverse=False)  

    return coeflist[:25]

def main():
    filename = "dataclass1.txt"
    outputname = "parced.txt"
    ethalon = 300
    linesym = 10

    outputfile = open(outputname, "w")

    no_eps = []

    parcefile = open(filename, "r")
    for i, line in enumerate(parcefile):
        res = []
        parceline = re.split('\s+', line)[:-1]
        if len(parceline) == 3:
            configs = get_config(parceline[2])
            for c in configs:
                res.append(c)
            res.append(parceline[1])
            res.append(parceline[0])
            res.append(str(int(parceline[0]) - ethalon))

            if (int(parceline[0]) - ethalon) == 0:
                no_eps.append(res)

            for j, r in enumerate(res):
                if (j == len(res) - 1):
                    outputfile.write(r + " ")
                else:
                    outputfile.write(r + " & ")

            if i != 0 and i == linesym:
                outputfile.write(" \\\\ \hline\n")
                linesym += 11
            else:
                outputfile.write("\\\\ \n")

    parcefile.close()
    outputfile.close()
    
    no_eps = coefsort(no_eps)
    for items in no_eps:
        for i, item in enumerate(items):
            if i != len(items) - 1:
                sys.stdout.write('{} & '.format(item))
            else:
                sys.stdout.write('{} \\\\ \n'.format(item))

if __name__ == "__main__":
    main()