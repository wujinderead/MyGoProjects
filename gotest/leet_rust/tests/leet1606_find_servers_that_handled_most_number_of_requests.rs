// https://leetcode.com/problems/find-servers-that-handled-most-number-of-requests/

// You have k servers numbered from 0 to k-1 that are being used to handle multiple requests
// simultaneously. Each server has infinite computational capacity but cannot handle more
// than one request at a time. The requests are assigned to servers according to a specific
// algorithm:
//   The iᵗʰ (0-indexed) request arrives.
//   If all servers are busy, the request is dropped (not handled at all).
//   If the (i % k)ᵗʰ server is available, assign the request to that server.
//   Otherwise, assign the request to the next available server (wrapping around
//     the list of servers and starting from 0 if necessary). For example, if the iᵗʰ
//     server is busy, try to assign the request to the (i+1)ᵗʰ server, then the (i+2)ᵗʰ
//     server, and so on.
// You are given a strictly increasing array arrival of positive integers, where arrival[i]
// represents the arrival time of the iᵗʰ request, and another array load, where load[i]
// represents the load of the iᵗʰ request (the time it takes to complete). Your goal is to
// find the busiest server(s). A server is considered busiest if it handled the most number
// of requests successfully among all the servers.
// Return a list containing the IDs (0-indexed) of the busiest server(s). You may return the
// IDs in any order.
// Example 1:
//   Input: k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3]
//   Output: [1]
//   Explanation:
//     All of the servers start out available.
//     The first 3 requests are handled by the first 3 servers in order.
//     Request 3 comes in. Server 0 is busy, so it's assigned to the next available server, which is 1.
//     Request 4 comes in. It cannot be handled since all servers are busy, so it is dropped.
//     Servers 0 and 2 handled one request each, while server 1 handled two requests.
//     Hence server 1 is the busiest server.
// Example 2:
//   Input: k = 3, arrival = [1,2,3,4], load = [1,2,1,2]
//   Output: [0]
//   Explanation:
//     The first 3 requests are handled by first 3 servers.
//     Request 3 comes in. It is handled by server 0 since the server is available.
//     Server 0 handled two requests, while servers 1 and 2 handled one request each.
//     Hence server 0 is the busiest server.
// Example 3:
//   Input: k = 3, arrival = [1,2,3], load = [10,12,11]
//   Output: [0,1,2]
//   Explanation: Each server handles a single request, so they are all considered the busiest.
// Constraints:
//   1 <= k <= 10⁵
//   1 <= arrival.length, load.length <= 10⁵
//   arrival.length == load.length
//   1 <= arrival[i], load[i] <= 10⁹
//   arrival is strictly increasing.

mod _find_servers_that_handled_most_number_of_requests {
    struct Solution{
        k: i32,
        arrival: Vec<i32>,
        load: Vec<i32>,
        ans: Vec<i32>,
    }

    use std::cmp::Reverse;
    use std::collections::{BinaryHeap, BTreeSet};
    use std::ops::Bound::{Included, Unbounded};
    impl Solution {
        pub fn busiest_servers(k: i32, arrival: Vec<i32>, load: Vec<i32>) -> Vec<i32> {
            // pre-check: if job number <= k, all job can run exactly once
            let kk = k as usize;
            if arrival.len() <= kk {
                return (0..arrival.len() as i32).collect();
            }

            let mut busy = vec![1; kk];  // all server can serve first k requests
            let mut heap = BinaryHeap::new(); // min-heap: unavailable servers, each item is (endTime, serverIndex)

            // https://leetcode.com/problems/find-servers-that-handled-most-number-of-requests/discuss/876883/Python-using-only-heaps
            // another solution: use only heap, do some math magic here
            let mut available = BTreeSet::new();  // the index of available servers
            for i in 0..kk {  // for first k server
                if arrival[i] + load[i] <= arrival[kk] {  // can finish before task[kk] arrive
                    available.insert(i);   // push to available
                } else {  // push to unavailable heap with endTime
                    heap.push(Reverse((arrival[i] + load[i], i)));
                }
            }

            // for each task
            for i in kk..arrival.len() {
                while let Some(&x) = heap.peek() {
                    let (available_time, ind) = x.0;
                    if available_time > arrival[i] {
                        break;
                    }
                    heap.pop();  // if the server available, pop from heap
                    available.insert(ind);  // and push to available
                }
                if available.len() == 0 {
                    continue;
                }
                // in available, find the first index that >= i%kk, or just the first
                if let Some(&ind) = available.range((Included(&(i%kk)), Unbounded)).next() {
                    // occupy ind
                    available.remove(&ind);
                    heap.push(Reverse((arrival[i] + load[i], ind)));
                    busy[ind] += 1;
                } else {
                    // occupy first server
                    let &ind = available.range((Unbounded::<usize>, Unbounded::<usize>)).next().unwrap();
                    available.remove(&ind);
                    heap.push(Reverse((arrival[i] + load[i], ind)));
                    busy[ind] += 1;
                }
            }

            // find most busy
            let &max = busy.iter().max().unwrap();
            return (0..busy.len())
                .into_iter()
                .filter(|&i| busy[i] == max)
                .map(|x| x as i32)
                .collect::<Vec<_>>();
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                k: 3,
                arrival: vec![1,2,3,4,5],
                load: vec! [5,2,3,3,3],
                ans: vec![1],
            },
            Solution {
                k: 3,
                arrival: vec![1,2,3,4,5,6],
                load: vec! [5,2,3,3,3,3],
                ans: vec![1,2],
            },
            Solution {
                k: 3,
                arrival: vec![1,2,3,4],
                load: vec! [1,2,1,2],
                ans: vec![0],
            },
            Solution {
                k: 3,
                arrival: vec![1,2,3],
                load: vec! [10,12,11],
                ans: vec![0,1,2],
            }
        ];
        for i in testcases {
            let ans = Solution::busiest_servers(i.k,  i.arrival, i.load);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}