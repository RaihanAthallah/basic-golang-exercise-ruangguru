# Jumlah Element Slice

## Description

Buatlah sebuah fungsi yang akan menerima sebuah slice dan akan mengembalikan jumlah element yang ada di dalam slice tersebut.

## Constraints

- slice tidak akan kosong
- slice bisa berisi tipe data apapun (interface{} / any)

## Example

```text
input: [1, 2, 3, 4, 5]
output: 5

input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
output: 10

input: [1, 1, 1, 5, 5, 5]
output: 6

input: ["Edo", "Budi", "Joko", "Tono"]
output: 4

input: ["Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"]
output: 8

input: [true, false, true, false, true, false]
output: 6
```

## Hint

- Gunakan `len` untuk menghitung jumlah element di dalam slice
