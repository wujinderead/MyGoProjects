// https://leetcode.com/problems/maximum-frequency-stack/

// Design a stack-like data structure to push elements to the stack and pop the
// most frequent element from the stack.
// Implement the FreqStack class:
//   FreqStack() constructs an empty frequency stack.
//   void push(int val) pushes an integer val onto the top of the stack.
//   int pop() removes and returns the most frequent element in the stack.
// If there is a tie for the most frequent element, the element closest to the 
// stack's top is removed and returned.
// Example 1:
//   Input
//     ["FreqStack", "push", "push", "push", "push", "push", "push", "pop", "pop", "pop", "pop"]
//     [[], [5], [7], [5], [7], [4], [5], [], [], [], []]
//   Output
//     [null, null, null, null, null, null, null, 5, 7, 5, 4]
//   Explanation
//     FreqStack freqStack = new FreqStack();
//     freqStack.push(5); // The stack is [5]
//     freqStack.push(7); // The stack is [5,7]
//     freqStack.push(5); // The stack is [5,7,5]
//     freqStack.push(7); // The stack is [5,7,5,7]
//     freqStack.push(4); // The stack is [5,7,5,7,4]
//     freqStack.push(5); // The stack is [5,7,5,7,4,5]
//     freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
//     freqStack.pop();   // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
//     freqStack.pop();   // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
//     freqStack.pop();   // return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the top. The stack becomes [5,7].
// Constraints:
//   0 <= val <= 10⁹
//   At most 2 * 10⁴ calls will be made to push and pop.
//   It is guaranteed that there will be at least one element in the stack before calling pop.

mod _maximum_frequency_stack {

    #[derive(Debug)]
    struct FreqStack {
        num_freq: HashMap<i32, i32>,
        freq_stack: HashMap<i32, Vec<i32>>,
        max_freq: i32,
    }

    // after push [5,7,5,7,4,5], the freq_stack looks like:
    // 3: [5]
    // 2: [5, 7]
    // 1: [5, 7, 4]
    use std::collections::HashMap;
    impl FreqStack {
        fn new() -> Self {
            Self {
                num_freq: HashMap::new(),
                freq_stack: HashMap::new(),
                max_freq: 0,
            }
        }

        fn push(&mut self, val: i32) {
            let f = self.num_freq.entry(val).or_insert(0);
            *f += 1;
            if *f > self.max_freq {
                self.max_freq = *f;
            }
            self.freq_stack.entry(*f).or_insert(Vec::new()).push(val);
        }

        fn pop(&mut self) -> i32 {
            let stack = self.freq_stack.get_mut(&self.max_freq).unwrap();
            let val = stack.pop().unwrap();
            if stack.len() == 0 {
                self.max_freq -= 1;
            }
            let f = self.num_freq.get_mut(&val).unwrap();
            *f -= 1;
            return val;
        }
    }

    #[test]
    fn test() {
        {
            let freq_stack = &mut FreqStack::new();
            println!("{:?}", freq_stack);
            freq_stack.push(5); // The stack is [5]
            println!("{:?}", freq_stack);
            freq_stack.push(7); // The stack is [5,7]
            println!("{:?}", freq_stack);
            freq_stack.push(5); // The stack is [5,7,5]
            println!("{:?}", freq_stack);
            freq_stack.push(7); // The stack is [5,7,5,7]
            println!("{:?}", freq_stack);
            freq_stack.push(4); // The stack is [5,7,5,7,4]
            println!("{:?}", freq_stack);
            freq_stack.push(5); // The stack is [5,7,5,7,4,5]
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());   // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());   // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());   // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());   // return 4, as 4, 5 and 7 is the most frequent, but 4 is closest to the top. The stack becomes [5,7].
            println!("{:?}", freq_stack);
        }
        {
            let freq_stack = &mut FreqStack::new();
            println!("{:?}", freq_stack);
            freq_stack.push(4);
            println!("{:?}", freq_stack);
            freq_stack.push(0);
            println!("{:?}", freq_stack);
            freq_stack.push(9);
            println!("{:?}", freq_stack);
            freq_stack.push(3);
            println!("{:?}", freq_stack);
            freq_stack.push(4);
            println!("{:?}", freq_stack);
            freq_stack.push(2);
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            freq_stack.push(6);
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            freq_stack.push(1);
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            freq_stack.push(1);
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            freq_stack.push(4);
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
            println!("{}", freq_stack.pop());
            println!("{:?}", freq_stack);
        }
    }
}
