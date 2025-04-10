const s = Date.now();
const n = 10_000_000;
const arr = new Array(n).fill(0).map((_, i) => i * 2);
const sum = arr.reduce((acc, v) => acc + v, 0);
console.log(`Memory Benchmark: Sum: ${sum} | Time: ${Date.now() - s}ms`);
