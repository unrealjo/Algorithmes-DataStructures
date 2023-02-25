# Quick Sort

## Description

Quick sort is a sorting algorithm that uses a divide-and-conquer approach. It selects a pivot element from the array and partitions the remaining elements into two sub-arrays, one with elements less than the pivot and one with elements greater than the pivot. The sub-arrays are then recursively sorted using the same process.

![Quick Sort](quick-sort.gif)

## Complexity

The algorithm is efficient, with an average time complexity of O(n log n) and a worst-case complexity of O(n^2), but is widely used due to its ease of implementation and efficiency in practice.

<table border="0"><tbody>
<tr><th><strong>Time Complexity</strong></th>
<td>&nbsp;</td>
</tr><tr><td>Best</td>
	<td>O(n*log n)</td>
</tr><tr><td>Worst</td>
	<td>O(n<sup>2</sup>)</td>
</tr><tr><td>Average</td>
	<td>O(n*log n)</td>
</tr><tr><th><strong>Space Complexity</strong></th>
	<td>O(log n)</td>
</tr><tr><th><strong>Stability</strong></th>
	<td>Yes</td>
</tr></tbody>
</table>
