console.log("Hello world");
console.log("argv=", process.argv);
console.log("typeof argv=", typeof process.argv);
let user_args = process.argv.slice(2);
console.log("user_args=", user_args);
console.log("typeof user_args=", typeof user_args);