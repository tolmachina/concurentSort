# Solution to a concurrent sorting problem #

Problem statement:
Sort a given array by dividing it into 4 subarrays, concurrently sorting each and merging them into one sorted aray.
The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

Solution using 4 goroutines running bubble sort on each of the subarrays, wait group to wait for goroutines to finish and custom merge with a particular counter type implemented.
