import numpy as np
import os

def generate(size, int_range):
    m = np.random.randint(1, int_range + 1, (size, size))
    m_sym = (m + m.T)

    for i in range(size):
        m_sym[i][i] = 0

    return m_sym



def main():
    size = 10
    range = 50
    matrix = generate(size, range)
    # print(matrix)

    save_path = '../meta/'
    file_name = 'graph-' + str(size) + '.txt'
    complete_name = os.path.join(save_path, file_name)
    file = open(complete_name, "w")


    for row in matrix:
        for item in row:
            file.write(str(item) + " ")
        file.write("\n")
    file.close()



if __name__ == "__main__":
    main()