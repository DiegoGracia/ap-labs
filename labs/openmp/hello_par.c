
#include "omp.h"
#include "logger.h"

void main(){
    #pragma omp parallel
    {
        int ID = omp_get_thread_num();
        infof(" Hello %d",ID);
        infof(" world %d \n", ID);
    }
}
