import re
import statistics

def main():
    vals = []
    file = open("compars-bin.dat", "r")
    for line in file:
        # print(re.split('\s+', line)[:-1])
        l = re.split('\s+', line)[:-1]
        # print(l)
        vals.append(int(l[1]))
        # print(vals[len(vals) - 1])
    
    file.close()

    print(statistics.median(vals))
    print(max(vals))
    print(min(vals))


if __name__ == "__main__":
    main()