# Convert Data

## Description

Buatlah sebuah fungsi yang akan mengubah data string menjadi struct User yang memiliki property `Name`, `Age`, dan `Address`. Data yang diberikan adalah data yang sudah diubah menjadi array of string, dimana tiap string berisi data user yang dipisahkan dengan koma (,) dan urutannya adalah `Name`, `Age`, dan `Address`.

Contoh:

```go
"Edo,20,Jakarta"
```

Function akan menerima 1 parameter, yaitu:

- `data` (string) - data yang akan kita convert menjadi struct User

## Example

```text
input: "Edo,20,Jakarta"
output: { name: "Edo", age: 20, address: "Jakarta" }

input: "Budi,30,Bandung"
output: { name: "Budi", age: 30, Address : "Bandung" }
```
