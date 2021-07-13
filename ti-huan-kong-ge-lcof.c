char* replaceSpace(char* s){
    int len = strlen(s);
    char* r = malloc(sizeof(char)*len*3+1);
    int i, j;
    for (i=0, j=0; i<len;++i) {
        if (s[i]!=' ') {
            r[j++] = s[i];
        } else {
            r[j++] = '%';
            r[j++] = '2';
            r[j++] = '0'; 
        }
    }
    r[j] = 0;
    return r;
}
