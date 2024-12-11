// JavaScript program to find second largest element in an array
// using Sorting

// function to find the second largest element
function getSecondLargest(arr) {
    let n = arr.length;

    // Sort the array in non-decreasing order
    arr.sort((a, b) => a - b);

    // start from second last element as last element is the largest
    for (let i = n - 2; i >= 0; i--) {

        // return the first element which is not equal to the 
        // largest element
        if (arr[i] !== arr[n - 1]) {
            return arr[i];
        }
    }

    // If no second largest element was found, return -1
    return -1;
}

const arr = [12, 35, 1, 10, 3, 1];
console.log(getSecondLargest(arr));