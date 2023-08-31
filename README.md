# linear-stats

This project is a Go program that reads numerical data from a file, calculates the Linear Regression Line, and the Pearson Correlation Coefficient for that data. It is an educational tool to help understand basic statistics and probability calculations.

## Usage/Examples

Place your data in a file where each number is on a new line. For example, you can create a file called data.txt and place your numerical data in it.

```bash
189
113
121
114
145
110
```

Run the program using:

```bash
go run . data.txt
```

The program will output the Linear Regression Line and Pearson Correlation Coefficient in 
the specified format:

```bash
Linear Regression Line: y = 0.054322x + 121.421200
Pearson Correlation Coefficient: 0.0023145821
```