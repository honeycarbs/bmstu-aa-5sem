#ifndef LEE_ROOT_H
#define LEE_ROOT_H

#include "../include/map.h"
#include <functional>
#include <thread>

int leeRootParallel(Map &map, const coord_t &src, coord_t &dest, int thread_count);

int leeRootSync(Map &map, const coord_t &src, coord_t &dest);

#endif//LEE_ROOT_H

