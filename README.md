# Sudoku Solver in Go

This is a simple Sudoku solver written in Go. It uses the following algorithm.

1. Iterate over all the emtpy cells until error occurs or solution is found.
2. For each emtpy cell, find the domain of that cell.
3. If any cell has empty domain, log that there is no solution
4. If any cell has domain of size 1, fill that cell with that value
5. If no single value cell is found, return with no solution
6. If all cells are filled, return with solution
