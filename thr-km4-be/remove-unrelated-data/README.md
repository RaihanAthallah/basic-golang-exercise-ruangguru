# Remove Unrelated Data

## Description

Buatlah sebuah fungsi yang akan menerima sebuah data map dan akan mengembalikan map baru yang beberapa key-value-nya telah dihapus karena tidak relevan dengan map tersebut.

Function akan menerima 2 parameter, yaitu:

- `mapData` (map) - map dengan key bertipe string dan value apapun (`interface{}` / `any`) yang akan dihapus key-value-nya
- `key` (string) - key yang akan dihapus dari map

## Constraints

- `mapData` tidak akan kosong
- `key` tidak akan berisi string kosong

## Example

```text
input: map[ name: "Edo", age: 20, address: "Jakarta"], 'address'
output: map[ name: "Edo", age: 20]


input: map[run: "lari", jump: "loncat", swim: "berenang"], 'jump'
output: map[run: "lari", swim: "berenang"]

input: map[satu: "ichi", dua: "ni", tiga: "san", empat: "yon", lima: "go"], 'tiga'
output: map[satu: "ichi", dua: "ni", empat: "yon", lima: "go"]
```

## Hint

- Gunakan `delete` untuk menghapus key-value dari map
