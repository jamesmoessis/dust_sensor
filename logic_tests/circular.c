/********************************************
 * Logic Practise
 * 
 * Dust Sensor Logic Test
 * 
 * Array that loops through itself infinitely
 * Takes average of last 10 measurements
 * Alerts if over threshold
 @Author: James Moessis
*********************************************/
#include <stdio.h>

int main(void) {
  long int circle[] = {0,0,0,0,0,0,0,0,0,0};/* size 10 */
  int sum;
  int i = 0;
  int circle_size = sizeof(circle)/sizeof(circle[0]);
  printf("Circle Size: %d", circle_size);
  while(1){
    printf("enter: ");
    scanf("%ld", &circle[i]);
    
    if(i >= 9) {
      i = 0;
    }
    else {
      i++;
    }
    
    int k = 0;
    sum = 0;
    for(k = 0; k < circle_size; k++) {
      sum = sum + circle[k];
    }
    
    float average;
    average = ( (float)sum ) / ( (float)circle_size );
    if (average >= 7) {
      printf("ALERT! ALERT!\ndust = %f\n", average);  
    } 
    else {
      printf("Nothing to worry about.\ndust = %f\n", average); 
    }  
  }
}
