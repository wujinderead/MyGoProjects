package main

// https://leetcode.com/problems/odd-even-jump/

// You are given an integer array A. From some starting index, you can make a series of jumps. 
// The (1st, 3rd, 5th, ...) jumps in the series are called odd numbered jumps, and the 
// (2nd, 4th, 6th, ...) jumps in the series are called even numbered jumps. 
// You may from index i jump forward to index j (with i < j) in the following way: 
//   During odd numbered jumps (ie. jumps 1, 3, 5, ...), you jump to the index j such that 
//     A[i] <= A[j] and A[j] is the smallest possible value. If there are multiple such indexes j, 
//     you can only jump to the smallest such index j. 
//   During even numbered jumps (ie. jumps 2, 4, 6, ...), you jump to the index j such that 
//     A[i] >= A[j] and A[j] is the largest possible value. If there are multiple such indexes j, 
//     you can only jump to the smallest such index j. 
//   It may be the case that for some index i, there are no legal jumps.
// A starting index is good if, starting from that index, you can reach the end of the array 
// (index A.length - 1) by jumping some number of times (possibly 0 or more than once.) 
// Return the number of good starting indexes. 
// Example 1: 
//   Input: [10,13,12,14,15]
//   Output: 2
//   Explanation: 
//     In total, there are 2 different starting indexes (i = 3, i = 4) where we can 
//     reach the end with some number of jumps.
// Example 2: 
//   Input: [2,3,1,1,4]
//   Output: 3
//   Explanation: 
//     In total, there are 3 different starting indexes (i = 1, i = 3, i = 4) where 
//     we can reach the end with some number of jumps.
// Example 3: 
//   Input: [5,1,3,4,2]
//   Output: 3
//   Explanation: 
//     We can reach the end from starting indexes 1, 2, and 4.
// Note: 
//   1 <= A.length <= 20000 
//   0 <= A[i] < 100000 

// basically, the task is for a[i], in the right of a[i], to find 
// the maximal value that is smaller than a[i], and, 
// to find the minimal value that is larger than a[i]. 
// or say, find the closest values for A[i]. 
// red-black tree is perfect for this problem, for a key, just find its predecessor and successor.
// since go do not have TreeMap, we use java for this problem.
// time O(nlogn).
class Solution {
    public int oddEvenJumps(int[] A) {
        int count = 1;
        TreeMap<Integer, Integer> map = new TreeMap<>();
        boolean[] odd = new boolean[A.length];
        boolean[] even = new boolean[A.length];
        odd[A.length-1] = even[A.length-1] = true;
        map.put(A[A.length-1], A.length-1);
        for (int i=A.length-2; i>=0; i--) {
            odd[i] = even[i] = false;
            Integer ind = map.put(A[i], i);  // A[i] is present, just update the index
            if (ind != null) {
                odd[i] = even[ind];
                even[i] = odd[ind];
                count += odd[i] ? 1 : 0;
                continue;
            }
            Map.Entry<Integer, Integer> e = null;
            if ((e = map.lowerEntry(A[i])) != null) {
                even[i] = odd[e.getValue()];
            }
            if ((e = map.higherEntry(A[i])) != null) {
                odd[i] = even[e.getValue()];
            }
            count += odd[i] ? 1 : 0;
        }
        return count;
    }
}

public class OddEvenJump {
    public static void main(String[] args) {
        Solution so = new Solution();
        System.out.println(so.oddEvenJumps(new int[]{10,13,12,14,15}));
        System.out.println(so.oddEvenJumps(new int[]{2,3,1,1,4}));
        System.out.println(so.oddEvenJumps(new int[]{5,1,3,4,2}));
        System.out.println(so.oddEvenJumps(new int[]{1,2,3,2,1,4,4,5}));
    }
}