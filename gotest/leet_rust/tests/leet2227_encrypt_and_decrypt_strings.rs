// https://leetcode.com/problems/encrypt-and-decrypt-strings/

// You are given a character array keys containing unique characters and a string array values
// containing strings of length 2. You are also given another string array dictionary that
// contains all permitted original strings after decryption.
// You should implement a data structure that can encrypt or decrypt a 0-indexed string.
// A string is encrypted with the following process:
//   For each character c in the string, we find the index i satisfying keys[i] == c in keys.
//   Replace c with values[i] in the string.
// A string is decrypted with the following process:
//   For each substring s of length 2 occurring at an even index in the string, we find an i
//     such that values[i] == s. If there are multiple valid i, we choose any one of them.
//     This means a string could have multiple possible strings it can decrypt to.
//   Replace s with keys[i] in the string.
// Implement the Encrypter class:
// Encrypter(char[] keys, String[] values, String[] dictionary) Initializes the 
//   Encrypter class with keys, values, and dictionary.
// String encrypt(String word1) Encrypts word1 with the encryption process 
//   described above and returns the encrypted string.
// int decrypt(String word2) Returns the number of possible strings word2 could 
//   decrypt to that also appear in dictionary.
// Example 1:
//   Input
//     ["Encrypter", "encrypt", "decrypt"]
//       [[['a', 'b', 'c', 'd'], ["ei", "zf", "ei", "am"], ["abcd", "acbd", "adbc",
//       "badc", "dacb", "cadb", "cbda", "abad"]], ["abcd"], ["eizfeiam"]]
//   Output
//     [null, "eizfeiam", 2]
// Constraints:
//   1 <= keys.length == values.length <= 26
//   values[i].length == 2
//   1 <= dictionary.length <= 100
//   1 <= dictionary[i].length <= 100
//   All keys[i] and dictionary[i] are unique.
//   1 <= word1.length <= 2000
//   1 <= word2.length <= 200
//   All word1[i] appear in keys.
//   word2.length is even.
//   keys, values[i], dictionary[i], word1, and word2 only contain lowercase English letters.
//   At most 200 calls will be made to encrypt and decrypt in total.

mod _encrypt_and_decrypt_strings {
    // just also encrypt the dictionary
    use std::collections::HashMap;
    struct Encrypter {
        kv: Vec<String>,
        dictionary: HashMap<String, i32>,
    }

    impl Encrypter {
        fn new(keys: Vec<char>, values: Vec<String>, dictionary: Vec<String>) -> Self {
            let mut kv = vec!["".to_string(); 26];
            for (i, &k) in keys.iter().enumerate() {
                kv[k as usize - 'a' as usize] = values[i].clone();
            }
            let mut s = Self {
                kv,
                dictionary: HashMap::new(),
            };
            for d in dictionary {
                let ss = s.encrypt(d);
                *s.dictionary.entry(ss).or_insert(0) += 1;
            }
            s
        }

        fn encrypt(&self, word1: String) -> String {
            let mut en = String::with_capacity(2*word1.len());
            for &c in word1.as_bytes() {
                en.push_str(self.kv[c as usize - b'a' as usize].as_str());
            }
            return en;
        }

        fn decrypt(&self, word2: String) -> i32 {
            return *self.dictionary.get(&word2).unwrap_or(&0);
        }
    }
    
    #[test]
    fn test() {
        let obj = Encrypter::new(
            vec!['a', 'b', 'c', 'd'],
            vec!["ei", "zf", "ei", "am"].iter().map(|s| s.to_string()).collect(),
            vec!["abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"].iter().map(|s| s.to_string()).collect());
        println!("{}", obj.encrypt("abcd".to_string()));
        println!("{}", obj.decrypt("eizfeiam".to_string()));
    } 
}