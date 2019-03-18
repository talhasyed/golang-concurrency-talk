interface FibCalcJobResult {
  id: number;
  value: number;
  result: number;
}

const fib = (n: number): number => {
  if (n <= 1) return n;

  return fib(n - 1) + fib(n - 2);
};

const randInt = (min: number, max: number): number => {
  min = Math.ceil(min);
  max = Math.floor(max);

  return Math.floor(Math.random() * (max - min + 1)) + min;
};

const calcFib = async (id: number, value: number): Promise<FibCalcJobResult> => {
  const result = fib(value);
  return { id: id, value: value, result: result };
};

const NumCalcs = 20;

async function main() {
  console.time("calculateFibs");

  const fibCalcJobs = [...Array(NumCalcs)].map((_, i) =>
    calcFib(i + 1, randInt(39, 40))
  );

  const results = await Promise.all(fibCalcJobs);
  results.forEach(result =>
    console.log(`[${result.id}] \tfib(${result.value}) \t${result.result}`)
  );

  console.timeEnd("calculateFibs");
}

main();
