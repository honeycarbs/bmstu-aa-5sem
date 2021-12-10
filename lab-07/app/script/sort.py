import re
import statistics

def main():
    vals = []
    file = open("compars-seg.dat", "r")
    for line in file:
        # print(re.split('\s+', line)[:-1])
        l = re.split('\s+', line)[:-1]
    #     # print(l)
        # l.reverse()
        vals.append(l)
        # print(vals[len(vals) - 1])
    
    file.close()

    vals.sort(key=lambda x: int(x[0]))
    for v in vals:
        print(v)

    file = open("compars-seg-r.dat", "w")
    for v in vals:
        file.write(str(v[0] + " " + str(v[1]) + "\n"))
    # print(statistics.median(vals))


if __name__ == "__main__":
    main()