import csv
import random
import os
import re

# def get_brute(word)

def main():
    file = open("../meta/opinions.csv", "r")
    csvreader = csv.reader(file)


    entrs = []
    for row in csvreader:
        entrs.append(row[0])
    file.close()

    driverpath = "../driver/main.go"
    runarg  = 'go run ' + driverpath

    dictname = "opinions.csv"

    # brute_file = open("brute-comaprs.txt", "w")
    # bin_file   = open("bin-comaprs.txt", "w")
    seg_file   = open("seg-compars.txt", "w")

    # for i in range(10):
    #     for j in range(3):
            # print("{}".format(runarg + " " + dictname + " " + str(j) + " " + entrs[i]))
    #         output = os.popen("{}".format(runarg + " " + dictname + " " + str(j) + " " + entrs[i])).readlines()
    #         output = re.split('\s+', output[0])[:-1]
    #         print(output)


    # for i in range(2):
    i = 2
    for j in range(len(entrs)):
        output = os.popen("{}".format(runarg + " " + dictname + " " + str(i) + " " + entrs[j])).readlines()
        output = re.split('\s+', output[0])[:-1]
        print(entrs[j] + " " + output[0])
        # if (i == 0):
        #     brute_file.write(entrs[j] + " " + output[0] + "\n")
        # elif (i == 1):
        #     bin_file.write(entrs[j] + " " + output[0] + "\n")
        # elif (i == 2):
        seg_file.write(entrs[j] + " " + output[0] + "\n")

    
    # brute_file.close()
    # bin_file.close()
    seg_file.close()


    # for entr in entrs:
    #     print(entr)

    # print(entrs)

if __name__ == "__main__":
    main()