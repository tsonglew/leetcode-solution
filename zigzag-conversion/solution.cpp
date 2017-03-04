#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    string convert(string s, int numRows) {
        string result;
        const int size = numRows;
        const int stringLength = static_cast<int>(s.length());
        vector<char>* arr[size];
        int colCount[size];
        for(int i=0; i<size; i++){
            arr[i] = new vector<char>;
            colCount[i] = 0;
        }
        int charsCount = 0;
        
        while(charsCount < stringLength){
            int rowCount = 0;
            while(rowCount < size&&charsCount < stringLength){
                if(rowCount%2==0&&colCount[rowCount]%2!=0){
                    arr[rowCount]->push_back(' ');
                    
                } else {
                    arr[rowCount]->push_back(s[charsCount]);
                    charsCount++;
                }
                colCount[rowCount]++;
                rowCount++;
            }
        }
        for(int i=0; i<size; i++){
            int col = 0;
            while(col<colCount[i]){
                char toAppend = (*arr[i])[col];
                if(toAppend!=' ')
                    result.append(&toAppend);
                col++;
            }
        }
        return result;
    }
};
