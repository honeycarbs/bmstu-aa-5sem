#include "../include/lee-root.h"

bool checkCoord(Map &map, coord_t pair, int value) {
    auto temp = map.getSlot(pair);
    return temp == map.UNSEEN || temp > value;
}

void parallelRoutine(CQueue &queue, Map &map, coord_t dest, bool &path_found) {
    while (true) {
        if (path_found) {
            if (queue.size() == 0) {
                return;
            }
            path_found = false;
        }
        auto opt = queue.pop();
        if (!opt)
            continue;

        auto pair = *opt;
        auto coord = pair.first;
        auto val = pair.second + 1;
        if (pair.first == dest) {
            map.setSlot(pair.first, pair.second);
            continue;
        }
        if (!checkCoord(map, coord, val)) {
            if (queue.size() == 0) {
                path_found = true;
            }
            continue;
        }

        map.setSlot(pair.first, pair.second);

        std::initializer_list<std::pair<coord_t, int>> list = {
                {{coord.first + 1, coord.second}, val},
                {{coord.first - 1, coord.second}, val},
                {{coord.first, coord.second + 1}, val},
                {{coord.first, coord.second - 1}, val}};
        queue.push(list);
    }
}

int leeRootParallel(Map &map, const coord_t &src, coord_t &dest, int thread_count) {
    CQueue queue;

    queue.push({src, 0});
    std::vector<std::thread> threads;
    bool path_found = false;
    threads.reserve(thread_count);
    for (int i = 0; i < thread_count; i++) {
        threads.emplace_back(
                std::thread(parallelRoutine, std::ref(queue), std::ref(map), dest, std::ref(path_found)));
    }
    for (int i = 0; i < thread_count; i++) {
        threads[i].join();
    }
    return map.getSlot(dest);
}

int leeRootSync(Map &map, const coord_t &src, coord_t &dest) {
    CQueue queue;

    queue.push({src, 0});

    bool f = false;
    parallelRoutine(queue, map, dest, f);

    return map.getSlot(dest);
}