const monsterCharacters = [
    {
        name: "Bulbasaur",
        type: "Grass / Poison",
        height: "0.7 m",
        weight: "6.9 kg",
        habitat: "Grassland",
    },
    {
        name: "Pikachu",
        type: "Electric",
        height: "0.4 m",
        weight: "6 kg",
        habitat: "Forest",
    },
    {
        name: "Charmander",
        type: "Fire",
        height: "0.6 m",
        weight: "8.5 kg",
        habitat: "Mountain",
    },

    {
        name: "Squirtle",
        type: "Water",
        height: "0.5 m",
        weight: "9 kg",
        habitat: "Water Edge",
    },
]

const tableBody = document.getElementById("table-body");
getData()

function getData() {
    tableBody.innerHTML = ""
    monsterCharacters.forEach((monster) => {
        const row = document.createElement("tr");

        for (const key in monster) {
            const cell = document.createElement("td");
            cell.innerText = monster[key];
            row.appendChild(cell);
        }
        tableBody.appendChild(row);
    });
}