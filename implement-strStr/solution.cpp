#include <iostream>
using namespace std;

class Solution {
public:
    int strStr(string haystack, string needle) {
        int hlen = static_cast<int>(haystack.length());
        int nlen = static_cast<int>(needle.length());
		return -1;
    }
};

class Ex {
public:
    string haystack;
    string needle;
};

int main() {
    Solution s;
    Ex e1, e2, e3, e4;
    e1.haystack = "hello";
    e1.needle = "ll";
    e2.haystack = "aaaaa";
    e2.needle = "bba";
    e3.haystack = "";
    e3.needle = "";
    e4.haystack = "a";
    e4.needle = "";
    cout << s.strStr(e1.haystack, e1.needle) << endl;
    cout << s.strStr(e2.haystack, e2.needle) << endl;
    cout << s.strStr(e3.haystack, e3.needle) << endl;
    cout << s.strStr(e4.haystack, e4.needle) << endl;
    
    return 0;
}
