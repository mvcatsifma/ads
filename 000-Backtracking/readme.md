# **Backtracking: Solution Steps**

### **1. Define the state**

Everything needed to describe where you are:
`i`, `subset`, `visited`, `board`, etc.

---

### **2. Base case**

If the state is complete, record the result and return.

---

### **3. Identify valid choices**

For subsets:

* include nums[i]
* exclude nums[i]

For permutations:

* any unused number

For N queens:

* any safe column

---

### **4. Apply → Recurse → Undo**

This is the engine of backtracking:

```code
for _, choice := range validChoices(state) {
    apply(&state, choice)   // move into next state
    backtrack(state)        // explore branch
    undo(&state, choice)    // restore state
}
```

Undo ensures the next choice starts clean.

---

# Updated Universal Template (Short)

```code
var backtrack func(state State)
backtrack = func(state State) {
    if isComplete(state) {
        record(state)
        return
    }

    for _, choice := range validChoices(state) {
        apply(&state, choice)
        backtrack(state)
        undo(&state, choice)
    }
}

backtrack(initialState)
```

---

# For subsets, the two choices are simply unrolled:

```code
// include
subset = append(subset, nums[i])
dfs(i+1)
subset = subset[:len(subset)-1]

// exclude
dfs(i+1)
```

Two choices → two branches → 2 recursive calls.

---

If you want, I can now show this template applied to *permutations*, *combinations*, or *N-Queens* in under 10 lines each.
*Backtracking: Solution Steps**

### **1. Define the state**

Everything needed to describe where you are:
`i`, `subset`, `visited`, `board`, etc.
Here is the **short, clean, correct** version.

---

# **Backtracking: Solution Steps**

### **1. Define the state**

Everything needed to describe where you are:
`i`, `subset`, `visited`, `board`, etc.

---

### **2. Base case**

If the state is complete, record the result and return.

---

### **3. Identify valid choices**

For subsets:

* include nums[i]
* exclude nums[i]

For permutations:

* any unused number

For N queens:

* any safe column

---

### **4. Apply → Recurse → Undo**

This is the engine of backtracking:

```code
for _, choice := range validChoices(state) {
    apply(&state, choice)   // move into next state
    backtrack(state)        // explore branch
    undo(&state, choice)    // restore state
}
```

Undo ensures the next choice starts clean.

---

# Updated Universal Template (Short)

```code
var backtrack func(state State)
backtrack = func(state State) {
    if isComplete(state) {
        record(state)
        return
    }

    for _, choice := range validChoices(state) {
        apply(&state, choice)
        backtrack(state)
        undo(&state, choice)
    }
}

backtrack(initialState)
```

---

# For subsets, the two choices are simply unrolled:

```code
// include
subset = append(subset, nums[i])
dfs(i+1)
subset = subset[:len(subset)-1]

// exclude
dfs(i+1)
```

Two choices → two branches → 2 recursive calls.

---

If you want, I can now show this template applied to *permutations*, *combinations*, or *N-Queens* in under 10 lines each.

---

### **2. Base case**

If the state is complete, record the result and return.

---

### **3. Identify valid choices**

For subsets:

* include nums[i]
* exclude nums[i]

For permutations:

* any unused number

For N queens:

* any safe column

---

### **4. Apply → Recurse → Undo**

This is the engine of backtracking:

```code
for _, choice := range validChoices(state) {
    apply(&state, choice)   // move into next state
    backtrack(state)        // explore branch
    undo(&state, choice)    // restore state
}
```

Undo ensures the next choice starts clean.

---

# Updated Universal Template (Short)

```code
var backtrack func(state State)
backtrack = func(state State) {
    if isComplete(state) {
        record(state)
        return
    }

    for _, choice := range validChoices(state) {
        apply(&state, choice)
        backtrack(state)
        undo(&state, choice)
    }
}

backtrack(initialState)
```

---

# For subsets, the two choices are simply unrolled:

```code
// include
subset = append(subset, nums[i])
dfs(i+1)
subset = subset[:len(subset)-1]

// exclude
dfs(i+1)
```

Two choices → two branches → 2 recursive calls.

---

If you want, I can now show this template applied to *permutations*, *combinations*, or *N-Queens* in under 10 lines each.
