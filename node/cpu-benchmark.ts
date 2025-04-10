// cpu-benchmark.js
const isPrime = (n: number): boolean => {
  if (n < 2) return false;
  for (let i = 2; i <= Math.sqrt(n); i++) {
    if (n % i === 0) return false;
  }
  return true;
};

const start = Date.now();
let count = 0;
const N = 1_000_000;
for (let i = 0; i < N; i++) {
  if (isPrime(i)) count++;
}
console.log(`Found ${count} primes in ${Date.now() - start}ms`);
