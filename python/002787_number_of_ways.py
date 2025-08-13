from dataclasses import dataclass
from typing import Set, List, Optional


from math import isqrt

MOD = 10**9 + 7

class Solution:
    def numberOfWays(self, n: int, x: int) -> int:
        # Precompute all x-th powers <= n
        powers = []
        i = 1
        while True:
            p = i ** x
            if p > n:
                break
            powers.append(p)
            i += 1

        # dp[s] = number of ways to make sum s using a subset of powers (each at most once)
        dp = [0] * (n + 1)
        dp[0] = 1

        # 0/1 knapsack: iterate sums backward to enforce uniqueness
        for p in powers:
            for s in range(n, p - 1, -1):
                dp[s] = (dp[s] + dp[s - p]) % MOD

        return dp[n]

if __name__ == "__main__":
    n = 10
    x = 2
    solution = Solution()
    result = solution.numberOfWays(n, x)
    print(f"Number of ways to express {n} as a sum of powers of {x}: {result}")
