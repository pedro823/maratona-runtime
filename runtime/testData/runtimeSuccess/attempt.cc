#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int n;
    cin >> n;
    vector<int> v;
    for (int i = 0; i < n; i++) {
        int temp;
        cin >> temp;
        v.push_back(temp);
    }

    sort(v.begin(), v.end());
    bool first = true;
    for (auto a : v) {
            if (!first) {
                cout << " ";
            }
            cout << a;
            first = false;
    }
    cout << endl;

    return 0;
}

//int main() {
//    for (int i = 0; i < 10; i++) {
//        if (i > 0) {
//            cout << " ";
//        }
//        cout << i;
//    }
//    cout << endl;
//    return 0;
//}
