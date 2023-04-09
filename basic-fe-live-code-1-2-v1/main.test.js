const { readFileSync } = require("fs");
const { JSDOM } = require("jsdom");
const { resolve } = require("path");

const html = readFileSync(resolve(__dirname, "index.html"), "utf8");

let windowJSDOM = null;

beforeEach(() => {
    const { window } = new JSDOM(html, {
        runScripts: "dangerously",
        includeNodeLocations: false,
        pretendToBeVisual: true,
    });

    windowJSDOM = window;
});

describe("Creates title and alternative title inside a div element with id 'pokemon-title'", () => {
    it("should render an 'h2' element with an 'id' 'title' and a text of 'Pokemon Series'", () => {
        const document = windowJSDOM.document;
        expect(document.getElementsByTagName("h2")[0].id).toBe("title");
        expect(document.getElementById("title").textContent).toBe(
            "Pokemon Series"
        );
    });

    it("should render an 'p' element with an 'id' 'title-alt' and a text of 'Pokemon Monster'", () => {
        const document = windowJSDOM.document;
        expect(document.querySelectorAll("p")[0].id).toBe("title-alt");
        expect(document.getElementById("title-alt").textContent).toBe(
            "Pokemon Monster"
        );
    });
});

describe("Creates an 'img' element inside a div element with class 'left'", () => {
    it("should render 1 element 'img' having attribute 'alt' with value 'pokemon-img'", () => {
        const document = windowJSDOM.document;
        expect(document.getElementsByTagName("img")).not.toBeNull();
        expect(document.getElementsByTagName("img")[0].alt).toBe(
            "pokemon-img"
        );
    });
});

describe("Create 2 'section' elements inside a div element with class 'right'", () => {
    it("should render 2 section elements which have id's of 'synopsis' and 'characters'", () => {
        const document = windowJSDOM.document;
        expect(
            document.getElementsByClassName("right")[0].children.length
        ).toBe(2);
        expect(document.getElementsByTagName("section")).not.toBeNull();
        expect(document.querySelectorAll("section")[0].id).toBe("synopsis");
        expect(document.querySelectorAll("section")[1].id).toBe("characters");
    });

    describe("'section' element with id 'synopsis'", () => {
        it("should render 'h3', 'p' and 'hr' (optional) elements", () => {
            const document = windowJSDOM.document;

            expect(
                document.querySelectorAll("section")[0].firstElementChild
                    .localName
            ).toBe("h3");

            expect(document.querySelectorAll("h3")[0].textContent).toBe(
                "Synopsis"
            );

            expect(
                document.querySelectorAll("section")[0].lastElementChild
                    .localName
            ).toBe("p");
        });
    });

    describe("'section' element with id 'characters'", () => {
        it("should render 'h3', 'table' and 'hr' (optional) elements", () => {
            const document = windowJSDOM.document;

            expect(
                document.querySelectorAll("section")[1].firstElementChild
                    .localName
            ).toBe("h3");

            expect(document.querySelectorAll("h3")[1].textContent).toBe(
                "Monster Characters"
            );

            expect(
                document.querySelectorAll("section")[1].lastElementChild
                    .localName
            ).toBe("table");

            expect(document.querySelector("table").children[0].localName).toBe(
                "thead"
            );
            expect(document.querySelector("table").children[0].id).toBe(
                "table-head"
            );

            expect(document.querySelector("table").children[1].localName).toBe(
                "tbody"
            );
            expect(document.querySelector("table").children[1].id).toBe(
                "table-body"
            );
        });
    });
});