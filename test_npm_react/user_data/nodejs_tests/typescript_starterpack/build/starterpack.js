"use strict";
// c.f. https://www.w3schools.com/typescript/index.php
// examples of type-hint syntax of typescript for various use-cases
function main() {
    console.log("\n>>>>>>>>>>> Main function starts\n");
    helperRunSection(sectionSimpleTypes, "Simple types");
    helperRunSection(sectionArrays, "Arrays");
    helperRunSection(sectionTuples, "Tuples");
    helperRunSection(sectionObjects, "Objects");
    helperRunSection(sectionEnums, "Enums");
    helperRunSection(sectionAliasesAndInterfaces, "Aliases and Interfaces");
    helperRunSection(sectionUnionTypes, "Union types");
    helperRunSection(sectionFunctions, "Functions");
    helperRunSection(sectionCasting, "Casting");
    helperRunSection(sectionClasses, "Classes");
    helperRunSection(sectionGenerics, "Generics");
    console.log("\n>>>>>>>>>>> Main function ends\n");
}
function sectionSimpleTypes() {
    // explicit type annotation
    let a1 = 10;
    let a2 = "Hello";
    let a3 = true;
    let a4 = BigInt(2);
    let a5 = 2n;
    let a6 = Symbol("test");
    let a7 = undefined;
    let a8 = null;
    console.log("Explicit variable type definition: ");
    console.log(a1, typeof (a1));
    console.log(a2, typeof (a2));
    console.log(a3, typeof (a3));
    console.log(a4, typeof (a4));
    console.log(a5, typeof (a5));
    console.log(a6, typeof (a6));
    console.log(a7, typeof (a7));
    console.log(a8, typeof (a8));
    console.log();
    // implicit type annotation
    let b1 = 20;
    let b2 = "Hello";
    let b3 = true;
    let b4 = BigInt(2);
    let b5 = 2n;
    let b6 = Symbol("test");
    let b7 = undefined;
    let b8 = null;
    console.log("Implicit variable type definition: ");
    console.log(b1, typeof (b1));
    console.log(b2, typeof (b2));
    console.log(b3, typeof (b3));
    console.log(b4, typeof (b4));
    console.log(b5, typeof (b5));
    console.log(b6, typeof (b6));
    console.log(b7, typeof (b7));
    console.log(b8, typeof (b8));
    let c1;
    // c1 = 8 // error
    // any - disables typechecking
    let d1 = 8;
    console.log(d1, typeof d1);
    d1 = "Hello";
    console.log(d1, typeof d1);
    let d2 = 8;
    console.log(d2, typeof d2);
    // d2 = "Hello" // error - type is infered to be number
    // unknown type - is not known @compile_time
    let d3;
    d3 = "Hello";
    //console.log(d3.length) // error, typescript does not know the type of d3
    console.log(d3.length); // use casting to access the length property
    // never type - throws error when it is defined
    let d4;
    // d4 = 9 // error
}
function sectionArrays() {
    // arrays
    const names = ["John", "Jane", "Mary"];
    names.push("Bob"); // ok, const identifier does not prevent this
    console.log(names, typeof names);
    // names = ["James", "Bond"] // error, const identifier prevents reassignment
    // preventing mutable actions on arrays
    const names2 = ["John", "Jane", "Mary"];
    //names2.push("James") // error, readonly identifier prevents this
    // names2[0] = "James" // error, readonly identifier prevents this
    console.log(names2, typeof names2);
    // type inference
    const numbers = [1, 2, 3];
    console.log(numbers, typeof numbers);
    // numbers.push("Whee") // error, numbers is implicitly of type number[]
    const numbers2 = [1, 2, "9"];
    numbers2.push("Whee"); // ok, numbers2 is implicitly of type any[]
    console.log(numbers2, typeof numbers2);
    const numbers3 = [1, 2, 3];
    numbers3.push("Whee"); // also ok
    console.log(numbers3, typeof numbers3);
}
function sectionTuples() {
    // tuples = (possibly mixed) typed arrays of fixed length
    let person = ["John", 19];
    // person = [19, "John"] // error, wrong order
    console.log(person, typeof person);
    // person.push(19) // this does not throw error, good practice is to make tuples readonly
    let person2 = ["John", 19];
    // person2.push(19) // error, readonly identifier prevents this
    console.log(person2, typeof person2);
    // named tuples - more verbose syntax to increase readability
    let person3 = ["John", 19];
    // person3.name = "Jane" // error, name is not property of person3, just string to improve readability
    console.log(person3, typeof person3);
    // tuples can also be destructured
    const [n, a] = person3;
    console.log(n, a);
}
function sectionObjects() {
    const person = { name: "John", age: 19 };
    person.name = "Jane"; // OK
    //person.name = 9                   // not OK
    // person.isHuman = true            // not OK
    console.log(person);
    // optional properties
    const person2 = { name: "John" };
    console.log(person2);
    person2.age = 19;
    // person2.isHuman = true           // not OK
    // type inference
    const person3 = { name: "John", age: 19 };
    // person3.name = 9                 // not OK
    // person3.isHuman = true           // not OK
    // index signatures -- any properties with specified type of keys and values
    const person4 = {};
    person4.age = 19;
    person4.fingerCount = 20;
    // person4.name = "John"            // not OK
    console.log(person4);
}
function sectionEnums() {
    // numeric enums
    let Direction;
    (function (Direction) {
        Direction[Direction["Up"] = 0] = "Up";
        Direction[Direction["Down"] = 1] = "Down";
        Direction[Direction["Left"] = 2] = "Left";
        Direction[Direction["Right"] = 3] = "Right";
    })(Direction || (Direction = {}));
    const dir = Direction.Up;
    console.log(dir);
    // numeric enums with values
    let StatusCode;
    (function (StatusCode) {
        StatusCode[StatusCode["NotFound"] = 404] = "NotFound";
        StatusCode[StatusCode["Ok"] = 200] = "Ok";
        StatusCode[StatusCode["Custom"] = 201] = "Custom";
        StatusCode[StatusCode["Error"] = 500] = "Error";
        StatusCode[StatusCode["Custom2"] = 501] = "Custom2";
    })(StatusCode || (StatusCode = {}));
    console.log(StatusCode.NotFound, StatusCode.Ok, StatusCode.Custom, StatusCode.Error, StatusCode.Custom2);
    // string enums
    let Direction2;
    (function (Direction2) {
        Direction2["Up"] = "UP";
        Direction2["Down"] = "DOWN";
        Direction2["Left"] = "LEFT";
        // Right // not allowed, all strings must be explicitly entered
        Direction2["Right"] = "RIGHT";
    })(Direction2 || (Direction2 = {}));
    const dir2 = Direction2.Up;
    console.log(dir2);
}
function sectionAliasesAndInterfaces() {
    // aliases
    // const age1 : AgeType = "9" // not allowed
    const age1 = 9;
    const person1 = { name: "John", age: age1 };
    console.log(age1, typeof age1);
    console.log(person1, typeof person1);
    const animal1 = { name: "Rex" };
    console.log(animal1, typeof animal1);
    const animal2 = { name: "Polly", canFly: true };
    console.log(animal2, typeof animal2);
    // const animal3 : Bird = {name: "Jerry"} // not allowed, canFly missing
}
function sectionUnionTypes() {
    let a1 = "Hello";
    console.log(a1, typeof a1);
    console.log(a1.toUpperCase());
    a1 = 9;
    console.log(a1, typeof a1);
    //a1 = true // not allowed
    // a1 = null  // not allowed
    // a1 = undefined // not allowed
    // a1 = {} // not allowed
    const fn1 = (x) => { console.log(x); };
    fn1(a1);
    // const fn2 : (a:number | string) => void = (x:string|number) => {console.log(x.toUpperCase())} // not allowed
}
function sectionFunctions() {
    // function with return value
    const fn1 = function () { return 9; };
    console.log(fn1());
    // function without return value
    const fn2 = function () { console.log(9); };
    fn2();
    // function with parameters
    const fn3 = function (a, b) { console.log(`${a}, ${b}`); };
    fn3(9, "hello");
    // arrow function
    const fn4 = (a) => { return a; };
    console.log(fn4(9));
    // optional parameters
    const fn5 = function (a, b) { console.log(`fn5: a = "${a}", b = "${b}"`); };
    fn5(9);
    fn5(9, "hello");
    // default values
    const fn6 = function (a, b = "world") { console.log(`fn6: a = "${a}", b = "${b}"`); };
    fn6(9);
    fn6(9, "hello");
    // named parameters
    const fn7 = function ({ a, b }) { console.log(`fn7: a = "${a}", b = "${b}"`); };
    fn7({ a: 9, b: "hello" });
    fn7({ b: "hello", a: 9 });
    // fn7({a:9}) // not allowed, b is required
    // rest parameters -- type must be always an array
    const fn8 = function (...args) { console.log(args); };
    fn8();
    fn8(9);
    fn8(9, 10);
    const fn9 = function (a, b) { console.log(`fn9: a = "${a}", b = "${b}"`); };
    fn9(9, "hello");
}
function sectionCasting() {
    // casting with as
    let a1 = 'hello';
    console.log(a1.length);
    a1 = 9;
    console.log(a1.length); // allowed, but prints undefined
    // casting with < >
    let a2 = 'hello';
    console.log(a2.length); // does not work in tsx files
    // force casting -- going via unknow intermediary cast
    let a3 = 'hello';
    // console.log(a3 as number) // not allowed
    console.log(a3); // allowed
}
function sectionClasses() {
    class Person {
        constructor(name) {
            this.name = "John";
            this.name = name;
        }
        greet(forName) {
            return `[${this.name} speaking]: Hello, ${forName}`;
        }
    }
    const p1 = new Person('John');
    console.log(p1.greet('Bob'));
    // default member visibility is public
    class Car {
        constructor(color, price, isTruck) {
            this.color = color;
            this.price = price;
            this.isTruck = isTruck;
        }
        getPrice() {
            return this.price;
        }
    }
    const c1 = new Car('red', 10000, true);
    console.log(c1.color);
    //console.log(c1.price) // not allowed
    console.log(c1.getPrice()); // allowed
    //console.log(c1.isTruck) // not allowed
    // shorthand : defining class parameters in constructor
    class Rectangle {
        constructor(width, height) {
            this.width = width;
            this.height = height;
            this.area = width * height; // must be also declared
        }
        getArea() {
            return this.area;
        }
    }
    const r1 = new Rectangle(20, 30);
    console.log(r1.width, r1.height, r1.getArea()); // allowed
    // readonly class members -- cannot be set after initialization
    class Line {
        constructor(x) {
            this.x = x;
        }
    }
    const l1 = new Line(10);
    class Duck {
        constructor(name) {
            this.name = name;
        }
        quack() {
            console.log(`[${this.name}]: quack`);
        }
    }
    const d1 = new Duck('Jacob');
    d1.quack(); // duck quacks
    class Bird {
        constructor(name) {
            this.name = name;
        }
        quack() {
            console.log(`[${this.name}]: quack`);
        }
        fly() {
            console.log(`[${this.name}]: fly`);
        }
    }
    const b1 = new Bird('Anna');
    b1.quack(); // duck quacks
    b1.fly(); // duck flies
    // inheritance with extends keyword
    // a class can extend another class
    // only one class can be extended
    class Animal {
        constructor(name) {
            this.name = name;
        }
        move() {
            console.log(`${this.name}: move`);
        }
    }
    class Dog extends Animal {
        constructor(name) {
            super(name);
        }
        bark() {
            console.log(`${this.name}: bark`);
        }
    }
    const d = new Dog('Rex');
    d.move(); // dog moves
    d.bark(); // dog barks
    // override keyword can be used to signify that a method is overriding parent method
    class Bird2 extends Animal {
        constructor(name) {
            super(name);
        }
        move() {
            // override the parent's method
            console.log(`${this.name}: fly`);
        }
    }
    const b = new Bird2('Polly');
    b.move(); // bird flies
    // abstract classes define abstract methods that are not implemented
    // child class must override these methods
    class AbstractAnimal {
        constructor(name) {
            this.name = name;
        }
    }
    class Cat extends AbstractAnimal {
        constructor(name) {
            super(name);
        }
        move() {
            // override the parent's method
            console.log(`${this.name}: walks sneakily`);
        }
    }
    const cat1 = new Cat("Kat");
    cat1.move(); // cat walks sneakily
    // const aa1 = new AbstractAnimal('Cat') // not allowed, abstract class cannot be instantiated 
}
function sectionGenerics() {
    // generics functions
    // assigning generic function to variable
    const makePair = function (value1, value2) {
        return [value1, value2];
    };
    console.log(makePair("hello", 3));
    // syntax without assigning to variable
    function makePair2(value1, value2) {
        return [value1, value2];
    }
    console.log(makePair2(true, false));
    // syntax with arrow function
    const makePair3 = (value1, value2) => {
        return [value1, value2];
    };
    console.log(makePair3("hello", 3));
    const x = { value: "hello" };
    console.log(x);
    // generic classes
    class WrappedValue2 {
        constructor(value) { }
    }
    const y = new WrappedValue2("hello");
    console.log(y);
    //const z : WrappedValue3 = {"hello"} // not allowed
    const z = { value: 3 };
    console.log(z);
    const z2 = { value: "hello" }; // this is OK
    console.log(z2);
    // const z3 : WrappedValue4<string> = {"hello"} // not allowed
    const z3 = { value: 10 }; // this is OK
    console.log(z3);
    const z4 = { value: null };
    console.log(z4);
    const z5 = { value: 9 };
    console.log(z5);
}
function helperRunSection(sectionFunction, sectionName) {
    console.log("\n***********************************************");
    console.log(`*********** Running section: ${sectionName} ***********`);
    console.log("***********************************************\n");
    sectionFunction();
}
main();
