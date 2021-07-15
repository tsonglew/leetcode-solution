typedef struct {
    int* inStack;
    int* ouStack;
    int inStackLen;
    int ouStackLen;
} CQueue;


CQueue* cQueueCreate() {
    CQueue* cq = malloc(sizeof(CQueue));
    cq->inStack = malloc(sizeof(int)*10000);
    cq->ouStack = malloc(sizeof(int)*10000);
    cq->inStackLen = 0;
    cq->ouStackLen = 0;
    return cq;
}

void cQueueAppendTail(CQueue* obj, int value) {
    obj->inStack[obj->inStackLen++] = value;
}

int cQueueDeleteHead(CQueue* obj) {
    if (obj->ouStackLen>0) {
        return obj->ouStack[--obj->ouStackLen];
    }
    while (obj->inStackLen>0) {
        obj->ouStack[obj->ouStackLen++] = obj->inStack[--obj->inStackLen];
    }
    if (obj->ouStackLen>0) {
        return obj->ouStack[--obj->ouStackLen];
    }
    return -1;
}

void cQueueFree(CQueue* obj) {
}

/**
 * Your CQueue struct will be instantiated and called as such:
 * CQueue* obj = cQueueCreate();
 * cQueueAppendTail(obj, value);
 
 * int param_2 = cQueueDeleteHead(obj);
 
 * cQueueFree(obj);
*/
