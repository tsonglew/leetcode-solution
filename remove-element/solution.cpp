#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        vector<int>::iterator i = nums.begin();
        while(i!=nums.end()) {
            if(*i==val){
                *i = nums.back();
                nums.pop_back();
            } else {
                ++i;
            }
        }
        return static_cast<int>(nums.size());
    }
};

template <class T>
void show_vector(vector<T>& v) {
        for (typename vector<T>::iterator i = v.begin(); i!=v.end();++i) {
                cout << *i << " ";
        }
        cout << endl;
}
