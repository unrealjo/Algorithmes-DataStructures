# Bubble Sort

## Description

Bubble sort is a simple sorting algorithm that works by repeatedly swapping adjacent elements in an array if they are in the wrong order. The algorithm starts at the beginning of the array and compares each pair of adjacent elements. If the elements are in the wrong order, they are swapped. This process is repeated until the entire array is sorted.

![Bubble sort](bubble-sort.gif)

## Complexity

The algorithm has a time complexity of O(n^2), making it inefficient for large arrays. However, it is easy to understand and implement, making it a good choice for small arrays or educational purposes. The algorithm can be optimized using the "short-bubble optimization", which stops the sorting process if no swaps are made during a single pass, reducing the number of steps required to sort the array.

<table border="0"><tbody>
<tr><th><strong>Time Complexity</strong></th>
<td>&nbsp;</td>
</tr><tr><td>Best</td>
	<td>O(n)</td>
</tr><tr><td>Worst</td>
	<td>O(n<sup>2</sup>)</td>
</tr><tr><td>Average</td>
	<td>O(n<sup>2</sup>)</td>
</tr><tr><th><strong>Space Complexity</strong></th>
	<td>O(1)</td>
</tr><tr><th><strong>Stability</strong></th>
	<td>Yes</td>
</tr></tbody>
</table>
