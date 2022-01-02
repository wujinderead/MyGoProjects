// https://leetcode.com/problems/destroying-asteroids/

// You are given an integer mass, which represents the original mass of a planet.
// You are further given an integer array asteroids, where asteroids[i] is the mass of the iᵗʰ asteroid.
// You can arrange for the planet to collide with the asteroids in any arbitrary order.
// If the mass of the planet is greater than or equal to the mass of the asteroid, the asteroid is
// destroyed and the planet gains the mass of the asteroid. Otherwise, the planet is destroyed.
// Return true if all asteroids can be destroyed. Otherwise, return false.
// Example 1:
//   Input: mass = 10, asteroids = [3,9,19,5,21]
//   Output: true
//   Explanation: One way to order the asteroids is [9,19,5,3,21]:
//     - The planet collides with the asteroid with a mass of 9. New planet mass: 10 + 9 = 19
//     - The planet collides with the asteroid with a mass of 19. New planet mass: 19 + 19 = 38
//     - The planet collides with the asteroid with a mass of 5. New planet mass: 38 + 5 = 43
//     - The planet collides with the asteroid with a mass of 3. New planet mass: 43 + 3 = 46
//     - The planet collides with the asteroid with a mass of 21. New planet mass: 46 + 21 = 67
//     All asteroids are destroyed.
// Example 2:
//   Input: mass = 5, asteroids = [4,9,23,4]
//   Output: false
//   Explanation:
//     The planet cannot ever gain enough mass to destroy the asteroid with a mass of 23.
//     After the planet destroys the other asteroids, it will have a mass of 5 + 4 + 9 + 4 = 22.
//     This is less than 23, so a collision would not destroy the last asteroid.
// Constraints:
//   1 <= mass <= 10⁵
//   1 <= asteroids.length <= 10⁵
//   1 <= asteroids[i] <= 10⁵

mod _destroying_asteroids {
    struct Solution{
        mass: i32,
        asteroids: Vec<i32>,
        ans: bool,
    }

    impl Solution {
        pub fn asteroids_destroyed(mass: i32, asteroids: Vec<i32>) -> bool {
            let mut copied = vec![0; asteroids.len()];
            copied.copy_from_slice(&asteroids);
            copied.sort();                    // sort asteroids
            let mut mass: i64 = mass as i64;  // use i64 to avoid overflow
            let mut i = 0;
            while i < copied.len() {
                if mass < (copied[i] as i64) {
                    return false;
                }
                mass += copied[i] as i64;
                i += 1;
            }
            return true
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                mass: 10,
                asteroids: vec![3,9,19,5,21],
                ans: true,
            },
            Solution {
                mass: 5,
                asteroids: vec![4,9,23,4],
                ans: false,
            },
        ];
        for i in testcases {
            let ans = Solution::asteroids_destroyed(i.mass, i.asteroids);
            println!("{}, {}", ans, i.ans);
        }
    } 
}