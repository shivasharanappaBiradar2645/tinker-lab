import numpy as np

def get_probs(state):
    a = alpha[state]
    return a / a.sum()


states = ["N","S","F"]

#the model to learn only from the observations, start with np.array([1, 1, 1]). This is called Laplace Smoothing.
alpha = {
        "N": np.array([95,4,1]),
        "S": np.array([30,50,20]),
        "F": np.array([2,8,90])
    }

observations = [
    ("N","N"),
    ("N","F"),
    ("N","S"),
    ("S","F"),
    ("S","F")
]

index = {"N":0, "S":1, "F":2}

for s, s2 in observations:
    alpha[s][index[s2]] += 1

for s in states:
    print(s, get_probs(s))


#if initial state is unknown ([0.33,0.33,0.33])

