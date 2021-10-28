#ifndef P_QUEUE_H
#define P_QUEUE_H

#include <list>
#include <mutex>
#include <optional>
#include <queue>
#include <stack>

template<typename T>
class PQueue {
public:
    PQueue() = default;
    PQueue(const PQueue<T> &) = delete ;
    PQueue& operator=(const PQueue<T> &) = delete ;

    virtual ~PQueue() = default;

    unsigned long size() const {
        std::lock_guard<std::mutex> lock(mutex_);
        return queue_.size();
    }

    std::optional<T> pop() {
        std::lock_guard<std::mutex> lock(mutex_);
        if (queue_.empty()) {
            return {};
        }
        _i++;
        T tmp;
        if (_i % 2) {
            tmp = queue_.front();
            queue_.pop_front();
        }
        else {
            tmp = queue_.back();
            queue_.pop_back();
        }
        return tmp;
    }

    void push(const T &item) {
        std::lock_guard<std::mutex> lock(mutex_);

        queue_.push_back(item);
    }

    template<class P>
    void push(const P &list) {
        std::lock_guard<std::mutex> lock(mutex_);
        for (auto &item: list) {
            queue_.push_front(item);
        }
    }

private:
    int _i = 0;
    std::list<T> queue_;
    mutable std::mutex mutex_;

    bool empty() const {
        return queue_.empty();
    }
};


#endif//P_QUEUE_H
