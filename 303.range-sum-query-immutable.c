
typedef struct {
	int* nums;
	} NumArray;

NumArray* numArrayCreate(int* nums, int numsSize) {
	int i;
	NumArray* na = malloc(sizeof(NumArray));
	na->nums = malloc(sizeof(int) * (numsSize + 1));

	int sum = 0;
	na->nums[0] = sum;
	for (i = 1; i <= numsSize; ++i) {
		sum += nums[i - 1];
		na->nums[i] = sum;
		}
	return na;
	}

int numArraySumRange(NumArray* obj, int left, int right) {
	return obj->nums[right + 1] - obj->nums[left];
	}

void numArrayFree(NumArray* obj) {
	free(obj->nums);
	free(obj);
	}
