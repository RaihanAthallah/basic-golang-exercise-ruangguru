# Remove Element Left & Right

## Deskripsi

Buatlah sebuah fungsi yang akan menerima sebuah array dan akan mengembalikan array baru yang merupakan hasil penghapusan element di awal atau di akhir array tersebut.

Function akan menerima 2 parameter, yaitu:

- `array` (array) - array yang akan dihapus elementnya
- `position` (string) - posisi element yang akan dihapus dari array, bisa berupa `left` atau `right`

Jika `position` adalah `left`, maka element di awal array akan dihapus. Jika `position` adalah `right`, maka element di akhir array akan dihapus.

## Constraints

- `array` tidak akan kosong
- `array` bisa berisi tipe data apapun ( `interface{}` / `any`)
- `position` hanya bisa berupa `left` atau `right`

## Contoh

```text
input: [1, 2, 3, 4, 5], 'left'
output: [2, 3, 4, 5]

input: [1, 2, 3, 4, 5], 'right'
output: [1, 2, 3, 4]

input: ["Edo", "Budi", "Joko", "Tono"], 'left'
output: ["Budi", "Joko", "Tono"]

input: ["Edo", "Budi", "Joko", "Tono"], 'right'
output: ["Edo", "Budi", "Joko"]
```
