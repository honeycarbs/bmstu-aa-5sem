# configs for ant algorithm 
# hella lot of files

import os
import itertools

def form_coefs():
    coefs_gen = []
    for i in range(11):
        stack = []
        alpha = round(0.0 + 0.1 * i, 2)
        betta = round(1 - alpha, 2)
        # stack.append([alpha, betta])

        stack.append([round(0.0 + 0.1 * j, 2) for j in range(11)])
        stack.append([3])                                          # q
        stack.append([0.5])                                        # init
        stack.append([0.7])                                        # eps


        cprod = list(itertools.product(*stack))

        coefs_cur = []
        for x in cprod:
            x = list(x)
            x.insert(0, alpha)
            x.insert(1, betta)
            coefs_cur.append(x)
        coefs_gen.append(coefs_cur)
    
    return coefs_gen


def main():
    stack = form_coefs()
    save_path = '../meta/'

    filenum = 0
    for coef_dump in stack:
        for filedata in coef_dump:
            file_name = 'config-' + str(filenum) + '.txt'
            complete_name = os.path.join(save_path, file_name)
            filenum += 1

            file = open(complete_name, "w")

            for coef in filedata:
                file.write(str(coef) + "\n")
            file.close()

if __name__ == "__main__":
    main()