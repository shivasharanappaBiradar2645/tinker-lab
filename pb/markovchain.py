import numpy as np 


def networkTrafficPred(n):
    states = ["L","M","H","O"]
    P = np.array([
    [0.6,  0.3, 0.1, 0.0],
    [0.2,  0.5, 0.2, 0.1],
    [0.1,  0.3, 0.4, 0.2],
    [0.05, 0.25,0.4, 0.3]
])
    state = np.array([1,0,0,0])

    for i in range(n):
        state = state @ P
    print(f"the highest probability after {n} steps is ",states[np.argmax(state)])

networkTrafficPred(5)

def FraudDetPred(n):
    states = ["Normal","suspicious","fraud"]
    P = np.array([
    [0.95, 0.04, 0.01],
    [0.30, 0.50, 0.20],
    [0.02, 0.08, 0.90]
])
    state = np.array([0,1,0])

    for i in range(n):
        state = state @ P
        state = state / state.sum() 
    print(f"the fraud probabilites after {n} is: ",state)

FraudDetPred(5)

