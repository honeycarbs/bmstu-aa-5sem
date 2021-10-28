#ifndef MAP_H
#define MAP_H

#include <mutex>
#include <string>
#include <vector>

#include "p-queue.h"

using coord_t = std::pair<int, int>;
using CQueue = PQueue<std::pair<coord_t, int>>;

class Map {
public:
    const int WALL = -2;
    const int UNSEEN = -1;

    explicit Map(std::string filename);
    int getSlot(int x, int y);
    int getSlot(std::pair<int, int> c);
    void setSlot(int x, int y, int state);
    void setSlot(std::pair<int, int> c, int state);
    void outputStdput();
    void clear();

    long width() { return width_; };
    long height() { return height_; };

private:
    long width_;
    long height_;
    std::vector<std::vector<int>> slots_;
    std::mutex mutex_;
};


#endif//MAP_H
