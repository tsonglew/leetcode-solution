/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) { 
    pairs = new Map()
    for (var i=0; i<nums.length; i++) {
         if (!pairs.has(nums[i])) {
             pairs.set(target-nums[i], i)
         } else {
             return [pairs.get(nums[i]), i]
         }
    }
};