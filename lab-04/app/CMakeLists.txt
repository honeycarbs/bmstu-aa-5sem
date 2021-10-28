cmake_minimum_required(VERSION 3.19)
project(app)

set(CMAKE_CXX_STANDARD 17)

add_executable(app main.cpp src/map.cpp include/map.h include/p-queue.h src/lee-root.cpp include/lee-root.h)

target_link_libraries(${PROJECT_NAME} pthread)