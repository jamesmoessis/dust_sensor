#include <stdio.h>

int main(void){
    printf("This code is built when dust_sensor repo gets committed to master.\n"
            "GitHub POSTs to Jenkins server, which triggers a build!");
    return 0;
}
    
