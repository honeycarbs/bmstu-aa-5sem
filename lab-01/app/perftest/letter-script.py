import string
import random

def stringGenerator(size):
    str = ''
    for _ in range(size):
        str += random.choice(string.ascii_lowercase)
    return str

log = open("generator-output.txt", "w+")

# sizes = [15, 20, 30, 50, 100, 200]
# for s in sizes:
#     log.write("============================ STRING OF SIZE " + str(s) + " ============================\n\n\n")
#     log.write(str(stringGenerator(s)) + "\n\n\n")
#     log.write(str(stringGenerator(s)) + "\n\n\n")

log.write(str(stringGenerator(150)) + "\n\n\n")
log.write(str(stringGenerator(150)) + "\n\n\n")

log.close()


# BenchmarkLevenshteinMIterative150-4        	    5690	    199897 ns/op
# BenchmarkLevenshteinMemoRec150-4           	    1632	    736246 ns/op
# BenchmarkDamerauIterative150-4             	    4716	    249735 ns/op
# BenchmarkDamerauLevenshteinRecursive5-4    	   51042	     22974 ns/op
# BenchmarkDamerauLevenshteinRecursive10-4   	       8	 139913520 ns/op