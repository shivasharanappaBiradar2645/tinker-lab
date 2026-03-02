import numpy as np

states = ["A","B"]
observations = ["H","T"]

pi = np.array([0.5, 0.5])

#transition matrix
A = np.array([
    [0.9, 0.1],   
    [0.2, 0.8]    
])
#emission matrix
B = np.array([
    [0.5, 0.5],   
    [0.9, 0.1]    
])

sequence = ["H", "H", "T","H","H","H","T"]
obs_idx = [0 if x=="H" else 1 for x in sequence]

T = len(obs_idx)
N = len(states)

dp = np.zeros((T, N))

for i in range(N):
    dp[0, i] = pi[i] * B[i, obs_idx[0]]

for t in range(1, T):
    for j in range(N):
        probs = dp[t-1] * A[:, j]
        dp[t, j] = np.max(probs) * B[j, obs_idx[t]]


print("DP Table:")
print(dp)

final_state = np.argmax(dp[T-1])
print("Most likely final state:", states[final_state])

