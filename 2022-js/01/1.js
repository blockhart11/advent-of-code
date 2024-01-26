import { input } from './input.js';
import { validate } from '@apollo/client';

let max = 0
let current = 0
input.split("\n").forEach((item) => {
    if (item === "") {
        if (current > max) {
            max = current
        }
        current = 0
    } else {
        current += parseInt(item)
    }
})

console.log(max)