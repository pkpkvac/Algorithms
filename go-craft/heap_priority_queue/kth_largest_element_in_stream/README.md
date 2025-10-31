Kth Largest Element in a Stream

Design a class to find the kth largest integer in a stream of values, including duplicates. E.g. the 2nd largest from [1, 2, 3, 3] is 3. The stream is not necessarily sorted.

Implement the following methods:

constructor(int k, int[] nums) Initializes the object given an integer k and the stream of integers nums.
int add(int val) Adds the integer val to the stream and returns the kth largest integer in the stream.
Example 1:

Input:
["KthLargest", [3, [1, 2, 3, 3]], "add", [3], "add", [5], "add", [6], "add", [7], "add", [8]]

Output:
[null, 3, 3, 3, 5, 6]

Explanation:
KthLargest kthLargest = new KthLargest(3, [1, 2, 3, 3]);
kthLargest.add(3); // return 3
kthLargest.add(5); // return 3
kthLargest.add(6); // return 3
kthLargest.add(7); // return 5
kthLargest.add(8); // return 6
Constraints:

1 <= k <= 1000
0 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-1000 <= val <= 1000
There will always be at least k integers in the stream when you search for the kth integer.

Hint 1
A brute force solution would involve sorting the array in every time a number is added using add(), and then returning the k-th largest element. This would take O(m \* nlogn) time, where m is the number of calls to add() and n is the total number of elements added. However, do we really need to track all the elements added, given that we only need the k-th largest element? Maybe you should think of a data structure which can maintain only the top k largest elements.

Hint 2
We can use a Min-Heap, which stores elements and keeps the smallest element at its top. When we add an element to the Min-Heap, it takes O(logk) time since we are storing k elements in it. Retrieving the top element (the smallest in the heap) takes O(1) time. How can this be useful for finding the k-th largest element?

Hint 3
The k-th largest element is the smallest element among the top k largest elements. This means we only need to maintain k elements in our Min-Heap to efficiently determine the k-th largest element. Whenever the size of the Min-Heap exceeds k, we remove the smallest element by popping from the heap. How do you implement this?

Hint 4
We initialize a Min-Heap with the elements of the input array. When the add() function is called, we insert the new element into the heap. If the heap size exceeds k, we remove the smallest element (the root of the heap). Finally, the top element of the heap represents the k-th largest element and is returned.
