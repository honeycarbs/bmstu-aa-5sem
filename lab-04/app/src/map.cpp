#include "../include/map.h"
#include <fstream>

#include <functional>
#include <iomanip>
#include <iostream>

#define KRED "\x1B[31m"
#define KGRN "\x1B[32m"
#define RESET "\x1B[0m"

using std::setw;


Map::Map(std::string filename) {
    std::string fullpath = "../meta/" + filename;
    std::ifstream ifs(fullpath);

    if (!ifs.is_open()) {
        throw std::exception();
    }
    ifs >> width_ >> height_;
    slots_.resize(height_);
    for (auto &row : slots_) {
        row.resize(width_);
    }

    long buf;

    for (size_t i = 0; i < width_; i++) {
        for (size_t j = 0; j < height_; j++) {
            ifs >> buf;
            switch (buf) {
                case 1:
                    slots_[i][j] = WALL;
                    break;
                case 0:
                    slots_[i][j] = UNSEEN;
                    break;
                default: {
                    ifs.close();
                    throw std::exception();
                }
            }
        }
    }
    ifs.close();
}

int Map::getSlot(int x, int y) {
    if (x < 0 || y < 0 || x >= width_ || y >= height_)
        return WALL;
    std::lock_guard<std::mutex> guard(mutex_);
    return slots_[x][y];
}
int Map::getSlot(std::pair<int, int> c) {
    return getSlot(c.first, c.second);
}

void Map::setSlot(int x, int y, int state) {
    if (x < 0 || y < 0 || x > width_ || y >= height_)
        throw std::exception();

    std::lock_guard<std::mutex> guard(mutex_);
    slots_[x][y] = state;
}

void Map::setSlot(std::pair<int, int> c, int state) {
    this->setSlot(c.first, c.second, state);
}

void Map::outputStdput() {
    std::lock_guard<std::mutex> guard(mutex_);
    for (const auto &row : slots_) {
        for (auto &val : row) {
            if (val == WALL) {
                std::cout << setw(10) << KRED << "x" << RESET;
            } else if (val == UNSEEN) {
                std::cout << setw(10) << KGRN << "c" << RESET;
            } else {
                std::cout << setw(11 - std::to_string(val).length()) << KGRN << val << RESET;
            }
        }
        std::cout << std::endl;
    }
}

void Map::clear() {
    for (auto &row : slots_) {
        for (auto &val : row) {
            if (val >= 0) {
                val = UNSEEN;
            }
        }
    }
}
