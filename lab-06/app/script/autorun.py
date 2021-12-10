import os
import re
from typing import get_args

def get_ans(daycounts, rightanswer, runarg, filename):
    optimal = -1
    optimal_days = -1
    ethalon = min(rightanswer)

    for days in daycounts:
            output = os.popen("{}".format(runarg + filename + " graph-10.txt {} -AUTOCONF".format(days))).readlines()
            output = re.split('\s+', output[0])[:-1]

            comp = []
            for num in output:
                comp.append(int(num))
            
            curr_solution = min(comp)
            if curr_solution == ethalon:
                found = True
                optimal_days = days
                return found, ethalon, days
            
            if optimal == -1 or curr_solution < optimal:
                optimal = curr_solution
                optimal_days = days
            

    return False, optimal, optimal_days + 1
    


def main():
    srcpath = '../src/main.go '
    runarg  = 'go run ' + srcpath
    rightanswer = [39, 39, 39, 39, 39, 39, 39, 39, 39, 39]
    daycounts = [i for i in range (10, 500, 1)]

    for i in range(121):
        filename = "config-{}.txt".format(i)
        tup = get_ans(daycounts, rightanswer, runarg, filename)
        found = tup[0]
        print("{} {} {}".format(tup[1], tup[2], filename))


if __name__ == "__main__":
    main()