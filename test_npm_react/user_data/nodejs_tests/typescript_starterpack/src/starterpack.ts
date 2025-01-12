// c.f. https://www.w3schools.com/typescript/index.php
// examples of type-hint syntax of typescript for various use-cases


function main() {
    console.log("\n>>>>>>>>>>> Main function starts\n")

    helperRunSection(sectionSimpleTypes, "Simple types")
    helperRunSection(sectionArrays, "Arrays")
    helperRunSection(sectionTuples, "Tuples")
    helperRunSection(sectionObjects, "Objects")
    helperRunSection(sectionEnums, "Enums")
    helperRunSection(sectionAliasesAndInterfaces, "Aliases and Interfaces")
    helperRunSection(sectionUnionTypes, "Union types")
    helperRunSection(sectionFunctions, "Functions")
    helperRunSection(sectionCasting, "Casting")
    helperRunSection(sectionClasses, "Classes")
    helperRunSection(sectionGenerics, "Generics")

    console.log("\n>>>>>>>>>>> Main function ends\n")
}

function sectionSimpleTypes() {
    // explicit type annotation
    let a1: number = 10
    let a2: string = "Hello"
    let a3: boolean = true
    let a4: bigint = BigInt(2)
    let a5: bigint = 2n
    let a6: symbol = Symbol("test")
    let a7: undefined = undefined
    let a8: null = null

    console.log("Explicit variable type definition: ")
    console.log(a1, typeof(a1))
    console.log(a2, typeof(a2))
    console.log(a3, typeof(a3))
    console.log(a4, typeof(a4))
    console.log(a5, typeof(a5))
    console.log(a6, typeof(a6))
    console.log(a7, typeof(a7))
    console.log(a8, typeof(a8))

    console.log()

    // implicit type annotation
    let b1 = 20
    let b2 = "Hello"
    let b3 = true
    let b4 = BigInt(2)
    let b5 = 2n
    let b6 = Symbol("test")
    let b7 = undefined
    let b8 = null     
    console.log("Implicit variable type definition: ")
    console.log(b1, typeof(b1))
    console.log(b2, typeof(b2))
    console.log(b3, typeof(b3))
    console.log(b4, typeof(b4))
    console.log(b5, typeof(b5))
    console.log(b6, typeof(b6))
    console.log(b7, typeof(b7))
    console.log(b8, typeof(b8))

    let c1: string
    // c1 = 8 // error

    // any - disables typechecking
    let d1 : any = 8
    console.log(d1, typeof d1)
    d1 = "Hello"
    console.log(d1, typeof d1)

    let d2 = 8
    console.log(d2, typeof d2)
    // d2 = "Hello" // error - type is infered to be number

    // unknown type - is not known @compile_time
    let d3: unknown
    d3 = "Hello"
    //console.log(d3.length) // error, typescript does not know the type of d3
    console.log((d3 as string).length) // use casting to access the length property

    // never type - throws error when it is defined
    let d4: never
    // d4 = 9 // error
}

function sectionArrays() {
    // arrays
    const names: string[] = ["John", "Jane", "Mary"]
    names.push("Bob") // ok, const identifier does not prevent this
    console.log(names, typeof names)

    // names = ["James", "Bond"] // error, const identifier prevents reassignment

    // preventing mutable actions on arrays
    const names2: readonly string[] = ["John", "Jane", "Mary"]
    //names2.push("James") // error, readonly identifier prevents this
    // names2[0] = "James" // error, readonly identifier prevents this
    console.log(names2, typeof names2)

    // type inference
    const numbers = [1, 2, 3]
    console.log(numbers, typeof numbers)
    // numbers.push("Whee") // error, numbers is implicitly of type number[]
    const numbers2 = [1, 2, "9"]
    numbers2.push("Whee") // ok, numbers2 is implicitly of type any[]
    console.log(numbers2, typeof numbers2)
    const numbers3: any[] = [1, 2, 3]
    numbers3.push("Whee") // also ok
    console.log(numbers3, typeof numbers3)

}

function sectionTuples(){
    // tuples = (possibly mixed) typed arrays of fixed length
    let person: [string, number] = ["John", 19]
    // person = [19, "John"] // error, wrong order
    console.log(person, typeof person)
    
    // person.push(19) // this does not throw error, good practice is to make tuples readonly

    let person2: readonly [string, number] = ["John", 19]
    // person2.push(19) // error, readonly identifier prevents this
    console.log(person2, typeof person2)

    // named tuples - more verbose syntax to increase readability
    let person3: [name: string, age: number] = ["John", 19]
    // person3.name = "Jane" // error, name is not property of person3, just string to improve readability
    console.log(person3, typeof person3)
    
    // tuples can also be destructured
    const [n, a] = person3
    console.log(n, a)
}

function sectionObjects(){
    const person: {name: string, age: number} = {name: "John", age: 19}
    person.name = "Jane"                // OK
    //person.name = 9                   // not OK
    // person.isHuman = true            // not OK
    console.log(person)

    // optional properties
    const person2: {name: string, age?: number} = {name: "John"} 
    console.log(person2)
    person2.age = 19
    // person2.isHuman = true           // not OK

    // type inference
    const person3 = {name: "John", age: 19}
    // person3.name = 9                 // not OK
    // person3.isHuman = true           // not OK

    // index signatures -- any properties with specified type of keys and values
    const person4 : {[index: string]: number} = {}
    person4.age = 19
    person4.fingerCount = 20
    // person4.name = "John"            // not OK
    console.log(person4)
}

function sectionEnums(){
    // numeric enums
    enum Direction {
        Up, // enumerates from 0 by default
        Down,
        Left,
        Right
    }
    const dir = Direction.Up
    console.log(dir)

    // numeric enums with values
    enum StatusCode {
        NotFound = 404,
        Ok = 200,
        Custom, // continues by incrementing since last value
        Error = 500, 
        Custom2
    }
    console.log(StatusCode.NotFound, StatusCode.Ok, StatusCode.Custom, StatusCode.Error, StatusCode.Custom2)

    // string enums
    enum Direction2 {
        Up = "UP",
        Down = "DOWN",
        Left = "LEFT",
        // Right // not allowed, all strings must be explicitly entered
        Right = "RIGHT"
    }
    const dir2 = Direction2.Up
    console.log(dir2)
}

function sectionAliasesAndInterfaces(){
    // aliases

    type AgeType = number
    type Person = { name: string, age: AgeType }

    // const age1 : AgeType = "9" // not allowed
    const age1 : AgeType = 9
    const person1 : Person = { name: "John", age: age1 }
    console.log(age1, typeof age1)
    console.log(person1, typeof person1)

    // interfaces - similar to aliases but apply only to objects
    interface Animal {
        name: string
    }

    const animal1 : Animal = {name: "Rex"}
    console.log(animal1, typeof animal1)

    // interfaces can extend other interfaces
    interface Bird extends Animal {
        canFly: boolean
    }
    
    const animal2 : Bird = {name: "Polly", canFly: true}
    console.log(animal2, typeof animal2)

    // const animal3 : Bird = {name: "Jerry"} // not allowed, canFly missing
    
}

function sectionUnionTypes (){
    let a1: number | string = "Hello"
    console.log(a1, typeof a1)
    console.log(a1.toUpperCase())
    a1 = 9
    console.log(a1, typeof a1)
    //a1 = true // not allowed
    // a1 = null  // not allowed
    // a1 = undefined // not allowed
    // a1 = {} // not allowed

    const fn1 : (a:number | string) => void = (x:number | string) => {console.log(x)}
    fn1(a1)
    // const fn2 : (a:number | string) => void = (x:string|number) => {console.log(x.toUpperCase())} // not allowed
}

function sectionFunctions() {
    // function with return value

    const fn1 = function(): number {return 9}
    console.log(fn1())

    // function without return value
    const fn2 = function():void {console.log(9)}
    fn2()

    // function with parameters
    const fn3 = function(a:number, b:string): void {console.log(`${a}, ${b}`)}
    fn3(9, "hello")

    // arrow function
    const fn4 = (a:number): number => {return a}
    console.log(fn4(9))

    // optional parameters
    const fn5 = function(a:number, b?:string):void {console.log(`fn5: a = "${a}", b = "${b}"`)}
    fn5(9)
    fn5(9,"hello")

    // default values
    const fn6 = function(a:number, b:string = "world"):void {console.log(`fn6: a = "${a}", b = "${b}"`)}
    fn6(9)
    fn6(9,"hello")

    // named parameters
    const fn7 = function({a, b} : {a:number, b:string}):void {console.log(`fn7: a = "${a}", b = "${b}"`)}
    fn7({a:9,b:"hello"})
    fn7({b:"hello", a:9})
    // fn7({a:9}) // not allowed, b is required

    // rest parameters -- type must be always an array
    const fn8 = function(...args:number[]):void {console.log(args)}
    fn8()
    fn8(9)
    fn8(9,10)

    // function type alias
    type fnType = (a:number, b:string)=>void;
    const fn9 :fnType = function(a,b) {console.log(`fn9: a = "${a}", b = "${b}"`)}
    fn9(9,"hello")
}

function sectionCasting() {
    // casting with as
    let a1: unknown = 'hello';
    console.log((a1 as string).length)

    a1 = 9
    console.log((a1 as string).length) // allowed, but prints undefined

    // casting with < >
    let a2 = 'hello'
    console.log((<string>a2).length) // does not work in tsx files

    // force casting -- going via unknow intermediary cast
    let a3 = 'hello'
    // console.log(a3 as number) // not allowed
    console.log((a3 as unknown) as number) // allowed
}

function sectionClasses(){
    class Person{
        name:string = "John"

        constructor(name:string){
            this.name = name
        }

        greet(forName:string): string{
            return `[${this.name} speaking]: Hello, ${forName}`
        }
    }

    const p1 = new Person('John')
    console.log(p1.greet('Bob'))

    // default member visibility is public
    class Car {
        public color: string
        private price: number
        protected isTruck: boolean
        public constructor(color:string, price:number, isTruck:boolean){
            this.color = color
            this.price = price
            this.isTruck = isTruck
        }

        public getPrice(): number{
            return this.price
        }
    }

    const c1 = new Car('red', 10000, true)
    console.log(c1.color)
    //console.log(c1.price) // not allowed
    console.log(c1.getPrice()) // allowed
    //console.log(c1.isTruck) // not allowed

    // shorthand : defining class parameters in constructor
    class Rectangle{
        private area: number

        constructor(public width:number, public height:number){
            this.area = width * height // must be also declared
        }

        public getArea(): number{
            return this.area
        }
    }

    const r1 = new Rectangle(20, 30)
    console.log(r1.width, r1.height, r1.getArea()) // allowed

    // readonly class members -- cannot be set after initialization
    class Line{
        constructor(public readonly x: number){}

        // also not allowed
        // public setX(x:number){
        //     this.x = x
        // }
    }

    const l1 = new Line(10)
    // l1.x = 20; // not allowed

    // inheritance with implements keyword
    // interfaces can be implemented by classes
    interface Quackable{
        quack():void
    }

    class Duck implements Quackable{
        constructor(public name:string){}
        quack(){
            console.log(`[${this.name}]: quack`)
        }
    }

    const d1 = new Duck('Jacob')
    d1.quack() // duck quacks

    // a class can implement more than one interface
    interface Flyable{
        fly():void
    }
    class Bird implements Quackable,Flyable{
        constructor(public name:string){}
        quack(){
            console.log(`[${this.name}]: quack`)
        }
        fly(){
            console.log(`[${this.name}]: fly`)
        }
    }
    const b1 = new Bird('Anna')
    b1.quack() // duck quacks
    b1.fly() // duck flies

    // inheritance with extends keyword
    // a class can extend another class
    // only one class can be extended
    class Animal{
        constructor(public name:string){}
        move(){
            console.log(`${this.name}: move`)
        }
    }
    class Dog extends Animal{
        constructor(name:string){
            super(name)
        }
        bark(){
            console.log(`${this.name}: bark`)
        }
    }
    const d = new Dog('Rex')
    d.move() // dog moves
    d.bark() // dog barks

    // override keyword can be used to signify that a method is overriding parent method
    class Bird2 extends Animal{
        constructor(name:string){
            super(name)
        }
        override move(){
            // override the parent's method
            console.log(`${this.name}: fly`)
        }
    }
    const b = new Bird2('Polly')
    b.move() // bird flies

    // abstract classes define abstract methods that are not implemented
    // child class must override these methods

    abstract class AbstractAnimal{
        constructor(public name:string){}
        abstract move():void;
    }

    class Cat extends AbstractAnimal{
        constructor(name:string){
            super(name)
        }
        override move(){
            // override the parent's method
            console.log(`${this.name}: walks sneakily`)
        }
    }

    const cat1 = new Cat("Kat")
    cat1.move() // cat walks sneakily

    // const aa1 = new AbstractAnimal('Cat') // not allowed, abstract class cannot be instantiated 
}

function sectionGenerics() {
    // generics functions
    
    // assigning generic function to variable
    const makePair = function <S, T>(value1:S, value2:T):[S,T]{
        return [value1,value2]
    }

    console.log(makePair<string, number>("hello", 3))

    // syntax without assigning to variable
    function makePair2 <S, T>(value1:S, value2:T):[S,T]{
        return [value1,value2]
    }

    console.log(makePair2<boolean, boolean>(true, false))

    // syntax with arrow function
    const makePair3 = <S, T>(value1:S, value2:T):[S,T] => {
        return [value1,value2]
    }
    console.log(makePair3<string, number>("hello", 3))

    // generic types
    type WrappedValue<T> = {value:T}

    const x: WrappedValue<string> = {value:"hello"}
    console.log(x)

    // generic classes
    class WrappedValue2<T>{
        constructor(value:T){}
    }

    const y = new WrappedValue2<string>("hello");
    console.log(y)

    // generics can have default type
    type WrappedValue3<T = number> = {value:T}
    //const z : WrappedValue3 = {"hello"} // not allowed
    const z: WrappedValue3 = {value:3}
    console.log(z)

    const z2: WrappedValue3<string> = {value:"hello"} // this is OK
    console.log(z2)

    // limiting generics type using extends syntax
    type WrappedValue4<T extends number | null> = {value:T}
    // const z3 : WrappedValue4<string> = {"hello"} // not allowed
    const z3: WrappedValue4<number> = {value:10} // this is OK
    console.log(z3)
    const z4: WrappedValue4<null> = {value:null}
    console.log(z4)
    const z5: WrappedValue4<number | null> = {value:9}
    console.log(z5)

}



function helperRunSection(sectionFunction: () => void, sectionName: string) {
    console.log("\n***********************************************")
    console.log(`*********** Running section: ${sectionName} ***********`)
    console.log("***********************************************\n")
    sectionFunction()
}

main();

