#include "include/lee-root.h"
#include <iomanip>
#include <iostream>

int main(int argc, char *argv[]) {
    if (argc == 7) {
        printf("                                                   __  _        ");
        printf("\n   ____ ___  ____ _____  ___     _________  ____  / /_(_)___  ____ _");
        printf("\n  / __ `__ \\/ __ `/_  / / _ \\   / ___/ __ \\/ __ \\/ __/ / __ \\/ __ `/");
        printf("\n / / / / / / /_/ / / /_/  __/  / /  / /_/ / /_/ / /_/ / / / / /_/ / ");
        printf("\n/_/ /_/ /_/\\__,_/ /___/\\___/  /_/   \\____/\\____/\\__/_/_/ /_/\\__, /  ");
        printf("\n                                                           /____/   \n");
        Map map(argv[1]);

        int x1 = atoi(argv[2]), y1 = atoi(argv[3]);
        int x2 = atoi(argv[4]), y2 = atoi(argv[5]);

        coord_t start = {x1, y1};
        coord_t finish = {x2, y2};

        int result;

        map.clear();
        result = leeRootSync(map, start, finish);
        std::cout << "Sync algorithm result :" << result << std::endl;

        map.clear();
        result = leeRootParallel(map, start, finish, 1);
        std::cout << "Parallel algorithm result :" << result << std::endl;
        map.outputStdput();
        return 0;
    } else if (argc == 1) {
        int count = 10;
        std::string sizes[2] = {"400", "500"};
        for (const auto &prefix : sizes) {
            std::string filename = "maze-" + prefix + ".txt";
            Map map(filename);
            coord_t start = {0, 0};
            coord_t finish = {std::stoi(prefix) - 1, std::stoi(prefix) - 1};
            double time_sum_sync = 0;
            double time_sums_parall[9] = {0, 0, 0, 0, 0, 0, 0, 0, 0};

            std::chrono::time_point<std::chrono::system_clock> now, end;

            for (int i = 0; i < count; i++) {
                map.clear();
                now = std::chrono::system_clock::now();
                leeRootSync(map, start, finish);
                end = std::chrono::system_clock::now();
                auto time =
                        std::chrono::duration_cast<std::chrono::milliseconds>(end - now).count();
                time_sum_sync += (double) time;
            }
            for (int i = 0; i < count; i++) {
                for (int j = 0; j < 9; j++) {
                    map.clear();
                    if (j == 0) {
                        now = std::chrono::system_clock::now();
                        leeRootParallel(map, start, finish, 1);
                        end = std::chrono::system_clock::now();
                    } else {
                        now = std::chrono::system_clock::now();
                        leeRootParallel(map, start, finish, j * 2);
                        end = std::chrono::system_clock::now();
                    }
                    auto time =
                            std::chrono::duration_cast<std::chrono::milliseconds>(end - now).count();
                    time_sums_parall[j] += (double) time;
                }
            }
            std::cout << prefix << " ";
            std::cout << time_sum_sync / count << " ";
            for (auto thread : time_sums_parall) {
                //printf("%3.0e ", thread / count);
                std::cout << thread / count << " ";
            }
            std::cout << std::endl;
        }
    }

    return 0;
}
