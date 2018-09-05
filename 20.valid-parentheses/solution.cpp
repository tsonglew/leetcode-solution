#include <iostream>
#include <stack>
using namespace std;

class Solution {
public:
    bool isValid(string s) {
        stack<char> con;
        for(string::iterator it=s.begin(); it!=s.end(); it++) {
            switch (*it) {
                case '{':
                case '(':
                case '[': con.push(*it); break;
                case ')': if(con.empty()||con.top()!='(') return false; con.pop();break;
                case '}': if(con.empty()||con.top()!='{') return false; con.pop();break;
                case ']': if(con.empty()||con.top()!='[') return false; con.pop();break;
            }
        }
        return con.empty();
    }
};
