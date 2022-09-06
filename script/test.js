"use strict";

//import { sayHi } from './module.js';

console.log('random = ', Math.random())

console.log("Hello from Js, " + u.Name + "!")
u.SetToken('dddd')

console.log('underscore random = ', _.random(0, 100))
//console.log('underscore now time = ', _.now())

var compiled = _.template("hello: <%= name %>");
console.log('underscore template = ', compiled({name: 'moe'}))

var func = function (greeting) {
    return greeting + ': ' + this.name
};
func = _.bind(func, {name: 'moe'}, 'hi');
console.log('underscore function = ', func())

function printMessage(msg) {
    console.log('js print message: ', msg)
}

sayHello("Xyzzy")
