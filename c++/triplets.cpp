#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<vector<int>> triplets(vector<int> arr, int sum){
    int n = arr.size();
    sort(arr.begin(), arr.end());
    vector<vector<int>> result;

    for(int i=0; i<=n-3;i++) {
        int j = i+1;
        int k = n-1;

        while(j < k) {
            int currentSum = arr[i];
            currentSum += arr[j];
            currentSum += arr[k];

            if(currentSum == sum) {
                result.push_back({ arr[i], arr[j], arr[k]});
                j++;
                k--;
            } else if(currentSum > sum) {
                k--;
            } else {
                j--;
            }
        }
    }

    return result;
}

//driver
int main() {
    vector<int> arr{1,2,3,4,5,6,7,8,9,15};
    int S = 18;
    auto result = triplets(arr, S);

    for (auto v : result){
        for(auto no : v) {
            cout << no << ','; 
        }

        cout << endl;
    }

    return 0;
}