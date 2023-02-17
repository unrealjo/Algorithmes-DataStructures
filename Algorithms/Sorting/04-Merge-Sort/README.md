# Selection Sort

## Description

Merge sort is a divide-and-conquer sorting algorithm that works by recursively dividing the input array into two halves, sorting each half, and merging the sorted halves back together. The algorithm divides the input array in half, and continues to divide each half until there are only single elements remaining, which are by definition sorted. The sorted halves are then merged back together by comparing the elements in each half and placing them in the correct order.

![Merge Sort](merge-sort.png)

## Complexity

The algorithm has a time complexity of O(n log n), making it an efficient sorting algorithm for large arrays. It is also a stable sorting algorithm, meaning that it preserves the relative order of equal elements in the input array.

<table border="0"><tbody>
<tr><th><strong>Time Complexity</strong></th>
<td>&nbsp;</td>
</tr><tr><td>Best</td>
	<td>O(n*log n)</td>
</tr><tr><td>Worst</td>
	<td>O(n*log n)</td>
</tr><tr><td>Average</td>
	<td>O(n*log n)</td>
</tr><tr><th><strong>Space Complexity</strong></th>
	<td>O(n)</td>
</tr><tr><th><strong>Stability</strong></th>
	<td>Yes</td>
</tr></tbody>
</table>
