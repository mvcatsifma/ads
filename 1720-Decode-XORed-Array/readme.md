#1720 Decode XORed Array
https://leetcode.com/problems/decode-xored-array/

```text
1 XOR 0 = 1
0 XOR 2 = 2
2 XOR 1 = 3

(1) 0000 0001
(0) 0000 0000
------------- XOR
(1) 0000 0001

(0) 0000 0000
(2) 0000 0010
------------- XOR
(2) 0000 0010

(2) 0000 0010
(1) 0000 0001
------------- XOR
(3) 0000 0011

encoded[i] = arr[i] XOR arr[1+1]
arr[1+1] = arr[i] XOR encoded[i]
```

See [Properties of bitwise operations](https://leetcode.com/explore/learn/card/bit-manipulation/669/bit-manipulation-concepts/4496/):
`XOR operation properties: a⊕0=a，a⊕a=0;`


