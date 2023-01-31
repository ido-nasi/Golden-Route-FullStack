# The Golden Route

## Task 1 part 4.c
`m = base mass`
`M = cargo mass`

The acceleration is defined as follows:
$$a = \frac{F}{M_T} => \frac{F}{m+M}$$

We also know the formula:
$$t = \frac{V}{a}\ => \frac{V}{\frac{F}{m+M}} => \frac{V(m+M)}{F}$$

We can formulate the max cargo mass with the max take off time:
$$t_{max} = \frac{V(m+M_{max})}{F}$$
$$t_{max}\times F = V(m+M_{max})$$
$$\frac{t_{max}\times F}{V} = m+M_{max}$$
$$M_{max}=\frac{t_{max}\times F}{V}-m$$

<p>
The known max take off time is 60 seconds. We can calculate the exact Max Mass for the cargo based on that:
</p>

$$M_{max}=\frac{60\times100,000}{140}-35,000=7857.142857[kg]$$


It means that for every `M`, larger than the `M_max`, the take off time will exceed its limits. To find the excess mass we can subtract the current mass with the max cargo mass: 

$$M_{excess} = M_{current} - M_{max}$$
