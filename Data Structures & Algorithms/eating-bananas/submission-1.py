class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        left = 1
        right = max(piles)

        while left < right:
            mid = (left + right) // 2
            total_hours = 0
            for p in piles:
            # 向上取整：ceil(p / mid)
                total_hours += math.ceil(float(p) / mid)

            if total_hours <= h:
            # 速度mid可行，尝试找更小k
                right = mid
            else:
            # 速度太慢，需要增大
                left = mid + 1
        return left
