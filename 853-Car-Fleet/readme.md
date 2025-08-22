## 853. Car Fleet

https://leetcode.com/problems/car-fleet/

## Title
**Car Fleet Problem**

## Description
Given `n` cars traveling towards a target destination, determine how many distinct car fleets will arrive at the target. Each car has an initial position and constant speed. Cars that meet or catch up during travel form a fleet and continue at the slower car's speed.

**Key Concepts:**
- Cars closer to the target with slower speeds can block faster cars behind them
- When a faster car catches up to a slower car, they form a fleet
- Fleets travel at the speed of the slowest car in the group

**Example:**
```
Target: 12, Positions: [10,8,0,5,3], Speeds: [2,4,1,1,3]

Car positions and times to reach target:
- Car at pos 10, speed 2: time = (12-10)/2 = 1.0
- Car at pos 8, speed 4: time = (12-8)/4 = 1.0  → Forms fleet with car above
- Car at pos 5, speed 1: time = (12-5)/1 = 7.0
- Car at pos 3, speed 3: time = (12-3)/3 = 3.0  → Forms fleet with car above
- Car at pos 0, speed 1: time = (12-0)/1 = 12.0

Result: 3 fleets
```

**Algorithm:** Sort cars by position (descending), calculate arrival times, and merge fleets when faster cars catch up to slower ones.

**Time Complexity:** O(n log n) due to sorting  
**Space Complexity:** O(n) for storing car pairs and times